# RequestTransferStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestTransferStatus** | **string** | 照会対象ステータス 半角数字 2:申請中、3:差戻、4:取下げ、5:期限切れ、8:承認取消/予約取消、 11:予約中、12:手続中、13:リトライ中、 20:手続済、22:資金返却、24:組戻手続中、25:組戻済、26:組戻不成立、 30:不能・組戻あり、 40:手続不成立 配列のため、複数設定した場合は対象のステータスをOR条件で検索します 省略した場合は全てを設定したものとみなします 22、24、25、26は振込状況照会でのみ指定可 30は総合振込状況照会でのみ指定可 レスポンスの場合はリクエストしたときと同じ内容となります  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

