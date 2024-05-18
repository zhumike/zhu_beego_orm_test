package tools

/**
 * @Author: zhenzhongzhu
 * @Description:
 * @File: rspbody
 * @Date: 2024/5/18 16:44
 */

type ResponseMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempy"`
}

/**
 * 成功时候返回
 * @Params data 返回的数据
 */
func Successed(data interface{}) *ResponseMsg {
	return &ResponseMsg{Success.Code, Success.Msg, data}
}

/**
 * 失败时候返回
 * @Params msg 返回的异常信息
 */
func Failures(msg string) *ResponseMsg {
	return &ResponseMsg{Failure.Code, msg, nil}
}

/**
 * 自定义返回
 * @Params code 返回的异常码
 * @Params msg 返回的异常信息
 */
func Customize(code int, msg string) *ResponseMsg {
	return &ResponseMsg{code, msg, nil}
}
