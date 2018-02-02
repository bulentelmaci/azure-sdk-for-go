package account

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"net/http"
	"time"
)

// Marker represents an opaque value used in paged responses.
type Marker struct {
	val *string
}

// NotDone returns true if the list enumeration should be started or is not yet complete. Specifically, NotDone returns true
// for a just-initialized (zero value) Marker indicating that you should make an initial request to get a result portion from
// the service. NotDone also returns true whenever the service returns an interim result portion. NotDone returns false only
// after the service has returned the final result portion.
func (m Marker) NotDone() bool {
	return m.val == nil || *m.val != ""
}

// UnmarshalXML implements the xml.Unmarshaler interface for Marker.
func (m *Marker) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var out string
	err := d.DecodeElement(&out, &start)
	m.val = &out
	return err
}

// DataLakeAnalyticsAccountStateType enumerates the values for data lake analytics account state.
type DataLakeAnalyticsAccountStateType string

const (
	// DataLakeAnalyticsAccountStateActive ...
	DataLakeAnalyticsAccountStateActive DataLakeAnalyticsAccountStateType = "active"
	// DataLakeAnalyticsAccountStateNone represents an empty DataLakeAnalyticsAccountStateType.
	DataLakeAnalyticsAccountStateNone DataLakeAnalyticsAccountStateType = ""
	// DataLakeAnalyticsAccountStateSuspended ...
	DataLakeAnalyticsAccountStateSuspended DataLakeAnalyticsAccountStateType = "suspended"
)

// DataLakeAnalyticsAccountStatusType enumerates the values for data lake analytics account status.
type DataLakeAnalyticsAccountStatusType string

const (
	// DataLakeAnalyticsAccountStatusCreating ...
	DataLakeAnalyticsAccountStatusCreating DataLakeAnalyticsAccountStatusType = "Creating"
	// DataLakeAnalyticsAccountStatusDeleted ...
	DataLakeAnalyticsAccountStatusDeleted DataLakeAnalyticsAccountStatusType = "Deleted"
	// DataLakeAnalyticsAccountStatusDeleting ...
	DataLakeAnalyticsAccountStatusDeleting DataLakeAnalyticsAccountStatusType = "Deleting"
	// DataLakeAnalyticsAccountStatusFailed ...
	DataLakeAnalyticsAccountStatusFailed DataLakeAnalyticsAccountStatusType = "Failed"
	// DataLakeAnalyticsAccountStatusNone represents an empty DataLakeAnalyticsAccountStatusType.
	DataLakeAnalyticsAccountStatusNone DataLakeAnalyticsAccountStatusType = ""
	// DataLakeAnalyticsAccountStatusPatching ...
	DataLakeAnalyticsAccountStatusPatching DataLakeAnalyticsAccountStatusType = "Patching"
	// DataLakeAnalyticsAccountStatusResuming ...
	DataLakeAnalyticsAccountStatusResuming DataLakeAnalyticsAccountStatusType = "Resuming"
	// DataLakeAnalyticsAccountStatusRunning ...
	DataLakeAnalyticsAccountStatusRunning DataLakeAnalyticsAccountStatusType = "Running"
	// DataLakeAnalyticsAccountStatusSucceeded ...
	DataLakeAnalyticsAccountStatusSucceeded DataLakeAnalyticsAccountStatusType = "Succeeded"
	// DataLakeAnalyticsAccountStatusSuspending ...
	DataLakeAnalyticsAccountStatusSuspending DataLakeAnalyticsAccountStatusType = "Suspending"
)

// OperationStatusType enumerates the values for operation status.
type OperationStatusType string

const (
	// OperationStatusFailed ...
	OperationStatusFailed OperationStatusType = "Failed"
	// OperationStatusInProgress ...
	OperationStatusInProgress OperationStatusType = "InProgress"
	// OperationStatusNone represents an empty OperationStatusType.
	OperationStatusNone OperationStatusType = ""
	// OperationStatusSucceeded ...
	OperationStatusSucceeded OperationStatusType = "Succeeded"
)

// AddDataLakeStoreParameters - Additional Data Lake Store parameters.
type AddDataLakeStoreParameters struct {
	// Properties - the properties for the Data Lake Store account being added.
	Properties DataLakeStoreAccountInfoProperties `json:"properties,omitempty"`
}

// AddStorageAccountParameters - Additional Azure Storage account parameters.
type AddStorageAccountParameters struct {
	// Properties - the properties for the Azure Storage account being added.
	Properties StorageAccountProperties `json:"properties,omitempty"`
}

// AzureAsyncOperationResult - The response body contains the status of the specified asynchronous operation,
// indicating whether it has succeeded, is inprogress, or has failed. Note that this status is distinct from the HTTP
// status code returned for the Get Operation Status operation itself. If the asynchronous operation succeeded, the
// response body includes the HTTP status code for the successful request. If the asynchronous operation failed, the
// response body includes the HTTP status code for the failed request and error information regarding the failure.
type AzureAsyncOperationResult struct {
	// Status - the status of the AzureAsuncOperation. Possible values include: 'InProgress', 'Succeeded', 'Failed', 'None'
	Status OperationStatusType `json:"status,omitempty"`
	Error  *Error              `json:"error,omitempty"`
}

// BlobContainer - Azure Storage blob container information.
type BlobContainer struct {
	rawResponse *http.Response
	// Name - the name of the blob container.
	Name *string `json:"name,omitempty"`
	// ID - the unique identifier of the blob container.
	ID *string `json:"id,omitempty"`
	// Type - the type of the blob container.
	Type *string `json:"type,omitempty"`
	// Properties - the properties of the blob container.
	Properties *BlobContainerProperties `json:"properties,omitempty"`
}

// Response returns the raw HTTP response object.
func (bc BlobContainer) Response() *http.Response {
	return bc.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (bc BlobContainer) StatusCode() int {
	return bc.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (bc BlobContainer) Status() string {
	return bc.rawResponse.Status
}

// BlobContainerProperties - Azure Storage blob container properties information.
type BlobContainerProperties struct {
	// LastModifiedTime - the last modified time of the blob container.
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`
}

// DataLakeAnalyticsAccount - A Data Lake Analytics account object, containing all information associated with the
// named Data Lake Analytics account.
type DataLakeAnalyticsAccount struct {
	rawResponse *http.Response
	// Location - the account regional location.
	Location *string `json:"location,omitempty"`
	// Name - the account name.
	Name *string `json:"name,omitempty"`
	// Type - the namespace and type of the account.
	Type *string `json:"type,omitempty"`
	// ID - the account subscription ID.
	ID *string `json:"id,omitempty"`
	// Tags - the value of custom properties.
	Tags map[string]string `json:"tags,omitempty"`
	// Properties - the properties defined by Data Lake Analytics all properties are specific to each resource provider.
	Properties *DataLakeAnalyticsAccountProperties `json:"properties,omitempty"`
}

// Response returns the raw HTTP response object.
func (dlaa DataLakeAnalyticsAccount) Response() *http.Response {
	return dlaa.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (dlaa DataLakeAnalyticsAccount) StatusCode() int {
	return dlaa.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (dlaa DataLakeAnalyticsAccount) Status() string {
	return dlaa.rawResponse.Status
}

// DataLakeAnalyticsAccountListDataLakeStoreResult - Data Lake Account list information.
type DataLakeAnalyticsAccountListDataLakeStoreResult struct {
	rawResponse *http.Response
	// Value - the results of the list operation
	Value []DataLakeStoreAccountInfo `json:"value,omitempty"`
	// Count - total number of results.
	Count *int32 `json:"count,omitempty"`
	// NextLink - the link (url) to the next page of results.
	NextLink Marker `json:"NextLink"`
}

// Response returns the raw HTTP response object.
func (dlaaldlsr DataLakeAnalyticsAccountListDataLakeStoreResult) Response() *http.Response {
	return dlaaldlsr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (dlaaldlsr DataLakeAnalyticsAccountListDataLakeStoreResult) StatusCode() int {
	return dlaaldlsr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (dlaaldlsr DataLakeAnalyticsAccountListDataLakeStoreResult) Status() string {
	return dlaaldlsr.rawResponse.Status
}

// DataLakeAnalyticsAccountListResult - DataLakeAnalytics Account list information.
type DataLakeAnalyticsAccountListResult struct {
	rawResponse *http.Response
	// Value - the results of the list operation
	Value []DataLakeAnalyticsAccount `json:"value,omitempty"`
	// NextLink - the link (url) to the next page of results.
	NextLink Marker `json:"NextLink"`
}

// Response returns the raw HTTP response object.
func (dlaalr DataLakeAnalyticsAccountListResult) Response() *http.Response {
	return dlaalr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (dlaalr DataLakeAnalyticsAccountListResult) StatusCode() int {
	return dlaalr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (dlaalr DataLakeAnalyticsAccountListResult) Status() string {
	return dlaalr.rawResponse.Status
}

// DataLakeAnalyticsAccountListStorageAccountsResult - Azure Storage Account list information.
type DataLakeAnalyticsAccountListStorageAccountsResult struct {
	rawResponse *http.Response
	// Value - the results of the list operation
	Value []StorageAccountInfo `json:"value,omitempty"`
	// Count - total number of results.
	Count *int32 `json:"count,omitempty"`
	// NextLink - the link (url) to the next page of results.
	NextLink Marker `json:"NextLink"`
}

// Response returns the raw HTTP response object.
func (dlaalsar DataLakeAnalyticsAccountListStorageAccountsResult) Response() *http.Response {
	return dlaalsar.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (dlaalsar DataLakeAnalyticsAccountListStorageAccountsResult) StatusCode() int {
	return dlaalsar.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (dlaalsar DataLakeAnalyticsAccountListStorageAccountsResult) Status() string {
	return dlaalsar.rawResponse.Status
}

// DataLakeAnalyticsAccountProperties - The account specific properties that are associated with an underlying Data
// Lake Analytics account.
type DataLakeAnalyticsAccountProperties struct {
	// ProvisioningState - the provisioning status of the Data Lake Analytics account. Possible values include: 'Failed', 'Creating', 'Running', 'Succeeded', 'Patching', 'Suspending', 'Resuming', 'Deleting', 'Deleted', 'None'
	ProvisioningState DataLakeAnalyticsAccountStatusType `json:"provisioningState,omitempty"`
	// State - the state of the Data Lake Analytics account. Possible values include: 'Active', 'Suspended', 'None'
	State DataLakeAnalyticsAccountStateType `json:"state,omitempty"`
	// DefaultDataLakeStoreAccount - the default data lake storage account associated with this Data Lake Analytics account.
	DefaultDataLakeStoreAccount *string `json:"defaultDataLakeStoreAccount,omitempty"`
	// MaxDegreeOfParallelism - the maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism *int32 `json:"maxDegreeOfParallelism,omitempty"`
	// MaxJobCount - the maximum supported jobs running under the account at the same time.
	MaxJobCount *int32 `json:"maxJobCount,omitempty"`
	// DataLakeStoreAccounts - the list of Data Lake storage accounts associated with this account.
	DataLakeStoreAccounts []DataLakeStoreAccountInfo `json:"dataLakeStoreAccounts,omitempty"`
	// StorageAccounts - the list of Azure Blob storage accounts associated with this account.
	StorageAccounts []StorageAccountInfo `json:"storageAccounts,omitempty"`
	// CreationTime - the account creation time.
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// LastModifiedTime - the account last modified time.
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`
	// Endpoint - the full CName endpoint for this account.
	Endpoint *string `json:"endpoint,omitempty"`
}

// DataLakeStoreAccountInfo - Data Lake Store account information.
type DataLakeStoreAccountInfo struct {
	rawResponse *http.Response
	// Name - the account name of the Data Lake Store account.
	Name string `json:"name,omitempty"`
	// Properties - the properties associated with this Data Lake Store account.
	Properties *DataLakeStoreAccountInfoProperties `json:"properties,omitempty"`
}

// Response returns the raw HTTP response object.
func (dlsai DataLakeStoreAccountInfo) Response() *http.Response {
	return dlsai.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (dlsai DataLakeStoreAccountInfo) StatusCode() int {
	return dlsai.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (dlsai DataLakeStoreAccountInfo) Status() string {
	return dlsai.rawResponse.Status
}

// DataLakeStoreAccountInfoProperties - Data Lake Store account properties information.
type DataLakeStoreAccountInfoProperties struct {
	// Suffix - the optional suffix for the Data Lake Store account.
	Suffix *string `json:"suffix,omitempty"`
}

// Error - Generic resource error information.
type Error struct {
	// Code - the HTTP status code or error code associated with this error
	Code *string `json:"code,omitempty"`
	// Message - the error message to display.
	Message *string `json:"message,omitempty"`
	// Target - the target of the error.
	Target *string `json:"target,omitempty"`
	// Details - the list of error details
	Details []ErrorDetails `json:"details,omitempty"`
	// InnerError - the inner exceptions or errors, if any
	InnerError *InnerError `json:"innerError,omitempty"`
}

// ErrorDetails - Generic resource error details information.
type ErrorDetails struct {
	// Code - the HTTP status code or error code associated with this error
	Code *string `json:"code,omitempty"`
	// Message - the error message localized based on Accept-Language
	Message *string `json:"message,omitempty"`
	// Target - the target of the particular error (for example, the name of the property in error).
	Target *string `json:"target,omitempty"`
}

// InnerError - Generic resource inner error information.
type InnerError struct {
	// Trace - the stack trace for the error
	Trace *string `json:"trace,omitempty"`
	// Context - the context for the error message
	Context *string `json:"context,omitempty"`
}

// ListBlobContainersResult - The list of blob containers associated with the storage account attached to the Data Lake
// Analytics account.
type ListBlobContainersResult struct {
	rawResponse *http.Response
	// Value - the results of the list operation
	Value []BlobContainer `json:"value,omitempty"`
	// NextLink - the link (url) to the next page of results.
	NextLink Marker `json:"NextLink"`
}

// Response returns the raw HTTP response object.
func (lbcr ListBlobContainersResult) Response() *http.Response {
	return lbcr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (lbcr ListBlobContainersResult) StatusCode() int {
	return lbcr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (lbcr ListBlobContainersResult) Status() string {
	return lbcr.rawResponse.Status
}

// ListSasTokensResult - The SAS response that contains the storage account, container and associated SAS token for
// connection use.
type ListSasTokensResult struct {
	rawResponse *http.Response
	Value       []SasTokenInfo `json:"value,omitempty"`
	// NextLink - the link (url) to the next page of results.
	NextLink Marker `json:"NextLink"`
}

// Response returns the raw HTTP response object.
func (lstr ListSasTokensResult) Response() *http.Response {
	return lstr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (lstr ListSasTokensResult) StatusCode() int {
	return lstr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (lstr ListSasTokensResult) Status() string {
	return lstr.rawResponse.Status
}

// SasTokenInfo - SAS token information.
type SasTokenInfo struct {
	// AccessToken - the access token for the associated Azure Storage Container.
	AccessToken *string `json:"accessToken,omitempty"`
}

// StorageAccountInfo - Azure Storage account information.
type StorageAccountInfo struct {
	rawResponse *http.Response
	// Name - the account name associated with the Azure storage account.
	Name string `json:"name,omitempty"`
	// Properties - the properties associated with this storage account.
	Properties StorageAccountProperties `json:"properties,omitempty"`
}

// Response returns the raw HTTP response object.
func (sai StorageAccountInfo) Response() *http.Response {
	return sai.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (sai StorageAccountInfo) StatusCode() int {
	return sai.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (sai StorageAccountInfo) Status() string {
	return sai.rawResponse.Status
}

// StorageAccountProperties - Azure Storage account properties information.
type StorageAccountProperties struct {
	// AccessKey - the access key associated with this Azure Storage account that will be used to connect to it.
	AccessKey string `json:"accessKey,omitempty"`
	// Suffix - the optional suffix for the Data Lake account.
	Suffix *string `json:"suffix,omitempty"`
}
