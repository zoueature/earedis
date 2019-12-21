/*
 + ------------------------------------------------+
 | Author: Zoueature                               |
 + ------------------------------------------------+
 | Email: zoueature@gmail.com                      |
 + ------------------------------------------------+
 | Date: 2019/12/16                                |
 + ------------------------------------------------+
 | Time: 21:48                                     |
 + ------------------------------------------------+
 | Description:                                    |
 + ------------------------------------------------+
*/

package conn

import (
	"errors"
	"github.com/zoueature/eapool"
	"io"
	"net"
	"strconv"
	"time"
)

type EaConn struct {
	Conn     *eapool.TcpPool
	maxConn  int
	idleConn int
}

type ConnectOption struct {
	Host     string
	Port     string
	MaxConn  int
	IdleConn int
	Timeout  time.Duration
}

func Connect(options *ConnectOption) (*EaConn, error) {
	redisCons, err := eapool.NewTCPPool(eapool.PoolArg{Max: options.MaxConn, Idle: options.IdleConn}, func() (conn net.Conn, e error) {
		conn, e = net.DialTimeout("tcp", options.Host+":"+options.Port, options.Timeout*time.Millisecond)
		return
	})
	if err != nil {
		return nil, err
	}
	connect := &EaConn{redisCons, options.MaxConn, options.IdleConn}
	return connect, nil
}

func (er *EaConn) SetIdleConn(idleConnNum int) {
	er.idleConn = idleConnNum
	er.Conn.SetIdleConn(idleConnNum)
}

func (er *EaConn) SetMaxConn(maxConnNum int) {
	er.maxConn = maxConnNum
	er.Conn.SetMaxConn(maxConnNum)
}

func (er *EaConn) Command(command string, parameters ...string) (string, error) {
	conn, err := er.Conn.GetConn(false)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	commandLen := len(command)
	allParaNum := len(parameters) + 1
	sendStr := "*" + strconv.Itoa(allParaNum) + "\r\n" + "$" + strconv.Itoa(commandLen) + "\r\n" + command + "\r\n"
	for _, param := range parameters {
		paramLen := len(param)
		sendStr += "$" + strconv.Itoa(paramLen) + "\r\n" + param + "\r\n"
	}
	size, err := conn.Conn.Write([]byte(sendStr))
	if err != nil {
		return "", err
	}
	if size != len(sendStr) {
		return "", errors.New("Send part of command string , all " + strconv.Itoa(len(sendStr)) + ", sent: " + strconv.Itoa(size))
	}
	response := make([]byte, 0)
	length := 0
	for {
		tmp := make([]byte, 64)
		size, err = conn.Conn.Read(tmp)
		response = append(response, tmp...)
		length += size
		if err == io.EOF || size < 64 {
			break
		}
	}
	return string(response[0:length]), nil
}
