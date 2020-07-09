# \PbxFieldDevicePreferenceApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PbxfielddevicepreferencesGet**](PbxFieldDevicePreferenceApi.md#PbxfielddevicepreferencesGet) | **Get** /pbxfielddevicepreferences/ | Get PbxFieldDevicePreference items
[**PbxfielddevicepreferencesIdGet**](PbxFieldDevicePreferenceApi.md#PbxfielddevicepreferencesIdGet) | **Get** /pbxfielddevicepreferences/{id} | Get a specific PbxFieldDevicePreference
[**PbxfielddevicepreferencesIdPatch**](PbxFieldDevicePreferenceApi.md#PbxfielddevicepreferencesIdPatch) | **Patch** /pbxfielddevicepreferences/{id} | Change a specific PbxFieldDevicePreference
[**PbxfielddevicepreferencesIdPut**](PbxFieldDevicePreferenceApi.md#PbxfielddevicepreferencesIdPut) | **Put** /pbxfielddevicepreferences/{id} | Replace/change a specific PbxFieldDevicePreference



## PbxfielddevicepreferencesGet

> []map[string]interface{} PbxfielddevicepreferencesGet(ctx, optional)

Get PbxFieldDevicePreference items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PbxfielddevicepreferencesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PbxfielddevicepreferencesGetOpts struct


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


## PbxfielddevicepreferencesIdGet

> map[string]interface{} PbxfielddevicepreferencesIdGet(ctx, id)

Get a specific PbxFieldDevicePreference

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


## PbxfielddevicepreferencesIdPatch

> map[string]interface{} PbxfielddevicepreferencesIdPatch(ctx, id, operation)

Change a specific PbxFieldDevicePreference

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


## PbxfielddevicepreferencesIdPut

> map[string]interface{} PbxfielddevicepreferencesIdPut(ctx, id, body)

Replace/change a specific PbxFieldDevicePreference

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

