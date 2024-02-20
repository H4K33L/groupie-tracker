package Groupie

/*import (
	"reflect"
	"fmt"
)*/

/*func UnitTest() bool {
	/*
	The function UnitTest is used to see if all 
	the function of api.go work well.
	--------------------------------------------
	input : none
	output : a boolean, false if the function 
	have problem and true if not.s
	--------------------------------------------
	

	// Unit test for GetApi function :
	result, err := GetApi("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Println(err)
		return false
	} else if reflect.TypeOf(result) == "api" {
		fmt.Println(reflect.TypeOf(result))
		return false
	}

	// Unit test for GetArtist function :
	result2, err := GetArtist(result.Artists)
	if err != nil {
		fmt.Println(err)
		return false
	} else if reflect.TypeOf(result2) == "" {
		fmt.Println(reflect.TypeOf(result2))
		return false
	}

	// Unit test for GetArtist function :
	result3, err := GetLinkInfos(result2[0].Relation)
	if err != nil {
		fmt.Println(err)
		return false
	} else if reflect.TypeOf(result3) == "information" {
		fmt.Println(reflect.TypeOf(result3))
		return false
	}

	// Unit test for Begin function :
	result4, err := Begin("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Println(err)
		return false
	} else if reflect.TypeOf(result4) == "" {
		fmt.Println(reflect.TypeOf(result4))
		return false
	}

	// all the test suceed
	return true
}*/