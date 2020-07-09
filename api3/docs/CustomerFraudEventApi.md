# \CustomerFraudEventApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerfraudeventsGet**](CustomerFraudEventApi.md#CustomerfraudeventsGet) | **Get** /customerfraudevents/ | Get CustomerFraudEvent items
[**CustomerfraudeventsIdGet**](CustomerFraudEventApi.md#CustomerfraudeventsIdGet) | **Get** /customerfraudevents/{id} | Get a specific CustomerFraudEvent



## CustomerfraudeventsGet

> []CustomerFraudEvent CustomerfraudeventsGet(ctx, optional)

Get CustomerFraudEvent items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerfraudeventsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CustomerfraudeventsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for fraud events belonging to a specific reseller | 
 **interval** | **optional.String**| Interval filter. values: day, month. default: month | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CustomerFraudEvent**](CustomerFraudEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerfraudeventsIdGet

> CustomerFraudEvent CustomerfraudeventsIdGet(ctx, id)

Get a specific CustomerFraudEvent

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CustomerFraudEvent**](CustomerFraudEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

