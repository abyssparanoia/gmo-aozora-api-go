# Transaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionDate** | **string** | 取引日 半角文字 YYYY-MM-DD形式  | [default to null]
**ValueDate** | **string** | 起算日 半角文字 YYYY-MM-DD形式  | [default to null]
**TransactionType** | **string** | 入払コード 半角数字 ・1&#x3D;入金 ・2&#x3D;出金  | [default to null]
**Amount** | **string** | 取引金額 半角数字　マイナス含む  | [default to null]
**Remarks** | **string** | 摘要内容 全半角文字 該当データがない場合は項目自体を設定しません  | [optional] [default to null]
**Balance** | **string** | 取引後残高 半角数字　マイナス含む  | [default to null]
**ItemKey** | **string** | 明細キー 半角数字 口座の科目が01&#x3D;普通預金（有利息）または02&#x3D;普通預金（決済用）の場合は口座ID毎に設定される明細キー（明細データtimestamp（μs）） 該当データがない場合は項目自体を設定しません | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

