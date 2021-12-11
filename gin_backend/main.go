package main

import (
	"gin_backend/model"
	"gin_backend/route"
)

func main() {
	r := route.Init()
	model.ConnectDB()
	r.Run(":8000")
}
