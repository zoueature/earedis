/*+---------------------------------+
 *| Author: Zoueature               |
 *+---------------------------------+
 *| Email: zoueature@gmail.com      |
 *+---------------------------------+
 *| Date: 2019-12-20 21:14          |
 *+---------------------------------+
 */

package protocol

import (
	"fmt"
	"testing"
)

func TestParseRedisProtocol(t *testing.T) {
	responseStr := "+OK\r\n"
	response := ParseRedisProtocol(responseStr)
	if response.Error() != nil || response.GetResult() != "OK" || response.ReplyType() != StatusReply {
		t.Error("parse status reply error")
	}
	responseStr = "-Error:unknown command\r\n"
	response = ParseRedisProtocol(responseStr)
	if response.Error() == nil || response.Error().Error() != "Error:unknown command" || response.ReplyType() != ErrorReply {
		t.Error("parse error reply error")
	}
	responseStr = ":2834\r\n"
	response = ParseRedisProtocol(responseStr)
	if response.Error() != nil || response.GetResult() != 2834 || response.ReplyType() != IntReply {
		t.Error("parse int reply error")
	}
	responseStr = "$1\r\n1"
	response = ParseRedisProtocol(responseStr)
	if response.Error() != nil || response.GetResult() != "1" || response.ReplyType() != BulkReply {
		fmt.Println(response.Error())
		t.Error("parse bulk reply error")
	}
	responseStr = "*3\r\n$1\r\n1\r\n$1\r\n2\r\n$2\r\n3\r\n"
	response = ParseRedisProtocol(responseStr)
	if response.Error() != nil || len(response.GetResult().([]string)) != 3 || response.ReplyType() != MultiBulkReply {
		fmt.Println(response.GetResult())
		t.Error("parse bulk reply error")
	}
}
