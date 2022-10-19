package main

import (
	"fmt"
)

func main() {

	names := []string{"shreeyash", "akhilesh", "sagar", "sanjay", "priti", "anil"}
    
	for index, name := range names {
		fmt.Println(index +1 ,"My name is "+ name)
	}

}