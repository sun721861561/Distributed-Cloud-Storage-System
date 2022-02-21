package common

// ErrorCode : 错误码
type ErrorCode int32

const (
	_ int32 = iota + 9999
	// StatusOK : Normal
	StatusOK
	// StatusParamInvalid : Invalid request parameter
	StatusParamInvalid
	// StatusServerError : Service Error
	StatusServerError
	// StatusRegisterFailed : Registration failed
	StatusRegisterFailed
	// StatusLoginFailed : Login failed
	StatusLoginFailed
	// StatusTokenInvalid : 10005 Invalid token
	StatusTokenInvalid
	// StatusUserNotExists: 10006 User does not exist
	StatusUserNotExists
)
