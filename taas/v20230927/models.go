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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AssociatedRHInstance struct {
	// ResultHouse 实例 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`
}

type AuthorizedTCRRegistryInfo struct {
	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespaces []*string `json:"Namespaces,omitnil" name:"Namespaces"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryId *string `json:"RegistryId,omitnil" name:"RegistryId"`

	// 服务凭证id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceAccountId *string `json:"ServiceAccountId,omitnil" name:"ServiceAccountId"`
}

type ClusterFilter struct {
	// name 需要过滤的字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// name 操作类型，例如equal、notEqual、in、notIn、like、notLike，默认equal
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ops *string `json:"Ops,omitnil" name:"Ops"`

	// values 字段的过滤值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

// Predefined struct for user
type CreateClusterRequestParams struct {
	// associatedRHInstance 关联的 ResultHouse 实例信息
	AssociatedRHInstance *AssociatedRHInstance `json:"AssociatedRHInstance,omitnil" name:"AssociatedRHInstance"`

	// authorizedTCRRegistryInfos 关联的tcr实例信息
	AuthorizedTCRRegistryInfos []*AuthorizedTCRRegistryInfo `json:"AuthorizedTCRRegistryInfos,omitnil" name:"AuthorizedTCRRegistryInfos"`

	// clusterName 集群名
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// description 集群描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// publicAccess 是否公网访问
	PublicAccess *bool `json:"PublicAccess,omitnil" name:"PublicAccess"`

	// tags 集群标签信息
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// thirdPartyServiceAccounts 关联的第三方凭证信息
	ThirdPartyServiceAccounts []*ThirdPartyServiceAccount `json:"ThirdPartyServiceAccounts,omitnil" name:"ThirdPartyServiceAccounts"`

	// vpcConfig 私有网络参数配置
	VpcConfig *VpcConfig `json:"VpcConfig,omitnil" name:"VpcConfig"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest
	
	// associatedRHInstance 关联的 ResultHouse 实例信息
	AssociatedRHInstance *AssociatedRHInstance `json:"AssociatedRHInstance,omitnil" name:"AssociatedRHInstance"`

	// authorizedTCRRegistryInfos 关联的tcr实例信息
	AuthorizedTCRRegistryInfos []*AuthorizedTCRRegistryInfo `json:"AuthorizedTCRRegistryInfos,omitnil" name:"AuthorizedTCRRegistryInfos"`

	// clusterName 集群名
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// description 集群描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// publicAccess 是否公网访问
	PublicAccess *bool `json:"PublicAccess,omitnil" name:"PublicAccess"`

	// tags 集群标签信息
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// thirdPartyServiceAccounts 关联的第三方凭证信息
	ThirdPartyServiceAccounts []*ThirdPartyServiceAccount `json:"ThirdPartyServiceAccounts,omitnil" name:"ThirdPartyServiceAccounts"`

	// vpcConfig 私有网络参数配置
	VpcConfig *VpcConfig `json:"VpcConfig,omitnil" name:"VpcConfig"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssociatedRHInstance")
	delete(f, "AuthorizedTCRRegistryInfos")
	delete(f, "ClusterName")
	delete(f, "Description")
	delete(f, "PublicAccess")
	delete(f, "Tags")
	delete(f, "ThirdPartyServiceAccounts")
	delete(f, "VpcConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterResponseParams struct {
	// appId
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// clusterId 集群 ID 唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// instance 实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instance *ReadOnlyCluster `json:"Instance,omitnil" name:"Instance"`

	// root account uin
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterResponseParams `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRHInstanceRequestParams struct {
	// autoVoucher
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// chargeType
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`

	// count
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// dealName
	DealName *string `json:"DealName,omitnil" name:"DealName"`

	// instanceName
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// instanceType
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// period
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// renewFlag
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// rootPassword
	RootPassword *string `json:"RootPassword,omitnil" name:"RootPassword"`

	// tags
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// volume
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`
}

type CreateRHInstanceRequest struct {
	*tchttp.BaseRequest
	
	// autoVoucher
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// chargeType
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`

	// count
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// dealName
	DealName *string `json:"DealName,omitnil" name:"DealName"`

	// instanceName
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// instanceType
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// period
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// renewFlag
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// rootPassword
	RootPassword *string `json:"RootPassword,omitnil" name:"RootPassword"`

	// tags
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// volume
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`
}

func (r *CreateRHInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRHInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoVoucher")
	delete(f, "ChargeType")
	delete(f, "Count")
	delete(f, "DealName")
	delete(f, "InstanceName")
	delete(f, "InstanceType")
	delete(f, "Period")
	delete(f, "RenewFlag")
	delete(f, "RootPassword")
	delete(f, "Tags")
	delete(f, "Volume")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRHInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRHInstanceResponseParams struct {
	// appId
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// instance 实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instance *ReadOnlyRHInstance `json:"Instance,omitnil" name:"Instance"`

	// instanceId  唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// root account uin
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRHInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateRHInstanceResponseParams `json:"Response"`
}

func (r *CreateRHInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRHInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterRequestParams struct {
	// clusterId 集群 ID 唯一ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest
	
	// clusterId 集群 ID 唯一ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterResponseParams struct {
	// appId
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// root account uin
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterResponseParams `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRHInstanceRequestParams struct {
	// instanceId  唯一ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DeleteRHInstanceRequest struct {
	*tchttp.BaseRequest
	
	// instanceId  唯一ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DeleteRHInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRHInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRHInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRHInstanceResponseParams struct {
	// appId
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// root account uin
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRHInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRHInstanceResponseParams `json:"Response"`
}

func (r *DeleteRHInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRHInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// clusterId 集群 ID 唯一ID
	ClusterIdList []*string `json:"ClusterIdList,omitnil" name:"ClusterIdList"`

	// 过滤条件
	Filters []*ClusterFilter `json:"Filters,omitnil" name:"Filters"`

	// 页码，从0开始
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// clusterId 集群 ID 唯一ID
	ClusterIdList []*string `json:"ClusterIdList,omitnil" name:"ClusterIdList"`

	// 过滤条件
	Filters []*ClusterFilter `json:"Filters,omitnil" name:"Filters"`

	// 页码，从0开始
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIdList")
	delete(f, "Filters")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// appId
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// 实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSet []*ReadOnlyCluster `json:"InstanceSet,omitnil" name:"InstanceSet"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// root account uin
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClustersResponseParams `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRHInstancesRequestParams struct {
	// 过滤条件
	Filters []*RHInstanceFilter `json:"Filters,omitnil" name:"Filters"`

	// instanceId  唯一ID
	InstanceIdList []*string `json:"InstanceIdList,omitnil" name:"InstanceIdList"`

	// 页码，从0开始
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeRHInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	Filters []*RHInstanceFilter `json:"Filters,omitnil" name:"Filters"`

	// instanceId  唯一ID
	InstanceIdList []*string `json:"InstanceIdList,omitnil" name:"InstanceIdList"`

	// 页码，从0开始
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeRHInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRHInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "InstanceIdList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRHInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRHInstancesResponseParams struct {
	// appId
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// 实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSet []*ReadOnlyRHInstance `json:"InstanceSet,omitnil" name:"InstanceSet"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// root account uin
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRHInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRHInstancesResponseParams `json:"Response"`
}

func (r *DescribeRHInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRHInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAttributesRequestParams struct {
	// associatedRHInstance
	AssociatedRHInstance *AssociatedRHInstance `json:"AssociatedRHInstance,omitnil" name:"AssociatedRHInstance"`

	// authorizedTCRRegistryInfos
	AuthorizedTCRRegistryInfos []*AuthorizedTCRRegistryInfo `json:"AuthorizedTCRRegistryInfos,omitnil" name:"AuthorizedTCRRegistryInfos"`

	// clusterId 集群 ID 唯一ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 需要修改的属性 clusterName
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// description
	Description *string `json:"Description,omitnil" name:"Description"`

	// publicAccess
	PublicAccess *bool `json:"PublicAccess,omitnil" name:"PublicAccess"`

	// tags
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// thirdPartyServiceAccounts
	ThirdPartyServiceAccounts []*ThirdPartyServiceAccount `json:"ThirdPartyServiceAccounts,omitnil" name:"ThirdPartyServiceAccounts"`
}

type ModifyClusterAttributesRequest struct {
	*tchttp.BaseRequest
	
	// associatedRHInstance
	AssociatedRHInstance *AssociatedRHInstance `json:"AssociatedRHInstance,omitnil" name:"AssociatedRHInstance"`

	// authorizedTCRRegistryInfos
	AuthorizedTCRRegistryInfos []*AuthorizedTCRRegistryInfo `json:"AuthorizedTCRRegistryInfos,omitnil" name:"AuthorizedTCRRegistryInfos"`

	// clusterId 集群 ID 唯一ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 需要修改的属性 clusterName
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// description
	Description *string `json:"Description,omitnil" name:"Description"`

	// publicAccess
	PublicAccess *bool `json:"PublicAccess,omitnil" name:"PublicAccess"`

	// tags
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// thirdPartyServiceAccounts
	ThirdPartyServiceAccounts []*ThirdPartyServiceAccount `json:"ThirdPartyServiceAccounts,omitnil" name:"ThirdPartyServiceAccounts"`
}

func (r *ModifyClusterAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssociatedRHInstance")
	delete(f, "AuthorizedTCRRegistryInfos")
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	delete(f, "Description")
	delete(f, "PublicAccess")
	delete(f, "Tags")
	delete(f, "ThirdPartyServiceAccounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAttributesResponseParams struct {
	// appId
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// root account uin
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyClusterAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterAttributesResponseParams `json:"Response"`
}

func (r *ModifyClusterAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRHInstanceAttributesRequestParams struct {
	// autoVoucher
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// instanceId  唯一ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// instanceName
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 需要修改的属性 instanceType
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// period
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// renewFlag
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// volume
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`
}

type ModifyRHInstanceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// autoVoucher
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// instanceId  唯一ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// instanceName
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 需要修改的属性 instanceType
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// period
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// renewFlag
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// volume
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`
}

func (r *ModifyRHInstanceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRHInstanceAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoVoucher")
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "InstanceType")
	delete(f, "Period")
	delete(f, "RenewFlag")
	delete(f, "Volume")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRHInstanceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRHInstanceAttributesResponseParams struct {
	// appId
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// root account uin
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRHInstanceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRHInstanceAttributesResponseParams `json:"Response"`
}

func (r *ModifyRHInstanceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRHInstanceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RHInstanceFilter struct {
	// name 需要过滤的字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// name 操作类型，例如equal、notEqual、in、notIn、like、notLike，默认equal
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ops *string `json:"Ops,omitnil" name:"Ops"`

	// values 字段的过滤值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type ReadOnlyCluster struct {
	// associatedRHInstance 关联的 ResultHouse 实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedRHInstance *AssociatedRHInstance `json:"AssociatedRHInstance,omitnil" name:"AssociatedRHInstance"`

	// authorizedTCRRegistryInfos 关联的tcr实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizedTCRRegistryInfos []*AuthorizedTCRRegistryInfo `json:"AuthorizedTCRRegistryInfos,omitnil" name:"AuthorizedTCRRegistryInfos"`

	// clusterId 集群 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// clusterName 集群名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// createdAt 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// description 集群描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// publicAccess 是否公网访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAccess *bool `json:"PublicAccess,omitnil" name:"PublicAccess"`

	// status 集群状态 \n- CLUSTER_PENDING  创建中\n\n- CLUSTER_RUNNING 运行中\n\n- CLUSTER_TERMINATING \n销毁中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// tags 集群标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// thirdPartyServiceAccounts 关联的第三方凭证信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThirdPartyServiceAccounts []*ThirdPartyServiceAccount `json:"ThirdPartyServiceAccounts,omitnil" name:"ThirdPartyServiceAccounts"`

	// updatedAt 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil" name:"UpdatedAt"`

	// vpcConfig 私有网络参数配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcConfig *VpcConfig `json:"VpcConfig,omitnil" name:"VpcConfig"`
}

type ReadOnlyRHInstance struct {
	// autoVoucher
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// chargeType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`

	// count
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// createdTime 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// dealName
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitnil" name:"DealName"`

	// expiredTime 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// instanceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// instanceName
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// instanceStatus
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitnil" name:"InstanceStatus"`

	// instanceType
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// period
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// renewFlag
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// rootPassword
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootPassword *string `json:"RootPassword,omitnil" name:"RootPassword"`

	// tags
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// volume
	// 注意：此字段可能返回 null，表示取不到有效值。
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`

	// volumeUsed
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeUsed *int64 `json:"VolumeUsed,omitnil" name:"VolumeUsed"`
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type ThirdPartyServiceAccount struct {
	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil" name:"Password"`

	// 业务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceAccountId *string `json:"ServiceAccountId,omitnil" name:"ServiceAccountId"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Username *string `json:"Username,omitnil" name:"Username"`
}

type VpcConfig struct {
	// 子网的 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 私有网络 的 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}