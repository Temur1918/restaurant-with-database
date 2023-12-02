package main

import (
	"restaurant/api"
	"restaurant/files"
)

func main() {
	files.CreateFiles()
	api.Router()
}
