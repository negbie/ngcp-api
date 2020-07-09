# \PhonebookEntryApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PhonebookentriesGet**](PhonebookEntryApi.md#PhonebookentriesGet) | **Get** /phonebookentries/ | Get PhonebookEntry items
[**PhonebookentriesIdDelete**](PhonebookEntryApi.md#PhonebookentriesIdDelete) | **Delete** /phonebookentries/{id} | Delete a specific PhonebookEntry
[**PhonebookentriesIdGet**](PhonebookEntryApi.md#PhonebookentriesIdGet) | **Get** /phonebookentries/{id} | Get a specific PhonebookEntry
[**PhonebookentriesIdPatch**](PhonebookEntryApi.md#PhonebookentriesIdPatch) | **Patch** /phonebookentries/{id} | Change a specific PhonebookEntry
[**PhonebookentriesIdPut**](PhonebookEntryApi.md#PhonebookentriesIdPut) | **Put** /phonebookentries/{id} | Replace/change a specific PhonebookEntry
[**PhonebookentriesPost**](PhonebookEntryApi.md#PhonebookentriesPost) | **Post** /phonebookentries/ | Create a new PhonebookEntry



## PhonebookentriesGet

> []PhonebookEntry PhonebookentriesGet(ctx, optional)

Get PhonebookEntry items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PhonebookentriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PhonebookentriesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for Phonebook entries belonging to a specific reseeller | 
 **customerId** | **optional.String**| Filter for Phonebook entries belonging to a specific contract | 
 **subscriberId** | **optional.String**| Filter for Phonebook entries belonging to a specific subscriber | 
 **number** | **optional.String**| Filter for Phonebook numbers with a specific number (wildcards possible) | 
 **name** | **optional.String**| Filter for Phonebook numbers with a specific name (wildcards possible) | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]PhonebookEntry**](PhonebookEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PhonebookentriesIdDelete

> PhonebookentriesIdDelete(ctx, id)

Delete a specific PhonebookEntry

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


## PhonebookentriesIdGet

> PhonebookEntry PhonebookentriesIdGet(ctx, id)

Get a specific PhonebookEntry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**PhonebookEntry**](PhonebookEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PhonebookentriesIdPatch

> PhonebookEntry PhonebookentriesIdPatch(ctx, id, operation)

Change a specific PhonebookEntry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**PhonebookEntry**](PhonebookEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PhonebookentriesIdPut

> PhonebookEntry PhonebookentriesIdPut(ctx, id, phonebookEntry)

Replace/change a specific PhonebookEntry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**phonebookEntry** | [**PhonebookEntry**](PhonebookEntry.md)|  | 

### Return type

[**PhonebookEntry**](PhonebookEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PhonebookentriesPost

> []PhonebookEntry PhonebookentriesPost(ctx, phonebookEntry)

Create a new PhonebookEntry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phonebookEntry** | [**PhonebookEntry**](PhonebookEntry.md)|  | 

### Return type

[**[]PhonebookEntry**](PhonebookEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

