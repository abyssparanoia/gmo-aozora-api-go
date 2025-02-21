/*
 * GMO Aozora Net Bank Open API
 *
 * <p>Ph2.5向けに作成したもの</p>
 *
 * API version: 1.1.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ErrorResponse struct {
	// エラーコード 半角英数文字
	ErrorCode string `json:"errorCode"`
	// エラーメッセージ 全半角英数記号文字
	ErrorMessage string `json:"errorMessage"`
}
