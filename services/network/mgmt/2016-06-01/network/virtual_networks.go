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

// VirtualNetworksClient is the network Client
type VirtualNetworksClient struct {
	ManagementClient
}

// NewVirtualNetworksClient creates an instance of the VirtualNetworksClient client.
func NewVirtualNetworksClient(p pipeline.Pipeline) VirtualNetworksClient {
	return VirtualNetworksClient{NewManagementClient(p)}
}

// CheckIPAddressAvailability checks whether a private Ip address is available for use.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is the name of the virtual network.
// IPAddress is the private IP address to be verified.
func (client VirtualNetworksClient) CheckIPAddressAvailability(ctx context.Context, resourceGroupName string, virtualNetworkName string, IPAddress *string) (*IPAddressAvailabilityResult, error) {
	req, err := client.checkIPAddressAvailabilityPreparer(resourceGroupName, virtualNetworkName, IPAddress)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.checkIPAddressAvailabilityResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*IPAddressAvailabilityResult), err
}

// checkIPAddressAvailabilityPreparer prepares the CheckIPAddressAvailability request.
func (client VirtualNetworksClient) checkIPAddressAvailabilityPreparer(resourceGroupName string, virtualNetworkName string, IPAddress *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/CheckIPAddressAvailability"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if IPAddress != nil {
		params.Set("ipAddress", *IPAddress)
	}
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// checkIPAddressAvailabilityResponder handles the response to the CheckIPAddressAvailability request.
func (client VirtualNetworksClient) checkIPAddressAvailabilityResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &IPAddressAvailabilityResult{rawResponse: resp.Response()}
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

// CreateOrUpdate the Put VirtualNetwork operation creates/updates a virtual network in the specified resource group.
// This method may poll for completion. Polling can be canceled by passing the cancel channel argument. The channel
// will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is the name of the virtual network.
// parameters is parameters supplied to the create/update Virtual Network operation
func (client VirtualNetworksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork) (*VirtualNetwork, error) {
	req, err := client.createOrUpdatePreparer(resourceGroupName, virtualNetworkName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VirtualNetwork), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworksClient) createOrUpdatePreparer(resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
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
func (client VirtualNetworksClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	result := &VirtualNetwork{rawResponse: resp.Response()}
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

// Delete the Delete VirtualNetwork operation deletes the specified virtual network This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is the name of the virtual network.
func (client VirtualNetworksClient) Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, virtualNetworkName)
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
func (client VirtualNetworksClient) deletePreparer(resourceGroupName string, virtualNetworkName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
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
func (client VirtualNetworksClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusAccepted, http.StatusNoContent, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// Get the Get VirtualNetwork operation retrieves information about the specified virtual network.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is the name of the virtual network. expand
// is expand references resources.
func (client VirtualNetworksClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, expand *string) (*VirtualNetwork, error) {
	req, err := client.getPreparer(resourceGroupName, virtualNetworkName, expand)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VirtualNetwork), err
}

// getPreparer prepares the Get request.
func (client VirtualNetworksClient) getPreparer(resourceGroupName string, virtualNetworkName string, expand *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	if expand != nil {
		params.Set("$expand", *expand)
	}
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client VirtualNetworksClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &VirtualNetwork{rawResponse: resp.Response()}
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

// List the list VirtualNetwork returns all Virtual Networks in a resource group
//
// resourceGroupName is the name of the resource group.
func (client VirtualNetworksClient) List(ctx context.Context, resourceGroupName string) (*VirtualNetworkListResult, error) {
	req, err := client.listPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VirtualNetworkListResult), err
}

// listPreparer prepares the List request.
func (client VirtualNetworksClient) listPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks"
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
func (client VirtualNetworksClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &VirtualNetworkListResult{rawResponse: resp.Response()}
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

// ListAll the list VirtualNetwork returns all Virtual Networks in a subscription
func (client VirtualNetworksClient) ListAll(ctx context.Context) (*VirtualNetworkListResult, error) {
	req, err := client.listAllPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listAllResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VirtualNetworkListResult), err
}

// listAllPreparer prepares the ListAll request.
func (client VirtualNetworksClient) listAllPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworks"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listAllResponder handles the response to the ListAll request.
func (client VirtualNetworksClient) listAllResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &VirtualNetworkListResult{rawResponse: resp.Response()}
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
