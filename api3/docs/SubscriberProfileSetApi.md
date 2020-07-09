# \SubscriberProfileSetApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriberprofilesetsGet**](SubscriberProfileSetApi.md#SubscriberprofilesetsGet) | **Get** /subscriberprofilesets/ | Get SubscriberProfileSet items
[**SubscriberprofilesetsIdDelete**](SubscriberProfileSetApi.md#SubscriberprofilesetsIdDelete) | **Delete** /subscriberprofilesets/{id} | Delete a specific SubscriberProfileSet
[**SubscriberprofilesetsIdGet**](SubscriberProfileSetApi.md#SubscriberprofilesetsIdGet) | **Get** /subscriberprofilesets/{id} | Get a specific SubscriberProfileSet
[**SubscriberprofilesetsIdPatch**](SubscriberProfileSetApi.md#SubscriberprofilesetsIdPatch) | **Patch** /subscriberprofilesets/{id} | Change a specific SubscriberProfileSet
[**SubscriberprofilesetsIdPut**](SubscriberProfileSetApi.md#SubscriberprofilesetsIdPut) | **Put** /subscriberprofilesets/{id} | Replace/change a specific SubscriberProfileSet
[**SubscriberprofilesetsPost**](SubscriberProfileSetApi.md#SubscriberprofilesetsPost) | **Post** /subscriberprofilesets/ | Create a new SubscriberProfileSet



## SubscriberprofilesetsGet

> []SubscriberProfileSet SubscriberprofilesetsGet(ctx, optional)

Get SubscriberProfileSet items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubscriberprofilesetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubscriberprofilesetsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for profile sets belonging to a specific reseller | 
 **name** | **optional.String**| Filter for profile sets with a specific name | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]SubscriberProfileSet**](SubscriberProfileSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberprofilesetsIdDelete

> SubscriberprofilesetsIdDelete(ctx, id)

Delete a specific SubscriberProfileSet

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


## SubscriberprofilesetsIdGet

> SubscriberProfileSet SubscriberprofilesetsIdGet(ctx, id)

Get a specific SubscriberProfileSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**SubscriberProfileSet**](SubscriberProfileSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberprofilesetsIdPatch

> SubscriberProfileSet SubscriberprofilesetsIdPatch(ctx, id, operation)

Change a specific SubscriberProfileSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**SubscriberProfileSet**](SubscriberProfileSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberprofilesetsIdPut

> SubscriberProfileSet SubscriberprofilesetsIdPut(ctx, id, subscriberProfileSet)

Replace/change a specific SubscriberProfileSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**subscriberProfileSet** | [**SubscriberProfileSet**](SubscriberProfileSet.md)|  | 

### Return type

[**SubscriberProfileSet**](SubscriberProfileSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberprofilesetsPost

> []SubscriberProfileSet SubscriberprofilesetsPost(ctx, subscriberProfileSet)

Create a new SubscriberProfileSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberProfileSet** | [**SubscriberProfileSet**](SubscriberProfileSet.md)|  | 

### Return type

[**[]SubscriberProfileSet**](SubscriberProfileSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

