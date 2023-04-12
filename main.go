package main

import (
	"fmt"

	. "github.com/wellfrog/gotest1/generated/example"
)

func main() {
	persionValidate()
	numbericValidate()

}

func numbericValidate() {
	p := new(Numerics)
	// uint64 id = 1 [(validate.rules).uint64.gt = 999];
	// Numerics
	//All numeric types (float, double, int32, int64, uint32, uint64 , sint32, sint64, fixed32, fixed64, sfixed32, sfixed64) share the same rules.

	//const: the field must be exactly the specified value.

	// x must equal 1.23 exactly
	// float x1 = 1 [(validate.rules).float.const = 1.23];
	p.X1 = 1.23
	//lt/lte/gt/gte: these inequalities (<, <=, >, >=, respectively) allow for deriving ranges in which the field must reside.

	// x must be less than 10
	// int32 x2 = 2 [(validate.rules).int32.lt = 10];
	p.X2 = 5

	// x must be greater than or equal to 20
	// uint64 x3 = 3 [(validate.rules).uint64.gte = 20];
	p.X3 = 20

	// x must be in the range [30, 40)
	// fixed32 x4 = 4 [(validate.rules).fixed32 = {gte:30, lt: 40}];
	p.X4 = 35
	//Inverting the values of lt(e) and gt(e) is valid and creates an exclusive range.

	// x must be outside the range [30, 40)
	// double x5 = 5 [(validate.rules).double = {lt:30, gte:40}];
	p.X5 = 35.22
	//in/not_in: these two rules permit specifying allow/denylists for the values of a field.

	// x must be either 1, 2, or 3
	// uint32 x6 = 6 [(validate.rules).uint32 = {in: [1,2,3]}];
	p.X6 = 1

	// x cannot be 0 nor 0.99
	// float x7 = 7 [(validate.rules).float = {not_in: [0, 0.99]}];
	p.X7 = 0
	//ignore_empty: this rule specifies that if field is empty or set to the default value, to ignore any validation rules. These are typically useful where being able to unset a field in an update request, or to skip validation for optional fields where switching to WKTs is not feasible.

	// uint32 x8 = 8 [(validate.rules).uint32 = {ignore_empty: true, gte: 200}];//unint32
	// p.X8=1
	err := p.ValidateAll()

	//err := p.Validate() // err: nil
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("person no err")
	}
}
func persionValidate() {
	p := new(Person)

	//err := p.Validate() // err: Id must be greater than 999
	p.Id = 1000

	//err = p.Validate() // err: Email must be a valid email address
	p.Email = "example@bufbuild.com"

	//err = p.Validate() // err: Name must match pattern '^[^\d\s]+( [^\d\s]+)*$'
	//p.Name = "Protocol Buffer"
	p.Name = "!1a"

	//err = p.Validate() // err: Home is required
	p.Home = &Person_Location{Lat: 37.7, Lng: 999}

	//err = p.Validate() // err: Home.Lng must be within [-180, 180]
	p.Home.Lng = -122.4

	// x must be 1.23
	p.X5 = 1.23
	// x must be less than 10
	p.X7 = 2
	// x must be greater than or equal to 20
	p.X8 = 20
	// x must be in the range [30, 40)
	p.X9 = 35
	err := p.Validate() // err: nil
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("person no err")
	}
}
