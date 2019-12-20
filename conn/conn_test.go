/*+---------------------------------+
 *| Author: Zoueature               |
 *+---------------------------------+
 *| Email: zoueature@gmail.com      |
 *+---------------------------------+
 *| Date: 2019-12-20 20:19          |
 *+---------------------------------+
 */

package conn

import "testing"

func TestConnect(t *testing.T) {
	connOpt := ConnectOption{
		"127.0.0.1",
		"6379",
		16,
		4,
		0.5,
	}
	_ := Connect(&connOpt)
}
