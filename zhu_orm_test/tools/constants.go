package tools

/**
 * @Author: zhenzhongzhu
 * @Description:
 * @File: constants
 * @Date: 2024/5/18 16:56
 */

type ErrCode struct {
	Code int
	Msg  string
}

func (err ErrCode) getErrCode() int {
	return err.Code
}

func (err ErrCode) getErrMsg() string {
	return err.Msg
}

var (
	Success = &ErrCode{Code: 200, Msg: "请求成功"}
	Failure = &ErrCode{Code: 500, Msg: "请求失败"}
)
