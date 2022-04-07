package main

import "github.com/hengfeiyang/lsmdb/internal/server"

func main() {
	server.New().Run(":8080")
}
