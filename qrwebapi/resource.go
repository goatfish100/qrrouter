package main

type Resource struct {
	Id          int    `json:"id" bson:"_id,omitempty"`
	Description string `json:"description"`
	Protected   bool   `json:"protected"`
	Action      string `json:"action"`
	Address     string `json:"address"`
}

type QRResource []Resource
