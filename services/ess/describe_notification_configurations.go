package ess

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeNotificationConfigurations invokes the ess.DescribeNotificationConfigurations API synchronously
// api document: https://help.aliyun.com/api/ess/describenotificationconfigurations.html
func (client *Client) DescribeNotificationConfigurations(request *DescribeNotificationConfigurationsRequest) (response *DescribeNotificationConfigurationsResponse, err error) {
	response = CreateDescribeNotificationConfigurationsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNotificationConfigurationsWithChan invokes the ess.DescribeNotificationConfigurations API asynchronously
// api document: https://help.aliyun.com/api/ess/describenotificationconfigurations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNotificationConfigurationsWithChan(request *DescribeNotificationConfigurationsRequest) (<-chan *DescribeNotificationConfigurationsResponse, <-chan error) {
	responseChan := make(chan *DescribeNotificationConfigurationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNotificationConfigurations(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeNotificationConfigurationsWithCallback invokes the ess.DescribeNotificationConfigurations API asynchronously
// api document: https://help.aliyun.com/api/ess/describenotificationconfigurations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNotificationConfigurationsWithCallback(request *DescribeNotificationConfigurationsRequest, callback func(response *DescribeNotificationConfigurationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNotificationConfigurationsResponse
		var err error
		defer close(result)
		response, err = client.DescribeNotificationConfigurations(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeNotificationConfigurationsRequest is the request struct for api DescribeNotificationConfigurations
type DescribeNotificationConfigurationsRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
}

// DescribeNotificationConfigurationsResponse is the response struct for api DescribeNotificationConfigurations
type DescribeNotificationConfigurationsResponse struct {
	*responses.BaseResponse
	RequestId                       string                          `json:"RequestId" xml:"RequestId"`
	NotificationConfigurationModels NotificationConfigurationModels `json:"NotificationConfigurationModels" xml:"NotificationConfigurationModels"`
}

// CreateDescribeNotificationConfigurationsRequest creates a request to invoke DescribeNotificationConfigurations API
func CreateDescribeNotificationConfigurationsRequest() (request *DescribeNotificationConfigurationsRequest) {
	request = &DescribeNotificationConfigurationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeNotificationConfigurations", "ess", "openAPI")
	return
}

// CreateDescribeNotificationConfigurationsResponse creates a response to parse from DescribeNotificationConfigurations response
func CreateDescribeNotificationConfigurationsResponse() (response *DescribeNotificationConfigurationsResponse) {
	response = &DescribeNotificationConfigurationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
