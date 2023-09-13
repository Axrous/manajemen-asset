package dto

import "time"

type ManageAssetRequest struct {
	Id                   string
	IdUser               string
	NikStaff             string
	SubmisstionDate      time.Time
	ReturnDate           time.Time
	ManageAssetDetailReq []ManageAssetDetailRequest
}

type ManageAssetDetailRequest struct {
	Id string
	IdManageAsset string
	IdAsset       string
	TotalItem     int
	Status        string
}