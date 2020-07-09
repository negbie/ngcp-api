# \SubscriberApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscribersGet**](SubscriberApi.md#SubscribersGet) | **Get** /subscribers/ | Get Subscriber items
[**SubscribersIdDelete**](SubscriberApi.md#SubscribersIdDelete) | **Delete** /subscribers/{id} | Delete a specific Subscriber
[**SubscribersIdGet**](SubscriberApi.md#SubscribersIdGet) | **Get** /subscribers/{id} | Get a specific Subscriber
[**SubscribersIdPatch**](SubscriberApi.md#SubscribersIdPatch) | **Patch** /subscribers/{id} | Change a specific Subscriber
[**SubscribersIdPut**](SubscriberApi.md#SubscribersIdPut) | **Put** /subscribers/{id} | Replace/change a specific Subscriber
[**SubscribersPost**](SubscriberApi.md#SubscribersPost) | **Post** /subscribers/ | Create a new Subscriber



## SubscribersGet

> []Subscriber SubscribersGet(ctx, optional)

Get Subscriber items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubscribersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubscribersGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **optional.String**| Search for subscribers having a specific subscriber profile | 
 **username** | **optional.String**| Search for specific SIP username | 
 **webusername** | **optional.String**| Search for specific webuser login credentials (exact match) | 
 **webpassword** | **optional.String**| Search for specific webuser login password (exact match) | 
 **domain** | **optional.String**| Filter for subscribers in specific domain | 
 **customerId** | **optional.String**| Filter for subscribers of a specific customer. | 
 **customerExternalId** | **optional.String**| Filter for subscribers of a specific customer external_id. | 
 **subscriberExternalId** | **optional.String**| Filter for subscribers by subscriber&#39;s external_id. | 
 **isPbxGroup** | **optional.String**| Filter for subscribers who are (not) pbx_groups. | 
 **isAdmin** | **optional.String**| Filter for subscribers who are (not) pbx subscriber admins. | 
 **isPbxPilot** | **optional.String**| Filter for subscribers who are pbx pilot subscribers. | 
 **alias** | **optional.String**| Filter for subscribers who has specified alias. | 
 **resellerId** | **optional.String**| Filter for subscribers of customers belonging to a specific reseller | 
 **contactId** | **optional.String**| Filter for subscribers of contracts with a specific contact id | 
 **createTimestampGt** | **optional.String**| Filter for subscriber with create_timestamp greater then specified value | 
 **createTimestampLt** | **optional.String**| Filter for subscriber with create_timestamp less then specified value | 
 **modifyTimestampGt** | **optional.String**| Filter for subscriber with modify_timestamp greater then specified value | 
 **modifyTimestampLt** | **optional.String**| Filter for subscriber with modify_timestamp less then specified value | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]Subscriber**](Subscriber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribersIdDelete

> SubscribersIdDelete(ctx, id)

Delete a specific Subscriber

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


## SubscribersIdGet

> Subscriber SubscribersIdGet(ctx, id)

Get a specific Subscriber

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribersIdPatch

> Subscriber SubscribersIdPatch(ctx, id, operation)

Change a specific Subscriber

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribersIdPut

> Subscriber SubscribersIdPut(ctx, id, subscriber)

Replace/change a specific Subscriber

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**subscriber** | [**Subscriber**](Subscriber.md)|  | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribersPost

> []Subscriber SubscribersPost(ctx, subscriber)

Create a new Subscriber

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriber** | [**Subscriber**](Subscriber.md)|  | 

### Return type

[**[]Subscriber**](Subscriber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

