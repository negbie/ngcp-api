# PbxDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **float32** | Geräte-Profil | [optional] 
**StationName** | **string** | Stations-Name | [optional] 
**Lines** | [**[]PbxDeviceLines**](PbxDevice_lines.md) | The lines for this pbx device. Required keys are &#39;linerange&#39; (name of range to use), &#39;key_num&#39; (key number in line range, starting from 0), &#39;type&#39; (one of &#39;private&#39;, &#39;shared&#39;, &#39;blf&#39;), &#39;subscriber_id&#39; (the subscriber mapped to this key). | [optional] 
**CustomerId** | **float32** | The customer contract this device is belonging to. | [optional] 
**Identifier** | **string** | MAC Adresse / Identifizierung | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


