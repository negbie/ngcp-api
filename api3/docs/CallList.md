# CallList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | **string** | Call direction, either \&quot;in\&quot; or \&quot;out\&quot; | [optional] 
**Clir** | **string** | CLIR indicator for the call | 
**MosAverageRoundtrip** | **string** | MOS average roundtrip. | [optional] 
**RatingStatus** | **string** | The status of the call rating, one of ok, unrated, failed. | [optional] 
**CallId** | **string** | The internal SIP Call-ID of the call. | [optional] 
**CustomerCost** | **float32** | The cost for the customer. | 
**CustomerFreeTime** | **float32** | The number of free seconds of the customer used for this call. | 
**MosAverage** | **string** | MOS average. | [optional] 
**Duration** | **float32** | The duration of the call. | [optional] 
**TotalCustomerCost** | **float32** | Total cost for the customer. VAT is included if applicable. | 
**Status** | **string** | The status of the call, one of ok, busy, noanswer, cancel, offline, timeout, other. | [optional] 
**Type** | **string** | The type of call, one of call, cfu, cfb, cft, cfna, cfs, cfr. | [optional] 
**OtherCli** | **string** | The CLI of the other party, or null if CLIR was active. For intra-PBX calls it is the PBX extension, for inter-PBX calls it is the value of the field specified by the alias_field parameter if available, otherwise the souce_cli or destination_user_in. CLI format is denormalized by caller-out rewrite rule of subscriber. | 
**MosAverageJitter** | **string** | MOS average jitter. | [optional] 
**OwnCli** | **string** | The CLI of the own party. For PBX subscribers it is always the PBX extension, otherwise the source_cli or destination_user_in. CLI format is denormalized by caller-out rewrite rule of subscriber. | [optional] 
**StartTime** | **string** | The timestamp of the call connection. | [optional] 
**IntraCustomer** | **bool** | Whether it is a call between subscribers of one single customer. | [optional] 
**MosAveragePacketloss** | **string** | MOS average packetloss. | [optional] 
**InitTime** | **string** | The timestamp of the call initiation. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


