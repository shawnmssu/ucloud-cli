//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet DescribeFirewallResource

package unet

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeFirewallResourceRequest is request schema for DescribeFirewallResource action
type DescribeFirewallResourceRequest struct {
	request.CommonBase

	// 防火墙ID
	FWId *string `required:"true"`

	// 返回数据长度，默认为20，最大10000000
	Limit *string `required:"false"`

	// 列表起始位置偏移量，默认为0
	Offset *string `required:"false"`
}

// DescribeFirewallResourceResponse is response schema for DescribeFirewallResource action
type DescribeFirewallResourceResponse struct {
	response.CommonBase

	// 资源列表，见 ResourceSet
	ResourceSet []ResourceSet

	// 防火墙已绑定资源的总数
	TotalCount int
}

// NewDescribeFirewallResourceRequest will create request of DescribeFirewallResource action.
func (c *UNetClient) NewDescribeFirewallResourceRequest() *DescribeFirewallResourceRequest {
	req := &DescribeFirewallResourceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeFirewallResource - 获取防火墙组所绑定资源的外网IP
func (c *UNetClient) DescribeFirewallResource(req *DescribeFirewallResourceRequest) (*DescribeFirewallResourceResponse, error) {
	var err error
	var res DescribeFirewallResourceResponse

	err = c.client.InvokeAction("DescribeFirewallResource", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
