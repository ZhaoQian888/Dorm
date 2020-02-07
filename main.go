package main

import "Dorm/conf"

func main() {
	router := conf.Start()
	router.Run()
}
