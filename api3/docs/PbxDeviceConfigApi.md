# \PbxDeviceConfigApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PbxdeviceconfigsGet**](PbxDeviceConfigApi.md#PbxdeviceconfigsGet) | **Get** /pbxdeviceconfigs/ | Get PbxDeviceConfig items
[**PbxdeviceconfigsIdGet**](PbxDeviceConfigApi.md#PbxdeviceconfigsIdGet) | **Get** /pbxdeviceconfigs/{id} | Get a specific PbxDeviceConfig
[**PbxdeviceconfigsIdPut**](PbxDeviceConfigApi.md#PbxdeviceconfigsIdPut) | **Put** /pbxdeviceconfigs/{id} | Replace/change a specific PbxDeviceConfig
[**PbxdeviceconfigsPost**](PbxDeviceConfigApi.md#PbxdeviceconfigsPost) | **Post** /pbxdeviceconfigs/ | Create a new PbxDeviceConfig



## PbxdeviceconfigsGet

> []PbxDeviceConfig PbxdeviceconfigsGet(ctx, optional)

Get PbxDeviceConfig items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PbxdeviceconfigsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PbxdeviceconfigsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **optional.String**| Filter for configs of a specific device model | 
 **version** | **optional.String**| Filter for configs by a specific version | 
 **contentType** | **optional.String**| Filter for configs by a specific content type | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]PbxDeviceConfig**](PbxDeviceConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdeviceconfigsIdGet

> PbxDeviceConfig PbxdeviceconfigsIdGet(ctx, id)

Get a specific PbxDeviceConfig

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**PbxDeviceConfig**](PbxDeviceConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdeviceconfigsIdPut

> PbxDeviceConfig PbxdeviceconfigsIdPut(ctx, id, pbxDeviceConfig)

Replace/change a specific PbxDeviceConfig

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**pbxDeviceConfig** | [**PbxDeviceConfig**](PbxDeviceConfig.md)|  | 

### Return type

[**PbxDeviceConfig**](PbxDeviceConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdeviceconfigsPost

> []PbxDeviceConfig PbxdeviceconfigsPost(ctx, pbxDeviceConfig)

Create a new PbxDeviceConfig

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbxDeviceConfig** | [**PbxDeviceConfig**](PbxDeviceConfig.md)|  | 

### Return type

[**[]PbxDeviceConfig**](PbxDeviceConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

