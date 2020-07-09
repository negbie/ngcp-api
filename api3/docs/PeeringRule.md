# PeeringRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallerPattern** | **string** | A PCRE regex matching against &#39;sip:user@domain&#39; (e.g. &#39;^sip:.+@example\\.org$&#39; matching the whole URI, or &#39;999&#39; matching if the URI contains &#39;999&#39;) | 
**Description** | **string** | string, rule description | 
**CalleePrefix** | **string** | Callee prefix, eg: 43 | 
**GroupId** | **float32** | The peering group this rule belongs to. | [optional] 
**CalleePattern** | **string** | A PCRE regex matching against the full Request-URI (e.g. &#39;^sip:.+@example\\.org$&#39; or &#39;^sip:431&#39;) | 
**Stopper** | **bool** | Stop processing of further rules if this rule matches. | 
**Enabled** | **bool** | Rule enabled state. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


