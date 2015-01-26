package models

type Campaign struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Campaigns []Campaign
