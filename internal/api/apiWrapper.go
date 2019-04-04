package api

import (
	"fmt"
	"time"
)

type ApiResultCommon struct {
	T       int64 `json:"t"`
	Success bool  `json:"success"`
	Status  bool  `json:"status"`
}

type ApiErrorResult struct {
	ApiResultCommon
	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

func Success(result string) []byte {
	nowMillis := time.Now().UnixNano() / 1e6
	ret := fmt.Sprintf(`{"t":%d,"success":true,"status":true,"result":%s}`, nowMillis, result)
	return []byte(ret)
}

func NewApiErrorResult(errorCode, errorMsg string) ApiErrorResult {
	nowMillis := time.Now().UnixNano() / 1e6
	return ApiErrorResult{
		ApiResultCommon: ApiResultCommon{
			T:       nowMillis,
			Success: false,
			Status:  true,
		},
		ErrorCode: errorCode,
		ErrorMsg:  errorMsg,
	}
}
