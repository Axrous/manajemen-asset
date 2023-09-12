package model

import "time"

type Asset struct {
	Id          string
	Category  	Category
	AssetType 	TypeAsset
	Name        string
	Amount      int
	Status      string
	EntryDate   time.Time
	ImgUrl		string
}

type AssetRequest struct {
	Id 			string
	CategoryId 	string
	AssetTypeId	string
	Name 		string
	Amount 		int
	Status		string
	EntryDate 	time.Time
	ImgUrl 		string
}