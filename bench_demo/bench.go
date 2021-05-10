package bench_demo

type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:LastName`
	Email     string `json:Email`
}

var employees = []Employee{
	{
		ID:        1,
		FirstName: "Jacqui",
		LastName:  "Lambrecht",
		Email:     "jlambrecht0@tinyurl.com",
	},
	{
		ID:        2,
		FirstName: "Sterne",
		LastName:  "Ruffell",
		Email:     "sruffell1@bigcartel.com",
	},
	{
		ID:        3,
		FirstName: "Karlie",
		LastName:  "Lufkin",
		Email:     "klufkin2@ihg.com",
	},
	{
		ID:        4,
		FirstName: "Othelia",
		LastName:  "Farran",
		Email:     "ofarran3@opera.com",
	},
	{
		ID:        5,
		FirstName: "Thatcher",
		LastName:  "Brasener",
		Email:     "tbrasener4@amazon.de",
	},
}

func ChangeName(id int) []Employee {
	var emp = employees

	for i, _ := range emp {
		if emp[i].ID == id {
			emp[i].FirstName = "Tony"
			break
		}

	}
	return emp
}

// func ChangeName(id int) []Employee {
// 	var emp = &employees

// 	for _, v := range *emp {
// 		if v.ID == id {
// 			v.FirstName = "Tony"
// 			break
// 		}

// 	}
// 	return *emp
// }
