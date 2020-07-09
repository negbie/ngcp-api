# CfSourceSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the source set | [optional] 
**SubscriberId** | **float32** | The subscriber id this source set belongs to | [optional] 
**IsRegex** | **bool** | A flag indicating, whether the numbers in this set are regular expressions. If true, all sources will be interepreted as perl compatible regular expressions and matched against the calling party number (in E164 format) of the calls. If false, the whole numbers are plainly matched while shell patterns like 431* or 49123[1-5]67 are possible. | 
**Mode** | **string** | Source set mode. If set to \&quot;blacklist\&quot; it enables forwarding for everything except numbers in the list, and \&quot;whitelist\&quot; only enables forwards for numbers defined in this list. This field is mandatory. | [optional] 
**Sources** | [**[]CfSourceSetSources**](CFSourceSet_sources.md) | An array of sources, each containing the key \&quot;source\&quot; which will be matched against the calling party number to determine whether to apply the callforward or not. \&quot;source\&quot; is the calling party number in E164 format to match. Regular expressions or shell patterns can be used depending on the is_regex flag. Use \&quot;anonymous\&quot; to match suppressed numbers. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


