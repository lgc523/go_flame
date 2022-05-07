package main

import "github.com/flamego/flamego"

func simple() {
	f := flamego.Classic()
	f.Get("/", func() string {
		return "Hello Flamego!\n"
	})
	f.Run()
}
