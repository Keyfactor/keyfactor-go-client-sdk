# CertificateCollectionsCertificateCollectionPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | Pointer to **int32** |  | [optional] 
**AccessControlList** | Pointer to [**[]CertificateCollectionsCertificateQueryAccessControl**](CertificateCollectionsCertificateQueryAccessControl.md) |  | [optional] 
**AssignableRoles** | Pointer to [**[]CertificateCollectionsAssignableQueryRole**](CertificateCollectionsAssignableQueryRole.md) |  | [optional] 

## Methods

### NewCertificateCollectionsCertificateCollectionPermissionsResponse

`func NewCertificateCollectionsCertificateCollectionPermissionsResponse() *CertificateCollectionsCertificateCollectionPermissionsResponse`

NewCertificateCollectionsCertificateCollectionPermissionsResponse instantiates a new CertificateCollectionsCertificateCollectionPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCollectionsCertificateCollectionPermissionsResponseWithDefaults

`func NewCertificateCollectionsCertificateCollectionPermissionsResponseWithDefaults() *CertificateCollectionsCertificateCollectionPermissionsResponse`

NewCertificateCollectionsCertificateCollectionPermissionsResponseWithDefaults instantiates a new CertificateCollectionsCertificateCollectionPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) GetQueryId() int32`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) GetQueryIdOk() (*int32, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) SetQueryId(v int32)`

SetQueryId sets QueryId field to given value.

### HasQueryId

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) HasQueryId() bool`

HasQueryId returns a boolean if a field has been set.

### GetAccessControlList

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) GetAccessControlList() []CertificateCollectionsCertificateQueryAccessControl`

GetAccessControlList returns the AccessControlList field if non-nil, zero value otherwise.

### GetAccessControlListOk

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) GetAccessControlListOk() (*[]CertificateCollectionsCertificateQueryAccessControl, bool)`

GetAccessControlListOk returns a tuple with the AccessControlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlList

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) SetAccessControlList(v []CertificateCollectionsCertificateQueryAccessControl)`

SetAccessControlList sets AccessControlList field to given value.

### HasAccessControlList

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) HasAccessControlList() bool`

HasAccessControlList returns a boolean if a field has been set.

### SetAccessControlListNil

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) SetAccessControlListNil(b bool)`

 SetAccessControlListNil sets the value for AccessControlList to be an explicit nil

### UnsetAccessControlList
`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) UnsetAccessControlList()`

UnsetAccessControlList ensures that no value is present for AccessControlList, not even an explicit nil
### GetAssignableRoles

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) GetAssignableRoles() []CertificateCollectionsAssignableQueryRole`

GetAssignableRoles returns the AssignableRoles field if non-nil, zero value otherwise.

### GetAssignableRolesOk

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) GetAssignableRolesOk() (*[]CertificateCollectionsAssignableQueryRole, bool)`

GetAssignableRolesOk returns a tuple with the AssignableRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableRoles

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) SetAssignableRoles(v []CertificateCollectionsAssignableQueryRole)`

SetAssignableRoles sets AssignableRoles field to given value.

### HasAssignableRoles

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) HasAssignableRoles() bool`

HasAssignableRoles returns a boolean if a field has been set.

### SetAssignableRolesNil

`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) SetAssignableRolesNil(b bool)`

 SetAssignableRolesNil sets the value for AssignableRoles to be an explicit nil

### UnsetAssignableRoles
`func (o *CertificateCollectionsCertificateCollectionPermissionsResponse) UnsetAssignableRoles()`

UnsetAssignableRoles ensures that no value is present for AssignableRoles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


