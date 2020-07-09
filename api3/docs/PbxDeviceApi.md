# \PbxDeviceApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PbxdevicesGet**](PbxDeviceApi.md#PbxdevicesGet) | **Get** /pbxdevices/ | Get PbxDevice items
[**PbxdevicesIdDelete**](PbxDeviceApi.md#PbxdevicesIdDelete) | **Delete** /pbxdevices/{id} | Delete a specific PbxDevice
[**PbxdevicesIdGet**](PbxDeviceApi.md#PbxdevicesIdGet) | **Get** /pbxdevices/{id} | Get a specific PbxDevice
[**PbxdevicesIdPatch**](PbxDeviceApi.md#PbxdevicesIdPatch) | **Patch** /pbxdevices/{id} | Change a specific PbxDevice
[**PbxdevicesIdPut**](PbxDeviceApi.md#PbxdevicesIdPut) | **Put** /pbxdevices/{id} | Replace/change a specific PbxDevice
[**PbxdevicesPost**](PbxDeviceApi.md#PbxdevicesPost) | **Post** /pbxdevices/ | Create a new PbxDevice



## PbxdevicesGet

> []PbxDevice PbxdevicesGet(ctx, optional)

Get PbxDevice items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PbxdevicesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PbxdevicesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **optional.String**| Search for PBX devices belonging to a specific customer | 
 **profileId** | **optional.String**| Search for PBX devices with a specific autoprovisioning device profile | 
 **identifier** | **optional.String**| Search for PBX devices matching an identifier/MAC pattern (wildcards possible) | 
 **stationName** | **optional.String**| Search for PBX devices matching a station_name pattern (wildcards possible) | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]PbxDevice**](PbxDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdevicesIdDelete

> PbxdevicesIdDelete(ctx, id)

Delete a specific PbxDevice

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


## PbxdevicesIdGet

> PbxDevice PbxdevicesIdGet(ctx, id)

Get a specific PbxDevice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**PbxDevice**](PbxDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdevicesIdPatch

> PbxDevice PbxdevicesIdPatch(ctx, id, operation)

Change a specific PbxDevice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**PbxDevice**](PbxDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdevicesIdPut

> PbxDevice PbxdevicesIdPut(ctx, id, pbxDevice)

Replace/change a specific PbxDevice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**pbxDevice** | [**PbxDevice**](PbxDevice.md)|  | 

### Return type

[**PbxDevice**](PbxDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdevicesPost

> []PbxDevice PbxdevicesPost(ctx, pbxDevice)

Create a new PbxDevice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbxDevice** | [**PbxDevice**](PbxDevice.md)|  | 

### Return type

[**[]PbxDevice**](PbxDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

