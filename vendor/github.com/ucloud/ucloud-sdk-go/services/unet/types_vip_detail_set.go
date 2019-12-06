// Code is generated by ucloud-model, DO NOT EDIT IT.

package unet

/*
VIPDetailSet - VIPDetailSet

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type VIPDetailSet struct {

	// 创建时间
	CreateTime int

	// VIP名称
	Name string

	// 真实主机ip
	RealIp string

	// VIP备注
	Remark string

	// 子网id
	SubnetId string

	// VIP所属业务组
	Tag string

	// 虚拟ip
	VIP string

	// 虚拟ip id
	VIPId string

	// VPC id
	VPCId string

	// 地域
	Zone string
}
