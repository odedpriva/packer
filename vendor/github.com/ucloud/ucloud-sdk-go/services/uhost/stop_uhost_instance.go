//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost StopUHostInstance

package uhost

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// StopUHostInstanceRequest is request schema for StopUHostInstance action
type StopUHostInstanceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// UHost实例ID 参见 [DescribeUHostInstance](describe_uhost_instance.html)
	UHostId *string `required:"true"`
}

// StopUHostInstanceResponse is response schema for StopUHostInstance action
type StopUHostInstanceResponse struct {
	response.CommonBase

	// UHost实例ID
	UhostId string
}

// NewStopUHostInstanceRequest will create request of StopUHostInstance action.
func (c *UHostClient) NewStopUHostInstanceRequest() *StopUHostInstanceRequest {
	req := &StopUHostInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// StopUHostInstance - 指停止处于运行状态的UHost实例，需指定数据中心及UhostID。
func (c *UHostClient) StopUHostInstance(req *StopUHostInstanceRequest) (*StopUHostInstanceResponse, error) {
	var err error
	var res StopUHostInstanceResponse

	err = c.Client.InvokeAction("StopUHostInstance", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
