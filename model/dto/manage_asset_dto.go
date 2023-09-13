package dto

import "time"

type ManageAssetRequest struct {
	Id                   string                     `json:"id"`
	IdUser               string                     `json:"id_user"`
	NikStaff             string                     `json:"nik_staff"`
	SubmisstionDate      time.Time                  `json:"submisstion_date"`
	ReturnDate           time.Time                  `json:"return_date"`
	Duration             int                        `json:"duration"`
	ManageAssetDetailReq []ManageAssetDetailRequest `json:"manage_asset_detail"`
}

type ManageAssetDetailRequest struct {
	Id            string
	IdManageAsset string
	IdAsset       string
	TotalItem     int
	Status        string
}
