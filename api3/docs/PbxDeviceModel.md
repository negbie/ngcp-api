# PbxDeviceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapConfigRedirectPolycomPassword** | **string** | Password used to configure bootstrap url on Polycom redirect server. | 
**BootstrapConfigRedirectPanasonicPassword** | **string** | Password used to configure bootstrap url on Panasonic redirect server. Obtained from Panasonic. | 
**Type** | **string** | Phone or the phone extension | [optional] 
**BootstrapConfigRedirectPolycomUser** | **string** | Username used to configure bootstrap url on Polycom redirect server. | 
**BootstrapConfigHttpSyncMethod** | **string** | The HTTP method to set the provisioning server (one of GET, POST). | 
**BootstrapConfigRedirectSnomUser** | **string** | Username used to configure bootstrap url on Snom redirect server. Obtained from Snom. | 
**BootstrapUri** | **string** | Custom provisioning server URI. | 
**ExtensionsNum** | **float32** | The max number of extension modules which can be attached to this device model. | 
**Linerange** | [**[]PbxDeviceModelLinerange**](PbxDeviceModel_linerange.md) | An array of line/key definitions for this device. Each element is a hash containing the keys name, can_private, can_shared, can_blf and keys (which in turn is an array of hashes having x, y and labelpos allowing top, bottom, left right). | [optional] 
**ResellerId** | **float32** | The reseller id this device model belongs to. | [optional] 
**Model** | **string** | The model name of this device. | [optional] 
**BootstrapConfigRedirectYealinkPassword** | **string** | Password used to configure bootstrap url on Yealink redirect server. Obtained from Yealink. | 
**BootstrapConfigHttpSyncParams** | **string** | The parameters appended to the sync URI when setting the provisioning server, e.g. server.uri/$MA. The server.uri variable is automatically expanded during provisioning time. | 
**BootstrapMethod** | **string** | Method to configure the provisioning server on the phone. One of http, redirect_panasonic, redirect_yealink, redirect_polycom, redirect_snom, redirect_grandstream. | [optional] 
**ConnectableModelsId** | **float32** | The extension modules which can be attached to this device model. | 
**BootstrapConfigRedirectPanasonicUser** | **string** | Username used to configure bootstrap url on Panasonic redirect server. Obtained from Panasonic. | 
**BootstrapConfigRedirectPolycomProfile** | **string** | Preliminary created in ZeroTouch Provisioning console Polycom ZTP profile. Refer to documentation. | 
**BootstrapConfigRedirectGrandstreamCid** | **string** | Customer ID for GAPS access. Obtained from Grandstream. | 
**BootstrapConfigRedirectGrandstreamKey** | **string** | Customer key for GAPS access. Obtained from Grandstream. | 
**BootstrapConfigHttpSyncUri** | **string** | The sync URI to set the provisioning server of the device (e.g. http://client.ip/admin/resync. The client.ip variable is automatically expanded during provisioning time. | 
**BootstrapConfigRedirectSnomPassword** | **string** | Password used to configure bootstrap url on Snom redirect server. Obtained from Snom. | 
**BootstrapConfigRedirectYealinkUser** | **string** | Username used to configure bootstrap url on Yealink redirect server. Obtained from Yealink. | 
**Vendor** | **string** | The vendor name of this device. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


