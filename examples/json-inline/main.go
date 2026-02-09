package main

import (
	"encoding/json"
	"fmt"
)

type Base struct {
	Type        string `json:"type"`
	*Coordinate `json:",inline"`
	*Name       `json:",inline"`
}

type Coordinate struct {
	Type string `json:"type"`

	X int `json:"x"`
	Y int `json:"y"`
}

type Name struct {
	Type string `json:"type"`

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	data1 := []byte(`{
		"type": "coordinate",
		"x": 1,
		"y": 2
	}`)
	data2 := []byte(`{
		"type": "name",
		"firstName": "John",
		"lastName": "Doe"
	}`)

	for _, data := range [][]byte{data1, data2} {
		var base Base
		err := json.Unmarshal(data, &base)
		fmt.Printf("Base: %+v %v\n", base, err)

		switch true {
		case base.Coordinate != nil:
			var coordinate Coordinate
			err := json.Unmarshal(data, &coordinate)
			fmt.Printf("Coordinate: %+v %v\n", coordinate, err)
		case base.Name != nil:
			var name Name
			err := json.Unmarshal(data, &name)
			fmt.Printf("Name: %+v %v\n", name, err)
		}
	}

}
