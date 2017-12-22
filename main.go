package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Integer ...
type Integer int32

func checkIf(i Integer) bool {
	return i > 10
}

func env2map(env []string) map[string]string {
	m := make(map[string]string, len(env))
	for _, s := range env {
		d := strings.Split(s, "=")
		m[d[0]] = d[1]
	}
	return m
}

func main() {
	var pwd, err = os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(pwd)

	var varMap = env2map(os.Environ())

	log.Println(varMap)

	// fmt.Println("Hello World")
	// hello()
	// fmt.Println(checkIf(11))

	// var num1 Integer = 123
	// var ptr = &num1

	// fmt.Printf("%x=%d", ptr, *ptr)
}

func hello() {
	fmt.Println("Hello World 2")
}
