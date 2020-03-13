package common

type ApiError struct {
	Code	  string
	Msg 	  string
}

/**
	Error Code
	(part)-(module)-(code).(level)
	part : [
		1 = SDK Error
		2 = Invalid Input
		3 = Service Error
		4 = Database Error
		5 = System Error
	]
	module : [00:99]
	code : [000:999]
	level : [0:9]
 */
var (
	SDK_INITIAL_ERROR = ApiError{"1-00-000.9", "SDK initial failed."}
	VERIFY_ERROR = ApiError{"2-00-000.1", "Params verify failed."}

	SYSTEM_ERROR = ApiError{"5-00-000.9", "System error."}
)