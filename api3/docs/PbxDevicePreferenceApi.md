# \PbxDevicePreferenceApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PbxdevicepreferencesGet**](PbxDevicePreferenceApi.md#PbxdevicepreferencesGet) | **Get** /pbxdevicepreferences/ | Get PbxDevicePreference items
[**PbxdevicepreferencesIdGet**](PbxDevicePreferenceApi.md#PbxdevicepreferencesIdGet) | **Get** /pbxdevicepreferences/{id} | Get a specific PbxDevicePreference
[**PbxdevicepreferencesIdPatch**](PbxDevicePreferenceApi.md#PbxdevicepreferencesIdPatch) | **Patch** /pbxdevicepreferences/{id} | Change a specific PbxDevicePreference
[**PbxdevicepreferencesIdPut**](PbxDevicePreferenceApi.md#PbxdevicepreferencesIdPut) | **Put** /pbxdevicepreferences/{id} | Replace/change a specific PbxDevicePreference



## PbxdevicepreferencesGet

> []map[string]interface{} PbxdevicepreferencesGet(ctx, optional)

Get PbxDevicePreference items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PbxdevicepreferencesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PbxdevicepreferencesGetOpts struct


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


## PbxdevicepreferencesIdGet

> map[string]interface{} PbxdevicepreferencesIdGet(ctx, id)

Get a specific PbxDevicePreference

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


## PbxdevicepreferencesIdPatch

> map[string]interface{} PbxdevicepreferencesIdPatch(ctx, id, operation)

Change a specific PbxDevicePreference

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


## PbxdevicepreferencesIdPut

> map[string]interface{} PbxdevicepreferencesIdPut(ctx, id, body)

Replace/change a specific PbxDevicePreference

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

