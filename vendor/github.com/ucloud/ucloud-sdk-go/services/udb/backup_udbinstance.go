//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB BackupUDBInstance

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// BackupUDBInstanceRequest is request schema for BackupUDBInstance action
type BackupUDBInstanceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// DB实例Id,该值可以通过DescribeUDBInstance获取
	DBId *string `required:"true"`

	// 备份名称
	BackupName *string `required:"true"`

	// 是否使用黑名单备份，默认false
	UseBlacklist *bool `required:"false"`
}

// BackupUDBInstanceResponse is response schema for BackupUDBInstance action
type BackupUDBInstanceResponse struct {
	response.CommonBase
}

// NewBackupUDBInstanceRequest will create request of BackupUDBInstance action.
func (c *UDBClient) NewBackupUDBInstanceRequest() *BackupUDBInstanceRequest {
	req := &BackupUDBInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// BackupUDBInstance - 备份UDB实例
func (c *UDBClient) BackupUDBInstance(req *BackupUDBInstanceRequest) (*BackupUDBInstanceResponse, error) {
	var err error
	var res BackupUDBInstanceResponse

	err = c.Client.InvokeAction("BackupUDBInstance", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
