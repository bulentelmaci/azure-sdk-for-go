package location

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
	"github.com/Azure/go-autorest/autorest"
)

// KeyType enumerates the values for key type.
type KeyType string

const (
	// Primary ...
	Primary KeyType = "primary"
	// Secondary ...
	Secondary KeyType = "secondary"
)

// BasedServicesAccount an Azure resource which represents access to a suite of Location Based Services REST APIs.
type BasedServicesAccount struct {
	autorest.Response `json:"-"`
	// ID - The fully qualified Location Based Services Account resource identifier.
	ID *string `json:"id,omitempty"`
	// Name - The name of the Location Based Services Account, which is unique within a Resource Group.
	Name *string `json:"name,omitempty"`
	// Type - Azure resource type.
	Type *string `json:"type,omitempty"`
	// LocationProperty - The location of the resource.
	LocationProperty *string `json:"location,omitempty"`
	// Tags - Gets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags *map[string]*string `json:"tags,omitempty"`
	// Sku - The SKU of this account.
	Sku *Sku `json:"sku,omitempty"`
}

// BasedServicesAccountCreateParameters parameters used to create a new Location Based Services Account.
type BasedServicesAccountCreateParameters struct {
	// LocationProperty - The location of the resource.
	LocationProperty *string `json:"location,omitempty"`
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags *map[string]*string `json:"tags,omitempty"`
	// Sku - The SKU of this account.
	Sku *Sku `json:"sku,omitempty"`
}

// BasedServicesAccountKeys the set of keys which can be used to access the Location Based Services REST APIs. Two keys
// are provided for key rotation without interruption.
type BasedServicesAccountKeys struct {
	autorest.Response `json:"-"`
	// ID - The full Azure resource identifier of the Location Based Services Account.
	ID *string `json:"id,omitempty"`
	// PrimaryKey - The primary key for accessing the Location Based Services REST APIs.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - The secondary key for accessing the Location Based Services REST APIs.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// BasedServicesAccounts a list of Location Based Services Accounts.
type BasedServicesAccounts struct {
	autorest.Response `json:"-"`
	// Value - A Location Based Services Account.
	Value *[]BasedServicesAccount `json:"value,omitempty"`
}

// BasedServicesAccountsMoveRequest the description of what resources to move between resource groups.
type BasedServicesAccountsMoveRequest struct {
	// TargetResourceGroup - The name of the destination resource group.
	TargetResourceGroup *string `json:"targetResourceGroup,omitempty"`
	// ResourceIds - A list of resource names to move from the source resource group.
	ResourceIds *[]string `json:"resourceIds,omitempty"`
}

// BasedServicesAccountUpdateParameters parameters used to update an existing Location Based Services Account.
type BasedServicesAccountUpdateParameters struct {
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags *map[string]*string `json:"tags,omitempty"`
	// Sku - The SKU of this account.
	Sku *Sku `json:"sku,omitempty"`
}

// BasedServicesKeySpecification whether the operation refers to the primary or secondary key.
type BasedServicesKeySpecification struct {
	// KeyType - Whether the operation refers to the primary or secondary key. Possible values include: 'Primary', 'Secondary'
	KeyType KeyType `json:"keyType,omitempty"`
}

// BasedServicesOperations the set of operations available for Location Based Services.
type BasedServicesOperations struct {
	autorest.Response `json:"-"`
	// Value - An operation available for Location Based Services.
	Value *[]BasedServicesOperationsValueItem `json:"value,omitempty"`
}

// BasedServicesOperationsValueItem ...
type BasedServicesOperationsValueItem struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The human-readable description of the operation.
	Display *BasedServicesOperationsValueItemDisplay `json:"display,omitempty"`
	// Origin - The origin of the operation.
	Origin *string `json:"origin,omitempty"`
}

// BasedServicesOperationsValueItemDisplay the human-readable description of the operation.
type BasedServicesOperationsValueItemDisplay struct {
	// Provider - Service provider: Microsoft Location Based Services.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - The action that users can perform, based on their permission level.
	Operation *string `json:"operation,omitempty"`
	// Description - The description of the operation.
	Description *string `json:"description,omitempty"`
}

// Error this object is returned when an error occurs in the Location Based Service API
type Error struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - If available, a human readable description of the error.
	Message *string `json:"message,omitempty"`
	// Target - If available, the component generating the error.
	Target *string `json:"target,omitempty"`
	// Details - If available, a list of additional details about the error.
	Details *[]ErrorDetailsItem `json:"details,omitempty"`
}

// ErrorDetailsItem ...
type ErrorDetailsItem struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - If available, a human readable description of the error.
	Message *string `json:"message,omitempty"`
	// Target - If available, the component generating the error.
	Target *string `json:"target,omitempty"`
}

// Resource an Azure resource
type Resource struct {
	// ID - The fully qualified Location Based Services Account resource identifier.
	ID *string `json:"id,omitempty"`
	// Name - The name of the Location Based Services Account, which is unique within a Resource Group.
	Name *string `json:"name,omitempty"`
	// Type - Azure resource type.
	Type *string `json:"type,omitempty"`
}

// Sku the SKU of the Location Based Services Account.
type Sku struct {
	// Name - The name of the SKU, in standard format (such as S0).
	Name *string `json:"name,omitempty"`
	// Tier - Gets the sku tier. This is based on the SKU name.
	Tier *string `json:"tier,omitempty"`
}