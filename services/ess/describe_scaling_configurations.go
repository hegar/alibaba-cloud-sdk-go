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

// DescribeScalingConfigurations invokes the ess.DescribeScalingConfigurations API synchronously
// api document: https://help.aliyun.com/api/ess/describescalingconfigurations.html
func (client *Client) DescribeScalingConfigurations(request *DescribeScalingConfigurationsRequest) (response *DescribeScalingConfigurationsResponse, err error) {
	response = CreateDescribeScalingConfigurationsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScalingConfigurationsWithChan invokes the ess.DescribeScalingConfigurations API asynchronously
// api document: https://help.aliyun.com/api/ess/describescalingconfigurations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScalingConfigurationsWithChan(request *DescribeScalingConfigurationsRequest) (<-chan *DescribeScalingConfigurationsResponse, <-chan error) {
	responseChan := make(chan *DescribeScalingConfigurationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScalingConfigurations(request)
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

// DescribeScalingConfigurationsWithCallback invokes the ess.DescribeScalingConfigurations API asynchronously
// api document: https://help.aliyun.com/api/ess/describescalingconfigurations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScalingConfigurationsWithCallback(request *DescribeScalingConfigurationsRequest, callback func(response *DescribeScalingConfigurationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScalingConfigurationsResponse
		var err error
		defer close(result)
		response, err = client.DescribeScalingConfigurations(request)
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

// DescribeScalingConfigurationsRequest is the request struct for api DescribeScalingConfigurations
type DescribeScalingConfigurationsRequest struct {
	*requests.RpcRequest
	OwnerId                    requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount       string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber                 requests.Integer `position:"Query" name:"PageNumber"`
	PageSize                   requests.Integer `position:"Query" name:"PageSize"`
	ScalingGroupId             string           `position:"Query" name:"ScalingGroupId"`
	ScalingConfigurationId1    string           `position:"Query" name:"ScalingConfigurationId.1"`
	ScalingConfigurationId2    string           `position:"Query" name:"ScalingConfigurationId.2"`
	ScalingConfigurationId3    string           `position:"Query" name:"ScalingConfigurationId.3"`
	ScalingConfigurationId4    string           `position:"Query" name:"ScalingConfigurationId.4"`
	ScalingConfigurationId5    string           `position:"Query" name:"ScalingConfigurationId.5"`
	ScalingConfigurationId6    string           `position:"Query" name:"ScalingConfigurationId.6"`
	ScalingConfigurationId7    string           `position:"Query" name:"ScalingConfigurationId.7"`
	ScalingConfigurationId8    string           `position:"Query" name:"ScalingConfigurationId.8"`
	ScalingConfigurationId9    string           `position:"Query" name:"ScalingConfigurationId.9"`
	ScalingConfigurationId10   string           `position:"Query" name:"ScalingConfigurationId.10"`
	ScalingConfigurationName1  string           `position:"Query" name:"ScalingConfigurationName.1"`
	ScalingConfigurationName2  string           `position:"Query" name:"ScalingConfigurationName.2"`
	ScalingConfigurationName3  string           `position:"Query" name:"ScalingConfigurationName.3"`
	ScalingConfigurationName4  string           `position:"Query" name:"ScalingConfigurationName.4"`
	ScalingConfigurationName5  string           `position:"Query" name:"ScalingConfigurationName.5"`
	ScalingConfigurationName6  string           `position:"Query" name:"ScalingConfigurationName.6"`
	ScalingConfigurationName7  string           `position:"Query" name:"ScalingConfigurationName.7"`
	ScalingConfigurationName8  string           `position:"Query" name:"ScalingConfigurationName.8"`
	ScalingConfigurationName9  string           `position:"Query" name:"ScalingConfigurationName.9"`
	ScalingConfigurationName10 string           `position:"Query" name:"ScalingConfigurationName.10"`
	OwnerAccount               string           `position:"Query" name:"OwnerAccount"`
}

// DescribeScalingConfigurationsResponse is the response struct for api DescribeScalingConfigurations
type DescribeScalingConfigurationsResponse struct {
	*responses.BaseResponse
	TotalCount            int                   `json:"TotalCount" xml:"TotalCount"`
	PageNumber            int                   `json:"PageNumber" xml:"PageNumber"`
	PageSize              int                   `json:"PageSize" xml:"PageSize"`
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	ScalingConfigurations ScalingConfigurations `json:"ScalingConfigurations" xml:"ScalingConfigurations"`
}

// CreateDescribeScalingConfigurationsRequest creates a request to invoke DescribeScalingConfigurations API
func CreateDescribeScalingConfigurationsRequest() (request *DescribeScalingConfigurationsRequest) {
	request = &DescribeScalingConfigurationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingConfigurations", "ess", "openAPI")
	return
}

// CreateDescribeScalingConfigurationsResponse creates a response to parse from DescribeScalingConfigurations response
func CreateDescribeScalingConfigurationsResponse() (response *DescribeScalingConfigurationsResponse) {
	response = &DescribeScalingConfigurationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
