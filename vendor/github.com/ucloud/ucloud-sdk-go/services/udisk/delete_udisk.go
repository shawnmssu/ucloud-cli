//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDisk DeleteUDisk

package udisk

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteUDiskRequest is request schema for DeleteUDisk action
type DeleteUDiskRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 要删除的UDisk的Id
	UDiskId *string `required:"true"`
}

// DeleteUDiskResponse is response schema for DeleteUDisk action
type DeleteUDiskResponse struct {
	response.CommonBase
}

// NewDeleteUDiskRequest will create request of DeleteUDisk action.
func (c *UDiskClient) NewDeleteUDiskRequest() *DeleteUDiskRequest {
	req := &DeleteUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteUDisk - 删除UDisk
func (c *UDiskClient) DeleteUDisk(req *DeleteUDiskRequest) (*DeleteUDiskResponse, error) {
	var err error
	var res DeleteUDiskResponse

	err = c.Client.InvokeAction("DeleteUDisk", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}