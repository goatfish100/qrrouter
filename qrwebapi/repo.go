package main

var currentId int

var qrresource QRResource

// Give us some seed data
func init() {
	RepoCreateResource(Resource{Description: "yahoo", Protected: "false", Action: "forward", Address: "http://www.yahoo.com"})
	RepoCreateResource(Resource{Description: "google", Protected: "false", Action: "forward", Address: "https://www.google.com"})
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

//this is bad, I don't think it passes race condtions
func RepoCreateResource(r Resource) Resource {
	currentId += 1
	r.Id = currentId
	qrresource = append(qrresource, r)
	return r
}

// //this is bad, I don't think it passes race condtions
// func RepoCreateTodo(t Todo) Todo {
// 	currentId += 1
// 	t.Id = currentId
// 	todos = append(todos, t)
// 	return t
// }
