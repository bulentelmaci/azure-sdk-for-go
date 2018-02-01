package network

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

// VirtualNetworkGatewayConnectionsClient is the network Client
type VirtualNetworkGatewayConnectionsClient struct {
	ManagementClient
}

// NewVirtualNetworkGatewayConnectionsClient creates an instance of the VirtualNetworkGatewayConnectionsClient client.
func NewVirtualNetworkGatewayConnectionsClient(p pipeline.Pipeline) VirtualNetworkGatewayConnectionsClient {
	return VirtualNetworkGatewayConnectionsClient{NewManagementClient(p)}
}

// CreateOrUpdate creates or updates a virtual network gateway connection in the specified resource group. This method
// may poll for completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to
// cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkGatewayConnectionName is the name of the virtual
// network gateway connection. parameters is parameters supplied to the create or update virtual network gateway
// connection operation.
func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection) (*VirtualNetworkGatewayConnection, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.VirtualNetworkGatewayConnectionPropertiesFormat", name: null, rule: true,
				chain: []constraint{{target: "parameters.VirtualNetworkGatewayConnectionPropertiesFormat.VirtualNetworkGateway1", name: null, rule: true,
					chain: []constraint{{target: "parameters.VirtualNetworkGatewayConnectionPropertiesFormat.VirtualNetworkGateway1.VirtualNetworkGatewayPropertiesFormat", name: null, rule: true, chain: nil}}},
					{target: "parameters.VirtualNetworkGatewayConnectionPropertiesFormat.VirtualNetworkGateway2", name: null, rule: false,
						chain: []constraint{{target: "parameters.VirtualNetworkGatewayConnectionPropertiesFormat.VirtualNetworkGateway2.VirtualNetworkGatewayPropertiesFormat", name: null, rule: true, chain: nil}}},
					{target: "parameters.VirtualNetworkGatewayConnectionPropertiesFormat.LocalNetworkGateway2", name: null, rule: false,
						chain: []constraint{{target: "parameters.VirtualNetworkGatewayConnectionPropertiesFormat.LocalNetworkGateway2.LocalNetworkGatewayPropertiesFormat", name: null, rule: true, chain: nil}}},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VirtualNetworkGatewayConnection), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworkGatewayConnectionsClient) createOrUpdatePreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-08-01")
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(parameters)
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
func (client VirtualNetworkGatewayConnectionsClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	result := &VirtualNetworkGatewayConnection{rawResponse: resp.Response()}
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

// Delete deletes the specified virtual network Gateway connection. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkGatewayConnectionName is the name of the virtual
// network gateway connection.
func (client VirtualNetworkGatewayConnectionsClient) Delete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, virtualNetworkGatewayConnectionName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// deletePreparer prepares the Delete request.
func (client VirtualNetworkGatewayConnectionsClient) deletePreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-08-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client VirtualNetworkGatewayConnectionsClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// Get gets the specified virtual network gateway connection by resource group.
//
// resourceGroupName is the name of the resource group. virtualNetworkGatewayConnectionName is the name of the virtual
// network gateway connection.
func (client VirtualNetworkGatewayConnectionsClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*VirtualNetworkGatewayConnection, error) {
	req, err := client.getPreparer(resourceGroupName, virtualNetworkGatewayConnectionName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VirtualNetworkGatewayConnection), err
}

// getPreparer prepares the Get request.
func (client VirtualNetworkGatewayConnectionsClient) getPreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-08-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client VirtualNetworkGatewayConnectionsClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &VirtualNetworkGatewayConnection{rawResponse: resp.Response()}
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

// GetSharedKey the Get VirtualNetworkGatewayConnectionSharedKey operation retrieves information about the specified
// virtual network gateway connection shared key through Network resource provider.
//
// resourceGroupName is the name of the resource group. virtualNetworkGatewayConnectionName is the virtual network
// gateway connection shared key name.
func (client VirtualNetworkGatewayConnectionsClient) GetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*ConnectionSharedKey, error) {
	req, err := client.getSharedKeyPreparer(resourceGroupName, virtualNetworkGatewayConnectionName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getSharedKeyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ConnectionSharedKey), err
}

// getSharedKeyPreparer prepares the GetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) getSharedKeyPreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-08-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getSharedKeyResponder handles the response to the GetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) getSharedKeyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ConnectionSharedKey{rawResponse: resp.Response()}
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

// List the List VirtualNetworkGatewayConnections operation retrieves all the virtual network gateways connections
// created.
//
// resourceGroupName is the name of the resource group.
func (client VirtualNetworkGatewayConnectionsClient) List(ctx context.Context, resourceGroupName string) (*VirtualNetworkGatewayConnectionListResult, error) {
	req, err := client.listPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VirtualNetworkGatewayConnectionListResult), err
}

// listPreparer prepares the List request.
func (client VirtualNetworkGatewayConnectionsClient) listPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-08-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client VirtualNetworkGatewayConnectionsClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &VirtualNetworkGatewayConnectionListResult{rawResponse: resp.Response()}
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

// ResetSharedKey the VirtualNetworkGatewayConnectionResetSharedKey operation resets the virtual network gateway
// connection shared key for passed virtual network gateway connection in the specified resource group through Network
// resource provider. This method may poll for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkGatewayConnectionName is the virtual network
// gateway connection reset shared key Name. parameters is parameters supplied to the begin reset virtual network
// gateway connection shared key operation through network resource provider.
func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey) (*ConnectionResetSharedKey, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.KeyLength", name: null, rule: true,
				chain: []constraint{{target: "parameters.KeyLength", name: inclusiveMaximum, rule: 128, chain: nil},
					{target: "parameters.KeyLength", name: inclusiveMinimum, rule: 1, chain: nil},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.resetSharedKeyPreparer(resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.resetSharedKeyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ConnectionResetSharedKey), err
}

// resetSharedKeyPreparer prepares the ResetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) resetSharedKeyPreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey/reset"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-08-01")
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(parameters)
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

// resetSharedKeyResponder handles the response to the ResetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) resetSharedKeyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &ConnectionResetSharedKey{rawResponse: resp.Response()}
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

// SetSharedKey the Put VirtualNetworkGatewayConnectionSharedKey operation sets the virtual network gateway connection
// shared key for passed virtual network gateway connection in the specified resource group through Network resource
// provider. This method may poll for completion. Polling can be canceled by passing the cancel channel argument. The
// channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkGatewayConnectionName is the virtual network
// gateway connection name. parameters is parameters supplied to the Begin Set Virtual Network Gateway connection
// Shared key operation throughNetwork resource provider.
func (client VirtualNetworkGatewayConnectionsClient) SetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey) (*ConnectionSharedKey, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.Value", name: null, rule: true, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.setSharedKeyPreparer(resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setSharedKeyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ConnectionSharedKey), err
}

// setSharedKeyPreparer prepares the SetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) setSharedKeyPreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-08-01")
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(parameters)
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

// setSharedKeyResponder handles the response to the SetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) setSharedKeyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusCreated, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ConnectionSharedKey{rawResponse: resp.Response()}
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
