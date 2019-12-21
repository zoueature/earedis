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
	statusReply    = 1
	errorReply     = 2
	intReply       = 3
	bulkReply      = 4
	multiBulkReply = 5
)

func ParseRedisProtocol(response string) Response {
	respArr := strings.Split(response, "\r\n")
	resultStr := respArr[0]
	replyType := getResponseType(resultStr[0:1])
	result := &RespString{}
	switch replyType {
	case statusReply:
		result.Result = resultStr[1:]
	case errorReply:
		result.err = errors.New(resultStr[1:])
	case intReply:
		resultInt := &RespInt{}
		num, err := strconv.Atoi(resultStr[1:])
		if err != nil {
			resultInt.err = err
		}
		resultInt.Result = num
		return resultInt
	case bulkReply:
		result := &RespString{}
		result.Result = respArr[1]
		return result
	case multiBulkReply:
		length, err := strconv.Atoi(resultStr[1:])
		if err != nil {
			return &Resp{err}
		}
		resultSlice := &RespStringSlice{Resp{nil}, make([]string, 0, length)}
		for index, value := range respArr {
			if (index%2) == 0 && index != 0 && value != "" {
				resultSlice.Result = append(resultSlice.Result, value)
			}
		}
		if len(resultSlice.Result) != length {
			return &Resp{errors.New("Not fill full field ")}
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
		replyType = statusReply
	case "-":
		replyType = errorReply
	case ":":
		replyType = intReply
	case "$":
		replyType = bulkReply
	case "*":
		replyType = multiBulkReply
	}
	return replyType
}
