package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		log.Fatalln(err)
		//		panic(err)
	}
}

func foo() {
	fmt.Println("When os.Exit(1) is called, deferred functions don't run")
}

/*
... the Fatal functions call os.Exit(1) after
writing the log message ...
*/
