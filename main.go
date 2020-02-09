package main

import (
	"Dorm/conf"
)

func main() {
	router := conf.Init()
	router.Run()
}
