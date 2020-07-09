# SubscriberRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | **string** | The SIP URI pointing to the current contact of the subscriber. Should be a full sip uri, sip:user@ip:port. | [optional] 
**Path** | **string** | Readonly lb/path field. Composed from \&quot;socket\&quot; and internal configuration information. | 
**Nat** | **bool** | The registered contact is detected as behind NAT. | 
**SubscriberId** | **float32** | The subscriber the contact belongs to. | [optional] 
**Socket** | **string** | Points to the LB interface from which the incoming calls to this registration should be sent out. | 
**Q** | **float32** | The contact priority for serial forking (float value, higher is stronger) between -1.00 to 1.00 | [optional] 
**Expires** | **string** | The expire timestamp of the registered contact. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


