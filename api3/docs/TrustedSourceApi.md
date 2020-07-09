# \TrustedSourceApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrustedsourcesGet**](TrustedSourceApi.md#TrustedsourcesGet) | **Get** /trustedsources/ | Get TrustedSource items
[**TrustedsourcesIdDelete**](TrustedSourceApi.md#TrustedsourcesIdDelete) | **Delete** /trustedsources/{id} | Delete a specific TrustedSource
[**TrustedsourcesIdGet**](TrustedSourceApi.md#TrustedsourcesIdGet) | **Get** /trustedsources/{id} | Get a specific TrustedSource
[**TrustedsourcesIdPatch**](TrustedSourceApi.md#TrustedsourcesIdPatch) | **Patch** /trustedsources/{id} | Change a specific TrustedSource
[**TrustedsourcesIdPut**](TrustedSourceApi.md#TrustedsourcesIdPut) | **Put** /trustedsources/{id} | Replace/change a specific TrustedSource
[**TrustedsourcesPost**](TrustedSourceApi.md#TrustedsourcesPost) | **Post** /trustedsources/ | Create a new TrustedSource



## TrustedsourcesGet

> []TrustedSource TrustedsourcesGet(ctx, optional)

Get TrustedSource items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TrustedsourcesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TrustedsourcesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriberId** | **optional.String**| Filter for trusted sources of a specific subscriber | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]TrustedSource**](TrustedSource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrustedsourcesIdDelete

> TrustedsourcesIdDelete(ctx, id)

Delete a specific TrustedSource

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


## TrustedsourcesIdGet

> TrustedSource TrustedsourcesIdGet(ctx, id)

Get a specific TrustedSource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**TrustedSource**](TrustedSource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrustedsourcesIdPatch

> TrustedSource TrustedsourcesIdPatch(ctx, id, operation)

Change a specific TrustedSource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**TrustedSource**](TrustedSource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrustedsourcesIdPut

> TrustedSource TrustedsourcesIdPut(ctx, id, trustedSource)

Replace/change a specific TrustedSource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**trustedSource** | [**TrustedSource**](TrustedSource.md)|  | 

### Return type

[**TrustedSource**](TrustedSource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrustedsourcesPost

> []TrustedSource TrustedsourcesPost(ctx, trustedSource)

Create a new TrustedSource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedSource** | [**TrustedSource**](TrustedSource.md)|  | 

### Return type

[**[]TrustedSource**](TrustedSource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

