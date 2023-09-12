package model

import "time"

type Asset struct {
	ID          string
	Category  	Category
	AssetType 	AssetType
	Name        string
	Amount      int
	Status      string
	EntryDate   time.Time
	ImgUrl		string
}

type AssetRequest struct {
	ID 			string
	CategoryId 	string
	AssetTypeId	string
	Name 		string
	Amount 		string
	Status		string
	EntryDate 	time.Time
	ImgUrl 		string
}