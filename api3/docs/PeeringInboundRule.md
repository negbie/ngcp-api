# PeeringInboundRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | **string** | A PCRE regex matching against the specified field (e.g. &#39;^sip:.+@example\\.org$&#39; or &#39;^sip:431&#39;) when matching against a full URI | 
**RejectReason** | **string** | If reject code is specified and the call is rejected, the reason in the response is taken from this value | 
**Field** | **string** | The field of the inbound SIP message to match the pattern against | 
**RejectCode** | **float32** | If specified, the call is rejected if the source IP of the request is found in a peering server of the group, but the given pattern does not match; Range of 400-699 | 
**Priority** | **float32** | The priority of this rule. Lower values mean higher priority. | [optional] 
**GroupId** | **float32** | The peering group this rule belongs to. | [optional] 
**Enabled** | **bool** | Rule enabled state. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


