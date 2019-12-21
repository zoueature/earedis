/*+---------------------------------+
 *| Author: Zoueature               |
 *+---------------------------------+
 *| Email: zoueature@gmail.com      |
 *+---------------------------------+
 *| Date: 2019-12-21 22:46          |
 *+---------------------------------+
 */

package earedis

import (
	"github.com/zoueature/earedis/conn"
	"testing"
)

func TestCommand(t *testing.T) {
	connOpt := conn.ConnectOption{
		Host:     "127.0.0.1",
		Port:     "6379",
		MaxConn:  16,
		IdleConn: 4,
		Timeout:  500,
	}
	eaRedis, err := Connect(&connOpt)
	if err != nil {
		t.Error("Connect error : " + err.Error())
		return
	}
	result, err := eaRedis.command("set", "test", "123")
	if err != nil {
		t.Error("set error : " + err.Error())
	}
	if result != "+OK\r\n" {
		t.Error("Set response error : " + result)
	}
}
