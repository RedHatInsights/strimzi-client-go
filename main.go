package main

import (
	"encoding/json"
	"fmt"
	"os"

	kafka "github.com/RedHatInsights/strimzi-client-go/apis/kafka.strimzi.io/v1beta2"
)

func main() {
	data, err := os.ReadFile("kafka.json")

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	k := kafka.KafkaList{}
	err = json.Unmarshal(data, &k)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}
}
