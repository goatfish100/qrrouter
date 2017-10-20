// qrhelperdatastructs - QR Code Router/Forwarded in Go part of the QR Helper Application
// Copyright (C) 2016  James LaPointe
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
package datastructs

import "gopkg.in/mgo.v2/bson"

type Org struct {
	//Id         string `json:"id" bson:"_id,omitempty"`
	ID         bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Orgname    string        `json:"orgName"`
	Address    string        `json:"address"`
	City       string        `json:"city"`
	State      string        `json:"state"`
	Postalcode string        `json:"postalCode"`
	Users      Users
}

type User struct {
	Username string `json:"userName"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Orgs []Org

type Users []User
type Resource struct {
	//ID          string `json:"id" bson:"_id,omitempty"`
	ID          bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Uuid        string        `json:"uuid"`
	OrgId       string        `json:"orgid" bson:"orgid,omitempty"`
	Description string        `json:"description"`
	Email       string        `json:"email"`
	Name        string        `json:"name"`
	Action      string        `json:"action"`
	Address     string        `json:"address"`
	AccessCount int64         `json:"accessCount"`
}

type Resources []Resource

type JSONSuccess struct {
	Success string `json:"success"`
}
type ResourceSuccess struct {
	Success string `json:"success"`
	Uuid    string `json:"uuid"`
}

type JSONFailure struct {
	Failure string `json:"failure"`
	Reason  string `json:"reason"`
	Error   string `json:"error"`
}
