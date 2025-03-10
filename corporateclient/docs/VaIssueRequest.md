# VaIssueRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaTypeCode** | **string** | 振込入金口座　種類コード 半角数字 ・1&#x3D;期限型 ・2&#x3D;継続型  | [default to null]
**IssueRequestCount** | **string** | 発行口座数 半角数字 1回のリクエストで1,000口座まで発行可能  | [default to null]
**RaId** | **string** | 入金口座ID 半角数字 科目コードが以下の口座IDのみ受け付けます ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用） 銀行契約の場合は必須  | [optional] [default to null]
**VaContractAuthKey** | **string** | 振込入金口座API認証情報 半角英数字 銀行契約の方はNULLを設定 提携企業の方が、契約された顧客の発行を依頼される場合は必須 提携企業以外の方が値を設定されている場合は「400 Bad Request」を返却  | [optional] [default to null]
**VaHolderNameKana** | **string** | 追加名義カナ 半角文字 振込入金口座名義（※）に任意の文字を追加する場合は指定してください ※ご登録されている「法人名カナ」  追加名義カナで指定できる文字数 &#x3D; 40 - 振込入金口座名義の文字数 合計が40文字を超える場合は、追加名義カナの後部から文字が削られます  ・使用可能な記号は「/」、「(」、「)」、「.」、「,」、「-」 ・一部の非許容文字は、許容文字に変換を行います ・濁点、半濁点は１文字として数えます ・追加名義カナを前につける場合、半角スペースは追加名義カナの右側に1文字のみ許容します ・追加名義カナを前につける場合、追加名義カナの左側に半角スペースは許容しません ・追加名義カナを前につける場合、追加名義カナの中に半角の連続スペースは許容しません | [optional] [default to null]
**VaHolderNamePos** | **string** | 追加名義位置 半角数字 追加名義カナを口座名義の前につけるか後ろにつけるかの指定 ・1&#x3D;通常（後ろにつける） ・2&#x3D;前につける 指定が無い場合は後ろにつけます  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

