# \PbxDeviceProfileApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PbxdeviceprofilesGet**](PbxDeviceProfileApi.md#PbxdeviceprofilesGet) | **Get** /pbxdeviceprofiles/ | Get PbxDeviceProfile items
[**PbxdeviceprofilesIdGet**](PbxDeviceProfileApi.md#PbxdeviceprofilesIdGet) | **Get** /pbxdeviceprofiles/{id} | Get a specific PbxDeviceProfile
[**PbxdeviceprofilesIdPatch**](PbxDeviceProfileApi.md#PbxdeviceprofilesIdPatch) | **Patch** /pbxdeviceprofiles/{id} | Change a specific PbxDeviceProfile
[**PbxdeviceprofilesIdPut**](PbxDeviceProfileApi.md#PbxdeviceprofilesIdPut) | **Put** /pbxdeviceprofiles/{id} | Replace/change a specific PbxDeviceProfile
[**PbxdeviceprofilesPost**](PbxDeviceProfileApi.md#PbxdeviceprofilesPost) | **Post** /pbxdeviceprofiles/ | Create a new PbxDeviceProfile



## PbxdeviceprofilesGet

> []PbxDeviceProfile PbxdeviceprofilesGet(ctx, optional)

Get PbxDeviceProfile items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PbxdeviceprofilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PbxdeviceprofilesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter for profiles matching a name pattern | 
 **configId** | **optional.String**| Filter for profiles matching a config_id | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]PbxDeviceProfile**](PbxDeviceProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdeviceprofilesIdGet

> PbxDeviceProfile PbxdeviceprofilesIdGet(ctx, id)

Get a specific PbxDeviceProfile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**PbxDeviceProfile**](PbxDeviceProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdeviceprofilesIdPatch

> PbxDeviceProfile PbxdeviceprofilesIdPatch(ctx, id, operation)

Change a specific PbxDeviceProfile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**PbxDeviceProfile**](PbxDeviceProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdeviceprofilesIdPut

> PbxDeviceProfile PbxdeviceprofilesIdPut(ctx, id, pbxDeviceProfile)

Replace/change a specific PbxDeviceProfile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**pbxDeviceProfile** | [**PbxDeviceProfile**](PbxDeviceProfile.md)|  | 

### Return type

[**PbxDeviceProfile**](PbxDeviceProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdeviceprofilesPost

> []PbxDeviceProfile PbxdeviceprofilesPost(ctx, pbxDeviceProfile)

Create a new PbxDeviceProfile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbxDeviceProfile** | [**PbxDeviceProfile**](PbxDeviceProfile.md)|  | 

### Return type

[**[]PbxDeviceProfile**](PbxDeviceProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

