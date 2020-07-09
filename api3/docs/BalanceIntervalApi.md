# \BalanceIntervalApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BalanceintervalsGet**](BalanceIntervalApi.md#BalanceintervalsGet) | **Get** /balanceintervals/ | Get BalanceInterval items
[**BalanceintervalsIdGet**](BalanceIntervalApi.md#BalanceintervalsIdGet) | **Get** /balanceintervals/{id} | Get a specific BalanceInterval



## BalanceintervalsGet

> []BalanceInterval BalanceintervalsGet(ctx, optional)

Get BalanceInterval items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BalanceintervalsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BalanceintervalsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for actual balance intervals of customers belonging to a specific reseller | 
 **contactId** | **optional.String**| Filter for contracts with a specific contact id | 
 **status** | **optional.String**| Filter for contracts with a specific status (except \&quot;terminated\&quot;) | 
 **externalId** | **optional.String**| Filter for contracts with a specific external id | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]BalanceInterval**](BalanceInterval.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BalanceintervalsIdGet

> BalanceInterval BalanceintervalsIdGet(ctx, id)

Get a specific BalanceInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**BalanceInterval**](BalanceInterval.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

