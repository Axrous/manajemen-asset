package model

import (
	"time"
)

type ManageAsset struct {
	Id             string
	User         UserCredentials `json:"user,omitempty"`
	Staff			Staff	`json:"staff,omitempty"`
	SubmissionDate time.Time `json:"submission_date,omitempty"`
	ReturnDate     time.Time `json:"return_date"`
	Detail	[]ManageDetailAsset `json:"detail,omitempty"`
}

type ManageDetailAsset struct {
	Id 				string	`json:"id,omitempty"`
	ManageAssetId 	string	`json:"id_manage_asset,omitempty"`
	Asset 			Asset	`json:"asset,omitempty"`
	TotalItem 		int		`json:"total_item,omitempty"`
	Status 			string	`json:"status,omitempty"`
}
