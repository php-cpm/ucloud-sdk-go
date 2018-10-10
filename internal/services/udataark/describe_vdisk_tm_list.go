//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDataArk DescribeVDiskTmList

package udataark

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeVDiskTmListRequest is request schema for DescribeVDiskTmList action
type DescribeVDiskTmListRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 磁盘ID
	VDiskId *string `required:"true"`

	// 按备份类型过滤，"all"表示获取所有备份 ，目前仅支持“all”
	SnapshotType *string `required:"true"`
}

// DescribeVDiskTmListResponse is response schema for DescribeVDiskTmList action
type DescribeVDiskTmListResponse struct {
	response.CommonBase

	// 返回数量
	TotalCount string

	// 备份链信息
	DataSet []DiskListDataSet

	// 备份链秒级信息
	FreshSet []DiskListFreshSet
}

// NewDescribeVDiskTmListRequest will create request of DescribeVDiskTmList action.
func (c *UDataArkClient) NewDescribeVDiskTmListRequest() *DescribeVDiskTmListRequest {
	req := &DescribeVDiskTmListRequest{}
	c.client.SetupRequest(req)
	return req
}

// DescribeVDiskTmList - 查询磁盘备份链信息
func (c *UDataArkClient) DescribeVDiskTmList(req *DescribeVDiskTmListRequest) (*DescribeVDiskTmListResponse, error) {
	var err error
	var res DescribeVDiskTmListResponse

	err = c.client.InvokeAction("DescribeVDiskTmList", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
