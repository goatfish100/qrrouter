package main

var currentId int

var qrresource QRResource
var orgs Orgs

//var jsonsucess JsonSucess

// Give us some seed data
func init() {
	RepoCreateResource(Resource{Description: "yahoo", Protected: false, Action: "forward", Address: "http://www.yahoo.com"})
	RepoCreateResource(Resource{Description: "google", Protected: false, Action: "forward", Address: "https://www.google.com"})

	// OrgCreate(Org{Orgname: "", Address: "123 H Street",
	// 	City: "asdf", State: "lala",
	// 	Postalcode: "asddd",
	// 	Users: {Username: "1", Email: "2",
	// 		Name: "2", Password: "3"}})
	// fmt.Println("inside init")
	OrgCreate(Org{
		Orgname:    "Rest Holdings",
		Address:    "123 H street",
		City:       "Culver",
		State:      "CA",
		Postalcode: "84109",
		Users: []User{{
			Username: "freegyg",
			Email:    "freddy@yahoo.com",
			Name:     "Freddy G Spot",
			Password: "lsls",
		}, {
			Username: "toyo",
			Email:    "lsl@yahoo.com",
			Name:     "asdf",
			Password: "asdf",
		},
		},
	})
	OrgCreate(Org{
		Orgname:    "awake Holdings",
		Address:    "123 H street",
		City:       "Culver",
		State:      "CA",
		Postalcode: "84109",
		Users: []User{{
			Username: "freegyg",
			Email:    "freddy@yahoo.com",
			Name:     "Freddy G Spot",
			Password: "lsls",
		}, {
			Username: "toyo",
			Email:    "lsl@yahoo.com",
			Name:     "asdf",
			Password: "asdf",
		},
		},
	})

}

func RepoFindResource(id int) Resource {
	for _, t := range qrresource {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Resource{}
}

func OrgCreate(o Org) Org {
	// Insert Datas
	c := session.DB("resources").C("orgusers")
	err = c.Insert(o)

	if err != nil {
		panic(err)
	}
	orgs = append(orgs, o)
	return o
}

//this is bad, I don't think it passes race condtions
func RepoCreateResource(r Resource) Resource {
	currentId += 1
	r.Id = currentId
	qrresource = append(qrresource, r)
	return r
}

//this is bad, I don't think it passes race condtions
func RepoCreateOrg(org Org) Org {
	orgs = append(orgs, org)
	return org
}