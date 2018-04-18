package storsimple

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// VolumeContainersClient is the client for the VolumeContainers methods of the Storsimple service.
type VolumeContainersClient struct {
	BaseClient
}

// NewVolumeContainersClient creates an instance of the VolumeContainersClient client.
func NewVolumeContainersClient(subscriptionID string) VolumeContainersClient {
	return NewVolumeContainersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVolumeContainersClientWithBaseURI creates an instance of the VolumeContainersClient client.
func NewVolumeContainersClientWithBaseURI(baseURI string, subscriptionID string) VolumeContainersClient {
	return VolumeContainersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the volume container.
//
// deviceName is the device name volumeContainerName is the name of the volume container. parameters is the volume
// container to be added or updated. resourceGroupName is the resource group name managerName is the manager name
func (client VolumeContainersClient) CreateOrUpdate(ctx context.Context, deviceName string, volumeContainerName string, parameters VolumeContainer, resourceGroupName string, managerName string) (result VolumeContainersCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.VolumeContainerProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.VolumeContainerProperties.EncryptionKey", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.VolumeContainerProperties.EncryptionKey.Value", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.VolumeContainerProperties.StorageAccountCredentialID", Name: validation.Null, Rule: true, Chain: nil},
				}}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple.VolumeContainersClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(ctx, deviceName, volumeContainerName, parameters, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VolumeContainersClient) CreateOrUpdatePreparer(ctx context.Context, deviceName string, volumeContainerName string, parameters VolumeContainer, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":          deviceName,
		"managerName":         managerName,
		"resourceGroupName":   resourceGroupName,
		"subscriptionId":      client.SubscriptionID,
		"volumeContainerName": volumeContainerName,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VolumeContainersClient) CreateOrUpdateSender(req *http.Request) (future VolumeContainersCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VolumeContainersClient) CreateOrUpdateResponder(resp *http.Response) (result VolumeContainer, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the volume container.
//
// deviceName is the device name volumeContainerName is the name of the volume container. resourceGroupName is the
// resource group name managerName is the manager name
func (client VolumeContainersClient) Delete(ctx context.Context, deviceName string, volumeContainerName string, resourceGroupName string, managerName string) (result VolumeContainersDeleteFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple.VolumeContainersClient", "Delete")
	}

	req, err := client.DeletePreparer(ctx, deviceName, volumeContainerName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VolumeContainersClient) DeletePreparer(ctx context.Context, deviceName string, volumeContainerName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":          deviceName,
		"managerName":         managerName,
		"resourceGroupName":   resourceGroupName,
		"subscriptionId":      client.SubscriptionID,
		"volumeContainerName": volumeContainerName,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VolumeContainersClient) DeleteSender(req *http.Request) (future VolumeContainersDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VolumeContainersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the specified volume container name.
//
// deviceName is the device name volumeContainerName is the name of the volume container. resourceGroupName is the
// resource group name managerName is the manager name
func (client VolumeContainersClient) Get(ctx context.Context, deviceName string, volumeContainerName string, resourceGroupName string, managerName string) (result VolumeContainer, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple.VolumeContainersClient", "Get")
	}

	req, err := client.GetPreparer(ctx, deviceName, volumeContainerName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VolumeContainersClient) GetPreparer(ctx context.Context, deviceName string, volumeContainerName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":          deviceName,
		"managerName":         managerName,
		"resourceGroupName":   resourceGroupName,
		"subscriptionId":      client.SubscriptionID,
		"volumeContainerName": volumeContainerName,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VolumeContainersClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VolumeContainersClient) GetResponder(resp *http.Response) (result VolumeContainer, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDevice gets all the volume containers in a device.
//
// deviceName is the device name resourceGroupName is the resource group name managerName is the manager name
func (client VolumeContainersClient) ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result VolumeContainerList, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple.VolumeContainersClient", "ListByDevice")
	}

	req, err := client.ListByDevicePreparer(ctx, deviceName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "ListByDevice", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "ListByDevice", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "ListByDevice", resp, "Failure responding to request")
	}

	return
}

// ListByDevicePreparer prepares the ListByDevice request.
func (client VolumeContainersClient) ListByDevicePreparer(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDeviceSender sends the ListByDevice request. The method will close the
// http.Response Body if it receives an error.
func (client VolumeContainersClient) ListByDeviceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByDeviceResponder handles the response to the ListByDevice request. The method always
// closes the http.Response Body.
func (client VolumeContainersClient) ListByDeviceResponder(resp *http.Response) (result VolumeContainerList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListMetricDefinition gets the metric definitions for the specified volume container.
//
// deviceName is the device name volumeContainerName is the volume container name. resourceGroupName is the resource
// group name managerName is the manager name
func (client VolumeContainersClient) ListMetricDefinition(ctx context.Context, deviceName string, volumeContainerName string, resourceGroupName string, managerName string) (result MetricDefinitionList, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple.VolumeContainersClient", "ListMetricDefinition")
	}

	req, err := client.ListMetricDefinitionPreparer(ctx, deviceName, volumeContainerName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "ListMetricDefinition", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListMetricDefinitionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "ListMetricDefinition", resp, "Failure sending request")
		return
	}

	result, err = client.ListMetricDefinitionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "ListMetricDefinition", resp, "Failure responding to request")
	}

	return
}

// ListMetricDefinitionPreparer prepares the ListMetricDefinition request.
func (client VolumeContainersClient) ListMetricDefinitionPreparer(ctx context.Context, deviceName string, volumeContainerName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":          deviceName,
		"managerName":         managerName,
		"resourceGroupName":   resourceGroupName,
		"subscriptionId":      client.SubscriptionID,
		"volumeContainerName": volumeContainerName,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/metricsDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListMetricDefinitionSender sends the ListMetricDefinition request. The method will close the
// http.Response Body if it receives an error.
func (client VolumeContainersClient) ListMetricDefinitionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListMetricDefinitionResponder handles the response to the ListMetricDefinition request. The method always
// closes the http.Response Body.
func (client VolumeContainersClient) ListMetricDefinitionResponder(resp *http.Response) (result MetricDefinitionList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListMetrics gets the metrics for the specified volume container.
//
// deviceName is the device name volumeContainerName is the volume container name. resourceGroupName is the resource
// group name managerName is the manager name filter is oData Filter options
func (client VolumeContainersClient) ListMetrics(ctx context.Context, deviceName string, volumeContainerName string, resourceGroupName string, managerName string, filter string) (result MetricList, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple.VolumeContainersClient", "ListMetrics")
	}

	req, err := client.ListMetricsPreparer(ctx, deviceName, volumeContainerName, resourceGroupName, managerName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "ListMetrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListMetricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "ListMetrics", resp, "Failure sending request")
		return
	}

	result, err = client.ListMetricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.VolumeContainersClient", "ListMetrics", resp, "Failure responding to request")
	}

	return
}

// ListMetricsPreparer prepares the ListMetrics request.
func (client VolumeContainersClient) ListMetricsPreparer(ctx context.Context, deviceName string, volumeContainerName string, resourceGroupName string, managerName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":          deviceName,
		"managerName":         managerName,
		"resourceGroupName":   resourceGroupName,
		"subscriptionId":      client.SubscriptionID,
		"volumeContainerName": volumeContainerName,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/metrics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListMetricsSender sends the ListMetrics request. The method will close the
// http.Response Body if it receives an error.
func (client VolumeContainersClient) ListMetricsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListMetricsResponder handles the response to the ListMetrics request. The method always
// closes the http.Response Body.
func (client VolumeContainersClient) ListMetricsResponder(resp *http.Response) (result MetricList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}