package main

import (
    "fmt"
	"strconv"
	"crypto/sha1"
	"encoding/hex"
)

//for step 1
func concat(key string, dif int) string {
	appender := "0"

	for dif !=0{
		key = appender + key
		dif--
	}

	return key
}

//for step2
func xor(kPlus string, ipad string) string{
	len1 := len(kPlus)
	var answer string
	var x string
	for i:= 0; i < len1; i++ {
        x = strconv.Itoa(int(kPlus[i]^ipad[i]))
		answer = answer + x
	}
	return answer
}
func main(){

	var difference int
	var key string
	var bits int
	var message, xor1, appended1, appended2 string
	ipad := "00110110"
	opad := "01011100"

	fmt.Print("Enter the message to be encrypted := ")
	fmt.Scanln(&message)

    fmt.Print("Enter your key (K) to be used in the algorithm := ")
	//user inputs the key
	fmt.Scanln(&key)

	fmt.Print("Enter the number of bits to be used for your K+ := ")
	//user inputs the number of bits
	fmt.Scanln(&bits)

	for bits>8 || bits<0{
		fmt.Println("Please enter a number between 0 and 9. Thank you. ")
		fmt.Print("Enter the number of bits to be used for your K+ := ")
		fmt.Scanln(&bits)
	}

	//Check if key is smaller than 8 bits
	if len(key) < bits{
		difference = bits - len(key)
		key = concat(key,difference)
	}
	xor1 = xor(key, ipad)
	
	//now append xor1 to the main message

	appended1 = xor1 + message

	hasher1 := sha1.New()
	hasher1.Write([]byte(appended1))
	sha1_hash := hex.EncodeToString(hasher1.Sum(nil))

	//find the xor between k+ and opad
	xor2:= xor(key, opad)
	appended2 = xor2 + sha1_hash 

	hasher2 := sha1.New()
	hasher2.Write([]byte(appended2))
	sha2_hash := hex.EncodeToString(hasher2.Sum(nil))
	
	fmt.Print("Your encrypted message is:= ")
	fmt.Println(sha2_hash)


}
