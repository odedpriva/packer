// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type DescribeQuotaRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 资源类型，取值范围：vpc、elastic_ip、subnet、security_group、vpcpeering、network_interface（配额只统计辅助网卡）  */
    Type string `json:"type"`

    /* type为vpc、elastic_ip、network_interface不设置, type为subnet、security_group、vpcpeering设置为vpcId (Optional) */
    ParentResourceId *string `json:"parentResourceId"`
}

/*
 * param regionId: Region ID (Required)
 * param type_: 资源类型，取值范围：vpc、elastic_ip、subnet、security_group、vpcpeering、network_interface（配额只统计辅助网卡） (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeQuotaRequest(
    regionId string,
    type_ string,
) *DescribeQuotaRequest {

	return &DescribeQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/quotas/",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Type: type_,
	}
}

/*
 * param regionId: Region ID (Required)
 * param type_: 资源类型，取值范围：vpc、elastic_ip、subnet、security_group、vpcpeering、network_interface（配额只统计辅助网卡） (Required)
 * param parentResourceId: type为vpc、elastic_ip、network_interface不设置, type为subnet、security_group、vpcpeering设置为vpcId (Optional)
 */
func NewDescribeQuotaRequestWithAllParams(
    regionId string,
    type_ string,
    parentResourceId *string,
) *DescribeQuotaRequest {

    return &DescribeQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quotas/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Type: type_,
        ParentResourceId: parentResourceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeQuotaRequestWithoutParam() *DescribeQuotaRequest {

    return &DescribeQuotaRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quotas/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeQuotaRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param type_: 资源类型，取值范围：vpc、elastic_ip、subnet、security_group、vpcpeering、network_interface（配额只统计辅助网卡）(Required) */
func (r *DescribeQuotaRequest) SetType(type_ string) {
    r.Type = type_
}

/* param parentResourceId: type为vpc、elastic_ip、network_interface不设置, type为subnet、security_group、vpcpeering设置为vpcId(Optional) */
func (r *DescribeQuotaRequest) SetParentResourceId(parentResourceId string) {
    r.ParentResourceId = &parentResourceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeQuotaRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeQuotaResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeQuotaResult `json:"result"`
}

type DescribeQuotaResult struct {
    Quota interface{} `json:"quota"`
}