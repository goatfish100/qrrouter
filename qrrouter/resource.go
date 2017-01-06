package main

type Resource struct {
	Id          string `json:"id" bson:"_id,omitempty"`
	Uuid        string `json:"uuid"`
	Description string `json:"Description"`
	Protected   string `json:"Protected"`
	Action      string `json:"Action"`
	Address     string `json:"Address"`
}
