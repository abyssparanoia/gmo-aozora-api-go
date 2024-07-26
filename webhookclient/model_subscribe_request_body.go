/*
 * GMO Aozora Net Bank Open API
 *
 * <p>Ph2.5向けに作成したもの</p>
 *
 * API version: 1.1.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SubscribeRequestBody struct {
	// 配信状態 半角数字 0=配信停止要求 1=配信開始要求
	SubscribeStatus string `json:"subscribeStatus"`
	// イベント種別のリスト
	EventTypes []EventType `json:"eventTypes"`
}
