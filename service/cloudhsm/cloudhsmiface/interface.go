// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudhsmiface provides an interface for the Amazon CloudHSM.
package cloudhsmiface

import (
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/service/cloudhsm"
)

// CloudHSMAPI is the interface type for cloudhsm.CloudHSM.
type CloudHSMAPI interface {
	CreateHapgRequest(*cloudhsm.CreateHapgInput) (*service.Request, *cloudhsm.CreateHapgOutput)

	CreateHapg(*cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error)

	CreateHsmRequest(*cloudhsm.CreateHsmInput) (*service.Request, *cloudhsm.CreateHsmOutput)

	CreateHsm(*cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error)

	CreateLunaClientRequest(*cloudhsm.CreateLunaClientInput) (*service.Request, *cloudhsm.CreateLunaClientOutput)

	CreateLunaClient(*cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error)

	DeleteHapgRequest(*cloudhsm.DeleteHapgInput) (*service.Request, *cloudhsm.DeleteHapgOutput)

	DeleteHapg(*cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error)

	DeleteHsmRequest(*cloudhsm.DeleteHsmInput) (*service.Request, *cloudhsm.DeleteHsmOutput)

	DeleteHsm(*cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error)

	DeleteLunaClientRequest(*cloudhsm.DeleteLunaClientInput) (*service.Request, *cloudhsm.DeleteLunaClientOutput)

	DeleteLunaClient(*cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error)

	DescribeHapgRequest(*cloudhsm.DescribeHapgInput) (*service.Request, *cloudhsm.DescribeHapgOutput)

	DescribeHapg(*cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error)

	DescribeHsmRequest(*cloudhsm.DescribeHsmInput) (*service.Request, *cloudhsm.DescribeHsmOutput)

	DescribeHsm(*cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error)

	DescribeLunaClientRequest(*cloudhsm.DescribeLunaClientInput) (*service.Request, *cloudhsm.DescribeLunaClientOutput)

	DescribeLunaClient(*cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error)

	GetConfigRequest(*cloudhsm.GetConfigInput) (*service.Request, *cloudhsm.GetConfigOutput)

	GetConfig(*cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error)

	ListAvailableZonesRequest(*cloudhsm.ListAvailableZonesInput) (*service.Request, *cloudhsm.ListAvailableZonesOutput)

	ListAvailableZones(*cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error)

	ListHapgsRequest(*cloudhsm.ListHapgsInput) (*service.Request, *cloudhsm.ListHapgsOutput)

	ListHapgs(*cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error)

	ListHsmsRequest(*cloudhsm.ListHsmsInput) (*service.Request, *cloudhsm.ListHsmsOutput)

	ListHsms(*cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error)

	ListLunaClientsRequest(*cloudhsm.ListLunaClientsInput) (*service.Request, *cloudhsm.ListLunaClientsOutput)

	ListLunaClients(*cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error)

	ModifyHapgRequest(*cloudhsm.ModifyHapgInput) (*service.Request, *cloudhsm.ModifyHapgOutput)

	ModifyHapg(*cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error)

	ModifyHsmRequest(*cloudhsm.ModifyHsmInput) (*service.Request, *cloudhsm.ModifyHsmOutput)

	ModifyHsm(*cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error)

	ModifyLunaClientRequest(*cloudhsm.ModifyLunaClientInput) (*service.Request, *cloudhsm.ModifyLunaClientOutput)

	ModifyLunaClient(*cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error)
}
