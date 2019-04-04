package api

import "time"

type ApiResultCommon struct {
	T       int64 `json:"t"`
	Success bool  `json:"success"`
	Status  bool  `json:"status"`
}

type ApiSuccessResult struct {
	ApiResultCommon
	Result interface{} `json:"result"`
}

type ApiErrorResult struct {
	ApiResultCommon
	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

func NewApiSuccessResult(result interface{}) ApiSuccessResult {
	nowMillis := time.Now().UnixNano() / 1e6
	apiRes := ApiSuccessResult{
		Result: result,
		ApiResultCommon: ApiResultCommon{
			T:       nowMillis,
			Success: true,
			Status:  true,
		},
	}
	return apiRes
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

func NewApiPredefinedErrorResult(errMsg ApiErrorMsg) ApiErrorResult {
	nowMillis := time.Now().UnixNano() / 1e6
	return ApiErrorResult{
		ApiResultCommon: ApiResultCommon{
			T:       nowMillis,
			Success: false,
		},
		ErrorCode: errMsg.ErrorCode,
		ErrorMsg:  errMsg.ErrorMsg,
	}
}
