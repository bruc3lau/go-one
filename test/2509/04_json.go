package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X, Y      int
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Value     float64 `json:"value,omitempty"`
	Ts        *int64  `json:"ts,omitempty"`
}

func int64Ptr(i int64) *int64 {
	return &i
}

func main() {
	t := new(int64)
	*t = 200
	p := Point{
		X:         10,
		Y:         20,
		Longitude: 123.456,
		Latitude:  78.90,
		//Ts: t,
		Ts: int64Ptr(1625079600),
	}
	//p.Ts = new(int64)
	//*p.Ts = 1625079600
	fmt.Println(p)
	marshal, _ := json.Marshal(p)
	fmt.Println(string(marshal))
}
