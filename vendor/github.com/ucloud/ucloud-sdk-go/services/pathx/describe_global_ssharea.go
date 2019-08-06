//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api PathX DescribeGlobalSSHArea

package pathx

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeGlobalSSHAreaRequest is request schema for DescribeGlobalSSHArea action
type DescribeGlobalSSHAreaRequest struct {
	request.CommonBase
}

// DescribeGlobalSSHAreaResponse is response schema for DescribeGlobalSSHArea action
type DescribeGlobalSSHAreaResponse struct {
	response.CommonBase

	// 支持GlobalSSH的地区
	AreaSet []GlobalSSHArea

	// 提示信息
	Message string
}

// NewDescribeGlobalSSHAreaRequest will create request of DescribeGlobalSSHArea action.
func (c *PathXClient) NewDescribeGlobalSSHAreaRequest() *DescribeGlobalSSHAreaRequest {
	req := &DescribeGlobalSSHAreaRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeGlobalSSHArea - 获取GlobalSSH覆盖的地区列表 用于控制显示哪些机房地域可以使用SSH特性
func (c *PathXClient) DescribeGlobalSSHArea(req *DescribeGlobalSSHAreaRequest) (*DescribeGlobalSSHAreaResponse, error) {
	var err error
	var res DescribeGlobalSSHAreaResponse

	err = c.Client.InvokeAction("DescribeGlobalSSHArea", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
