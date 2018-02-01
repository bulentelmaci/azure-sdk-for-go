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

// RouteFiltersClient is the network Client
type RouteFiltersClient struct {
	ManagementClient
}

// NewRouteFiltersClient creates an instance of the RouteFiltersClient client.
func NewRouteFiltersClient(p pipeline.Pipeline) RouteFiltersClient {
	return RouteFiltersClient{NewManagementClient(p)}
}

// CreateOrUpdate creates or updates a route filter in a specified resource group. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. routeFilterName is the name of the route filter.
// routeFilterParameters is parameters supplied to the create or update route filter operation.
func (client RouteFiltersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter) (*RouteFilter, error) {
	req, err := client.createOrUpdatePreparer(resourceGroupName, routeFilterName, routeFilterParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RouteFilter), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RouteFiltersClient) createOrUpdatePreparer(resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-06-01")
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(routeFilterParameters)
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
func (client RouteFiltersClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	result := &RouteFilter{rawResponse: resp.Response()}
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

// Delete deletes the specified route filter. This method may poll for completion. Polling can be canceled by passing
// the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. routeFilterName is the name of the route filter.
func (client RouteFiltersClient) Delete(ctx context.Context, resourceGroupName string, routeFilterName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, routeFilterName)
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
func (client RouteFiltersClient) deletePreparer(resourceGroupName string, routeFilterName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-06-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client RouteFiltersClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusAccepted, http.StatusOK, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// Get gets the specified route filter.
//
// resourceGroupName is the name of the resource group. routeFilterName is the name of the route filter. expand is
// expands referenced express route bgp peering resources.
func (client RouteFiltersClient) Get(ctx context.Context, resourceGroupName string, routeFilterName string, expand *string) (*RouteFilter, error) {
	req, err := client.getPreparer(resourceGroupName, routeFilterName, expand)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RouteFilter), err
}

// getPreparer prepares the Get request.
func (client RouteFiltersClient) getPreparer(resourceGroupName string, routeFilterName string, expand *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-06-01")
	if expand != nil {
		params.Set("$expand", *expand)
	}
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client RouteFiltersClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RouteFilter{rawResponse: resp.Response()}
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

// List gets all route filters in a subscription.
func (client RouteFiltersClient) List(ctx context.Context) (*RouteFilterListResult, error) {
	req, err := client.listPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RouteFilterListResult), err
}

// listPreparer prepares the List request.
func (client RouteFiltersClient) listPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeFilters"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-06-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client RouteFiltersClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RouteFilterListResult{rawResponse: resp.Response()}
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

// ListByResourceGroup gets all route filters in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client RouteFiltersClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (*RouteFilterListResult, error) {
	req, err := client.listByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByResourceGroupResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RouteFilterListResult), err
}

// listByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client RouteFiltersClient) listByResourceGroupPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-06-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listByResourceGroupResponder handles the response to the ListByResourceGroup request.
func (client RouteFiltersClient) listByResourceGroupResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RouteFilterListResult{rawResponse: resp.Response()}
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

// Update updates a route filter in a specified resource group. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. routeFilterName is the name of the route filter.
// routeFilterParameters is parameters supplied to the update route filter operation.
func (client RouteFiltersClient) Update(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters PatchRouteFilter) (*RouteFilter, error) {
	req, err := client.updatePreparer(resourceGroupName, routeFilterName, routeFilterParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RouteFilter), err
}

// updatePreparer prepares the Update request.
func (client RouteFiltersClient) updatePreparer(resourceGroupName string, routeFilterName string, routeFilterParameters PatchRouteFilter) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	req, err := pipeline.NewRequest("PATCH", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-06-01")
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(routeFilterParameters)
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
func (client RouteFiltersClient) updateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RouteFilter{rawResponse: resp.Response()}
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
