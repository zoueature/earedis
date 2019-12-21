/*+---------------------------------+
 *| Author: Zoueature               |
 *+---------------------------------+
 *| Email: zoueature@gmail.com      |
 *+---------------------------------+
 *| Date: 2019-12-21 21:53          |
 *+---------------------------------+
 */

package earedis

import (
	"errors"
	"github.com/zoueature/earedis/conn"
	"io"
	"strconv"
)

type EaRedis struct {
	conn      *conn.EaConn
	logSwitch bool
}

func Connect(opt *conn.ConnectOption) (*EaRedis, error) {
	connect, err := conn.Connect(opt)
	if err != nil {
		return nil, err
	}
	eaRedis := new(EaRedis)
	eaRedis.conn = connect
	return eaRedis, nil
}

func (redis *EaRedis) command(command string, parameters ...string) (string, error) {
	conn, err := redis.conn.Conn.GetConn(false)
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
