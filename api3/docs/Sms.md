# Sms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The content of the message | [optional] 
**Direction** | **string** | Whether the logged message is sent, received or forwarded | 
**Caller** | **string** | A valid caller number in the E164 format. Must be valid according to the preferences alloed_clis, user_cli, cli | 
**Callee** | **string** | A valid callee number in the E164 format | [optional] 
**Reason** | **string** | An error message in case of a failed transmission | 
**SubscriberId** | **float32** | The subscriber id this journal entry belongs to | [optional] 
**Status** | **string** | Whether the message has been sent successfully | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


