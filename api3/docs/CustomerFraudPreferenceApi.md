# \CustomerFraudPreferenceApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerfraudpreferencesGet**](CustomerFraudPreferenceApi.md#CustomerfraudpreferencesGet) | **Get** /customerfraudpreferences/ | Get CustomerFraudPreference items
[**CustomerfraudpreferencesIdGet**](CustomerFraudPreferenceApi.md#CustomerfraudpreferencesIdGet) | **Get** /customerfraudpreferences/{id} | Get a specific CustomerFraudPreference
[**CustomerfraudpreferencesIdPatch**](CustomerFraudPreferenceApi.md#CustomerfraudpreferencesIdPatch) | **Patch** /customerfraudpreferences/{id} | Change a specific CustomerFraudPreference
[**CustomerfraudpreferencesIdPut**](CustomerFraudPreferenceApi.md#CustomerfraudpreferencesIdPut) | **Put** /customerfraudpreferences/{id} | Replace/change a specific CustomerFraudPreference



## CustomerfraudpreferencesGet

> []CustomerFraudPreference CustomerfraudpreferencesGet(ctx, optional)

Get CustomerFraudPreference items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerfraudpreferencesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CustomerfraudpreferencesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for fraud preferences of contracts belonging to a specific reseller | 
 **contactId** | **optional.String**| Filter for fraud preferences of contracts with a specific contact id | 
 **notify** | **optional.String**| Filter for fraud preferences of contracts containing a specific notification email address | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CustomerFraudPreference**](CustomerFraudPreference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerfraudpreferencesIdGet

> CustomerFraudPreference CustomerfraudpreferencesIdGet(ctx, id)

Get a specific CustomerFraudPreference

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CustomerFraudPreference**](CustomerFraudPreference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerfraudpreferencesIdPatch

> CustomerFraudPreference CustomerfraudpreferencesIdPatch(ctx, id, operation)

Change a specific CustomerFraudPreference

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**CustomerFraudPreference**](CustomerFraudPreference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerfraudpreferencesIdPut

> CustomerFraudPreference CustomerfraudpreferencesIdPut(ctx, id, customerFraudPreference)

Replace/change a specific CustomerFraudPreference

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**customerFraudPreference** | [**CustomerFraudPreference**](CustomerFraudPreference.md)|  | 

### Return type

[**CustomerFraudPreference**](CustomerFraudPreference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

