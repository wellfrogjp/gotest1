package main

import (
	"fmt"

	. "github.com/wellfrog/gotest1/generated/example"
)

func main1() {
	p := new(Person)
	p.Id = 1000
	err := p.Validate() // err: Id must be greater than 999
	//fmt.Println(err)

	fmt.Println(err)
	err = p.Validate() // err: Email must be a valid email address
	p.Email = "example@@bufbuild.com"

	err = p.Validate() // err: Name must match pattern '^[^\d\s]+( [^\d\s]+)*$'
	p.Name = "Protocol Buffer"

	err = p.Validate() // err: Home is required
	p.Home = &Person_Location{Lat: 37.7, Lng: 999}

	err = p.Validate() // err: Home.Lng must be within [-180, 180]
	p.Home.Lng = -122.4

	err = p.Validate() // err: nil
	fmt.Println(err)
}
