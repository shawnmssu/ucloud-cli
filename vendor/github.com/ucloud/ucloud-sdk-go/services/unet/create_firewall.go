//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet CreateFirewall

package unet

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// CreateFirewallRequest is request schema for CreateFirewall action
type CreateFirewallRequest struct {
	request.CommonBase

	// 防火墙规则，例如：TCP|22|192.168.1.1/22|DROP|LOW，第一个参数代表协议：第二个参数代表端口号，第三个参数为ip，第四个参数为ACCEPT（接受）和DROP（拒绝），第五个参数优先级：HIGH（高），MEDIUM（中），LOW（低）
	Rule []string `required:"true"`

	// 防火墙名称， 默认为Firewall
	Name *string `required:"false"`

	// 防火墙业务组，默认为Default
	Tag *string `required:"false"`

	// 防火墙描述，默认为空
	Remark *string `required:"false"`
}

// CreateFirewallResponse is response schema for CreateFirewall action
type CreateFirewallResponse struct {
	response.CommonBase

	// 防火墙 ID
	FWId string
}

// NewCreateFirewallRequest will create request of CreateFirewall action.
func (c *UNetClient) NewCreateFirewallRequest() *CreateFirewallRequest {
	req := &CreateFirewallRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CreateFirewall - 创建防火墙
func (c *UNetClient) CreateFirewall(req *CreateFirewallRequest) (*CreateFirewallResponse, error) {
	var err error
	var res CreateFirewallResponse

	err = c.client.InvokeAction("CreateFirewall", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
