package main

import "honnef.co/go/js/dom"

func main() {
	d := dom.GetWindow().Document()
	p := d.GetElementByID("pali").(*dom.HTMLInputElement)

	p.Value = "Hello Pali"
}
