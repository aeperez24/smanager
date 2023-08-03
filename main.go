package main

import (
	"smanager/server"
)

func main() {

	server.NewServer().Run(":8181")
}
