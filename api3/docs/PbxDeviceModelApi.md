# \PbxDeviceModelApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PbxdevicemodelsGet**](PbxDeviceModelApi.md#PbxdevicemodelsGet) | **Get** /pbxdevicemodels/ | Get PbxDeviceModel items
[**PbxdevicemodelsIdGet**](PbxDeviceModelApi.md#PbxdevicemodelsIdGet) | **Get** /pbxdevicemodels/{id} | Get a specific PbxDeviceModel
[**PbxdevicemodelsIdPatch**](PbxDeviceModelApi.md#PbxdevicemodelsIdPatch) | **Patch** /pbxdevicemodels/{id} | Change a specific PbxDeviceModel
[**PbxdevicemodelsIdPut**](PbxDeviceModelApi.md#PbxdevicemodelsIdPut) | **Put** /pbxdevicemodels/{id} | Replace/change a specific PbxDeviceModel
[**PbxdevicemodelsPost**](PbxDeviceModelApi.md#PbxdevicemodelsPost) | **Post** /pbxdevicemodels/ | Create a new PbxDeviceModel



## PbxdevicemodelsGet

> []PbxDeviceModel PbxdevicemodelsGet(ctx, optional)

Get PbxDeviceModel items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PbxdevicemodelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PbxdevicemodelsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for models belonging to a certain reseller | 
 **vendor** | **optional.String**| Filter for vendor matching a vendor name pattern | 
 **model** | **optional.String**| Filter for models matching a model name pattern | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]PbxDeviceModel**](PbxDeviceModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdevicemodelsIdGet

> PbxDeviceModel PbxdevicemodelsIdGet(ctx, id)

Get a specific PbxDeviceModel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**PbxDeviceModel**](PbxDeviceModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdevicemodelsIdPatch

> PbxDeviceModel PbxdevicemodelsIdPatch(ctx, id, operation)

Change a specific PbxDeviceModel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**PbxDeviceModel**](PbxDeviceModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdevicemodelsIdPut

> PbxDeviceModel PbxdevicemodelsIdPut(ctx, id, pbxDeviceModel)

Replace/change a specific PbxDeviceModel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**pbxDeviceModel** | [**PbxDeviceModel**](PbxDeviceModel.md)|  | 

### Return type

[**PbxDeviceModel**](PbxDeviceModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PbxdevicemodelsPost

> []PbxDeviceModel PbxdevicemodelsPost(ctx, pbxDeviceModel)

Create a new PbxDeviceModel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbxDeviceModel** | [**PbxDeviceModel**](PbxDeviceModel.md)|  | 

### Return type

[**[]PbxDeviceModel**](PbxDeviceModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

