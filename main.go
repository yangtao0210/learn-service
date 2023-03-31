package main

import "learn_code/db"

func main() {
	api := &db.Apis{
		ServiceName:   "service",
		FirstOperator: "jacktaoyang",
		IsDeleted:     0,
	}
	db.UpdateApi(api)
}
