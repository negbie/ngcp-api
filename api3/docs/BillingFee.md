# BillingFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingZoneId** | **float32** | The billing zone id this fee belongs to. | [optional] 
**OffpeakInitInterval** | **float32** | The length of the first interval during offpeak hours in seconds (e.g. 60). | [optional] 
**OffpeakInitRate** | **float32** | The cost per second of the first interval during offpeak hours in cents (e.g. 0.70 cents). | 
**Direction** | **string** | The call direction when to apply this fee (either for inbound or outbound calls). | [optional] 
**OnpeakFollowInterval** | **float32** | The length of each following interval during onpeak hours in seconds (e.g. 30). | [optional] 
**Destination** | **string** | A string (eg. 431001), string prefix (eg. 43) or PCRE regular expression (eg. ^.+$) to match the called number or sip uri. | [optional] 
**PurgeExisting** | **bool** | If fees are uploaded via text/csv bulk upload, this option defines whether to purge any existing fees for the given billing profile before inserting the new ones. | 
**UseFreeTime** | **bool** | Whether free minutes may be used when calling this destination. | 
**OnpeakInitInterval** | **float32** | The length of the first interval during onpeak hours in seconds (e.g. 60). | [optional] 
**OnpeakFollowRate** | **float32** | The cost per second of each following interval during onpeak hours in cents (e.g. 0.90 cents). | 
**Source** | **string** | A string (eg. 431001), string prefix (eg. 43) or PCRE regular expression (eg. ^.+$) to match the calling number or sip uri. | 
**OffpeakFollowInterval** | **float32** | The length of each following interval during offpeak hours in seconds (e.g. 30). | [optional] 
**BillingProfileId** | **float32** | The billing profile this billing fee belongs to. | [optional] 
**OnpeakInitRate** | **float32** | The cost per second of the first interval during onpeak hours (e.g. 0.90 cent). | 
**OffpeakFollowRate** | **float32** | The cost per second of each following interval during offpeak hours in cents (e.g. 0.70 cents). | 
**MatchMode** | **string** | The mode how the the fee&#39;s source/destination has to match a call&#39;s source/destination. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


