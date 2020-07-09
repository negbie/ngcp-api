# \MailToFaxSettingApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MailtofaxsettingsGet**](MailToFaxSettingApi.md#MailtofaxsettingsGet) | **Get** /mailtofaxsettings/ | Get MailToFaxSetting items
[**MailtofaxsettingsIdGet**](MailToFaxSettingApi.md#MailtofaxsettingsIdGet) | **Get** /mailtofaxsettings/{id} | Get a specific MailToFaxSetting
[**MailtofaxsettingsIdPatch**](MailToFaxSettingApi.md#MailtofaxsettingsIdPatch) | **Patch** /mailtofaxsettings/{id} | Change a specific MailToFaxSetting
[**MailtofaxsettingsIdPut**](MailToFaxSettingApi.md#MailtofaxsettingsIdPut) | **Put** /mailtofaxsettings/{id} | Replace/change a specific MailToFaxSetting



## MailtofaxsettingsGet

> []MailToFaxSetting MailtofaxsettingsGet(ctx, optional)

Get MailToFaxSetting items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MailtofaxsettingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MailtofaxsettingsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **optional.String**| Filter for items (subscribers) with active mail to fax settings | 
 **secretKeyRenew** | **optional.String**| Filter for items (subscribers) where secret_key_renew field matches given pattern | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]MailToFaxSetting**](MailToFaxSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MailtofaxsettingsIdGet

> MailToFaxSetting MailtofaxsettingsIdGet(ctx, id)

Get a specific MailToFaxSetting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**MailToFaxSetting**](MailToFaxSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MailtofaxsettingsIdPatch

> MailToFaxSetting MailtofaxsettingsIdPatch(ctx, id, operation)

Change a specific MailToFaxSetting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**MailToFaxSetting**](MailToFaxSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MailtofaxsettingsIdPut

> MailToFaxSetting MailtofaxsettingsIdPut(ctx, id, mailToFaxSetting)

Replace/change a specific MailToFaxSetting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**mailToFaxSetting** | [**MailToFaxSetting**](MailToFaxSetting.md)|  | 

### Return type

[**MailToFaxSetting**](MailToFaxSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

