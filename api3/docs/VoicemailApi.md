# \VoicemailApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VoicemailsGet**](VoicemailApi.md#VoicemailsGet) | **Get** /voicemails/ | Get Voicemail items
[**VoicemailsIdDelete**](VoicemailApi.md#VoicemailsIdDelete) | **Delete** /voicemails/{id} | Delete a specific Voicemail
[**VoicemailsIdGet**](VoicemailApi.md#VoicemailsIdGet) | **Get** /voicemails/{id} | Get a specific Voicemail
[**VoicemailsIdPatch**](VoicemailApi.md#VoicemailsIdPatch) | **Patch** /voicemails/{id} | Change a specific Voicemail
[**VoicemailsIdPut**](VoicemailApi.md#VoicemailsIdPut) | **Put** /voicemails/{id} | Replace/change a specific Voicemail



## VoicemailsGet

> []Voicemail VoicemailsGet(ctx, optional)

Get Voicemail items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VoicemailsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VoicemailsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tz** | **optional.String**| Format timestamp according to the optional time zone provided here, e.g. Europe/Berlin. | 
 **useOwnerTz** | **optional.String**| Format timestamp according to the filtered customer&#39;s/subscribers&#39;s inherited time zone. | 
 **subscriberId** | **optional.String**| Filter for voicemails belonging to a specific subscriber | 
 **folder** | **optional.String**| Filter for voicemails in a specific folder (one of INBOX, Old, Friends, Family, Cust1 to Cust4) | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]Voicemail**](Voicemail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoicemailsIdDelete

> VoicemailsIdDelete(ctx, id)

Delete a specific Voicemail

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


## VoicemailsIdGet

> Voicemail VoicemailsIdGet(ctx, id)

Get a specific Voicemail

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Voicemail**](Voicemail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoicemailsIdPatch

> Voicemail VoicemailsIdPatch(ctx, id, operation)

Change a specific Voicemail

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**Voicemail**](Voicemail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoicemailsIdPut

> Voicemail VoicemailsIdPut(ctx, id, voicemail)

Replace/change a specific Voicemail

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**voicemail** | [**Voicemail**](Voicemail.md)|  | 

### Return type

[**Voicemail**](Voicemail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

