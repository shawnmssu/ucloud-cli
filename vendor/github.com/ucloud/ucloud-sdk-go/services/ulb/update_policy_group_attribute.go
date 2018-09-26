//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api ULB UpdatePolicyGroupAttribute

package ulb

import (
	"github.com/ucloud/ucloud-sdk-go/sdk"
	"github.com/ucloud/ucloud-sdk-go/sdk/request"
	"github.com/ucloud/ucloud-sdk-go/sdk/response"
)

// UpdatePolicyGroupAttributeRequest is request schema for UpdatePolicyGroupAttribute action
type UpdatePolicyGroupAttributeRequest struct {
	request.CommonBase

	// 内容转发策略组ID
	GroupId *string `required:"true"`

	// 修改策略转发组名称
	GroupName *string `required:"false"`
}

// UpdatePolicyGroupAttributeResponse is response schema for UpdatePolicyGroupAttribute action
type UpdatePolicyGroupAttributeResponse struct {
	response.CommonBase
}

// NewUpdatePolicyGroupAttributeRequest will create request of UpdatePolicyGroupAttribute action.
func (c *ULBClient) NewUpdatePolicyGroupAttributeRequest() *UpdatePolicyGroupAttributeRequest {
	cfg := c.client.GetConfig()

	return &UpdatePolicyGroupAttributeRequest{
		CommonBase: request.CommonBase{
			Region:    sdk.String(cfg.Region),
			ProjectId: sdk.String(cfg.ProjectId),
		},
	}
}

// UpdatePolicyGroupAttribute - 更新内容转发策略组属性
func (c *ULBClient) UpdatePolicyGroupAttribute(req *UpdatePolicyGroupAttributeRequest) (*UpdatePolicyGroupAttributeResponse, error) {
	var err error
	var res UpdatePolicyGroupAttributeResponse

	err = c.client.InvokeAction("UpdatePolicyGroupAttribute", req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
