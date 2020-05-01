package main

import (
	"fmt"
	"io/ioutil"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
	"os"
	"strings"
)

func readKing() []byte { // More NinjaJc01 code, ty bb

	buff, err := ioutil.ReadFile(kingPath)
	if err != nil {
		fmt.Println(err)
	}
	return buff

}

func getFlagArray(channel chan<- []string) { // Returns flag array to the /api/get endpoint

	encoded := make(chan []string) // Open return channel

	go generateFlags(flags, encoded) // Generate flags, place them into files as determined by the map file

	flagsChan := <- encoded	// Get generated flags

	channel <- flagsChan // Return flags


}

func generateFlags(amount int, channel chan<- []string) {

	hasher := md5.New() // Setup MD5 hasher
	rand.Seed(time.Now().UnixNano()) // Set the seed for the random strings
	var flags []string // Initialize flags slice

	for i := 1; i <= amount; i++ {
		hasher.Write([]byte(randomString(10))) // MD5 hash a 10 char string

		flag := "THM{" + hex.EncodeToString(hasher.Sum(nil)) + "}" // Wrap the MD5 hash in "THM{*flag*}"

		buff, err := os.Open(mapPath) // Open the map file to read
		if err != nil {
			fmt.Println(err)
		}
		defer buff.Close()

		rawBytes, err := ioutil.ReadAll(buff) // Read each line of the map file
		if err != nil {
			fmt.Println(err)
		}
		lines := strings.Split(string(rawBytes), "\n")

		for num, line := range lines {
			if num == i {
				writeFlag(flag, strings.TrimRight(line, "\r\n")) // Write the flag to the flag locations as deterrmined by the map file
			}
		}

		flags = append(flags, flag)	// Append generated flag to the flag slice
	}

	channel <- flags // Send flags slice over return channel

}

func randomInt(min, max int) int { // Random Integer function, used to generate a Random String
    return min + rand.Intn(max-min)
}

func randomString(len int) string {	// Random String function, used as food for the MD5 hasher
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(randomInt(65, 90))
    }
    return string(bytes)
}


func writeFlag(flag string, path string) { // Function that opens the flag files specified in the map file and write the flag to them

	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(flag)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func deleteMap(channel chan<- bool) { // Function that deletes the map file on /api/delete

	err := os.Remove(mapPath)
	if err != nil {
		fmt.Println(err)
		channel <- delCheck
	}

	delCheck = true
	channel <- delCheck

}