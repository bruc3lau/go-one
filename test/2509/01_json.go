package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("test/2509/.json")
	if err != nil {
		panic(err)
	}
	var root map[string]interface{}
	err = json.Unmarshal(file, &root)
	if err != nil {
		panic(err)
	}
	for _, v := range root {
		if graphicList, ok := v.(map[string]interface{})["graphic_list"]; ok {
			for _, graphic := range graphicList.([]interface{}) {
				if graphic, ok := graphic.(map[string]interface{}); ok {
					//fmt.Println(idx, graphic)
					if property, ok := graphic["property"].(map[string]interface{}); ok {
						//fmt.Println(property)
						lot := property["parking_no"]
						fmt.Println(lot.(string))
					}
				}
			}
		}
	}
}
