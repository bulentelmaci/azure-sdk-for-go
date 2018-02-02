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
	"encoding/xml"
	"reflect"
	"time"
	"unsafe"
)

const (
	rfc3339Format = "2006-01-02T15:04:05.0000000Z07:00"
)

// used to convert times from UTC to GMT before sending across the wire
var gmt = time.FixedZone("GMT", 0)

// internal type used for marshalling time in RFC1123 format
type timeRFC1123 struct {
	time.Time
}

// MarshalText implements the encoding.TextMarshaler interface for timeRFC1123.
func (t timeRFC1123) MarshalText() ([]byte, error) {
	return []byte(t.Format(time.RFC1123)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for timeRFC1123.
func (t *timeRFC1123) UnmarshalText(data []byte) (err error) {
	t.Time, err = time.Parse(time.RFC1123, string(data))
	return
}

// internal type used for marshalling time in RFC3339 format
type timeRFC3339 struct {
	time.Time
}

// MarshalText implements the encoding.TextMarshaler interface for timeRFC3339.
func (t timeRFC3339) MarshalText() ([]byte, error) {
	return []byte(t.Format(rfc3339Format)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for timeRFC3339.
func (t *timeRFC3339) UnmarshalText(data []byte) (err error) {
	t.Time, err = time.Parse(rfc3339Format, string(data))
	return
}

// internal type used for marshalling
type dataLakeAnalyticsAccountPropertiesBasic struct {
	AccountID         *uuid.UUID                         `json:"accountId,omitempty"`
	ProvisioningState DataLakeAnalyticsAccountStatusType `json:"provisioningState,omitempty"`
	State             DataLakeAnalyticsAccountStateType  `json:"state,omitempty"`
	CreationTime      *timeRFC3339                       `json:"creationTime,omitempty"`
	LastModifiedTime  *timeRFC3339                       `json:"lastModifiedTime,omitempty"`
	Endpoint          *string                            `json:"endpoint,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for DataLakeAnalyticsAccountPropertiesBasic.
func (dlaapb DataLakeAnalyticsAccountPropertiesBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if reflect.TypeOf((*DataLakeAnalyticsAccountPropertiesBasic)(nil)).Elem().Size() != reflect.TypeOf((*dataLakeAnalyticsAccountPropertiesBasic)(nil)).Elem().Size() {
		panic("size mismatch between DataLakeAnalyticsAccountPropertiesBasic and dataLakeAnalyticsAccountPropertiesBasic")
	}
	dlaapb2 := (*dataLakeAnalyticsAccountPropertiesBasic)(unsafe.Pointer(&dlaapb))
	return e.EncodeElement(*dlaapb2, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for DataLakeAnalyticsAccountPropertiesBasic.
func (dlaapb *DataLakeAnalyticsAccountPropertiesBasic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if reflect.TypeOf((*DataLakeAnalyticsAccountPropertiesBasic)(nil)).Elem().Size() != reflect.TypeOf((*dataLakeAnalyticsAccountPropertiesBasic)(nil)).Elem().Size() {
		panic("size mismatch between DataLakeAnalyticsAccountPropertiesBasic and dataLakeAnalyticsAccountPropertiesBasic")
	}
	dlaapb2 := (*dataLakeAnalyticsAccountPropertiesBasic)(unsafe.Pointer(dlaapb))
	return d.DecodeElement(dlaapb2, &start)
}

// internal type used for marshalling
type storageContainerProperties struct {
	LastModifiedTime *timeRFC3339 `json:"lastModifiedTime,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for StorageContainerProperties.
func (scp StorageContainerProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if reflect.TypeOf((*StorageContainerProperties)(nil)).Elem().Size() != reflect.TypeOf((*storageContainerProperties)(nil)).Elem().Size() {
		panic("size mismatch between StorageContainerProperties and storageContainerProperties")
	}
	scp2 := (*storageContainerProperties)(unsafe.Pointer(&scp))
	return e.EncodeElement(*scp2, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for StorageContainerProperties.
func (scp *StorageContainerProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if reflect.TypeOf((*StorageContainerProperties)(nil)).Elem().Size() != reflect.TypeOf((*storageContainerProperties)(nil)).Elem().Size() {
		panic("size mismatch between StorageContainerProperties and storageContainerProperties")
	}
	scp2 := (*storageContainerProperties)(unsafe.Pointer(scp))
	return d.DecodeElement(scp2, &start)
}
