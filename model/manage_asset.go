package model

import (
	"time"
)

type ManageAsset struct {
	Id             string
	User         UserCredentials `json:"omitempty"`
	Staff			Staff
	SubmissionDate time.Time
	ReturnDate     time.Time
	Detail	[]ManageDetailAsset
}

type ManageDetailAsset struct {
	Id 				string
	ManageAssetId 	string
	Asset 			Asset `json:"omitempty"`
	TotalItem 		int
	Status 			string
}

