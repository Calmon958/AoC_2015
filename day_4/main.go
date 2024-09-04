package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func genMD5(str string) int {
	number := 1
	for {
		input := fmt.Sprintf("%s%d", str, number)
		wordHash := md5.New()
		wordHash.Write([]byte(input))
		hashBytes := wordHash.Sum(nil)
		hashString := hex.EncodeToString(hashBytes)
		// return hashString
if strings.HasPrefix(hashString, "000000"){
	return number
} 
number++		
	}


}

func genMD5Part1(str string) int {
	number := 1
	for {
		input := fmt.Sprintf("%s%d", str, number)
		wordHash := md5.New()
		wordHash.Write(([]byte(input)))
		hashBytes := wordHash.Sum(nil)
		hashString := hex.EncodeToString(hashBytes)
		if strings.HasPrefix(hashString, "00000") {
			return number
		}
		number ++ 
	}
}

func main() {
	fmt.Println(genMD5("yzbqklnj"))
	fmt.Println(genMD5("what do you mean"))
	fmt.Println(genMD5("hey there"))

}
