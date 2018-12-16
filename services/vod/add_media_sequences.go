package vod

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

// AddMediaSequences invokes the vod.AddMediaSequences API synchronously
// api document: https://help.aliyun.com/api/vod/addmediasequences.html
func (client *Client) AddMediaSequences(request *AddMediaSequencesRequest) (response *AddMediaSequencesResponse, err error) {
	response = CreateAddMediaSequencesResponse()
	err = client.DoAction(request, response)
	return
}

// AddMediaSequencesWithChan invokes the vod.AddMediaSequences API asynchronously
// api document: https://help.aliyun.com/api/vod/addmediasequences.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddMediaSequencesWithChan(request *AddMediaSequencesRequest) (<-chan *AddMediaSequencesResponse, <-chan error) {
	responseChan := make(chan *AddMediaSequencesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddMediaSequences(request)
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

// AddMediaSequencesWithCallback invokes the vod.AddMediaSequences API asynchronously
// api document: https://help.aliyun.com/api/vod/addmediasequences.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddMediaSequencesWithCallback(request *AddMediaSequencesRequest, callback func(response *AddMediaSequencesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddMediaSequencesResponse
		var err error
		defer close(result)
		response, err = client.AddMediaSequences(request)
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

// AddMediaSequencesRequest is the request struct for api AddMediaSequences
type AddMediaSequencesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	MediaURL             string `position:"Query" name:"MediaURL"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaSequences       string `position:"Query" name:"MediaSequences"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	MediaType            string `position:"Query" name:"MediaType"`
}

// AddMediaSequencesResponse is the response struct for api AddMediaSequences
type AddMediaSequencesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddMediaSequencesRequest creates a request to invoke AddMediaSequences API
func CreateAddMediaSequencesRequest() (request *AddMediaSequencesRequest) {
	request = &AddMediaSequencesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "AddMediaSequences", "vod", "openAPI")
	return
}

// CreateAddMediaSequencesResponse creates a response to parse from AddMediaSequences response
func CreateAddMediaSequencesResponse() (response *AddMediaSequencesResponse) {
	response = &AddMediaSequencesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
