package model

import "time"

type Asset struct {
	Id          string `json:"id"`
	Category  	Category `json:"category,omitempty"`
	AssetType 	TypeAsset `json:"asset_type,omitempty"`
	Name        string `json:"name"`
	Available      int `json:"available,omitempty"`
	Total			int `json:"total,omitempty"`
	Status      string `json:"status,omitempty"`
	EntryDate   time.Time `json:"entryDate,omitempty"`
	ImgUrl		string `json:"imgUrl,omitempty"`
}

type AssetRequest struct {
	Id 			string `json:"id"`
	CategoryId 	string `json:"category_id"`
	AssetTypeId	string `json:"asset_type_id"`
	Name 		string `json:"name"`
	Total 		int		`json:"total"`
	Available	int		`json:"available"`
	Status		string `json:"status"`
	EntryDate 	time.Time
	ImgUrl 		string `json:"img_url"`
}