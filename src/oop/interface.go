package main

import "fmt"

type Person struct {
	Name string
}

type notifier interface {
	notice()
}

type informer interface {
	inform()
}

type tell interface {
	notice()
	inform()
}

func (p Person) notice() {
	fmt.Println(p.Name, "noticing")
}

func (p Person) inform() {
	fmt.Println(p.Name, "informing")
}
