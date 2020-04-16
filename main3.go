package main

import (
	"fmt"
	//"reflect"
	"strconv"
	"crypto/sha1"
	//"encoding/base64"
	"encoding/hex"
)

func main(){

	kPlus := "01010101"
	ipad  := "11111110"
	var x string
	var answer string
	len1 := 8
	for i:= 0; i < len1; i++ {
		//fmt.Println(reflect.TypeOf(string(kPlus[i] ^ ipad[i])))
		//fmt.Print(answer)
		//fmt.Println(reflect.TypeOf(strconv.Itoa(int(kPlus[i]^ipad[i]))))
		//fmt.Print(x)
		 x = strconv.Itoa(int(kPlus[i]^ipad[i]))
		 answer = answer + x
	}
	
	//fmt.Print(answer)
	hasher := sha1.New()
	//sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	hasher.Write([]byte(answer))
    sha1_hash := hex.EncodeToString(hasher.Sum(nil))

	fmt.Print(sha1_hash)
}