/*+---------------------------------+
 *| Author: Zoueature               |
 *+---------------------------------+
 *| Email: zoueature@gmail.com      |
 *+---------------------------------+
 *| Date: 2019-12-21 22:53          |
 *+---------------------------------+
 */

package earedis

import (
	"github.com/zoueature/earedis/conn"
	"strconv"
	"testing"
)

func getTestRedis() *EaRedis {
	connOpt := conn.ConnectOption{
		Host:     "127.0.0.1",
		Port:     "6379",
		MaxConn:  16,
		IdleConn: 4,
		Timeout:  500,
	}
	eaRedis, err := Connect(&connOpt)
	if err != nil {
		return nil
	}
	return eaRedis
}

func TestEaRedis_Set(t *testing.T) {
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
	result := eaRedis.Set("test", "this is a test")
	if result.err != nil {
		t.Error("Set error : " + result.err.Error())
	}
	if result.result != true {
		t.Error("Set error , result is not true")
	}
	if result.cmd != "SET test this is a test" {
		t.Error("Set error, command error : " + result.cmd)
	}
	if result.response != "OK" {
		t.Error("Set error, response is not ok, is " + result.response)
	}
	getResult := eaRedis.GET("test")
	if getResult.response != "this is a test" {
		t.Error("Set is not successfully, value is :" + getResult.response)
	}
}

func TestEaRedis_SetNx(t *testing.T) {
	eaRedis := getTestRedis()
	if eaRedis == nil {
		return
	}
	result := eaRedis.SetNx("test", "123")
	if result.response != 0 {
		t.Error("Setnx key error, response is " + strconv.Itoa(result.response) + ",need fail but success")
	}
	if result.result == true {
		t.Error("Setnx key error, result is true, need fail but success")
	}
	succResult := eaRedis.SetNx("succSetNx", "1234567")
	if succResult.response != 1 {
		t.Error("Setnx key error, response is " + strconv.Itoa(result.response) + ",need success but fail")
	}
	if succResult.result == false {
		t.Error("Setnx key error, result is false, need fail but success")
	}
	getResult := eaRedis.GET("succSetNx")
	if getResult.response != "1234567" {
		t.Error("Setnx error, response success but value is not set success")
	}
}
