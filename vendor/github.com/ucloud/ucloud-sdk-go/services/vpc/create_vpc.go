//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api VPC CreateVPC

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// CreateVPCRequest is request schema for CreateVPC action
type CreateVPCRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// VPC名称
	Name *string `required:"true"`

	// VPC网段
	Network []string `required:"true"`

	// 业务组名称
	Tag *string `required:"false"`

	// 备注
	Remark *string `required:"false"`

	// VPC类型
	Type *int `required:"false"`
}

// CreateVPCResponse is response schema for CreateVPC action
type CreateVPCResponse struct {
	response.CommonBase

	// VPC资源Id
	VPCId string
}

// NewCreateVPCRequest will create request of CreateVPC action.
func (c *VPCClient) NewCreateVPCRequest() *CreateVPCRequest {
	req := &CreateVPCRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CreateVPC - 创建VPC
func (c *VPCClient) CreateVPC(req *CreateVPCRequest) (*CreateVPCResponse, error) {
	var err error
	var res CreateVPCResponse

	err = c.Client.InvokeAction("CreateVPC", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
