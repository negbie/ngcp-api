# \PbxDeviceFirmwareApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PbxdevicefirmwaresGet**](PbxDeviceFirmwareApi.md#PbxdevicefirmwaresGet) | **Get** /pbxdevicefirmwares/ | Get PbxDeviceFirmware items
[**PbxdevicefirmwaresIdGet**](PbxDeviceFirmwareApi.md#PbxdevicefirmwaresIdGet) | **Get** /pbxdevicefirmwares/{id} | Get a specific PbxDeviceFirmware
[**PbxdevicefirmwaresIdPut**](PbxDeviceFirmwareApi.md#PbxdevicefirmwaresIdPut) | **Put** /pbxdevicefirmwares/{id} | Replace/change a specific PbxDeviceFirmware
[**PbxdevicefirmwaresPost**](PbxDeviceFirmwareApi.md#PbxdevicefirmwaresPost) | **Post** /pbxdevicefirmwares/ | Create a new PbxDeviceFirmware



## PbxdevicefirmwaresGet

> []PbxDeviceFirmware PbxdevicefirmwaresGet(ctx, optional)

Get PbxDeviceFirmware items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PbxdevicefirmwaresGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PbxdevicefirmwaresGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **optional.String**| Filter for firmwares of a specific device model | 
 **version** | **optional.String**| Filter for firmwares by a specific version | 
 **filename** | **optional.String**| Filter for firmwares by a specific file name | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]PbxDeviceFirmware**](PbxDeviceFirmware.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdevicefirmwaresIdGet

> PbxDeviceFirmware PbxdevicefirmwaresIdGet(ctx, id)

Get a specific PbxDeviceFirmware

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**PbxDeviceFirmware**](PbxDeviceFirmware.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdevicefirmwaresIdPut

> PbxDeviceFirmware PbxdevicefirmwaresIdPut(ctx, id, pbxDeviceFirmware)

Replace/change a specific PbxDeviceFirmware

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**pbxDeviceFirmware** | [**PbxDeviceFirmware**](PbxDeviceFirmware.md)|  | 

### Return type

[**PbxDeviceFirmware**](PbxDeviceFirmware.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdevicefirmwaresPost

> []PbxDeviceFirmware PbxdevicefirmwaresPost(ctx, pbxDeviceFirmware)

Create a new PbxDeviceFirmware

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbxDeviceFirmware** | [**PbxDeviceFirmware**](PbxDeviceFirmware.md)|  | 

### Return type

[**[]PbxDeviceFirmware**](PbxDeviceFirmware.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

