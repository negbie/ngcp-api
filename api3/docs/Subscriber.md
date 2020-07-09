# Subscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username for SIP and XMPP services. | [optional] 
**IsPbxPilot** | **bool** | Whether this subscriber is used as PBX pilot subscriber. | 
**ProfileSetId** | **float32** | The profile set defining the possible feature sets for this subscriber. | 
**Password** | **string** | The password to authenticate for SIP and XMPP services. | [optional] 
**Webusername** | **string** | The username to log into the CSC Panel. | 
**Administrative** | **bool** | Whether the subscriber can configure other subscribers within his Customer account. | 
**IsPbxGroup** | **bool** | Whether this subscriber is used as PBX group. | 
**Status** | **string** | The status of the subscriber (one of \&quot;active\&quot;, \&quot;locked\&quot;, \&quot;terminated\&quot;). | 
**PbxExtension** | **string** | The PBX extension used for short dialling. If provided, the primary number will automatically be derived from the pilot subscriber&#39;s primary number suffixed by this extension. | 
**DisplayName** | **string** | The person&#39;s name, which is then used in XMPP contact lists or auto-provisioned phones, and which can be used as network-provided display name in SIP calls. | 
**Webpassword** | **string** | The password to log into the CSC Panel. | 
**PrimaryNumber** | [**SubscriberPrimaryNumber**](Subscriber_primary_number.md) |  | 
**Lock** | **string** | The lock level of the subscriber. | 
**DomainId** | **float32** | The domain id this subscriber belongs to. | [optional] 
**PbxGroupIds** | **[]map[string]interface{}** | An array of PBX group ids this subscriber belongs to. | 
**Email** | **string** | The email address of the subscriber. | 
**ExternalId** | **string** | A non-unique external ID e.g., provided by a 3rd party provisioning | 
**Domain** | **string** | The domain name this subscriber belongs to. | 
**PbxGroupmemberIdsId** | **[]map[string]interface{}** | An array of PBX subscriber ids belonging to this group. | 
**AliasNumbers** | [**[]SubscriberAliasNumbers**](Subscriber_alias_numbers.md) | Additional E.164 numbers (each containing a cc, ac and sn attribute) mapped to this subscriber for inbound calls. | 
**Timezone** | **string** | The timezone of the subscriber. See http://search.cpan.org/dist/DateTime-TimeZone/lib/DateTime/TimeZone/Catalog.pm for a full list of valid timezone names. | 
**ProfileId** | **float32** | The profile defining the actual feature set for this subscriber. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


