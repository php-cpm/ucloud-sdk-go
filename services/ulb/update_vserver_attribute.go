//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api ULB UpdateVServerAttribute

package ulb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UpdateVServerAttributeRequest is request schema for UpdateVServerAttribute action
type UpdateVServerAttributeRequest struct {
	request.CommonBase

	// 负载均衡实例ID
	ULBId *string `required:"true"`

	// VServer实例ID
	VServerId *string `required:"true"`

	// VServer实例名称，若无此字段则不做修改
	VServerName *string `required:"false"`

	// VServer协议类型，请求代理只支持修改为 HTTP/HTTPS，报文转发VServer只支持修改为 TCP/UDP
	Protocol *string `required:"false"`

	// VServer负载均衡算法，ConsistentHash，SourcePort，ConsistentHashPort 只在报文转发中使用；Roundrobin和Source在请求代理和报文转发中使用。
	Method *string `required:"false"`

	// VServer会话保持模式，若无此字段则不做修改。枚举值：None：关闭；ServerInsert：自动生成KEY；UserDefined：用户自定义KEY。
	PersistenceType *string `required:"false"`

	// 根据PersistenceType确定: None或ServerInsert, 此字段无意义; UserDefined, 则此字段传入用户自定义会话保持String. 若无此字段则不做修改
	PersistenceInfo *string `required:"false"`

	// 请求代理的VServer下表示空闲连接的回收时间，单位：秒，取值范围：时(0，86400]，默认值为60；报文转发的VServer下表示回话保持的时间，单位：秒，取值范围：[60，900]，0 表示禁用连接保持
	ClientTimeout *int `required:"false"`

	// 健康检查的类型，Port:端口,Path:路径
	MonitorType *string `required:"false"`

	// MonitorType 为 Path 时指定健康检查发送请求时HTTP HEADER 里的域名
	Domain *string `required:"false"`

	// MonitorType 为 Path 时指定健康检查发送请求时的路径，默认为 /
	Path *string `required:"false"`
}

// UpdateVServerAttributeResponse is response schema for UpdateVServerAttribute action
type UpdateVServerAttributeResponse struct {
	response.CommonBase
}

// NewUpdateVServerAttributeRequest will create request of UpdateVServerAttribute action.
func (c *ULBClient) NewUpdateVServerAttributeRequest() *UpdateVServerAttributeRequest {
	req := &UpdateVServerAttributeRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// UpdateVServerAttribute - 更新VServer实例属性
func (c *ULBClient) UpdateVServerAttribute(req *UpdateVServerAttributeRequest) (*UpdateVServerAttributeResponse, error) {
	var err error
	var res UpdateVServerAttributeResponse

	err = c.client.InvokeAction("UpdateVServerAttribute", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
