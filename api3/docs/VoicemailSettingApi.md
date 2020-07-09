# \VoicemailSettingApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VoicemailsettingsGet**](VoicemailSettingApi.md#VoicemailsettingsGet) | **Get** /voicemailsettings/ | Get VoicemailSetting items
[**VoicemailsettingsIdGet**](VoicemailSettingApi.md#VoicemailsettingsIdGet) | **Get** /voicemailsettings/{id} | Get a specific VoicemailSetting
[**VoicemailsettingsIdPatch**](VoicemailSettingApi.md#VoicemailsettingsIdPatch) | **Patch** /voicemailsettings/{id} | Change a specific VoicemailSetting
[**VoicemailsettingsIdPut**](VoicemailSettingApi.md#VoicemailsettingsIdPut) | **Put** /voicemailsettings/{id} | Replace/change a specific VoicemailSetting



## VoicemailsettingsGet

> []VoicemailSetting VoicemailsettingsGet(ctx, optional)

Get VoicemailSetting items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VoicemailsettingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VoicemailsettingsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriberId** | **optional.String**| Filter for voicemail settings of a specific subscriber | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]VoicemailSetting**](VoicemailSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoicemailsettingsIdGet

> VoicemailSetting VoicemailsettingsIdGet(ctx, id)

Get a specific VoicemailSetting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**VoicemailSetting**](VoicemailSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoicemailsettingsIdPatch

> VoicemailSetting VoicemailsettingsIdPatch(ctx, id, operation)

Change a specific VoicemailSetting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**VoicemailSetting**](VoicemailSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoicemailsettingsIdPut

> VoicemailSetting VoicemailsettingsIdPut(ctx, id, voicemailSetting)

Replace/change a specific VoicemailSetting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**voicemailSetting** | [**VoicemailSetting**](VoicemailSetting.md)|  | 

### Return type

[**VoicemailSetting**](VoicemailSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

