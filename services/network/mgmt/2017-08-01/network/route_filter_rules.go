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

// RouteFilterRulesClient is the network Client
type RouteFilterRulesClient struct {
	ManagementClient
}

// NewRouteFilterRulesClient creates an instance of the RouteFilterRulesClient client.
func NewRouteFilterRulesClient(p pipeline.Pipeline) RouteFilterRulesClient {
	return RouteFilterRulesClient{NewManagementClient(p)}
}

// CreateOrUpdate creates or updates a route in the specified route filter. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. routeFilterName is the name of the route filter. ruleName is
// the name of the route filter rule. routeFilterRuleParameters is parameters supplied to the create or update route
// filter rule operation.
func (client RouteFilterRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters RouteFilterRule) (*RouteFilterRule, error) {
	if err := validate([]validation{
		{targetValue: routeFilterRuleParameters,
			constraints: []constraint{{target: "routeFilterRuleParameters.RouteFilterRulePropertiesFormat", name: null, rule: false,
				chain: []constraint{{target: "routeFilterRuleParameters.RouteFilterRulePropertiesFormat.RouteFilterRuleType", name: null, rule: true, chain: nil},
					{target: "routeFilterRuleParameters.RouteFilterRulePropertiesFormat.Communities", name: null, rule: true, chain: nil},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(resourceGroupName, routeFilterName, ruleName, routeFilterRuleParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RouteFilterRule), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RouteFilterRulesClient) createOrUpdatePreparer(resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters RouteFilterRule) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-08-01")
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(routeFilterRuleParameters)
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
func (client RouteFilterRulesClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	result := &RouteFilterRule{rawResponse: resp.Response()}
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

// Delete deletes the specified rule from a route filter. This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. routeFilterName is the name of the route filter. ruleName is
// the name of the rule.
func (client RouteFilterRulesClient) Delete(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, routeFilterName, ruleName)
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
func (client RouteFilterRulesClient) deletePreparer(resourceGroupName string, routeFilterName string, ruleName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}"
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
func (client RouteFilterRulesClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusAccepted, http.StatusOK, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// Get gets the specified rule from a route filter.
//
// resourceGroupName is the name of the resource group. routeFilterName is the name of the route filter. ruleName is
// the name of the rule.
func (client RouteFilterRulesClient) Get(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (*RouteFilterRule, error) {
	req, err := client.getPreparer(resourceGroupName, routeFilterName, ruleName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RouteFilterRule), err
}

// getPreparer prepares the Get request.
func (client RouteFilterRulesClient) getPreparer(resourceGroupName string, routeFilterName string, ruleName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}"
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
func (client RouteFilterRulesClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RouteFilterRule{rawResponse: resp.Response()}
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

// ListByRouteFilter gets all RouteFilterRules in a route filter.
//
// resourceGroupName is the name of the resource group. routeFilterName is the name of the route filter.
func (client RouteFilterRulesClient) ListByRouteFilter(ctx context.Context, resourceGroupName string, routeFilterName string) (*RouteFilterRuleListResult, error) {
	req, err := client.listByRouteFilterPreparer(resourceGroupName, routeFilterName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByRouteFilterResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RouteFilterRuleListResult), err
}

// listByRouteFilterPreparer prepares the ListByRouteFilter request.
func (client RouteFilterRulesClient) listByRouteFilterPreparer(resourceGroupName string, routeFilterName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-08-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listByRouteFilterResponder handles the response to the ListByRouteFilter request.
func (client RouteFilterRulesClient) listByRouteFilterResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RouteFilterRuleListResult{rawResponse: resp.Response()}
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

// Update updates a route in the specified route filter. This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. routeFilterName is the name of the route filter. ruleName is
// the name of the route filter rule. routeFilterRuleParameters is parameters supplied to the update route filter rule
// operation.
func (client RouteFilterRulesClient) Update(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters PatchRouteFilterRule) (*RouteFilterRule, error) {
	req, err := client.updatePreparer(resourceGroupName, routeFilterName, ruleName, routeFilterRuleParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RouteFilterRule), err
}

// updatePreparer prepares the Update request.
func (client RouteFilterRulesClient) updatePreparer(resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters PatchRouteFilterRule) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}"
	req, err := pipeline.NewRequest("PATCH", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-08-01")
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(routeFilterRuleParameters)
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
func (client RouteFilterRulesClient) updateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RouteFilterRule{rawResponse: resp.Response()}
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
