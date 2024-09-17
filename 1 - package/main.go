package main

import (
	"app/hello"
	"fmt"

	"github.com/badoux/checkmail"
)



func main()  {
	hello.Say()

	err := checkmail.ValidateFormat("itallomkspop@gmail.com")

	fmt.Println(err)
	
}