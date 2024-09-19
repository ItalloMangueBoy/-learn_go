package main

import "fmt"

func main() {
	user := map[string]string{
		"nome":       "Itallo",
		"soberenome": "Andre Quaresma",
	}

	delete(user, "soberenome")
	
	fmt.Println(user)
}
