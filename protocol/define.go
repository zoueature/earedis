/*+---------------------------------+
 *| Author: Zoueature               |
 *+---------------------------------+
 *| Email: zoueature@gmail.com      |
 *+---------------------------------+
 *| Date: 2019-12-20 20:28          |
 *+---------------------------------+
 */

package protocol

type Response interface {
	Error() error
	GetResult() interface{}
	ReplyType() int
}

type Resp struct {
	err error
	replyType int
}

func (resp *Resp) GetResult() interface{} {
	return nil
}

//字符串结果
type RespString struct {
	Resp
	Result string
}

//整形结果
type RespInt struct {
	Resp
	Result int
}

//字符串切片结果
type RespStringSlice struct {
	Resp
	Result []string
}

//map结果
type RespMapString struct {
	Resp
	Result map[string]string
}

func (resp *Resp) Error() error {
	return resp.err
}

func (resp *Resp) ReplyType() int {
	return resp.replyType
}

func (resp *RespInt) GetResult() interface{} {
	return resp.Result
}

func (resp *RespString) GetResult() interface{} {
	return resp.Result
}

func (resp *RespStringSlice) GetResult() interface{} {
	return resp.Result
}

func (resp *RespMapString) GetResult() interface{} {
	return resp.Result
}
