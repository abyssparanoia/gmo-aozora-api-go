/*
 * GMO Aozora Net Bank Open API
 *
 * <p>Ph2.5向けに作成したもの</p>
 *
 * API version: 1.1.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SubscribeResponse struct {
	// イベント種別 半角英数記号文字 va-deposit-transaction = 振込入金口座への入金明細通知
	EventTypes string `json:"eventTypes"`
}
