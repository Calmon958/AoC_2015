# MD5 hashing

MD5 stands for Message Digest Algorithm which is a hash function which uses 128-bit hash value.It's mostly used in Cryptography since it's a cryptographic hash function.

MD5 takes in a message and changes it to a 16-byte fixed length.

## Working theory
MD5 works in the following way, the following are just examples to help relate with the code:

## 1. Create a writing pad
`input := md5.New()`
So basically we create a place to write out MD5 value. There is a package called `encoding` which holds the algorithm for getting the MD5 value of a given message.

## 2. Enter values
`input.Write([]byte(input))`
We write the values we get to input as a slice of byte.

## 3. Add the total values
`hashSum := hash.Sum(nil)`
Append the data which is raw 16 bytes data

## 4. Number
So esse