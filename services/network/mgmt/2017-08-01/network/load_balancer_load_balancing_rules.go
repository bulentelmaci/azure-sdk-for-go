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
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// LoadBalancerLoadBalancingRulesClient is the network Client
type LoadBalancerLoadBalancingRulesClient struct {
	ManagementClient
}

// NewLoadBalancerLoadBalancingRulesClient creates an instance of the LoadBalancerLoadBalancingRulesClient client.
func NewLoadBalancerLoadBalancingRulesClient(p pipeline.Pipeline) LoadBalancerLoadBalancingRulesClient {
	return LoadBalancerLoadBalancingRulesClient{NewManagementClient(p)}
}

// Get gets the specified load balancer load balancing rule.
//
// resourceGroupName is the name of the resource group. loadBalancerName is the name of the load balancer.
// loadBalancingRuleName is the name of the load balancing rule.
func (client LoadBalancerLoadBalancingRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string) (*LoadBalancingRule, error) {
	req, err := client.getPreparer(resourceGroupName, loadBalancerName, loadBalancingRuleName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*LoadBalancingRule), err
}

// getPreparer prepares the Get request.
func (client LoadBalancerLoadBalancingRulesClient) getPreparer(resourceGroupName string, loadBalancerName string, loadBalancingRuleName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules/{loadBalancingRuleName}"
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
func (client LoadBalancerLoadBalancingRulesClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &LoadBalancingRule{rawResponse: resp.Response()}
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

// List gets all the load balancing rules in a load balancer.
//
// resourceGroupName is the name of the resource group. loadBalancerName is the name of the load balancer.
func (client LoadBalancerLoadBalancingRulesClient) List(ctx context.Context, resourceGroupName string, loadBalancerName string) (*LoadBalancerLoadBalancingRuleListResult, error) {
	req, err := client.listPreparer(resourceGroupName, loadBalancerName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*LoadBalancerLoadBalancingRuleListResult), err
}

// listPreparer prepares the List request.
func (client LoadBalancerLoadBalancingRulesClient) listPreparer(resourceGroupName string, loadBalancerName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules"
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
func (client LoadBalancerLoadBalancingRulesClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &LoadBalancerLoadBalancingRuleListResult{rawResponse: resp.Response()}
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
