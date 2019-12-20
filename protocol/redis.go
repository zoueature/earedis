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

import "strings"

const (
	statusReply    = 1
	errorReply     = 2
	intReply       = 3
	bulkReply      = 4
	multiBulkReply = 5
)

func ParseRedisProtocol(response string) interface{} {
	respArr := strings.Split(response, "\r\n")
	for _, resp := range respArr {
		responseType := getResponseType(resp)
		switch responseType {
		case statusReply:

		}
	}
	return nil
}

func getResponseType(resp string) int {
	respType := resp[0:1]
	switch respType {
	case "+":
		return statusReply
	case "-":
		return errorReply
	case ":":
		return intReply
	case "$":
		return bulkReply
	case "*":
		return multiBulkReply
	}
	return 0
}
