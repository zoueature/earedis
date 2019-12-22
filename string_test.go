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
	getResult := eaRedis.Get("test")
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
	getResult := eaRedis.Get("succSetNx")
	if getResult.response != "1234567" {
		t.Error("Setnx error, response success but value is not set success")
	}
}

func TestEaRedis_Incr(t *testing.T) {
	eaRedis := getTestRedis()
	if eaRedis == nil {
		return
	}
	setResult := eaRedis.Set("test", "-2")
	if !setResult.result {
		return
	}
	result := eaRedis.Incr("test")
	if result.err != nil {
		t.Error("incr error, err : " + result.err.Error())
	}
	if !result.result {
		t.Error("incr error, result is false")
	}
	if result.response != -1 {
		t.Error("incr error, response is error , expect -1 but return " + strconv.Itoa(result.response))
	}
	if result.cmd != "INCR test" {
		t.Error("Incr cmd error, expect INCR test but " + result.cmd + " given")
	}
	eaRedis.Set("test", "cadca")
	errResult := eaRedis.Incr("test")
	if errResult.err == nil {
		t.Error("Incr error, expect error but succ")
	}
	if errResult.result {
		t.Error("Incr error, expect error but succ")
	}
	if errResult.response != 0 {
		t.Error("Incr error, expect error but succ")
	}
}

func TestEaRedis_IncrBy(t *testing.T) {
	eaRedis := getTestRedis()
	if eaRedis == nil {
		return
	}
	setResult := eaRedis.Set("test", "-2")
	if !setResult.result {
		return
	}
	result := eaRedis.IncrBy("test", 1)
	if result.err != nil {
		t.Error("incr error, err : " + result.err.Error())
	}
	if !result.result {
		t.Error("incr error, result is false")
	}
	if result.response != -1 {
		t.Error("incr error, response is error , expect -1 but return " + strconv.Itoa(result.response))
	}
	if result.cmd != "INCRBY test 1" {
		t.Error("Incr cmd error, expect INCR test but " + result.cmd + " given")
	}
	eaRedis.Set("test", "cadca")
	errResult := eaRedis.IncrBy("test", 1)
	if errResult.err == nil {
		t.Error("Incr error, expect error but succ")
	}
	if errResult.result {
		t.Error("Incr error, expect error but succ")
	}
	if errResult.response != 0 {
		t.Error("Incr error, expect error but succ")
	}
}

func TestEaRedis_IncrByFloat(t *testing.T) {
	eaRedis := getTestRedis()
	if eaRedis == nil {
		return
	}
	setResult := eaRedis.Set("test", "-2")
	if !setResult.result {
		return
	}
	result := eaRedis.IncrByFloat("test", 1.1)
	if result.err != nil {
		t.Error("incr error, err : " + result.err.Error())
	}
	if !result.result {
		t.Error("incr error, result is false")
	}
	if result.response != -0.9 {
		t.Error("incr error, response is error , expect -1 but return " + strconv.FormatFloat(result.response, 'f', -1, 64))
	}
	if result.cmd != "INCRBYFLOAT test 1.1" {
		t.Error("Incr cmd error, expect INCR test but " + result.cmd + " given")
	}
	eaRedis.Set("test", "cadca")
	errResult := eaRedis.IncrBy("test", 1)
	if errResult.err == nil {
		t.Error("Incr error, expect error but succ")
	}
	if errResult.result {
		t.Error("Incr error, expect error but succ")
	}
	if errResult.response != 0 {
		t.Error("Incr error, expect error but succ")
	}
}

func TestEaRedis_Decr(t *testing.T) {
	eaRedis := getTestRedis()
	if eaRedis == nil {
		return
	}
	setResult := eaRedis.Set("test", "-2")
	if !setResult.result {
		return
	}
	result := eaRedis.Decr("test")
	if result.err != nil {
		t.Error("incr error, err : " + result.err.Error())
	}
	if !result.result {
		t.Error("incr error, result is false")
	}
	if result.response != -3 {
		t.Error("incr error, response is error , expect -1 but return " + strconv.Itoa(result.response))
	}
	if result.cmd != "DECR test" {
		t.Error("Incr cmd error, expect INCR test but " + result.cmd + " given")
	}
	eaRedis.Set("test", "cadca")
	errResult := eaRedis.Incr("test")
	if errResult.err == nil {
		t.Error("Incr error, expect error but succ")
	}
	if errResult.result {
		t.Error("Incr error, expect error but succ")
	}
	if errResult.response != 0 {
		t.Error("Incr error, expect error but succ")
	}
}

func TestEaRedis_DecrBy(t *testing.T) {
	eaRedis := getTestRedis()
	if eaRedis == nil {
		return
	}
	setResult := eaRedis.Set("test", "-2")
	if !setResult.result {
		return
	}
	result := eaRedis.DecrBy("test", 4)
	if result.err != nil {
		t.Error("incr error, err : " + result.err.Error())
	}
	if !result.result {
		t.Error("incr error, result is false")
	}
	if result.response != -6 {
		t.Error("incr error, response is error , expect -1 but return " + strconv.Itoa(result.response))
	}
	if result.cmd != "DECRBY test 4" {
		t.Error("Incr cmd error, expect INCR test but " + result.cmd + " given")
	}
	eaRedis.Set("test", "cadca")
	errResult := eaRedis.Incr("test")
	if errResult.err == nil {
		t.Error("Incr error, expect error but succ")
	}
	if errResult.result {
		t.Error("Incr error, expect error but succ")
	}
	if errResult.response != 0 {
		t.Error("Incr error, expect error but succ")
	}
}

func TestEaRedis_MSet(t *testing.T) {
	eaRedis := getTestRedis()
	if eaRedis == nil {
		return
	}
	result := eaRedis.MSet(map[string]string{"a":"123", "b":"456"})
	if result.err != nil {
		t.Error("mset error")
	}
	if !result.result {
		t.Error("mset error")
	}
	if result.response != "OK" {
		t.Error("Mset error")
	}
	if result.cmd != "MSET a 123 b 456" {
		t.Error("mset error")
	}
	a := eaRedis.Get("a")
	if a.response != "123" {
		t.Error("mset error")
	}
	b := eaRedis.Get("b")
	if b.response != "456" {
		t.Error("mset error")
	}

}

func TestEaRedis_MSetNx(t *testing.T) {
	eaRedis := getTestRedis()
	if eaRedis == nil {
		return
	}
	result := eaRedis.MSetNx(map[string]string{"a":"123", "b":"456"})
	if result.err != nil {
		t.Error("mset error")
	}
	if result.result {
		t.Error("mset error")
	}
	if result.response != 0 {
		t.Error("Msetnx error")
	}
	if result.cmd != "MSETNX a 123 b 456" {
		t.Error("mset error :" + result.cmd)
	}
	succResult := eaRedis.MSetNx(map[string]string{"c":"123345"})
	if succResult.response != 1 {
		t.Error("msetnx error")
	}
	if !succResult.result {
		t.Error("msetnx error")
	}
	if succResult.cmd != "MSETNX c 123345" {
		t.Error("msetnx error")
	}
	c := eaRedis.Get("c")
	if c.response != "123345" {
		t.Error("mset error")
	}

}

func TestEaRedis_MGet(t *testing.T) {
	eaRedis := getTestRedis()
	if eaRedis == nil {
		return
	}
	eaRedis.MSet(map[string]string{"a":"1", "b":"2", "c":"3"})
	result := eaRedis.MGet([]string{"a", "b", "c", "d"})
	if result.err != nil {
		t.Error("mget error")
	}
	if !result.result {
		t.Error("mget error")
	}
	if result.cmd != "MGET a b c d" {
		t.Error("mget error")
	}
	if len(result.response) != 3 {
		t.Error("mget error")
	}
	if result.response["a"] != "1" {
		t.Error("mget error")
	}
	if result.response["b"] != "2" {
		t.Error("mget error")
	}
	if result.response["c"] != "3" {
		t.Error("mget error")
	}
	if value, ok := result.response["d"]; ok {
		t.Error("mget error, d is not defined but return " + value)
	}
}