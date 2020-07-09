# \PbxDeviceProfilePreferenceApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PbxdeviceprofilepreferencesGet**](PbxDeviceProfilePreferenceApi.md#PbxdeviceprofilepreferencesGet) | **Get** /pbxdeviceprofilepreferences/ | Get PbxDeviceProfilePreference items
[**PbxdeviceprofilepreferencesIdGet**](PbxDeviceProfilePreferenceApi.md#PbxdeviceprofilepreferencesIdGet) | **Get** /pbxdeviceprofilepreferences/{id} | Get a specific PbxDeviceProfilePreference
[**PbxdeviceprofilepreferencesIdPatch**](PbxDeviceProfilePreferenceApi.md#PbxdeviceprofilepreferencesIdPatch) | **Patch** /pbxdeviceprofilepreferences/{id} | Change a specific PbxDeviceProfilePreference
[**PbxdeviceprofilepreferencesIdPut**](PbxDeviceProfilePreferenceApi.md#PbxdeviceprofilepreferencesIdPut) | **Put** /pbxdeviceprofilepreferences/{id} | Replace/change a specific PbxDeviceProfilePreference



## PbxdeviceprofilepreferencesGet

> []map[string]interface{} PbxdeviceprofilepreferencesGet(ctx, optional)

Get PbxDeviceProfilePreference items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PbxdeviceprofilepreferencesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PbxdeviceprofilepreferencesGetOpts struct


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


## PbxdeviceprofilepreferencesIdGet

> map[string]interface{} PbxdeviceprofilepreferencesIdGet(ctx, id)

Get a specific PbxDeviceProfilePreference

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


## PbxdeviceprofilepreferencesIdPatch

> map[string]interface{} PbxdeviceprofilepreferencesIdPatch(ctx, id, operation)

Change a specific PbxDeviceProfilePreference

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


## PbxdeviceprofilepreferencesIdPut

> map[string]interface{} PbxdeviceprofilepreferencesIdPut(ctx, id, body)

Replace/change a specific PbxDeviceProfilePreference

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

