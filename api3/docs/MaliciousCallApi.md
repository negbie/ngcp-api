# \MaliciousCallApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MaliciouscallsGet**](MaliciousCallApi.md#MaliciouscallsGet) | **Get** /maliciouscalls/ | Get MaliciousCall items
[**MaliciouscallsIdDelete**](MaliciousCallApi.md#MaliciouscallsIdDelete) | **Delete** /maliciouscalls/{id} | Delete a specific MaliciousCall
[**MaliciouscallsIdGet**](MaliciousCallApi.md#MaliciouscallsIdGet) | **Get** /maliciouscalls/{id} | Get a specific MaliciousCall



## MaliciouscallsGet

> []MaliciousCall MaliciouscallsGet(ctx, optional)

Get MaliciousCall items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MaliciouscallsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MaliciouscallsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for malicious calls belonging to a specific reseller | 
 **callid** | **optional.String**| Filter by the call id | 
 **caller** | **optional.String**| Filter by the caller number | 
 **callee** | **optional.String**| Filter by the callee number | 
 **startLe** | **optional.String**| Filter by records with lower or equal than the specified time stamp. | 
 **startGe** | **optional.String**| Filter by records with greater or equal the specified time stamp. | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]MaliciousCall**](MaliciousCall.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaliciouscallsIdDelete

> MaliciouscallsIdDelete(ctx, id)

Delete a specific MaliciousCall

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaliciouscallsIdGet

> MaliciousCall MaliciouscallsIdGet(ctx, id)

Get a specific MaliciousCall

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**MaliciousCall**](MaliciousCall.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

