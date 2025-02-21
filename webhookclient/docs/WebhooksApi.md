# {{classname}}

All URIs are relative to *https://api.gmo-aozora.com/ganb/api/webhooks/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsUsing**](WebhooksApi.md#AccountsUsing) | **Post** /subscribe | 通知配信制御
[**AccountsUsingGET**](WebhooksApi.md#AccountsUsingGET) | **Get** /unsentlist/va-deposit-transaction | 振込入金口座未送信明細取得
[**SubscribeStatusUsingGET**](WebhooksApi.md#SubscribeStatusUsingGET) | **Get** /subscribe-status | 通知配信状態取得

# **AccountsUsing**
> SubscribeResponse AccountsUsing(ctx, body, xAccessToken)
通知配信制御

指定したイベント通知に対応するイベント通知（WebHook）の配信開始、配信停止をコントロールします

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscribeRequestBody**](SubscribeRequestBody.md)| HTTPリクエストボディ | 
  **xAccessToken** | **string**| アクセストークン  minLength: 1 maxLength: 128  | 

### Return type

[**SubscribeResponse**](SubscribeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountsUsingGET**
> VaDepositTransactionUnsentResponse AccountsUsingGET(ctx, xAccessToken)
振込入金口座未送信明細取得

配信停止状態となっている場合、本APIを利用することで未送信または送信エラーとなっている、振込入金口座の入金明細を一括で取得することができます 通常、未送信または送信エラーとなっている明細は配信再開後に通知されますが、本APIで取得された明細は配信済みとなるため、配信再開後には通知されません 未送信または送信エラーとなっている明細が無い場合は404 Not Foundを返却します <p><font color=\"red\">※法人口座および個人事業主口座のみ対象となり、個人口座は対象外となります</font></p> <h4 style='margin-top:30px; border-left: solid 4px #1B2F48; padding: 0.1em 0.5em; color:#1B2F48;'>詳細説明</h4> <div style='margin:10px;'>   <p style='font-weight:bold; color:#616161;'>取得上限件数</p>   <p style='padding-left:20px;'>500件</p>   <p style='padding-left:20px;'>取得できる明細数が500に満たないときは取得できる明細のみを返却します</p>   <p style='padding-left:20px;'>取得できた明細数が500件の場合、まだ取得できる明細が残っている可能性がありますので、</p>   <p style='padding-left:20px;'>「404：Not Found」が返却されるまで、リクエストを繰り返してください。</p> </div> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccessToken** | **string**| アクセストークン  minLength: 1 maxLength: 128  | 

### Return type

[**VaDepositTransactionUnsentResponse**](vaDepositTransactionUnsentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscribeStatusUsingGET**
> SubscribeStatusResponse SubscribeStatusUsingGET(ctx, xAccessToken)
通知配信状態取得

各種イベント通知（Webhook）の配信状態を取得できます。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xAccessToken** | **string**| アクセストークン  minLength: 1 maxLength: 128  | 

### Return type

[**SubscribeStatusResponse**](SubscribeStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

