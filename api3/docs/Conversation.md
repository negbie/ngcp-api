# Conversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The conversation event type: call/voicemail/sms/fax/xmpp. | [optional] 
**Status** | **string** | Status of the conversation. Possible values are: \&quot;ok\&quot;,\&quot;busy\&quot;,\&quot;noanswer\&quot;,\&quot;cancel\&quot;,\&quot;offline\&quot;,\&quot;timeout\&quot;,\&quot;other\&quot;. | [optional] 
**Context** | **string** | TBD | 
**Currency** | **string** | Call cost currency. | 
**Caller** | **string** | Conversation initiator. | [optional] 
**Duration** | **string** | Conversation duration. | [optional] 
**TotalCustomerCost** | **float32** | Call cost with applied customer VAT. | 
**CallType** | **string** | One of the \&quot;call\&quot;,\&quot;cfu\&quot;,\&quot;cft\&quot;,\&quot;cfb\&quot;,\&quot;cfna\&quot;,\&quot;cfs\&quot;,\&quot;cfr\&quot;. | [optional] 
**VoicemailSubscriberId** | **float32** | The subscriber id the message belongs to. | 
**Callee** | **string** | Conversation receiver. | [optional] 
**StartTime** | **string** | The timestamp of the conversation event. | [optional] 
**Pages** | **float32** | Number of the pages in the fax document. | 
**Folder** | **string** | The folder the message is currently in (one of INBOX, Old, Work, Friends, Family, Cust1-Cust6). | 
**RatingStatus** | **string** | Status of the rate processing for the conversation. Possible values are: \&quot;unrated\&quot;,\&quot;ok\&quot;,\&quot;failed\&quot;. | [optional] 
**Direction** | **string** | Conversation direction. | [optional] 
**Id** | **float32** | The original conversation record id - cdr id/voicemail id/sms id/fax journal record id/prosody message archive mgmt (mam) record id. | [optional] 
**SubscriberId** | **float32** | Subscriber who can manage fax record. | 
**Filename** | **string** | Filename of the fax document. | 
**CallId** | **string** | Call id. | [optional] 
**CustomerCost** | **float32** | Call cost. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


