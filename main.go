package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	client := http.Client{}

	for {
		err := compare(&client)
		if err != nil {
			fmt.Println(err.Error())
		}

		time.Sleep(time.Second * 100)
	}

}

func compare(client *http.Client) error {
	resp, err := client.Get("https://oklama.com/theheart")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	readData, err := os.ReadFile("dump.txt")
	if err != nil {
		readData = nil
	}

	if bytes.Compare(readData, bodyData) != 0 {

		// NEW PAGE!!!

		fmt.Println("NEW PAGE")
		os.WriteFile("dump.txt", bodyData, 0666)
	} else {
		fmt.Println("")
	}

	return nil
}
