# \LnpNumberApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LnpnumbersGet**](LnpNumberApi.md#LnpnumbersGet) | **Get** /lnpnumbers/ | Get LnpNumber items
[**LnpnumbersIdDelete**](LnpNumberApi.md#LnpnumbersIdDelete) | **Delete** /lnpnumbers/{id} | Delete a specific LnpNumber
[**LnpnumbersIdGet**](LnpNumberApi.md#LnpnumbersIdGet) | **Get** /lnpnumbers/{id} | Get a specific LnpNumber
[**LnpnumbersIdPatch**](LnpNumberApi.md#LnpnumbersIdPatch) | **Patch** /lnpnumbers/{id} | Change a specific LnpNumber
[**LnpnumbersIdPut**](LnpNumberApi.md#LnpnumbersIdPut) | **Put** /lnpnumbers/{id} | Replace/change a specific LnpNumber
[**LnpnumbersPost**](LnpNumberApi.md#LnpnumbersPost) | **Post** /lnpnumbers/ | Create a new LnpNumber



## LnpnumbersGet

> []LnpNumber LnpnumbersGet(ctx, optional)

Get LnpNumber items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LnpnumbersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LnpnumbersGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **carrierId** | **optional.String**| Filter for LNP numbers belonging to a specific LNP carrier | 
 **number** | **optional.String**| Filter for LNP numbers with a specific number (wildcards possible) | 
 **actual** | **optional.String**| Filter for LNP numbers valid at the given timestamp (YYYY-MM-DD HH:mm:ss) or the current time | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]LnpNumber**](LnpNumber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LnpnumbersIdDelete

> LnpnumbersIdDelete(ctx, id)

Delete a specific LnpNumber

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


## LnpnumbersIdGet

> LnpNumber LnpnumbersIdGet(ctx, id)

Get a specific LnpNumber

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**LnpNumber**](LnpNumber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LnpnumbersIdPatch

> LnpNumber LnpnumbersIdPatch(ctx, id, operation)

Change a specific LnpNumber

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**LnpNumber**](LnpNumber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LnpnumbersIdPut

> LnpNumber LnpnumbersIdPut(ctx, id, lnpNumber)

Replace/change a specific LnpNumber

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**lnpNumber** | [**LnpNumber**](LnpNumber.md)|  | 

### Return type

[**LnpNumber**](LnpNumber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LnpnumbersPost

> []LnpNumber LnpnumbersPost(ctx, lnpNumber)

Create a new LnpNumber

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lnpNumber** | [**LnpNumber**](LnpNumber.md)|  | 

### Return type

[**[]LnpNumber**](LnpNumber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

