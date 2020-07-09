# \SpeedDialApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpeeddialsGet**](SpeedDialApi.md#SpeeddialsGet) | **Get** /speeddials/ | Get SpeedDial items
[**SpeeddialsIdGet**](SpeedDialApi.md#SpeeddialsIdGet) | **Get** /speeddials/{id} | Get a specific SpeedDial
[**SpeeddialsIdPatch**](SpeedDialApi.md#SpeeddialsIdPatch) | **Patch** /speeddials/{id} | Change a specific SpeedDial
[**SpeeddialsIdPut**](SpeedDialApi.md#SpeeddialsIdPut) | **Put** /speeddials/{id} | Replace/change a specific SpeedDial



## SpeeddialsGet

> []SpeedDial SpeeddialsGet(ctx, optional)

Get SpeedDial items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpeeddialsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SpeeddialsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nonempty** | **optional.String**| Filter for subscribers with nonempty speeddials | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]SpeedDial**](SpeedDial.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpeeddialsIdGet

> SpeedDial SpeeddialsIdGet(ctx, id)

Get a specific SpeedDial

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**SpeedDial**](SpeedDial.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpeeddialsIdPatch

> SpeedDial SpeeddialsIdPatch(ctx, id, operation)

Change a specific SpeedDial

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**SpeedDial**](SpeedDial.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpeeddialsIdPut

> SpeedDial SpeeddialsIdPut(ctx, id, speedDial)

Replace/change a specific SpeedDial

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**speedDial** | [**SpeedDial**](SpeedDial.md)|  | 

### Return type

[**SpeedDial**](SpeedDial.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

