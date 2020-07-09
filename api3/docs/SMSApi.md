# \SMSApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmsGet**](SMSApi.md#SmsGet) | **Get** /sms/ | Get SMS items
[**SmsIdGet**](SMSApi.md#SmsIdGet) | **Get** /sms/{id} | Get a specific SMS
[**SmsPost**](SMSApi.md#SmsPost) | **Post** /sms/ | Create a new SMS



## SmsGet

> []Sms SmsGet(ctx, optional)

Get SMS items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SmsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriberId** | **optional.String**| Filter for messages belonging to a specific subscriber | 
 **customerId** | **optional.String**| Filter for messages belonging to a specific customer | 
 **resellerId** | **optional.String**| Filter for messages belonging to a specific reseller | 
 **timeGe** | **optional.String**| Filter for messages sent later or equal the specified time stamp. | 
 **timeLe** | **optional.String**| Filter for messages sent earlier or equal the specified time stamp. | 
 **direction** | **optional.String**| Filter for messages sent (\&quot;out\&quot;), received (\&quot;in\&quot;) or forwarded (\&quot;forward\&quot;). | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]Sms**](SMS.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmsIdGet

> Sms SmsIdGet(ctx, id)

Get a specific SMS

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Sms**](SMS.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmsPost

> []Sms SmsPost(ctx, sms)

Create a new SMS

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sms** | [**Sms**](Sms.md)|  | 

### Return type

[**[]Sms**](SMS.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

