/*+---------------------------------+
 *| Author: Zoueature               |
 *+---------------------------------+
 *| Email: zoueature@gmail.com      |
 *+---------------------------------+
 *| Date: 2019-12-21 22:05          |
 *+---------------------------------+
 */

package earedis

import (
	"errors"
	"github.com/zoueature/earedis/protocol"
	"strconv"
	"strings"
)

const (
	redisSET               = "SET"
	redisSETNX             = "SETNX"
	redisSETEX             = "SETEX"
	redisPSETEX            = "PSETEX"
	redisGET               = "GET"
	redisGETSET            = "GETSET"
	redisSTRLEN            = "STRLEN"
	redisAPPEND            = "APPEND"
	redisSETRANGE          = "SETRANGE"
	redisGETRANGE          = "GETRANGE"
	redisINCR              = "INCR"
	redisINCRBY            = "INCRBY"
	redisINCRBYFLOAT       = "INCRBYFLOAT"
	redisDECR              = "DECR"
	redisDECRBY            = "DECRBY"
	redisMSET              = "MSET"
	redisMSETNX            = "MSETNX"
	redisMGET              = "MGET"
	redisHSET              = "HSET"
	redisHSETNX            = "HSETNX"
	redisHGET              = "HGET"
	redisHEXISTS           = "HEXISTS"
	redisHDEL              = "HDEL"
	redisHLEN              = "HLEN"
	redisHSTRLEN           = "HSTRLEN"
	redisHINCRBY           = "HINCRBY"
	redisHINCRBYFLOAT      = "HINCRBYFLOAT"
	redisHMSET             = "HMSET"
	redisHMGET             = "HMGET"
	redisHKEYS             = "HKEYS"
	redisHVALS             = "HVALS"
	redisHGETALL           = "HGETALL"
	redisHSCAN             = "HSCAN"
	redisLPUSH             = "LPUSH"
	redisLPUSHX            = "LPUSHX"
	redisRPUSH             = "RPUSH"
	redisRPUSHX            = "RPUSHX"
	redisLPOP              = "LPOP"
	redisRPOP              = "RPOP"
	redisRPOPLPUSH         = "RPOPLPUSH"
	redisLREM              = "LREM"
	redisLLEN              = "LLEN"
	redisLINDEX            = "LINDEX"
	redisLINSERT           = "LINSERT"
	redisLSET              = "LSET"
	redisLRANGE            = "LRANGE"
	redisLTRIM             = "LTRIM"
	redisBLPOP             = "BLPOP"
	redisBRPOP             = "BRPOP"
	redisBRPOPLPUSH        = "BRPOPLPUSH"
	redisSADD              = "SADD"
	redisSISMEMBER         = "SISMEMBER"
	redisSPOP              = "SPOP"
	redisSRANDMEMBER       = "SRANDMEMBER"
	redisSREM              = "SREM"
	redisSMOVE             = "SMOVE"
	redisSCARD             = "SCARD"
	redisSMEMBERS          = "SMEMBERS"
	redisSSCAN             = "SSCAN"
	redisSINTER            = "SINTER"
	redisSINTERSTORE       = "SINTERSTORE"
	redisSUNION            = "SUNION"
	redisSUNIONSTORE       = "SUNIONSTORE"
	redisSDIFF             = "SDIFF"
	redisSDIFFSTORE        = "SDIFFSTORE"
	redisZADD              = "ZADD"
	redisZSCORE            = "ZSCORE"
	redisZINCRBY           = "ZINCRBY"
	redisZCARD             = "ZCARD"
	redisZCOUNT            = "ZCOUNT"
	redisZRANGE            = "ZRANGE"
	redisZREVRANGE         = "ZREVRANGE"
	redisZRANGEBYSCORE     = "ZRANGEBYSCORE"
	redisZREVRANGEBYSCORE  = "ZREVRANGEBYSCORE"
	redisZRANK             = "ZRANK"
	redisZREVRANK          = "ZREVRANK"
	redisZREM              = "ZREM"
	redisZREMRANGEBYRANK   = "ZREMRANGEBYRANK"
	redisZREMRANGEBYSCORE  = "ZREMRANGEBYSCORE"
	redisZRANGEBYLEX       = "ZRANGEBYLEX"
	redisZLEXCOUNT         = "ZLEXCOUNT"
	redisZREMRANGEBYLEX    = "ZREMRANGEBYLEX"
	redisZSCAN             = "ZSCAN"
	redisZUNIONSTORE       = "ZUNIONSTORE"
	redisZINTERSTORE       = "ZINTERSTORE"
	redisPFADD             = "PFADD"
	redisPFCOUNT           = "PFCOUNT"
	redisPFMERGE           = "PFMERGE"
	redisGEOADD            = "GEOADD"
	redisGEOPOS            = "GEOPOS"
	redisGEODIST           = "GEODIST"
	redisGEORADIUS         = "GEORADIUS"
	redisGEORADIUSBYMEMBER = "GEORADIUSBYMEMBER"
	redisGEOHASH           = "GEOHASH"
	redisSETBIT            = "SETBIT"
	redisGETBIT            = "GETBIT"
	redisBITCOUNT          = "BITCOUNT"
	redisBITPOS            = "BITPOS"
	redisBITOP             = "BITOP"
	redisBITFIELD          = "BITFIELD"
	redisEXISTS            = "EXISTS"
	redisTYPE              = "TYPE"
	redisRENAME            = "RENAME"
	redisRENAMENX          = "RENAMENX"
	redisMOVE              = "MOVE"
	redisDEL               = "DEL"
	redisRANDOMKEY         = "RANDOMKEY"
	redisDBSIZE            = "DBSIZE"
	redisKEYS              = "KEYS"
	redisSCAN              = "SCAN"
	redisSORT              = "SORT"
	redisFLUSHDB           = "FLUSHDB"
	redisFLUSHALL          = "FLUSHALL"
	redisSELECT            = "SELECT"
	redisSWAPDB            = "SWAPDB"
	redisEXPIRE            = "EXPIRE"
	redisEXPIREAT          = "EXPIREAT"
	redisTTL               = "TTL"
	redisPERSIST           = "PERSIST"
	redisPEXPIRE           = "PEXPIRE"
	redisPEXPIREAT         = "PEXPIREAT"
	redisPTTL              = "PTTL"
	redisMULTI             = "MULTI"
	redisEXEC              = "EXEC"
	redisDISCARD           = "DISCARD"
	redisWATCH             = "WATCH"
	redisUNWATCH           = "UNWATCH"
	redisPUBLISH           = "PUBLISH"
	redisSUBSCRIBE         = "SUBSCRIBE"
	redisPSUBSCRIBE        = "PSUBSCRIBE"
	redisUNSUBSCRIBE       = "UNSUBSCRIBE"
	redisPUNSUBSCRIBE      = "PUNSUBSCRIBE"
	redisPUBSUB            = "PUBSUB"
	redisAUTH              = "AUTH"
	redisINFO              = "INFO"
)

type Result struct {
	err    error
	cmd    string //the redis command
	result bool   //the command exec result , success is true and fail is false
}

type StringResult struct {
	Result
	response string //redis response
}

type StringMapResult struct {
	Result
	response map[string]string
}

type IntResult struct {
	Result
	response int //redis response
}

type FloatResult struct {
	Result
	response float64
}

func commandLog(parameters ...string) string {
	command := strings.Join(parameters, " ")
	return command
}

// ---------- base result method ----------------

func (result *Result) Error() error {
	return result.err
}

func (result *Result) Success() bool {
	return result.result
}

func (result *Result) Command() string {
	return result.cmd
}

// ---------- string result method ---------------
func (result *StringResult) Value() string {
	return result.response
}

func (result *StringResult) parseProtocol(str string, succIdentify string) {
	response := protocol.ParseRedisProtocol(str)
	if response.Error() != nil {
		result.err = response.Error()
		result.result = false
		return
	}
	responseStr, ok := response.GetResult().(string)
	if !ok {
		result.err = errors.New("Error response type ")
		result.result = false
		return
	}
	result.response = responseStr
	if (succIdentify != "" && responseStr == succIdentify) || succIdentify == "" {
		result.result = true
	}
}

// ------------ integer result method --------------
func (result *IntResult) Value() int {
	return result.response
}

func (result *IntResult) parseProtocol(str string) {
	response := protocol.ParseRedisProtocol(str)
	if response.Error() != nil {
		result.err = response.Error()
		result.result = false
		return
	}
	responseInt, ok := response.GetResult().(int)
	if !ok {
		result.err = errors.New("Error response type ")
		result.result = false
		return
	}
	result.response = responseInt
}

// ------------ float result method ------------
func (result *FloatResult) Value() float64 {
	return result.response
}

func (result *FloatResult) parseProtocol(str string) {
	response := protocol.ParseRedisProtocol(str)
	if response.Error() != nil {
		result.err = response.Error()
		result.result = false
		return
	}
	responseStr, ok := response.GetResult().(string)
	if !ok {
		result.err = errors.New("Error response type ")
		result.result = false
		return
	}
	responseFloat, err := strconv.ParseFloat(responseStr, 64)
	if err != nil {
		result.err = err
		result.result = false
		return
	}
	result.response = responseFloat
}

// ------------- string map result method -------------

func (result *StringMapResult) Value() map[string]string {
	return result.response
}

func (result *StringMapResult) parseProtocol(str string, keys []string) {
	response := protocol.ParseRedisProtocol(str)
	if response.Error() != nil {
		result.err = response.Error()
		result.result = false
		return
	}
	responseStrSlice, ok := response.GetResult().([]string)
	if !ok {
		result.err = errors.New("Error response type ")
		result.result = false
		return
	}
	result.response = make(map[string]string)
	for index, key := range keys {
		var value string
		if index >= len(responseStrSlice) {
			continue
		} else {
			value = responseStrSlice[index]
		}
		result.response[key] = value
	}
}