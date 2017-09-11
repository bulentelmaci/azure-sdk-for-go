package datamasking

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// RulesClient is the provides create, read, update and delete functionality
// for Azure SQL Database datamasking policies and rules.
type RulesClient struct {
	ManagementClient
}

// NewRulesClient creates an instance of the RulesClient client.
func NewRulesClient(subscriptionID string) RulesClient {
	return NewRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRulesClientWithBaseURI creates an instance of the RulesClient client.
func NewRulesClientWithBaseURI(baseURI string, subscriptionID string) RulesClient {
	return RulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a database data masking rule.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. databaseName is the name
// of the database. dataMaskingPolicyName is the name of the database for which
// the data masking rule applies. dataMaskingRuleName is the name of the data
// masking rule. parameters is the required parameters for creating or updating
// a data masking rule.
func (client RulesClient) CreateOrUpdate(resourceGroupName string, serverName string, databaseName string, dataMaskingPolicyName string, dataMaskingRuleName string, parameters Rule) (result Rule, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.RuleProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.RuleProperties.SchemaName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.RuleProperties.TableName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.RuleProperties.ColumnName", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "datamasking.RulesClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, serverName, databaseName, dataMaskingPolicyName, dataMaskingRuleName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamasking.RulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamasking.RulesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamasking.RulesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RulesClient) CreateOrUpdatePreparer(resourceGroupName string, serverName string, databaseName string, dataMaskingPolicyName string, dataMaskingRuleName string, parameters Rule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":          autorest.Encode("path", databaseName),
		"dataMaskingPolicyName": autorest.Encode("path", dataMaskingPolicyName),
		"dataMaskingRuleName":   autorest.Encode("path", dataMaskingRuleName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"serverName":            autorest.Encode("path", serverName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataMaskingPolicies/{dataMaskingPolicyName}/rules/{dataMaskingRuleName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RulesClient) CreateOrUpdateResponder(resp *http.Response) (result Rule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDatabase gets a list of database data masking rules.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. databaseName is the name
// of the database. dataMaskingPolicyName is the name of the database for which
// the data masking rule applies.
func (client RulesClient) ListByDatabase(resourceGroupName string, serverName string, databaseName string, dataMaskingPolicyName string) (result RuleListResult, err error) {
	req, err := client.ListByDatabasePreparer(resourceGroupName, serverName, databaseName, dataMaskingPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamasking.RulesClient", "ListByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamasking.RulesClient", "ListByDatabase", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamasking.RulesClient", "ListByDatabase", resp, "Failure responding to request")
	}

	return
}

// ListByDatabasePreparer prepares the ListByDatabase request.
func (client RulesClient) ListByDatabasePreparer(resourceGroupName string, serverName string, databaseName string, dataMaskingPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":          autorest.Encode("path", databaseName),
		"dataMaskingPolicyName": autorest.Encode("path", dataMaskingPolicyName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"serverName":            autorest.Encode("path", serverName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataMaskingPolicies/{dataMaskingPolicyName}/rules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByDatabaseSender sends the ListByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client RulesClient) ListByDatabaseSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByDatabaseResponder handles the response to the ListByDatabase request. The method always
// closes the http.Response Body.
func (client RulesClient) ListByDatabaseResponder(resp *http.Response) (result RuleListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}