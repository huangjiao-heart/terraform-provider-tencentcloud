// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20230927

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-09-27"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("taas", APIVersion, "CreateCluster")
    
    
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCluster
// CreateCluster create instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    return c.CreateClusterWithContext(context.Background(), request)
}

// CreateCluster
// CreateCluster create instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRHInstanceRequest() (request *CreateRHInstanceRequest) {
    request = &CreateRHInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("taas", APIVersion, "CreateRHInstance")
    
    
    return
}

func NewCreateRHInstanceResponse() (response *CreateRHInstanceResponse) {
    response = &CreateRHInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRHInstance
// CreateRHInstance create instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) CreateRHInstance(request *CreateRHInstanceRequest) (response *CreateRHInstanceResponse, err error) {
    return c.CreateRHInstanceWithContext(context.Background(), request)
}

// CreateRHInstance
// CreateRHInstance create instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) CreateRHInstanceWithContext(ctx context.Context, request *CreateRHInstanceRequest) (response *CreateRHInstanceResponse, err error) {
    if request == nil {
        request = NewCreateRHInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRHInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRHInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("taas", APIVersion, "DeleteCluster")
    
    
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCluster
// DeleteCluster delete instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    return c.DeleteClusterWithContext(context.Background(), request)
}

// DeleteCluster
// DeleteCluster delete instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DeleteClusterWithContext(ctx context.Context, request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRHInstanceRequest() (request *DeleteRHInstanceRequest) {
    request = &DeleteRHInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("taas", APIVersion, "DeleteRHInstance")
    
    
    return
}

func NewDeleteRHInstanceResponse() (response *DeleteRHInstanceResponse) {
    response = &DeleteRHInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRHInstance
// DeleteRHInstance delete instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DeleteRHInstance(request *DeleteRHInstanceRequest) (response *DeleteRHInstanceResponse, err error) {
    return c.DeleteRHInstanceWithContext(context.Background(), request)
}

// DeleteRHInstance
// DeleteRHInstance delete instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DeleteRHInstanceWithContext(ctx context.Context, request *DeleteRHInstanceRequest) (response *DeleteRHInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteRHInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRHInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRHInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("taas", APIVersion, "DescribeClusters")
    
    
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusters
// DescribeCluster read instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// DescribeCluster read instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRHInstancesRequest() (request *DescribeRHInstancesRequest) {
    request = &DescribeRHInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("taas", APIVersion, "DescribeRHInstances")
    
    
    return
}

func NewDescribeRHInstancesResponse() (response *DescribeRHInstancesResponse) {
    response = &DescribeRHInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRHInstances
// DescribeRHInstance read instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeRHInstances(request *DescribeRHInstancesRequest) (response *DescribeRHInstancesResponse, err error) {
    return c.DescribeRHInstancesWithContext(context.Background(), request)
}

// DescribeRHInstances
// DescribeRHInstance read instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeRHInstancesWithContext(ctx context.Context, request *DescribeRHInstancesRequest) (response *DescribeRHInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeRHInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRHInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRHInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterAttributesRequest() (request *ModifyClusterAttributesRequest) {
    request = &ModifyClusterAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("taas", APIVersion, "ModifyClusterAttributes")
    
    
    return
}

func NewModifyClusterAttributesResponse() (response *ModifyClusterAttributesResponse) {
    response = &ModifyClusterAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterAttributes
// ModifyClusterAttributes modify instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ModifyClusterAttributes(request *ModifyClusterAttributesRequest) (response *ModifyClusterAttributesResponse, err error) {
    return c.ModifyClusterAttributesWithContext(context.Background(), request)
}

// ModifyClusterAttributes
// ModifyClusterAttributes modify instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ModifyClusterAttributesWithContext(ctx context.Context, request *ModifyClusterAttributesRequest) (response *ModifyClusterAttributesResponse, err error) {
    if request == nil {
        request = NewModifyClusterAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRHInstanceAttributesRequest() (request *ModifyRHInstanceAttributesRequest) {
    request = &ModifyRHInstanceAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("taas", APIVersion, "ModifyRHInstanceAttributes")
    
    
    return
}

func NewModifyRHInstanceAttributesResponse() (response *ModifyRHInstanceAttributesResponse) {
    response = &ModifyRHInstanceAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRHInstanceAttributes
// ModifyRHInstanceAttributes modify instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ModifyRHInstanceAttributes(request *ModifyRHInstanceAttributesRequest) (response *ModifyRHInstanceAttributesResponse, err error) {
    return c.ModifyRHInstanceAttributesWithContext(context.Background(), request)
}

// ModifyRHInstanceAttributes
// ModifyRHInstanceAttributes modify instance
//
// 可能返回的错误码:
//  FAILEDOPERATION_HTTPREQUESTERROR = "FailedOperation.HttpRequestError"
//  INTERNALERROR_CAPISERVICEFAILED = "InternalError.CapiServiceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ModifyRHInstanceAttributesWithContext(ctx context.Context, request *ModifyRHInstanceAttributesRequest) (response *ModifyRHInstanceAttributesResponse, err error) {
    if request == nil {
        request = NewModifyRHInstanceAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRHInstanceAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRHInstanceAttributesResponse()
    err = c.Send(request, response)
    return
}
