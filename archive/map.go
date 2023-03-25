package main

import "fmt"

func main(){

// map creation  

// Option 1  
//  var colors map[string]string

// Option 2  
//  colors := make(map[string]string)


// Option 3  
  colors := map[string]string{
    "red":    "#ff000",
    "green":  "#4bf745",
  }  

// adding values on map
  colors["white"] = "#ffAAAW"
  colors["blue"] = "#ffAAAB"

// delete the value on map
  delete(colors,"white")
  
  fmt.Println(colors)
}
