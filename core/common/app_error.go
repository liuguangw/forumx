package common

//AppError 表示应用出现的错误
type AppError struct {
	Code    int    //错误码
	Message string //错误消息
}

//Error 实现error接口
func (err *AppError) Error() string {
	return err.Message
}

//NewAppError 构造一个错误对象
func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}
