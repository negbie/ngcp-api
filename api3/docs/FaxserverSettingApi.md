# \FaxserverSettingApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FaxserversettingsGet**](FaxserverSettingApi.md#FaxserversettingsGet) | **Get** /faxserversettings/ | Get FaxserverSetting items
[**FaxserversettingsIdGet**](FaxserverSettingApi.md#FaxserversettingsIdGet) | **Get** /faxserversettings/{id} | Get a specific FaxserverSetting
[**FaxserversettingsIdPatch**](FaxserverSettingApi.md#FaxserversettingsIdPatch) | **Patch** /faxserversettings/{id} | Change a specific FaxserverSetting
[**FaxserversettingsIdPut**](FaxserverSettingApi.md#FaxserversettingsIdPut) | **Put** /faxserversettings/{id} | Replace/change a specific FaxserverSetting



## FaxserversettingsGet

> []FaxserverSetting FaxserversettingsGet(ctx, optional)

Get FaxserverSetting items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FaxserversettingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FaxserversettingsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameOrPassword** | **optional.String**| Filter for items (subscribers) where name or password field match given pattern | 
 **active** | **optional.String**| Filter for items (subscribers) with active faxserver settings | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]FaxserverSetting**](FaxserverSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FaxserversettingsIdGet

> FaxserverSetting FaxserversettingsIdGet(ctx, id)

Get a specific FaxserverSetting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**FaxserverSetting**](FaxserverSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FaxserversettingsIdPatch

> FaxserverSetting FaxserversettingsIdPatch(ctx, id, operation)

Change a specific FaxserverSetting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**FaxserverSetting**](FaxserverSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FaxserversettingsIdPut

> FaxserverSetting FaxserversettingsIdPut(ctx, id, faxserverSetting)

Replace/change a specific FaxserverSetting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**faxserverSetting** | [**FaxserverSetting**](FaxserverSetting.md)|  | 

### Return type

[**FaxserverSetting**](FaxserverSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

