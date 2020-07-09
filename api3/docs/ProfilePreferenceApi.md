# \ProfilePreferenceApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProfilepreferencesGet**](ProfilePreferenceApi.md#ProfilepreferencesGet) | **Get** /profilepreferences/ | Get ProfilePreference items
[**ProfilepreferencesIdGet**](ProfilePreferenceApi.md#ProfilepreferencesIdGet) | **Get** /profilepreferences/{id} | Get a specific ProfilePreference
[**ProfilepreferencesIdPatch**](ProfilePreferenceApi.md#ProfilepreferencesIdPatch) | **Patch** /profilepreferences/{id} | Change a specific ProfilePreference
[**ProfilepreferencesIdPut**](ProfilePreferenceApi.md#ProfilepreferencesIdPut) | **Put** /profilepreferences/{id} | Replace/change a specific ProfilePreference



## ProfilepreferencesGet

> []map[string]interface{} ProfilepreferencesGet(ctx, optional)

Get ProfilePreference items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProfilepreferencesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProfilepreferencesGetOpts struct


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


## ProfilepreferencesIdGet

> map[string]interface{} ProfilepreferencesIdGet(ctx, id)

Get a specific ProfilePreference

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


## ProfilepreferencesIdPatch

> map[string]interface{} ProfilepreferencesIdPatch(ctx, id, operation)

Change a specific ProfilePreference

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


## ProfilepreferencesIdPut

> map[string]interface{} ProfilepreferencesIdPut(ctx, id, body)

Replace/change a specific ProfilePreference

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

