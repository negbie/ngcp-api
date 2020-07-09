# \CustomerLocationApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerlocationsGet**](CustomerLocationApi.md#CustomerlocationsGet) | **Get** /customerlocations/ | Get CustomerLocation items
[**CustomerlocationsIdDelete**](CustomerLocationApi.md#CustomerlocationsIdDelete) | **Delete** /customerlocations/{id} | Delete a specific CustomerLocation
[**CustomerlocationsIdGet**](CustomerLocationApi.md#CustomerlocationsIdGet) | **Get** /customerlocations/{id} | Get a specific CustomerLocation
[**CustomerlocationsIdPatch**](CustomerLocationApi.md#CustomerlocationsIdPatch) | **Patch** /customerlocations/{id} | Change a specific CustomerLocation
[**CustomerlocationsIdPut**](CustomerLocationApi.md#CustomerlocationsIdPut) | **Put** /customerlocations/{id} | Replace/change a specific CustomerLocation
[**CustomerlocationsPost**](CustomerLocationApi.md#CustomerlocationsPost) | **Post** /customerlocations/ | Create a new CustomerLocation



## CustomerlocationsGet

> []CustomerLocation CustomerlocationsGet(ctx, optional)

Get CustomerLocation items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerlocationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CustomerlocationsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **optional.String**| Filter for customer locations containing a specific IP address | 
 **name** | **optional.String**| Filter for customer locations matching a name pattern | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CustomerLocation**](CustomerLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerlocationsIdDelete

> CustomerlocationsIdDelete(ctx, id)

Delete a specific CustomerLocation

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


## CustomerlocationsIdGet

> CustomerLocation CustomerlocationsIdGet(ctx, id)

Get a specific CustomerLocation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CustomerLocation**](CustomerLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerlocationsIdPatch

> CustomerLocation CustomerlocationsIdPatch(ctx, id, operation)

Change a specific CustomerLocation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**CustomerLocation**](CustomerLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerlocationsIdPut

> CustomerLocation CustomerlocationsIdPut(ctx, id, customerLocation)

Replace/change a specific CustomerLocation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**customerLocation** | [**CustomerLocation**](CustomerLocation.md)|  | 

### Return type

[**CustomerLocation**](CustomerLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerlocationsPost

> []CustomerLocation CustomerlocationsPost(ctx, customerLocation)

Create a new CustomerLocation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerLocation** | [**CustomerLocation**](CustomerLocation.md)|  | 

### Return type

[**[]CustomerLocation**](CustomerLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

