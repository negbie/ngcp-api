# \SubscriberProfileApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriberprofilesGet**](SubscriberProfileApi.md#SubscriberprofilesGet) | **Get** /subscriberprofiles/ | Get SubscriberProfile items
[**SubscriberprofilesIdDelete**](SubscriberProfileApi.md#SubscriberprofilesIdDelete) | **Delete** /subscriberprofiles/{id} | Delete a specific SubscriberProfile
[**SubscriberprofilesIdGet**](SubscriberProfileApi.md#SubscriberprofilesIdGet) | **Get** /subscriberprofiles/{id} | Get a specific SubscriberProfile
[**SubscriberprofilesIdPatch**](SubscriberProfileApi.md#SubscriberprofilesIdPatch) | **Patch** /subscriberprofiles/{id} | Change a specific SubscriberProfile
[**SubscriberprofilesIdPut**](SubscriberProfileApi.md#SubscriberprofilesIdPut) | **Put** /subscriberprofiles/{id} | Replace/change a specific SubscriberProfile
[**SubscriberprofilesPost**](SubscriberProfileApi.md#SubscriberprofilesPost) | **Post** /subscriberprofiles/ | Create a new SubscriberProfile



## SubscriberprofilesGet

> []SubscriberProfile SubscriberprofilesGet(ctx, optional)

Get SubscriberProfile items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubscriberprofilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubscriberprofilesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileSetId** | **optional.String**| Filter for profiles  belonging to a specific profile set | 
 **name** | **optional.String**| Filter for profile with a specific name | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]SubscriberProfile**](SubscriberProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberprofilesIdDelete

> SubscriberprofilesIdDelete(ctx, id)

Delete a specific SubscriberProfile

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


## SubscriberprofilesIdGet

> SubscriberProfile SubscriberprofilesIdGet(ctx, id)

Get a specific SubscriberProfile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**SubscriberProfile**](SubscriberProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberprofilesIdPatch

> SubscriberProfile SubscriberprofilesIdPatch(ctx, id, operation)

Change a specific SubscriberProfile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**SubscriberProfile**](SubscriberProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberprofilesIdPut

> SubscriberProfile SubscriberprofilesIdPut(ctx, id, subscriberProfile)

Replace/change a specific SubscriberProfile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**subscriberProfile** | [**SubscriberProfile**](SubscriberProfile.md)|  | 

### Return type

[**SubscriberProfile**](SubscriberProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberprofilesPost

> []SubscriberProfile SubscriberprofilesPost(ctx, subscriberProfile)

Create a new SubscriberProfile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberProfile** | [**SubscriberProfile**](SubscriberProfile.md)|  | 

### Return type

[**[]SubscriberProfile**](SubscriberProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

