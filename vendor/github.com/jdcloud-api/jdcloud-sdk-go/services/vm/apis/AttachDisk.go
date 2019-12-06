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

type AttachDiskRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 云主机ID  */
    InstanceId string `json:"instanceId"`

    /* 云硬盘ID  */
    DiskId string `json:"diskId"`

    /* 数据盘的逻辑挂载点[vda,vdb,vdc,vdd,vde,vdf,vdg,vdh,vdi]，挂载系统盘时vda必传 (Optional) */
    DeviceName *string `json:"deviceName"`

    /* 自动随主机删除此云硬盘，默认为False。仅按配置计费云硬盘支持修改此参数，包年包月云硬盘默认为False且不可修改。如果是共享型云硬盘，此参数无效。 (Optional) */
    AutoDelete *bool `json:"autoDelete"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 云主机ID (Required)
 * param diskId: 云硬盘ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAttachDiskRequest(
    regionId string,
    instanceId string,
    diskId string,
) *AttachDiskRequest {

	return &AttachDiskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:attachDisk",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        DiskId: diskId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 云主机ID (Required)
 * param diskId: 云硬盘ID (Required)
 * param deviceName: 数据盘的逻辑挂载点[vda,vdb,vdc,vdd,vde,vdf,vdg,vdh,vdi]，挂载系统盘时vda必传 (Optional)
 * param autoDelete: 自动随主机删除此云硬盘，默认为False。仅按配置计费云硬盘支持修改此参数，包年包月云硬盘默认为False且不可修改。如果是共享型云硬盘，此参数无效。 (Optional)
 */
func NewAttachDiskRequestWithAllParams(
    regionId string,
    instanceId string,
    diskId string,
    deviceName *string,
    autoDelete *bool,
) *AttachDiskRequest {

    return &AttachDiskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:attachDisk",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        DiskId: diskId,
        DeviceName: deviceName,
        AutoDelete: autoDelete,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAttachDiskRequestWithoutParam() *AttachDiskRequest {

    return &AttachDiskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:attachDisk",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *AttachDiskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 云主机ID(Required) */
func (r *AttachDiskRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param diskId: 云硬盘ID(Required) */
func (r *AttachDiskRequest) SetDiskId(diskId string) {
    r.DiskId = diskId
}

/* param deviceName: 数据盘的逻辑挂载点[vda,vdb,vdc,vdd,vde,vdf,vdg,vdh,vdi]，挂载系统盘时vda必传(Optional) */
func (r *AttachDiskRequest) SetDeviceName(deviceName string) {
    r.DeviceName = &deviceName
}

/* param autoDelete: 自动随主机删除此云硬盘，默认为False。仅按配置计费云硬盘支持修改此参数，包年包月云硬盘默认为False且不可修改。如果是共享型云硬盘，此参数无效。(Optional) */
func (r *AttachDiskRequest) SetAutoDelete(autoDelete bool) {
    r.AutoDelete = &autoDelete
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AttachDiskRequest) GetRegionId() string {
    return r.RegionId
}

type AttachDiskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AttachDiskResult `json:"result"`
}

type AttachDiskResult struct {
}