# {{classname}}

All URIs are relative to *https://api.gmo-aozora.com/ganb/api/webhooks/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Subscribe**](WebhooksApi.md#Subscribe) | **Post** /subscribe | 通知配信制御
[**SubscribeStatusUsingGET**](WebhooksApi.md#SubscribeStatusUsingGET) | **Get** /subscribe-status | 通知配信状態取得

# **Subscribe**
> Subscribe(ctx, body)
通知配信制御

指定したイベント通知に対応するイベント通知（WebHook）の配信開始、配信停止をコントロールします

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscribeRequestBody**](SubscribeRequestBody.md)| HTTPリクエストボディ | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscribeStatusUsingGET**
> SubscribeStatusResponse SubscribeStatusUsingGET(ctx, )
通知配信状態取得

各種イベント通知（Webhook）の配信状態を取得できます。

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SubscribeStatusResponse**](SubscribeStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

