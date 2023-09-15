package model

import "time"

type Staff struct {
	Nik_Staff    string    `json:"nik_staff"`
	Name         string    `json:"name"`
	Phone_number string    `json:"phone_number"`
	Address      string    `json:"address"`
	Birth_date   time.Time `json:"birth_date"`
	Img_url      string    `json:"img_url"`
	Divisi       string    `json:"divisi"`
}
