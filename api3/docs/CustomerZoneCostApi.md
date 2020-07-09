# \CustomerZoneCostApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerzonecostsGet**](CustomerZoneCostApi.md#CustomerzonecostsGet) | **Get** /customerzonecosts/ | Get CustomerZoneCost items
[**CustomerzonecostsIdGet**](CustomerZoneCostApi.md#CustomerzonecostsIdGet) | **Get** /customerzonecosts/{id} | Get a specific CustomerZoneCost



## CustomerzonecostsGet

> []map[string]interface{} CustomerzonecostsGet(ctx, optional)

Get CustomerZoneCost items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerzonecostsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CustomerzonecostsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **optional.String**| Filter for a specific customer. | 
 **start** | **optional.String**| Filter for a specific start time in format YYYY-MM-DDThhmmss. | 
 **end** | **optional.String**| Filter for a specific end time in format YYYY-MM-DDThhmmss. | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerzonecostsIdGet

> map[string]interface{} CustomerzonecostsIdGet(ctx, id)

Get a specific CustomerZoneCost

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

