package model

import "time"

type Asset struct {
	Id          string `json:"id"`
	Category  	Category `json:"category"`
	AssetType 	TypeAsset `json:"assetType"`
	Name        string `json:"name"`
	Amount      int `json:"amount"`
	Status      string `json:"status"`
	EntryDate   time.Time `json:"entryDate"`
	ImgUrl		string `json:"imgUrl"`
}

type AssetRequest struct {
	Id 			string `json:"id"`
	CategoryId 	string `json:"categoryId"`
	AssetTypeId	string `json:"assetTypeId"`
	Name 		string `json:"name"`
	Amount 		int		`json:"amount"`
	Status		string `json:"status"`
	EntryDate 	time.Time
	ImgUrl 		string `json:"imgUrl"`
}