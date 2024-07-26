# VaDepositTransactionsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RaId** | **string** | 入金口座ID 半角数字 入金先の口座を識別するID  | [default to null]
**RaBranchCode** | **string** | 入金口座　支店コード 半角数字  | [default to null]
**RaBranchNameKana** | **string** | 入金口座　支店名カナ 半角文字  | [default to null]
**RaAccountNumber** | **string** | 入金口座　口座番号 半角数字  | [default to null]
**RaHolderName** | **string** | 入金口座　口座名義（漢字） 全角文字  | [default to null]
**DateFrom** | **string** | 対象期間From 半角文字 YYYY-MM-DD形式 リクエストに対象期間From、Toが設定されていない場合は当日日付が設定されます  | [default to null]
**DateTo** | **string** | 対象期間To 半角文字 YYYY-MM-DD形式 リクエストに対象期間From、Toが設定されていない場合は当日日付が設定されます  | [default to null]
**BaseDate** | **string** | 基準日 半角文字 入金明細を照会した基準日を示します YYYY-MM-DD形式  | [default to null]
**BaseTime** | **string** | 基準時刻 半角文字 入金明細を照会した基準時刻を示します HH:MM:SS+09:00形式  | [default to null]
**HasNext** | **bool** | 次明細フラグ ・true&#x3D;次明細あり ・false&#x3D;次明細なし  | [default to null]
**NextItemKey** | **string** | 次明細キー 半角数字 次明細フラグがfalseの場合は、項目自体を設定しません  | [optional] [default to null]
**Count** | **string** | 明細取得件数 半角数字  | [default to null]
**VaTransactions** | [**[]VaTransaction**](VaTransaction.md) | 振込入金口座入金明細情報リスト 該当する情報が無い場合は、空のリストを返却します  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

