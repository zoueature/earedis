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
	"net"
	"strconv"
	"time"
)

type EaRedis struct {
	conn     *eapool.TcpPool
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

func Connect(options *ConnectOption) *EaRedis {
	redisCons, _ := eapool.NewTCPPool(eapool.PoolArg{Max: options.MaxConn, Idle: options.IdleConn}, func() (conn net.Conn, e error) {
		conn, e = net.DialTimeout("tcp", options.Host+":"+options.Port, options.Timeout*time.Second)
		return
	})
	connect := &EaRedis{redisCons, options.MaxConn, options.IdleConn}
	return connect
}

func (er *EaRedis) SetIdleConn(idleConnNum int) {
	er.idleConn = idleConnNum
	er.conn.SetIdleConn(idleConnNum)
}

func (er *EaRedis) SetMaxConn(maxConnNum int) {
	er.maxConn = maxConnNum
	er.conn.SetMaxConn(maxConnNum)
}

func (er *EaRedis) Command(command string, parameters ...string) (interface{}, error) {
	conn, err := er.conn.GetConn(false)
	if err != nil {
		return nil, err
	}
	commandLen := len(command)
	allParaNum := len(parameters) + 1
	sendStr := "*" + strconv.Itoa(allParaNum) + "\r\n" + "$" + strconv.Itoa(commandLen) + "\r\n" + command + "\r\n"
	for _, param := range parameters {
		paramLen := len(param)
		sendStr += "$" + strconv.Itoa(paramLen) + "\r\n" + param + "\r\n"
	}
	size, err := conn.Conn.Write([]byte(sendStr))
	if err != nil {
		return nil, err
	}
	if size != len(sendStr) {
		return nil, errors.New("Send part of command string , all " + strconv.Itoa(len(sendStr)) + ", sent: " + strconv.Itoa(size))
	}

	var response []byte
	_, _ = conn.Conn.Read(response)
	return response, nil
}
