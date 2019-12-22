/*+---------------------------------+
 *| Author: Zoueature               |
 *+---------------------------------+
 *| Email: zoueature@gmail.com      |
 *+---------------------------------+
 *| Date: 2019-12-21 22:23          |
 *+---------------------------------+
 */

package earedis

import (
	"strconv"
)

const (
	responseOK = "OK"
)

func (redis *EaRedis) Set(key, value string) *StringResult {
	result := new(StringResult)
	result.cmd = commandLog(redisSET, key, value)
	res, err := redis.command(redisSET, key, value)
	if err != nil {
		result.err = err
		return result
	}
	result.parseProtocol(res, responseOK)
	return result
}

func (redis *EaRedis) SetNx(key, value string) *IntResult {
	result := new(IntResult)
	result.cmd = commandLog(redisSETNX, key, value)
	res, err := redis.command(redisSETNX, key, value)
	if err != nil {
		result.err = err
		return result
	}
	result.parseProtocol(res)
	if result.response > 0 {
		result.result = true
	}
	return result
}

func (redis *EaRedis) SetEx(key, value string, expireTimeSeconds int) *StringResult {
	result := new(StringResult)
	expireTimeStr := strconv.Itoa(expireTimeSeconds)
	result.cmd = commandLog(redisSETEX, key, value, expireTimeStr)
	res, err := redis.command(redisSETEX, key, value, expireTimeStr)
	if err != nil {
		result.err = err
		return result
	}
	result.parseProtocol(res, responseOK)
	return result
}

func (redis *EaRedis) PSetEx(key, value string, expireTimeMilliseconds int) *StringResult {
	result := new(StringResult)
	expireTimeStr := strconv.Itoa(expireTimeMilliseconds)
	result.cmd = commandLog(redisPSETEX, key, value, expireTimeStr)
	res, err := redis.command(redisPSETEX, key, value, expireTimeStr)
	if err != nil {
		result.err = err
		return result
	}
	result.parseProtocol(res, responseOK)
	return result
}

func (redis *EaRedis) Get(key string) *StringResult {
	result := new(StringResult)
	result.cmd = commandLog(redisGET, key)
	res, err := redis.command(redisGET, key)
	if err != nil {
		result.err = err
		return result
	}
	result.parseProtocol(res, "")
	return result
}
func (redis *EaRedis) GetSet(key, value string) *StringResult {
	result := new(StringResult)
	result.cmd = commandLog(redisGETSET, key, value)
	res, err := redis.command(redisGETSET, key, value)
	if err != nil {
		result.err = err
		return result
	}
	result.parseProtocol(res, "")
	return result
}
func (redis *EaRedis) StrLen(key string) *IntResult {
	result := new(IntResult)
	result.cmd = commandLog(redisSTRLEN, key)
	res, err := redis.command(redisSTRLEN, key)
	if err != nil {
		result.err = err
		return result
	}
	result.result = true
	result.parseProtocol(res)
	return result

}
func (redis *EaRedis) Append(key, value string) *IntResult {
	result := new(IntResult)
	result.cmd = commandLog(redisAPPEND, key, value)
	res, err := redis.command(redisAPPEND, key, value)
	if err != nil {
		result.err = nil
		return result
	}
	result.parseProtocol(res)
	if result.response > 0 {
		result.result = true
	}
	return result
}
func (redis *EaRedis) SetRange(key string, offset int, value string) *IntResult {
	result := new(IntResult)
	offsetStr := strconv.Itoa(offset)
	result.cmd = commandLog(redisSETRANGE, key, offsetStr, value)
	res, err := redis.command(redisSETRANGE, key, offsetStr, value)
	if err != nil {
		result.err = err
		return result
	}
	result.parseProtocol(res)
	if result.response > 0 {
		result.result = true
	}
	return result
}
func (redis *EaRedis) GetRange(key string, start, end int) *StringResult {
	result := new(StringResult)
	startStr := strconv.Itoa(start)
	endStr := strconv.Itoa(end)
	result.cmd = commandLog(redisGETRANGE, key, startStr, endStr)
	res, err := redis.command(redisGETRANGE, key, startStr, endStr)
	if err != nil {
		result.err = err
		return result
	}
	result.parseProtocol(res, "")
	return result
}
func (redis *EaRedis) Incr(key string) *IntResult {
	result := new(IntResult)
	result.cmd = commandLog(redisINCR, key)
	res, err := redis.command(redisINCR, key)
	if err != nil {
		result.err = err
		return result
	}
	result.result = true
	result.parseProtocol(res)
	return result
}
func (redis *EaRedis) IncrBy(key string, step int) *IntResult {
	result := new(IntResult)
	stepStr := strconv.Itoa(step)
	result.cmd = commandLog(redisINCRBY, key, stepStr)
	res, err := redis.command(redisINCRBY, key, stepStr)
	if err != nil {
		result.err = err
		return result
	}
	result.result = true
	result.parseProtocol(res)
	return result
}
func (redis *EaRedis) IncrByFloat(key string, step float64) *FloatResult {
	result := new(FloatResult)
	stepStr := strconv.FormatFloat(step, 'f', -1, 64)
	result.cmd = commandLog(redisINCRBYFLOAT, key, stepStr)
	res, err := redis.command(redisINCRBYFLOAT, key, stepStr)
	if err != nil {
		result.err = err
		return result
	}
	result.result = true
	result.parseProtocol(res)
	return result
}

func (redis *EaRedis) Decr(key string) *IntResult {
	result := new(IntResult)
	result.cmd = commandLog(redisDECR, key)
	res, err := redis.command(redisDECR, key)
	if err != nil {
		result.err = err
		return result
	}
	result.result = true
	result.parseProtocol(res)
	return result
}
func (redis *EaRedis) DecrBy(key string, step int) *IntResult {
	result := new(IntResult)
	stepStr := strconv.Itoa(step)
	result.cmd = commandLog(redisDECRBY, key, stepStr)
	res, err := redis.command(redisDECRBY, key, stepStr)
	if err != nil {
		result.err = err
		return result
	}
	result.result = true
	result.parseProtocol(res)
	return result
}
func (redis *EaRedis) MSet(values map[string]string) *StringResult {
	setValues := make([]string, 1, len(values) * 2 + 1)
	setValues[0] = redisMSET
	for setKey, setValue := range values {
		setValues = append(setValues, setKey, setValue)
	}
	result := new(StringResult)
	result.cmd = commandLog(setValues...)
	res, err := redis.command(redisMSET, setValues[1:]...)
	if err != nil {
		result.err = err
		return result
	}
	result.parseProtocol(res, responseOK)
	return result

}
func (redis *EaRedis) MSetNx(values map[string]string) *IntResult {
	setValues := make([]string, 1, len(values) * 2 + 1)
	setValues[0] = redisMSETNX
	for setKey, setValue := range values {
		setValues = append(setValues, setKey, setValue)
	}
	result := new(IntResult)
	result.cmd = commandLog(setValues...)
	res, err := redis.command(redisMSETNX, setValues[1:]...)
	if err != nil {
		result.err = err
		return result
	}
	result.parseProtocol(res)
	if result.response > 0 {
		result.result = true
	}
	return result
}
func (redis *EaRedis) MGet(keys []string) *StringMapResult {
	result := new(StringMapResult)
	tmpSlice := make([]string, 1, len(keys) + 1)
	tmpSlice[0] = redisMGET
	tmpSlice = append(tmpSlice, keys...)
	result.cmd = commandLog(tmpSlice...)
	res, err := redis.command(redisMGET, keys...)
	if err != nil {
		result.err = err
		return result
	}
	result.result = true
	result.parseProtocol(res, keys)
	return result
}
