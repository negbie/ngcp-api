# \PreferencesMetaEntryApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreferencesmetaentriesGet**](PreferencesMetaEntryApi.md#PreferencesmetaentriesGet) | **Get** /preferencesmetaentries/ | Get PreferencesMetaEntry items
[**PreferencesmetaentriesIdDelete**](PreferencesMetaEntryApi.md#PreferencesmetaentriesIdDelete) | **Delete** /preferencesmetaentries/{id} | Delete a specific PreferencesMetaEntry
[**PreferencesmetaentriesIdGet**](PreferencesMetaEntryApi.md#PreferencesmetaentriesIdGet) | **Get** /preferencesmetaentries/{id} | Get a specific PreferencesMetaEntry
[**PreferencesmetaentriesIdPatch**](PreferencesMetaEntryApi.md#PreferencesmetaentriesIdPatch) | **Patch** /preferencesmetaentries/{id} | Change a specific PreferencesMetaEntry
[**PreferencesmetaentriesIdPut**](PreferencesMetaEntryApi.md#PreferencesmetaentriesIdPut) | **Put** /preferencesmetaentries/{id} | Replace/change a specific PreferencesMetaEntry
[**PreferencesmetaentriesPost**](PreferencesMetaEntryApi.md#PreferencesmetaentriesPost) | **Post** /preferencesmetaentries/ | Create a new PreferencesMetaEntry



## PreferencesmetaentriesGet

> []PreferencesMetaEntry PreferencesmetaentriesGet(ctx, optional)

Get PreferencesMetaEntry items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PreferencesmetaentriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PreferencesmetaentriesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attribute** | **optional.String**| Filter for dynamic preference with a specific name | 
 **modelId** | **optional.String**| Filter for dynamic preference relevant to the spcified pbx device model id | 
 **resellerId** | **optional.String**| Filter for dynamic preference relevant to the spcified reseller id | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]PreferencesMetaEntry**](PreferencesMetaEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreferencesmetaentriesIdDelete

> PreferencesmetaentriesIdDelete(ctx, id)

Delete a specific PreferencesMetaEntry

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


## PreferencesmetaentriesIdGet

> PreferencesMetaEntry PreferencesmetaentriesIdGet(ctx, id)

Get a specific PreferencesMetaEntry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**PreferencesMetaEntry**](PreferencesMetaEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreferencesmetaentriesIdPatch

> PreferencesMetaEntry PreferencesmetaentriesIdPatch(ctx, id, operation)

Change a specific PreferencesMetaEntry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**PreferencesMetaEntry**](PreferencesMetaEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreferencesmetaentriesIdPut

> PreferencesMetaEntry PreferencesmetaentriesIdPut(ctx, id, preferencesMetaEntry)

Replace/change a specific PreferencesMetaEntry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**preferencesMetaEntry** | [**PreferencesMetaEntry**](PreferencesMetaEntry.md)|  | 

### Return type

[**PreferencesMetaEntry**](PreferencesMetaEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreferencesmetaentriesPost

> []PreferencesMetaEntry PreferencesmetaentriesPost(ctx, preferencesMetaEntry)

Create a new PreferencesMetaEntry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**preferencesMetaEntry** | [**PreferencesMetaEntry**](PreferencesMetaEntry.md)|  | 

### Return type

[**[]PreferencesMetaEntry**](PreferencesMetaEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

