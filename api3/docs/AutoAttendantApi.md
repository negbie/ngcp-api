# \AutoAttendantApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoattendantsGet**](AutoAttendantApi.md#AutoattendantsGet) | **Get** /autoattendants/ | Get AutoAttendant items
[**AutoattendantsIdGet**](AutoAttendantApi.md#AutoattendantsIdGet) | **Get** /autoattendants/{id} | Get a specific AutoAttendant
[**AutoattendantsIdPatch**](AutoAttendantApi.md#AutoattendantsIdPatch) | **Patch** /autoattendants/{id} | Change a specific AutoAttendant
[**AutoattendantsIdPut**](AutoAttendantApi.md#AutoattendantsIdPut) | **Put** /autoattendants/{id} | Replace/change a specific AutoAttendant



## AutoattendantsGet

> []AutoAttendant AutoattendantsGet(ctx, optional)

Get AutoAttendant items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutoattendantsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AutoattendantsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]AutoAttendant**](AutoAttendant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoattendantsIdGet

> AutoAttendant AutoattendantsIdGet(ctx, id)

Get a specific AutoAttendant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**AutoAttendant**](AutoAttendant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoattendantsIdPatch

> AutoAttendant AutoattendantsIdPatch(ctx, id, operation)

Change a specific AutoAttendant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**AutoAttendant**](AutoAttendant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoattendantsIdPut

> AutoAttendant AutoattendantsIdPut(ctx, id, autoAttendant)

Replace/change a specific AutoAttendant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**autoAttendant** | [**AutoAttendant**](AutoAttendant.md)|  | 

### Return type

[**AutoAttendant**](AutoAttendant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

