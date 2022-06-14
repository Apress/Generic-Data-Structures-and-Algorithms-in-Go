package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func main() {
	name1 := "Richard"
	name2 := "Richards"

	md5hash := md5.Sum([]byte(name1))
	sha256hash := sha256.Sum256([]byte(name1))
	fmt.Println("   MD5: ", md5hash)
	fmt.Println("SHA256: ", sha256hash)

	md5hash = md5.Sum([]byte(name2))
	sha256hash = sha256.Sum256([]byte(name2))
	fmt.Println("   MD5: ", md5hash)
	fmt.Println("SHA256: ", sha256hash)
}
/* Output
 MD5:  [197 28 139 189 158 140 139 196 144 66 204 213 211 233 134 77]
SHA256:  [29 235 10 59 134 117 13 14 74 76 33 220 150 1 115 105 84 174 92 202 198 84 197 127 61 69 86 58 31 89 89 152]
   MD5:  [166 13 63 148 118 202 25 29 165 242 21 183 0 101 165 76]
SHA256:  [107 180 140 197 199 134 66 52 247 101 104 172 63 77 46 205 135 103 147 106 45 109 84 183 195 48 107 144 11 99 127 198]
*/