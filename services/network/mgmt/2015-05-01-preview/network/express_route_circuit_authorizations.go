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

// ExpressRouteCircuitAuthorizationsClient is the network Client
type ExpressRouteCircuitAuthorizationsClient struct {
	ManagementClient
}

// NewExpressRouteCircuitAuthorizationsClient creates an instance of the ExpressRouteCircuitAuthorizationsClient
// client.
func NewExpressRouteCircuitAuthorizationsClient(p pipeline.Pipeline) ExpressRouteCircuitAuthorizationsClient {
	return ExpressRouteCircuitAuthorizationsClient{NewManagementClient(p)}
}

// CreateOrUpdate the Put Authorization operation creates/updates an authorization in thespecified ExpressRouteCircuits
// This method may poll for completion. Polling can be canceled by passing the cancel channel argument. The channel
// will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the express route circuit.
// authorizationName is the name of the authorization. authorizationParameters is parameters supplied to the
// create/update ExpressRouteCircuitAuthorization operation
func (client ExpressRouteCircuitAuthorizationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization) (*ExpressRouteCircuitAuthorization, error) {
	req, err := client.createOrUpdatePreparer(resourceGroupName, circuitName, authorizationName, authorizationParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuitAuthorization), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRouteCircuitAuthorizationsClient) createOrUpdatePreparer(resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(authorizationParameters)
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
func (client ExpressRouteCircuitAuthorizationsClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusCreated, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuitAuthorization{rawResponse: resp.Response()}
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

// Delete the delete authorization operation deletes the specified authorization from the specified
// ExpressRouteCircuit. This method may poll for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the express route circuit.
// authorizationName is the name of the authorization.
func (client ExpressRouteCircuitAuthorizationsClient) Delete(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, circuitName, authorizationName)
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
func (client ExpressRouteCircuitAuthorizationsClient) deletePreparer(resourceGroupName string, circuitName string, authorizationName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"
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
func (client ExpressRouteCircuitAuthorizationsClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusNoContent, http.StatusAccepted, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// Get the GET authorization operation retrieves the specified authorization from the specified ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the express route circuit.
// authorizationName is the name of the authorization.
func (client ExpressRouteCircuitAuthorizationsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string) (*ExpressRouteCircuitAuthorization, error) {
	req, err := client.getPreparer(resourceGroupName, circuitName, authorizationName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuitAuthorization), err
}

// getPreparer prepares the Get request.
func (client ExpressRouteCircuitAuthorizationsClient) getPreparer(resourceGroupName string, circuitName string, authorizationName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"
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
func (client ExpressRouteCircuitAuthorizationsClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuitAuthorization{rawResponse: resp.Response()}
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

// List the List authorization operation retrieves all the authorizations in an ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the curcuit.
func (client ExpressRouteCircuitAuthorizationsClient) List(ctx context.Context, resourceGroupName string, circuitName string) (*AuthorizationListResult, error) {
	req, err := client.listPreparer(resourceGroupName, circuitName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*AuthorizationListResult), err
}

// listPreparer prepares the List request.
func (client ExpressRouteCircuitAuthorizationsClient) listPreparer(resourceGroupName string, circuitName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations"
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
func (client ExpressRouteCircuitAuthorizationsClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &AuthorizationListResult{rawResponse: resp.Response()}
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
