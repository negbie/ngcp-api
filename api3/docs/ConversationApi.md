# \ConversationApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConversationsGet**](ConversationApi.md#ConversationsGet) | **Get** /conversations/ | Get Conversation items
[**ConversationsIdGet**](ConversationApi.md#ConversationsIdGet) | **Get** /conversations/{id} | Get a specific Conversation



## ConversationsGet

> []Conversation ConversationsGet(ctx, optional)

Get Conversation items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConversationsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tz** | **optional.String**| Format timestamp according to the optional time zone provided here, e.g. Europe/Berlin. | 
 **useOwnerTz** | **optional.String**| Format timestamp according to the filtered customer&#39;s/subscribers&#39;s inherited time zone. | 
 **subscriberId** | **optional.String**| Filter for conversation events of a specific subscriber. Either this or customer_id filter is mandatory if called by admin, reseller or subscriberadmin. | 
 **customerId** | **optional.String**| Filter for conversation events for a specific customer. Either this or subscriber_id filter is mandatory if called by admin, reseller or subscriberadmin. | 
 **direction** | **optional.String**| Filter for conversation events with a specific direction. One of \&quot;in\&quot;, \&quot;out\&quot;. Voicemails are considered as incoming only. | 
 **status** | **optional.String**| todo | 
 **type_** | **optional.String**| Filter for conversation events of given types (\&quot;call\&quot;, \&quot;voicemail\&quot;, \&quot;sms\&quot;, \&quot;fax\&quot;, \&quot;xmpp\&quot;). Multiple types can be included by concatenating type strings, eg. \&quot;?type&#x3D;call-voicemial\&quot;. | 
 **from** | **optional.String**| Filter for conversation events starting greater or equal the specified time stamp. | 
 **to** | **optional.String**| Filter for conversation events starting lower or equal the specified time stamp. | 
 **faxNumberRewriteMode** | **optional.String**| Force the fax numbers normalization logic (available: &#39;default&#39;, &#39;extended&#39;). | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]Conversation**](Conversation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConversationsIdGet

> Conversation ConversationsIdGet(ctx, id)

Get a specific Conversation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Conversation**](Conversation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

