package main

import(
	"encoding/json"
	"fmt"
)

type Envelope struct { // Base struct

	Type 	string
	Msg 	interface{}

}

type Flags struct { // Flags Struct, used to pack flags into JSON to be used in response to THM Server POST request.

	Data 	[]string

}

type Delete struct {
	Data 	bool
}

func packFlags(arr *[]string, channel chan<- []byte) { // Array passed in using a pointer, call using packFlags(&flagArray)

	data := Envelope{
		Type: "response",
		Msg: Flags{
			Data: (*arr),
		},
	}

	var outputData []byte
	outputData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	channel <- outputData
	
}


func packDelete(state bool, channel chan<- []byte) {

	data := Envelope{
		Type: "delete",
		Msg: Delete{
			Data: state,
		},
	}

	var outputData []byte
	outputData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	channel <- outputData

}