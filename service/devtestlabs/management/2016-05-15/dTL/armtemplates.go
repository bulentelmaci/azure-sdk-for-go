package dtl

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ArmTemplatesClient is the the DevTest Labs Client.
type ArmTemplatesClient struct {
	ManagementClient
}

// NewArmTemplatesClient creates an instance of the ArmTemplatesClient client.
func NewArmTemplatesClient(subscriptionID string) ArmTemplatesClient {
	return NewArmTemplatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewArmTemplatesClientWithBaseURI creates an instance of the
// ArmTemplatesClient client.
func NewArmTemplatesClientWithBaseURI(baseURI string, subscriptionID string) ArmTemplatesClient {
	return ArmTemplatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get azure resource manager template.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. artifactSourceName is the name of the artifact source. name is the
// name of the azure Resource Manager template. expand is specify the $expand
// query. Example: 'properties($select=displayName)'
func (client ArmTemplatesClient) Get(resourceGroupName string, labName string, artifactSourceName string, name string, expand string) (result ArmTemplate, err error) {
	req, err := client.GetPreparer(resourceGroupName, labName, artifactSourceName, name, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ArmTemplatesClient) GetPreparer(resourceGroupName string, labName string, artifactSourceName string, name string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"artifactSourceName": autorest.Encode("path", artifactSourceName),
		"labName":            autorest.Encode("path", labName),
		"name":               autorest.Encode("path", name),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ArmTemplatesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ArmTemplatesClient) GetResponder(resp *http.Response) (result ArmTemplate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list azure resource manager templates in a given artifact source.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. artifactSourceName is the name of the artifact source. expand is
// specify the $expand query. Example: 'properties($select=displayName)' filter
// is the filter to apply to the operation. top is the maximum number of
// resources to return from the operation. orderby is the ordering expression
// for the results, using OData notation.
func (client ArmTemplatesClient) List(resourceGroupName string, labName string, artifactSourceName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationArmTemplate, err error) {
	req, err := client.ListPreparer(resourceGroupName, labName, artifactSourceName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ArmTemplatesClient) ListPreparer(resourceGroupName string, labName string, artifactSourceName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"artifactSourceName": autorest.Encode("path", artifactSourceName),
		"labName":            autorest.Encode("path", labName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ArmTemplatesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ArmTemplatesClient) ListResponder(resp *http.Response) (result ResponseWithContinuationArmTemplate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ArmTemplatesClient) ListNextResults(lastResults ResponseWithContinuationArmTemplate) (result ResponseWithContinuationArmTemplate, err error) {
	req, err := lastResults.ResponseWithContinuationArmTemplatePreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArmTemplatesClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client ArmTemplatesClient) ListComplete(resourceGroupName string, labName string, artifactSourceName string, expand string, filter string, top *int32, orderby string, cancel <-chan struct{}) (<-chan ArmTemplate, <-chan error) {
	resultChan := make(chan ArmTemplate)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceGroupName, labName, artifactSourceName, expand, filter, top, orderby)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}