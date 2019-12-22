/*
 + ------------------------------------------------+
 | Author: Zoueature                               |
 + ------------------------------------------------+
 | Email: zoueature@gmail.com                      |
 + ------------------------------------------------+
 | Date: 2019/12/20                                |
 + ------------------------------------------------+
 | Time: 16:39                                     |
 + ------------------------------------------------+
 | Description:                                    |
 + ------------------------------------------------+
*/

package protocol

import (
	"errors"
	"strconv"
	"strings"
)

const (
	StatusReply    = 1
	ErrorReply     = 2
	IntReply       = 3
	BulkReply      = 4
	MultiBulkReply = 5
)

func ParseRedisProtocol(response string) Response {
	respArr := strings.Split(response, "\r\n")
	respArr = respArr[0:len(respArr)-1]
	resultStr := respArr[0]
	replyType := getResponseType(resultStr[0:1])
	result := &RespString{}
	result.replyType = replyType
	switch replyType {
	case StatusReply:
		result.Result = resultStr[1:]
	case ErrorReply:
		result.err = errors.New(resultStr[1:])
	case IntReply:
		resultInt := &RespInt{}
		resultInt.replyType = replyType
		num, err := strconv.Atoi(resultStr[1:])
		if err != nil {
			resultInt.err = err
		}
		resultInt.Result = num
		return resultInt
	case BulkReply:
		result.replyType = replyType
		result.Result = respArr[1]
		return result
	case MultiBulkReply:
		length, err := strconv.Atoi(resultStr[1:])
		if err != nil {
			return &Resp{err:err, replyType:replyType}
		}
		resultSlice := &RespStringSlice{Resp{err:nil, replyType:replyType}, make([]string, 0, length)}
		keys := make([]string, 0, length)
		for index, value := range respArr {
			if (index % 2 ) == 1 {
				keys = append(keys, value)
			} else if (index % 2) == 0 && index != 0 {
				resultSlice.Result = append(resultSlice.Result, value)
			}
		}
		for index, value := range keys {
			valueLen := value[1:]
			vLen, _ := strconv.Atoi(valueLen)
			if vLen > 0 {
				resultSlice.Result[index] = resultSlice.Result[index][0:vLen]
			} else if vLen == 0 {
				resultSlice.Result[index] = ""
			}
		}
		return resultSlice
	default:
		result.err = errors.New("Error reply type ")

	}
	return result
}

func getResponseType(t string) int {
	var replyType int
	switch t {
	case "+":
		replyType = StatusReply
	case "-":
		replyType = ErrorReply
	case ":":
		replyType = IntReply
	case "$":
		replyType = BulkReply
	case "*":
		replyType = MultiBulkReply
	}
	return replyType
}
