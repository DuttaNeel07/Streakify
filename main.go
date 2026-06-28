package main

import(
	"flag";
 	"fmt"
)

func scan(path string){
	fmt.Printf(path)
}

func stats(email string){
	fmt.Printf(email)
}

func main(){

	var email string
	var folder string

	flag.StringVar(&folder, "add", "", "Use this to add new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "abc@email.com", "Use this to add email to display stats")
	flag.Parse()

	if folder != "" {

		scan(folder)
		return
	}

	stats(email)

}