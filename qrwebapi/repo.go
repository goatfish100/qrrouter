package main

var currentId int

var qrresource QRResource
var orgs Orgs

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
		Id:         2,
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
}

// 	orgs := &Org{
// 		Id:         2,
// 		Orgname:    "Rest Holdings",
// 		Address:    "123 H street",
// 		City:       "Culver",
// 		State:      "CA",
// 		Postalcode: "84109",
// 		Users: []User{{
// 			Username: "freegyg",
// 			Email:    "freddy@yahoo.com",
// 			Name:     "Freddy G Spot",
// 			Password: "lsls",
// 		}, {
// 			Username: "toyo",
// 			Email:    "lsl@yahoo.com",
// 			Name:     "asdf",
// 			Password: "asdf",
// 		},
// 		},
// 	}
// 	fmt.Println(orgs)
// }

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
	currentId += 1
	o.Id = currentId

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
func RepoCreateOrg(org Org) Resource {
	currentId += 1
	r.Id = currentId
	orgs = append(orgs, org)
	return r
}

// //this is bad, I don't think it passes race condtions
// func RepoCreateTodo(t Todo) Todo {
// 	currentId += 1
// 	t.Id = currentId
// 	todos = append(todos, t)
// 	return t
// }
