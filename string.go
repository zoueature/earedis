/*+---------------------------------+
 *| Author: Zoueature               |
 *+---------------------------------+
 *| Email: zoueature@gmail.com      |
 *+---------------------------------+
 *| Date: 2019-12-21 22:23          |
 *+---------------------------------+
 */

package earedis

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
	result.parseProtocol(res)
	if result.response == responseOK {
		result.result = true
	}
	return result
}

func (redis *EaRedis) SetNx(key, value string) *IntResult {
	result := new(IntResult)
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

func (redis *EaRedis) SetEx(key, value string) {
}

func (redis *EaRedis) PSetEx(key, value string) {
}

func (redis *EaRedis) GET(key string) *StringResult {
	result := new(StringResult)
	res, err := redis.command(redisGET, key)
	if err != nil {
		result.err = err
		return result
	}
	result.result = true
	result.parseProtocol(res)
	return result
}
func (redis *EaRedis) GETSET(key, value string) {
}
func (redis *EaRedis) STRLEN(key, value string) {
}
func (redis *EaRedis) APPEND(key, value string) {
}
func (redis *EaRedis) SETRANGE(key, value string) {
}
func (redis *EaRedis) GETRANGE(key, value string) {
}
func (redis *EaRedis) INCR(key, value string) {
}
func (redis *EaRedis) INCRBY(key, value string) {
}
func (redis *EaRedis) INCRBYFLOAT(key, value string) {
}
func (redis *EaRedis) DECR(key, value string) {
}
func (redis *EaRedis) DECRBY(key, value string) {
}
func (redis *EaRedis) MSET(key, value string) {
}
func (redis *EaRedis) MSETNX(key, value string) {
}
func (redis *EaRedis) MGET(key, value string) {
}
