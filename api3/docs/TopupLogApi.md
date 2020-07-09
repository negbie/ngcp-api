# \TopupLogApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TopuplogsGet**](TopupLogApi.md#TopuplogsGet) | **Get** /topuplogs/ | Get TopupLog items
[**TopuplogsIdGet**](TopupLogApi.md#TopuplogsIdGet) | **Get** /topuplogs/{id} | Get a specific TopupLog



## TopuplogsGet

> []TopupLog TopuplogsGet(ctx, optional)

Get TopupLog items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TopuplogsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TopuplogsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for top-up requests for customers/subscribers of a specific reseller. | 
 **requestToken** | **optional.String**| Filter for top-up requests with the given request_token. | 
 **timestampFrom** | **optional.String**| Filter for top-up requests performed after or at the given time stamp. | 
 **timestampTo** | **optional.String**| Filter for top-up requests performed before or at the given time stamp. | 
 **contractId** | **optional.String**| Filter for top-up requests of a specific customer contract. | 
 **subscriberId** | **optional.String**| Filter for top-up requests of a specific subscriber. | 
 **voucherId** | **optional.String**| Filter for top-up requests with a specific voucher. | 
 **outcome** | **optional.String**| Filter for top-up requests by outcome. | 
 **amountAbove** | **optional.String**| Filter for top-up requests with an amount greater than or equal to the given value in USD/EUR/etc. | 
 **amountBelow** | **optional.String**| Filter for top-up requests with an amount less than or equal to the given value in USD/EUR/etc. | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]TopupLog**](TopupLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TopuplogsIdGet

> TopupLog TopuplogsIdGet(ctx, id)

Get a specific TopupLog

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**TopupLog**](TopupLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

