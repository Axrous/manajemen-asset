package model

import "time"

type Staff struct {
	Nik_Staff    string    `json:"nik_staff"`
	Name         string    `json:"name"`
	Phone_number string    `json:"phone_number,omitempty"`
	Address      string    `json:"address,omitempty"`
	Birth_date   time.Time `json:"birth_date,omitempty"`
	Img_url      string    `json:"img_url,omitempty"`
	Divisi       string    `json:"divisi,omitempty"`
}
