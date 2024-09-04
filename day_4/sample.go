package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)


// understanding of how MD
func md5Gen(str string) int {
	number := 1
	for{
		input := fmt.Sprintf("%s%d", str, number)
	hash := md5.New()          // blank page for writing the data
	hash.Write([]byte(input))    // writing the data in bytes
	hashBytes := hash.Sum(nil) // finalize the appending of data using the nil value
	// 16-byte slice representing the MD5 hash as raw bytes

	hashString := hex.EncodeToString(hashBytes)
	/*
	   	hex. is a package used to convert byte slices to hex strings

	   converts the hash byte slice to a string using EncodetoString(<hashByte slice>)
	*/
	if strings.HasPrefix(hashString, "00000"){
		return number
	}
	number++
	}
}

// func main() {
// 	a := fmt.Sprintf(md5Gen("abdce609043"))
// 	for i := 0; i < len(a); i++ {
// 		if a[i] >= 0 && a[i] < 5 {
// 			a[i] == 0
// 		} else {
// 			continue
// 		}
// 	}
// 	fmt.Println(a)
// }
