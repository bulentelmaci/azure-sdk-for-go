package compute

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
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// SnapshotsClient is the compute Client
type SnapshotsClient struct {
	ManagementClient
}

// NewSnapshotsClient creates an instance of the SnapshotsClient client.
func NewSnapshotsClient(p pipeline.Pipeline) SnapshotsClient {
	return SnapshotsClient{NewManagementClient(p)}
}

// CreateOrUpdate creates or updates a snapshot. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. snapshotName is the name of the snapshot within the given
// subscription and resource group. snapshot is snapshot object supplied in the body of the Put disk operation.
func (client SnapshotsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, snapshotName string, snapshot Snapshot) (*Snapshot, error) {
	if err := validate([]validation{
		{targetValue: snapshot,
			constraints: []constraint{{target: "snapshot.DiskProperties", name: null, rule: false,
				chain: []constraint{{target: "snapshot.DiskProperties.CreationData", name: null, rule: true,
					chain: []constraint{{target: "snapshot.DiskProperties.CreationData.ImageReference", name: null, rule: false,
						chain: []constraint{{target: "snapshot.DiskProperties.CreationData.ImageReference.ID", name: null, rule: true, chain: nil}}},
					}},
					{target: "snapshot.DiskProperties.EncryptionSettings", name: null, rule: false,
						chain: []constraint{{target: "snapshot.DiskProperties.EncryptionSettings.DiskEncryptionKey", name: null, rule: false,
							chain: []constraint{{target: "snapshot.DiskProperties.EncryptionSettings.DiskEncryptionKey.SourceVault", name: null, rule: true, chain: nil},
								{target: "snapshot.DiskProperties.EncryptionSettings.DiskEncryptionKey.SecretURL", name: null, rule: true, chain: nil},
							}},
							{target: "snapshot.DiskProperties.EncryptionSettings.KeyEncryptionKey", name: null, rule: false,
								chain: []constraint{{target: "snapshot.DiskProperties.EncryptionSettings.KeyEncryptionKey.SourceVault", name: null, rule: true, chain: nil},
									{target: "snapshot.DiskProperties.EncryptionSettings.KeyEncryptionKey.KeyURL", name: null, rule: true, chain: nil},
								}},
						}},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(resourceGroupName, snapshotName, snapshot)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Snapshot), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SnapshotsClient) createOrUpdatePreparer(resourceGroupName string, snapshotName string, snapshot Snapshot) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(snapshot)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// createOrUpdateResponder handles the response to the CreateOrUpdate request.
func (client SnapshotsClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &Snapshot{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Delete deletes a snapshot. This method may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. snapshotName is the name of the snapshot within the given
// subscription and resource group.
func (client SnapshotsClient) Delete(ctx context.Context, resourceGroupName string, snapshotName string) (*OperationStatusResponse, error) {
	req, err := client.deletePreparer(resourceGroupName, snapshotName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*OperationStatusResponse), err
}

// deletePreparer prepares the Delete request.
func (client SnapshotsClient) deletePreparer(resourceGroupName string, snapshotName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client SnapshotsClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	result := &OperationStatusResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Get gets information about a snapshot.
//
// resourceGroupName is the name of the resource group. snapshotName is the name of the snapshot within the given
// subscription and resource group.
func (client SnapshotsClient) Get(ctx context.Context, resourceGroupName string, snapshotName string) (*Snapshot, error) {
	req, err := client.getPreparer(resourceGroupName, snapshotName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Snapshot), err
}

// getPreparer prepares the Get request.
func (client SnapshotsClient) getPreparer(resourceGroupName string, snapshotName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client SnapshotsClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Snapshot{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GrantAccess grants access to a snapshot. This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. snapshotName is the name of the snapshot within the given
// subscription and resource group. grantAccessData is access data object supplied in the body of the get snapshot
// access operation.
func (client SnapshotsClient) GrantAccess(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData GrantAccessData) (*AccessURI, error) {
	if err := validate([]validation{
		{targetValue: grantAccessData,
			constraints: []constraint{{target: "grantAccessData.DurationInSeconds", name: null, rule: true, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.grantAccessPreparer(resourceGroupName, snapshotName, grantAccessData)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.grantAccessResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*AccessURI), err
}

// grantAccessPreparer prepares the GrantAccess request.
func (client SnapshotsClient) grantAccessPreparer(resourceGroupName string, snapshotName string, grantAccessData GrantAccessData) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}/beginGetAccess"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(grantAccessData)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// grantAccessResponder handles the response to the GrantAccess request.
func (client SnapshotsClient) grantAccessResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &AccessURI{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// List lists snapshots under a subscription.
func (client SnapshotsClient) List(ctx context.Context) (*SnapshotList, error) {
	req, err := client.listPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*SnapshotList), err
}

// listPreparer prepares the List request.
func (client SnapshotsClient) listPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/snapshots"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client SnapshotsClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &SnapshotList{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListByResourceGroup lists snapshots under a resource group.
//
// resourceGroupName is the name of the resource group.
func (client SnapshotsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (*SnapshotList, error) {
	req, err := client.listByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByResourceGroupResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*SnapshotList), err
}

// listByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client SnapshotsClient) listByResourceGroupPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listByResourceGroupResponder handles the response to the ListByResourceGroup request.
func (client SnapshotsClient) listByResourceGroupResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &SnapshotList{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// RevokeAccess revokes access to a snapshot. This method may poll for completion. Polling can be canceled by passing
// the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. snapshotName is the name of the snapshot within the given
// subscription and resource group.
func (client SnapshotsClient) RevokeAccess(ctx context.Context, resourceGroupName string, snapshotName string) (*OperationStatusResponse, error) {
	req, err := client.revokeAccessPreparer(resourceGroupName, snapshotName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.revokeAccessResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*OperationStatusResponse), err
}

// revokeAccessPreparer prepares the RevokeAccess request.
func (client SnapshotsClient) revokeAccessPreparer(resourceGroupName string, snapshotName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}/endGetAccess"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// revokeAccessResponder handles the response to the RevokeAccess request.
func (client SnapshotsClient) revokeAccessResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &OperationStatusResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Update updates (patches) a snapshot. This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. snapshotName is the name of the snapshot within the given
// subscription and resource group. snapshot is snapshot object supplied in the body of the Patch snapshot operation.
func (client SnapshotsClient) Update(ctx context.Context, resourceGroupName string, snapshotName string, snapshot SnapshotUpdate) (*Snapshot, error) {
	req, err := client.updatePreparer(resourceGroupName, snapshotName, snapshot)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Snapshot), err
}

// updatePreparer prepares the Update request.
func (client SnapshotsClient) updatePreparer(resourceGroupName string, snapshotName string, snapshot SnapshotUpdate) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	req, err := pipeline.NewRequest("PATCH", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(snapshot)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// updateResponder handles the response to the Update request.
func (client SnapshotsClient) updateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &Snapshot{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}
