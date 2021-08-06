package main

import "fmt"

func main() {
  
  // declare map
	countries := map[string]string{
		"uk": "United Kingdom",
		"us": "United states",
		"fr": "France",
		"ch": "China"}

  //print 
	fmt.Println(countries["uk"])
}
