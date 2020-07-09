# \InterceptionApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InterceptionsGet**](InterceptionApi.md#InterceptionsGet) | **Get** /interceptions/ | Get Interception items
[**InterceptionsIdDelete**](InterceptionApi.md#InterceptionsIdDelete) | **Delete** /interceptions/{id} | Delete a specific Interception
[**InterceptionsIdGet**](InterceptionApi.md#InterceptionsIdGet) | **Get** /interceptions/{id} | Get a specific Interception
[**InterceptionsIdPatch**](InterceptionApi.md#InterceptionsIdPatch) | **Patch** /interceptions/{id} | Change a specific Interception
[**InterceptionsIdPut**](InterceptionApi.md#InterceptionsIdPut) | **Put** /interceptions/{id} | Replace/change a specific Interception
[**InterceptionsPost**](InterceptionApi.md#InterceptionsPost) | **Post** /interceptions/ | Create a new Interception



## InterceptionsGet

> []Interception InterceptionsGet(ctx, optional)

Get Interception items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InterceptionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InterceptionsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **liid** | **optional.String**| Filter for interceptions of a specific interception id | 
 **number** | **optional.String**| Filter for interceptions of a specific number (in E.164 format) | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]Interception**](Interception.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InterceptionsIdDelete

> InterceptionsIdDelete(ctx, id)

Delete a specific Interception

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


## InterceptionsIdGet

> Interception InterceptionsIdGet(ctx, id)

Get a specific Interception

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Interception**](Interception.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InterceptionsIdPatch

> Interception InterceptionsIdPatch(ctx, id, operation)

Change a specific Interception

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**Interception**](Interception.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InterceptionsIdPut

> Interception InterceptionsIdPut(ctx, id, interception)

Replace/change a specific Interception

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**interception** | [**Interception**](Interception.md)|  | 

### Return type

[**Interception**](Interception.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InterceptionsPost

> []Interception InterceptionsPost(ctx, interception)

Create a new Interception

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interception** | [**Interception**](Interception.md)|  | 

### Return type

[**[]Interception**](Interception.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

