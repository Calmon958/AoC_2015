package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	condition "santa/day_2/aspects"
)

func main() {
	file, err := os.Open("wrap.txt")
	if err != nil {
		fmt.Println("Error openning file:", err)
		return
	}
	rib := 0

	defer file.Close()
	scanner := bufio.NewScanner(file)
	area := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "x")
		if len(parts) == 3 {
			a, err1 := strconv.Atoi(parts[0])
			b, err2 := strconv.Atoi(parts[1])
			c, err3 := strconv.Atoi(parts[2])
			if err1 != nil || err2 != nil || err3 != nil {
				fmt.Println("Error converting string to int:", err1, err2, err3)
			}
			area += condition.Wrap(a, b, c)
			rib += condition.Ribbon(a, b, c)
		}
	}

	fmt.Println("Wrapping paper length:", area)
	fmt.Println("Ribbon length:", rib)
}
