# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingProfileId** | **float32** | The billing profile id used to charge this contract, which will become active immediately. This field is required if the profile definition mode is not defined or the &#39;id&#39; mode is used. | 
**BillingProfileDefinition** | **string** | Explicitly declare the way how you want to set billing profiles for this API call. | 
**ExternalId** | **string** | A non-unique external ID e.g., provided by a 3rd party provisioning | 
**Type** | **string** | Either \&quot;sippeering\&quot; or \&quot;reseller\&quot;. | [optional] 
**Status** | **string** | The status of the contract. | [optional] 
**BillingProfiles** | [**[]ContractBillingProfiles**](Contract_billing_profiles.md) | The billing profile / billing network interval schedule used to charge this contract can be specified. It is represented by an array of objects, each containing the keys \&quot;start\&quot;, \&quot;stop\&quot;, \&quot;profile_id\&quot; and \&quot;network_id\&quot; (/api/customers/ only). When POSTing, it has to contain a single interval with empty \&quot;start\&quot; and \&quot;stop\&quot; fields. Only intervals beginning in the future can be updated afterwards. This field is required if the &#39;profiles&#39; profile definition mode is used. | 
**ContactId** | **float32** | The contact id this contract belongs to. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


