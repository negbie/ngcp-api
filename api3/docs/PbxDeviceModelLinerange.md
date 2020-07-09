# PbxDeviceModelLinerange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanBlf** | **bool** | Lines/Keys in this range can be used as Busy Lamp Field. Value is accessible in the config template via phone.lineranges.lines.can_blf | 
**CanPrivate** | **bool** | Lines/Keys in this range can be used as regular phone lines. Value is accessible in the config template via phone.lineranges.lines.can_private | 
**Name** | **string** | The Name of this range, e.g. Phone Keys or Attendant Console 1 Keys, accessible in the config template array via phone.lineranges.name | 
**CanShared** | **bool** | Lines/Keys in this range can be used as shared lines. Value is accessible in the config template via phone.lineranges.lines.can_shared | 
**Keys** | [**[]PbxDeviceModelKeys**](PbxDeviceModel_keys.md) | The position of the keys on the front image. Attributes are x, y, labelpos (how the label for the key is displayed in the web interface, relative to the given coordinates; one of top, bottom, left, right). | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


