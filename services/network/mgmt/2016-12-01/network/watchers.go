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

// WatchersClient is the network Client
type WatchersClient struct {
	ManagementClient
}

// NewWatchersClient creates an instance of the WatchersClient client.
func NewWatchersClient(p pipeline.Pipeline) WatchersClient {
	return WatchersClient{NewManagementClient(p)}
}

// CreateOrUpdate creates or updates a network watcher in the specified resource group.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the network watcher.
// parameters is parameters that define the network watcher resource.
func (client WatchersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters Watcher) (*Watcher, error) {
	req, err := client.createOrUpdatePreparer(resourceGroupName, networkWatcherName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Watcher), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WatchersClient) createOrUpdatePreparer(resourceGroupName string, networkWatcherName string, parameters Watcher) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}"
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
func (client WatchersClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	result := &Watcher{rawResponse: resp.Response()}
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

// Delete deletes the specified network watcher resource. This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the network watcher.
func (client WatchersClient) Delete(ctx context.Context, resourceGroupName string, networkWatcherName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, networkWatcherName)
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
func (client WatchersClient) deletePreparer(resourceGroupName string, networkWatcherName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}"
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
func (client WatchersClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// Get gets the specified network watcher by resource group.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the network watcher.
func (client WatchersClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string) (*Watcher, error) {
	req, err := client.getPreparer(resourceGroupName, networkWatcherName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Watcher), err
}

// getPreparer prepares the Get request.
func (client WatchersClient) getPreparer(resourceGroupName string, networkWatcherName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}"
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
func (client WatchersClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Watcher{rawResponse: resp.Response()}
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

// GetFlowLogStatus queries status of flow log on a specified resource. This method may poll for completion. Polling
// can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the network watcher resource group. networkWatcherName is the name of the network
// watcher resource. parameters is parameters that define a resource to query flow log status.
func (client WatchersClient) GetFlowLogStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters FlowLogStatusParameters) (*FlowLogInformation, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.TargetResourceID", name: null, rule: true, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.getFlowLogStatusPreparer(resourceGroupName, networkWatcherName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getFlowLogStatusResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FlowLogInformation), err
}

// getFlowLogStatusPreparer prepares the GetFlowLogStatus request.
func (client WatchersClient) getFlowLogStatusPreparer(resourceGroupName string, networkWatcherName string, parameters FlowLogStatusParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/queryFlowLogStatus"
	req, err := pipeline.NewRequest("POST", u, nil)
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

// getFlowLogStatusResponder handles the response to the GetFlowLogStatus request.
func (client WatchersClient) getFlowLogStatusResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &FlowLogInformation{rawResponse: resp.Response()}
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

// GetNextHop gets the next hop from the specified VM. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the network watcher.
// parameters is parameters that define the source and destination endpoint.
func (client WatchersClient) GetNextHop(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters NextHopParameters) (*NextHopResult, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.TargetResourceID", name: null, rule: true, chain: nil},
				{target: "parameters.SourceIPAddress", name: null, rule: true, chain: nil},
				{target: "parameters.DestinationIPAddress", name: null, rule: true, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.getNextHopPreparer(resourceGroupName, networkWatcherName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getNextHopResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*NextHopResult), err
}

// getNextHopPreparer prepares the GetNextHop request.
func (client WatchersClient) getNextHopPreparer(resourceGroupName string, networkWatcherName string, parameters NextHopParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/nextHop"
	req, err := pipeline.NewRequest("POST", u, nil)
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

// getNextHopResponder handles the response to the GetNextHop request.
func (client WatchersClient) getNextHopResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &NextHopResult{rawResponse: resp.Response()}
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

// GetTopology gets the current network topology by resource group.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the network watcher.
// parameters is parameters that define the representation of topology.
func (client WatchersClient) GetTopology(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters TopologyParameters) (*Topology, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.TargetResourceGroupName", name: null, rule: true, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.getTopologyPreparer(resourceGroupName, networkWatcherName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getTopologyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Topology), err
}

// getTopologyPreparer prepares the GetTopology request.
func (client WatchersClient) getTopologyPreparer(resourceGroupName string, networkWatcherName string, parameters TopologyParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/topology"
	req, err := pipeline.NewRequest("POST", u, nil)
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

// getTopologyResponder handles the response to the GetTopology request.
func (client WatchersClient) getTopologyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Topology{rawResponse: resp.Response()}
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

// GetTroubleshooting initiate troubleshooting on a specified resource This method may poll for completion. Polling can
// be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the network watcher resource.
// parameters is parameters that define the resource to troubleshoot.
func (client WatchersClient) GetTroubleshooting(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters TroubleshootingParameters) (*TroubleshootingResult, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.TargetResourceID", name: null, rule: true, chain: nil},
				{target: "parameters.TroubleshootingProperties", name: null, rule: true,
					chain: []constraint{{target: "parameters.TroubleshootingProperties.StorageID", name: null, rule: true, chain: nil},
						{target: "parameters.TroubleshootingProperties.StoragePath", name: null, rule: true, chain: nil},
					}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getTroubleshootingPreparer(resourceGroupName, networkWatcherName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getTroubleshootingResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*TroubleshootingResult), err
}

// getTroubleshootingPreparer prepares the GetTroubleshooting request.
func (client WatchersClient) getTroubleshootingPreparer(resourceGroupName string, networkWatcherName string, parameters TroubleshootingParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/troubleshoot"
	req, err := pipeline.NewRequest("POST", u, nil)
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

// getTroubleshootingResponder handles the response to the GetTroubleshooting request.
func (client WatchersClient) getTroubleshootingResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &TroubleshootingResult{rawResponse: resp.Response()}
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

// GetTroubleshootingResult get the last completed troubleshooting result on a specified resource This method may poll
// for completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the network watcher resource.
// parameters is parameters that define the resource to query the troubleshooting result.
func (client WatchersClient) GetTroubleshootingResult(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters QueryTroubleshootingParameters) (*TroubleshootingResult, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.TargetResourceID", name: null, rule: true, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.getTroubleshootingResultPreparer(resourceGroupName, networkWatcherName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getTroubleshootingResultResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*TroubleshootingResult), err
}

// getTroubleshootingResultPreparer prepares the GetTroubleshootingResult request.
func (client WatchersClient) getTroubleshootingResultPreparer(resourceGroupName string, networkWatcherName string, parameters QueryTroubleshootingParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/queryTroubleshootResult"
	req, err := pipeline.NewRequest("POST", u, nil)
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

// getTroubleshootingResultResponder handles the response to the GetTroubleshootingResult request.
func (client WatchersClient) getTroubleshootingResultResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &TroubleshootingResult{rawResponse: resp.Response()}
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

// GetVMSecurityRules gets the configured and effective security group rules on the specified VM. This method may poll
// for completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the network watcher.
// parameters is parameters that define the VM to check security groups for.
func (client WatchersClient) GetVMSecurityRules(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters SecurityGroupViewParameters) (*SecurityGroupViewResult, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.TargetResourceID", name: null, rule: true, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.getVMSecurityRulesPreparer(resourceGroupName, networkWatcherName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getVMSecurityRulesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*SecurityGroupViewResult), err
}

// getVMSecurityRulesPreparer prepares the GetVMSecurityRules request.
func (client WatchersClient) getVMSecurityRulesPreparer(resourceGroupName string, networkWatcherName string, parameters SecurityGroupViewParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/securityGroupView"
	req, err := pipeline.NewRequest("POST", u, nil)
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

// getVMSecurityRulesResponder handles the response to the GetVMSecurityRules request.
func (client WatchersClient) getVMSecurityRulesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &SecurityGroupViewResult{rawResponse: resp.Response()}
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

// List gets all network watchers by resource group.
//
// resourceGroupName is the name of the resource group.
func (client WatchersClient) List(ctx context.Context, resourceGroupName string) (*WatcherListResult, error) {
	req, err := client.listPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*WatcherListResult), err
}

// listPreparer prepares the List request.
func (client WatchersClient) listPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers"
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
func (client WatchersClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &WatcherListResult{rawResponse: resp.Response()}
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

// ListAll gets all network watchers by subscription.
func (client WatchersClient) ListAll(ctx context.Context) (*WatcherListResult, error) {
	req, err := client.listAllPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listAllResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*WatcherListResult), err
}

// listAllPreparer prepares the ListAll request.
func (client WatchersClient) listAllPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkWatchers"
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
func (client WatchersClient) listAllResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &WatcherListResult{rawResponse: resp.Response()}
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

// SetFlowLogConfiguration configures flow log on a specified resource. This method may poll for completion. Polling
// can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the network watcher resource group. networkWatcherName is the name of the network
// watcher resource. parameters is parameters that define the configuration of flow log.
func (client WatchersClient) SetFlowLogConfiguration(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters FlowLogInformation) (*FlowLogInformation, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.TargetResourceID", name: null, rule: true, chain: nil},
				{target: "parameters.FlowLogProperties", name: null, rule: true,
					chain: []constraint{{target: "parameters.FlowLogProperties.StorageID", name: null, rule: true, chain: nil},
						{target: "parameters.FlowLogProperties.Enabled", name: null, rule: true, chain: nil},
					}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setFlowLogConfigurationPreparer(resourceGroupName, networkWatcherName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setFlowLogConfigurationResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*FlowLogInformation), err
}

// setFlowLogConfigurationPreparer prepares the SetFlowLogConfiguration request.
func (client WatchersClient) setFlowLogConfigurationPreparer(resourceGroupName string, networkWatcherName string, parameters FlowLogInformation) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/configureFlowLog"
	req, err := pipeline.NewRequest("POST", u, nil)
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

// setFlowLogConfigurationResponder handles the response to the SetFlowLogConfiguration request.
func (client WatchersClient) setFlowLogConfigurationResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &FlowLogInformation{rawResponse: resp.Response()}
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

// VerifyIPFlow verify IP flow from the specified VM to a location given the currently configured NSG rules. This
// method may poll for completion. Polling can be canceled by passing the cancel channel argument. The channel will be
// used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the network watcher.
// parameters is parameters that define the IP flow to be verified.
func (client WatchersClient) VerifyIPFlow(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters VerificationIPFlowParameters) (*VerificationIPFlowResult, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.TargetResourceID", name: null, rule: true, chain: nil},
				{target: "parameters.LocalPort", name: null, rule: true, chain: nil},
				{target: "parameters.RemotePort", name: null, rule: true, chain: nil},
				{target: "parameters.LocalIPAddress", name: null, rule: true, chain: nil},
				{target: "parameters.RemoteIPAddress", name: null, rule: true, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.verifyIPFlowPreparer(resourceGroupName, networkWatcherName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.verifyIPFlowResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VerificationIPFlowResult), err
}

// verifyIPFlowPreparer prepares the VerifyIPFlow request.
func (client WatchersClient) verifyIPFlowPreparer(resourceGroupName string, networkWatcherName string, parameters VerificationIPFlowParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/ipFlowVerify"
	req, err := pipeline.NewRequest("POST", u, nil)
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

// verifyIPFlowResponder handles the response to the VerifyIPFlow request.
func (client WatchersClient) verifyIPFlowResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &VerificationIPFlowResult{rawResponse: resp.Response()}
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
