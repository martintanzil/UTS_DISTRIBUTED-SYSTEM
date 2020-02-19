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
		countWords := strings.Split(scanner.Text(), " ")
		m := make(map[string]int) 
    	for _, words := range countWords { 
        m[words] += 1 
    } 
	fmt.Println(m)
	}
}