/*+---------------------------------+
 *| Author: Zoueature               |
 *+---------------------------------+
 *| Email: zoueature@gmail.com      |
 *+---------------------------------+
 *| Date: 2019-12-20 20:19          |
 *+---------------------------------+
 */

package conn

import (
	"fmt"
	"github.com/zoueature/earedis/protocol"
	"log"
	"testing"
)

func TestConnect(t *testing.T) {
	connOpt := ConnectOption{
		"127.0.0.1",
		"6379",
		16,
		4,
		500,
	}
	conn, err := Connect(&connOpt)
	if err != nil {
		log.Fatalln(err.Error())
	}
	response, err := conn.Command("hgetall", "hash")
	result := protocol.ParseRedisProtocol(response)
	fmt.Println(result.Error(), result.GetResult())
}
