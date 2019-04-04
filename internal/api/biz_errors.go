package api

type ApiErrorMsg struct {
	ErrorCode, ErrorMsg string
}

var (
	UNKNOWN_ERROR      = ApiErrorMsg{"UNKNOWN_ERROR", "远程服务调用失败，未知异常!"}
	PARAM_ERROR_CODE   = ApiErrorMsg{"API_PARAM_ILLEGAL", "api参数不合法"}
	DB_TEMP_ADD_FAILED = ApiErrorMsg{"DB_TEMP_ADD_FAILED", "数据新增失败, 请稍后重试"}

	ASSIGNEE_NOT_EXIST = ApiErrorMsg{"ASSIGNEE_NOT_EXIST", "分配的用户不存在"}
	RECORD_NOT_EXIST   = ApiErrorMsg{"RECORD_NOT_EXIST", "记录不存在"}
	PERMISSION_DENIED  = ApiErrorMsg{"PERMISSION_DENIED", "没有权限"}
)
