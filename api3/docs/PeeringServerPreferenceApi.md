# \PeeringServerPreferenceApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PeeringserverpreferencesGet**](PeeringServerPreferenceApi.md#PeeringserverpreferencesGet) | **Get** /peeringserverpreferences/ | Get PeeringServerPreference items
[**PeeringserverpreferencesIdGet**](PeeringServerPreferenceApi.md#PeeringserverpreferencesIdGet) | **Get** /peeringserverpreferences/{id} | Get a specific PeeringServerPreference
[**PeeringserverpreferencesIdPatch**](PeeringServerPreferenceApi.md#PeeringserverpreferencesIdPatch) | **Patch** /peeringserverpreferences/{id} | Change a specific PeeringServerPreference
[**PeeringserverpreferencesIdPut**](PeeringServerPreferenceApi.md#PeeringserverpreferencesIdPut) | **Put** /peeringserverpreferences/{id} | Replace/change a specific PeeringServerPreference



## PeeringserverpreferencesGet

> []map[string]interface{} PeeringserverpreferencesGet(ctx, optional)

Get PeeringServerPreference items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PeeringserverpreferencesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PeeringserverpreferencesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## PeeringserverpreferencesIdGet

> map[string]interface{} PeeringserverpreferencesIdGet(ctx, id)

Get a specific PeeringServerPreference

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


## PeeringserverpreferencesIdPatch

> map[string]interface{} PeeringserverpreferencesIdPatch(ctx, id, operation)

Change a specific PeeringServerPreference

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringserverpreferencesIdPut

> map[string]interface{} PeeringserverpreferencesIdPut(ctx, id, body)

Replace/change a specific PeeringServerPreference

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**body** | **map[string]interface{}**|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

