package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		m := make(map[string]int) 
    	for _, a := range words { 
        m[a] += 1 
    } 
	fmt.Println(m)
	}
}