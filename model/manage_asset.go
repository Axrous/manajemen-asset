package model

import (
	"time"
)

type ManageAsset struct {
	Id             string
	User         UserCredentials
	Staff			Staff
	SubmissionDate time.Time
	ReturnDate     time.Time
	Detail	[]ManageDetailAsset
}

type ManageDetailAsset struct {
	Id 				string
	ManageAsset 	ManageAsset
	Asset 			Asset
	TotalItem 		int
	Status 			string
}

