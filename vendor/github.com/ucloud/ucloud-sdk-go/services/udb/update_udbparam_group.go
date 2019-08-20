//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB UpdateUDBParamGroup

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UpdateUDBParamGroupRequest is request schema for UpdateUDBParamGroup action
type UpdateUDBParamGroupRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 配置参数组id，使用DescribeUDBParamGroup获得
	GroupId *int `required:"true"`

	// 参数名称
	Key *string `required:"true"`

	// 参数值
	Value *string `required:"true"`

	// 该配置文件是否是地域级别配置文件，默认是false
	RegionFlag *bool `required:"false"`
}

// UpdateUDBParamGroupResponse is response schema for UpdateUDBParamGroup action
type UpdateUDBParamGroupResponse struct {
	response.CommonBase
}

// NewUpdateUDBParamGroupRequest will create request of UpdateUDBParamGroup action.
func (c *UDBClient) NewUpdateUDBParamGroupRequest() *UpdateUDBParamGroupRequest {
	req := &UpdateUDBParamGroupRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// UpdateUDBParamGroup - 更新UDB配置参数项
func (c *UDBClient) UpdateUDBParamGroup(req *UpdateUDBParamGroupRequest) (*UpdateUDBParamGroupResponse, error) {
	var err error
	var res UpdateUDBParamGroupResponse

	err = c.Client.InvokeAction("UpdateUDBParamGroup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}