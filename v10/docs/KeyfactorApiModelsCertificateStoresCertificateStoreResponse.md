# KeyfactorApiModelsCertificateStoresCertificateStoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string indicating the GUID of the certificate store within Keyfactor Command. This ID is automatically set by Keyfactor Command. | [optional] 
**ContainerId** | Pointer to **int64** | An integer indicating the ID of the certificate store&#39;s associated certificate store container, if applicable (see GET Certificate Store Containers). | [optional] 
**DisplayName** | Pointer to **string** | A string indicating the display name of the certificate store. | [optional] 
**ClientMachine** | Pointer to **string** | The string value of the client machine. The value for this will vary depending on the certificate store type. For example, for a Java keystore or an F5 device, it is the hostname of the machine on which the store is located, but for an Amazon Web Services store, it is the FQDN of the Keyfactor Command Windows Orchestrator. See Adding or Modifying a Certificate Store in the Keyfactor Command Reference Guide for more information. | [optional] 
**Storepath** | Pointer to **string** | A string indicating the path to the certificate store on the target. The format for this path will vary depending on the certificate store type. For example, for a Java keystore, this will be a file path (e.g. /opt/myapp/store.jks), but for an F5 device, this will be a partition name on the device (e.g. Common). See Adding or Modifying a Certificate Store in the Keyfactor Command Reference Guide for more information. The maximum number of characters supported in this field is 722. | [optional] 
**CertStoreInventoryJobId** | Pointer to **string** | A string indicating the GUID that identifies the inventory job for the certificate store in the Keyfactor Command database. This will be null if an inventory schedule is not set for the certificate store. | [optional] 
**CertStoreType** | Pointer to **int64** | An integer indicating the ID of the certificate store type, as defined in Keyfactor Command, for this certificate store. (0-Javakeystore,2-PEMFile, 3-F5SSLProfiles,4-IISRoots, 5-NetScaler, 6-IISPersonal, 7-F5WebServer, 8-IISRevoked, 9-F5WebServerREST, 10-F5SSLProfilesREST, 11-F5CABundlesREST, 100-AmazonWebServices, 101-FileTransferProtocol) | [optional] 
**Approved** | Pointer to **bool** | A Boolean that indicates whether a certificate store is approved (true) or not (false). If a certificate store is approved, it can be used and updated. A certificate store that has been discovered using the discover feature but not yet marked as approved will be false here. | [optional] 
**CreateIfMissing** | Pointer to **bool** | A Boolean that indicates whether a new certificate store should be created with the information provided (true) or not (false). This option is only valid for Java keystores and any custom certificate store types you have defined to support this functionality. | [optional] 
**Properties** | Pointer to **string** | Some types of certificate stores have additional properties that are stored in this parameter. The data is stored in a series of, typically, key value pairs that define the property name and value (see GET Certificate Store Types for more information).  As of Keyfactor Command v10, this parameter is used to store certificate store server usernames, server passwords, and the UseSSL flag. Built-in certificate stores that typically require configuration of certificate store server parameters include NetScaler and F5 stores. The legacy methods for managing certificate store server credentials have been deprecated but are retained for backwards compatiblity. For more information, see POST Certificate Stores Server.  When reading this field, the values are returned as simple key value pairs, with the values being individual values. When writing, the values are specified as objects, though they are typically single values.  | [optional] 
**AgentId** | Pointer to **string** | A string indicating the Keyfactor Command GUID of the orchestrator for this store. | [optional] 
**AgentAssigned** | Pointer to **bool** | A Boolean that indicates whether there is an orchestrator assigned to this certificate store (true) or not (false). | [optional] 
**ContainerName** | Pointer to **string** | A string indicating the name of the certificate store&#39;s associated container, if applicable. | [optional] 
**InventorySchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**ReenrollmentStatus** | Pointer to [**ModelsReenrollmentStatus**](ModelsReenrollmentStatus.md) |  | [optional] 
**SetNewPasswordAllowed** | Pointer to **bool** | A Boolean that indicates whether the store password can be changed (true) or not (false). | [optional] 

## Methods

### NewKeyfactorApiModelsCertificateStoresCertificateStoreResponse

`func NewKeyfactorApiModelsCertificateStoresCertificateStoreResponse() *KeyfactorApiModelsCertificateStoresCertificateStoreResponse`

NewKeyfactorApiModelsCertificateStoresCertificateStoreResponse instantiates a new KeyfactorApiModelsCertificateStoresCertificateStoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsCertificateStoresCertificateStoreResponseWithDefaults

`func NewKeyfactorApiModelsCertificateStoresCertificateStoreResponseWithDefaults() *KeyfactorApiModelsCertificateStoresCertificateStoreResponse`

NewKeyfactorApiModelsCertificateStoresCertificateStoreResponseWithDefaults instantiates a new KeyfactorApiModelsCertificateStoresCertificateStoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContainerId

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetClientMachine

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### GetStorepath

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetStorepath() string`

GetStorepath returns the Storepath field if non-nil, zero value otherwise.

### GetStorepathOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetStorepathOk() (*string, bool)`

GetStorepathOk returns a tuple with the Storepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorepath

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetStorepath(v string)`

SetStorepath sets Storepath field to given value.

### HasStorepath

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasStorepath() bool`

HasStorepath returns a boolean if a field has been set.

### GetCertStoreInventoryJobId

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetCertStoreInventoryJobId() string`

GetCertStoreInventoryJobId returns the CertStoreInventoryJobId field if non-nil, zero value otherwise.

### GetCertStoreInventoryJobIdOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetCertStoreInventoryJobIdOk() (*string, bool)`

GetCertStoreInventoryJobIdOk returns a tuple with the CertStoreInventoryJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreInventoryJobId

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetCertStoreInventoryJobId(v string)`

SetCertStoreInventoryJobId sets CertStoreInventoryJobId field to given value.

### HasCertStoreInventoryJobId

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasCertStoreInventoryJobId() bool`

HasCertStoreInventoryJobId returns a boolean if a field has been set.

### GetCertStoreType

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetCertStoreType() int64`

GetCertStoreType returns the CertStoreType field if non-nil, zero value otherwise.

### GetCertStoreTypeOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetCertStoreTypeOk() (*int64, bool)`

GetCertStoreTypeOk returns a tuple with the CertStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreType

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetCertStoreType(v int64)`

SetCertStoreType sets CertStoreType field to given value.

### HasCertStoreType

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasCertStoreType() bool`

HasCertStoreType returns a boolean if a field has been set.

### GetApproved

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetCreateIfMissing

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetCreateIfMissing() bool`

GetCreateIfMissing returns the CreateIfMissing field if non-nil, zero value otherwise.

### GetCreateIfMissingOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetCreateIfMissingOk() (*bool, bool)`

GetCreateIfMissingOk returns a tuple with the CreateIfMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfMissing

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetCreateIfMissing(v bool)`

SetCreateIfMissing sets CreateIfMissing field to given value.

### HasCreateIfMissing

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasCreateIfMissing() bool`

HasCreateIfMissing returns a boolean if a field has been set.

### GetProperties

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetAgentId

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAgentAssigned

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetAgentAssigned() bool`

GetAgentAssigned returns the AgentAssigned field if non-nil, zero value otherwise.

### GetAgentAssignedOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetAgentAssignedOk() (*bool, bool)`

GetAgentAssignedOk returns a tuple with the AgentAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAssigned

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetAgentAssigned(v bool)`

SetAgentAssigned sets AgentAssigned field to given value.

### HasAgentAssigned

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasAgentAssigned() bool`

HasAgentAssigned returns a boolean if a field has been set.

### GetContainerName

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetInventorySchedule

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetInventorySchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetInventorySchedule returns the InventorySchedule field if non-nil, zero value otherwise.

### GetInventoryScheduleOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetInventoryScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetInventoryScheduleOk returns a tuple with the InventorySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySchedule

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetInventorySchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetInventorySchedule sets InventorySchedule field to given value.

### HasInventorySchedule

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasInventorySchedule() bool`

HasInventorySchedule returns a boolean if a field has been set.

### GetReenrollmentStatus

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetReenrollmentStatus() ModelsReenrollmentStatus`

GetReenrollmentStatus returns the ReenrollmentStatus field if non-nil, zero value otherwise.

### GetReenrollmentStatusOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetReenrollmentStatusOk() (*ModelsReenrollmentStatus, bool)`

GetReenrollmentStatusOk returns a tuple with the ReenrollmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReenrollmentStatus

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetReenrollmentStatus(v ModelsReenrollmentStatus)`

SetReenrollmentStatus sets ReenrollmentStatus field to given value.

### HasReenrollmentStatus

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasReenrollmentStatus() bool`

HasReenrollmentStatus returns a boolean if a field has been set.

### GetSetNewPasswordAllowed

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetSetNewPasswordAllowed() bool`

GetSetNewPasswordAllowed returns the SetNewPasswordAllowed field if non-nil, zero value otherwise.

### GetSetNewPasswordAllowedOk

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) GetSetNewPasswordAllowedOk() (*bool, bool)`

GetSetNewPasswordAllowedOk returns a tuple with the SetNewPasswordAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetNewPasswordAllowed

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) SetSetNewPasswordAllowed(v bool)`

SetSetNewPasswordAllowed sets SetNewPasswordAllowed field to given value.

### HasSetNewPasswordAllowed

`func (o *KeyfactorApiModelsCertificateStoresCertificateStoreResponse) HasSetNewPasswordAllowed() bool`

HasSetNewPasswordAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


