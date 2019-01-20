package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Page struct {
	Title string
	Body  []byte //a []byte rather than string because that is the type expected by the io libraries we will use
}

//This method's signature reads: "This is a method named save that takes as its receiver p,
//a pointer to Page . It takes no parameters, and returns a value of type error."
func (p *Page) save() error {
	filename := p.Title + ".txt"
	readWritePerm := os.FileMode(0600) // Convert 0600 to filemode
	return ioutil.WriteFile(filename, p.Body, readWritePerm)
}

/*
Functions can return multiple values. The standard library function io.ReadFile returns
[]byte and error. In loadPage, error isn't being handled yet; the "blank identifier"
represented by the underscore (_) symbol is used to throw away the error
return value (in essence, assigning the value to nothing)
*/
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
