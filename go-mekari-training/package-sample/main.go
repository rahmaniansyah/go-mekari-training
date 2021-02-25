package main

import "github.com/rahmaniansyah/go-mekari-training/package-sample/animal"

func main(){
	b := animal.Bird{Name: "Raven"}
	b.Fly(true)
	b.Show()
}