/*
 * GMO Aozora Net Bank Open API
 *
 * <p>Ph2.5向けに作成したもの</p>
 *
 * API version: 1.1.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VaDepositTransactionMessage struct {
	// メッセージID 半角英数字 メッセージを一意に識別するID
	MessageId string `json:"messageId"`
	// イベント生成日時 半角文字 ISO8601 時差(offset)も表記 YYYY-MM-DDTHH:MM:SS+09:00形式
	Timestamp     string         `json:"timestamp"`
	Account       *Account       `json:"account"`
	VaTransaction *VaTransaction `json:"vaTransaction"`
}
