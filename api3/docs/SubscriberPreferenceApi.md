# \SubscriberPreferenceApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriberpreferencesGet**](SubscriberPreferenceApi.md#SubscriberpreferencesGet) | **Get** /subscriberpreferences/ | Get SubscriberPreference items
[**SubscriberpreferencesIdGet**](SubscriberPreferenceApi.md#SubscriberpreferencesIdGet) | **Get** /subscriberpreferences/{id} | Get a specific SubscriberPreference
[**SubscriberpreferencesIdPatch**](SubscriberPreferenceApi.md#SubscriberpreferencesIdPatch) | **Patch** /subscriberpreferences/{id} | Change a specific SubscriberPreference
[**SubscriberpreferencesIdPut**](SubscriberPreferenceApi.md#SubscriberpreferencesIdPut) | **Put** /subscriberpreferences/{id} | Replace/change a specific SubscriberPreference



## SubscriberpreferencesGet

> []map[string]interface{} SubscriberpreferencesGet(ctx, optional)

Get SubscriberPreference items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubscriberpreferencesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubscriberpreferencesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for subscribers of customers belonging to a specific reseller | 
 **contactId** | **optional.String**| Filter for subscribers of contracts with a specific contact id | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberpreferencesIdGet

> map[string]interface{} SubscriberpreferencesIdGet(ctx, id)

Get a specific SubscriberPreference

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberpreferencesIdPatch

> map[string]interface{} SubscriberpreferencesIdPatch(ctx, id, operation)

Change a specific SubscriberPreference

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberpreferencesIdPut

> map[string]interface{} SubscriberpreferencesIdPut(ctx, id, body)

Replace/change a specific SubscriberPreference

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**body** | **map[string]interface{}**|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

