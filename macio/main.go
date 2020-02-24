// Copyright 2020
// Author: Venunadh Veeralanka(vvnadh@gmail.com)

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//HeaderBased Authentication
//https://api.macaddress.io/v1?output=json&search=44:38:39:ff:ef:57
//API Key via X-Authentication-Token
func main() {

	baseURL := "https://api.macaddress.io/v1?"
	apiKey := flag.String("apikey", "", "a string : Mandatory")
	macAddress := flag.String("macaddress", "44:38:39:ff:ef:57", "a string")
	flag.Parse()

	if len(*apiKey) == 0 {
		fmt.Println("ERROR: Mandatory flag apikey is missing")
		os.Exit(-1)
	}

	outputParam := "output=json"
	searchParam := "&search=" + *macAddress

	urlParts := []string{baseURL, outputParam, searchParam}
	url := strings.Join(urlParts, "")

	fmt.Println("DEBUG: Making HTTP Get for ", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Authentication-Token", *apiKey)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("ERROR: while querying {} there is an error:{}", baseURL, err.Error)
		panic(err)
		return
	}
	defer resp.Body.Close()
	resB, err := ioutil.ReadAll(resp.Body)
	s := string(resB)
	if err != nil {
		fmt.Printf("ERROR: Error while reading body\n\n", s)
		panic(err)
	} else {
		fmt.Printf("DEBUG: Response body=%v\n\n", s)
	}

	parseJSONByteSlice(resB)
}

/*
//Query-based Authentication
//https://api.macaddress.io/v1?apiKey=XYZ&output=json&search=44:38:39:ff:ef:57
func main1() {

	baseURL := "https://api.macaddress.io/v1?"

	apiKey := flag.String("apikey", "at_KXsmxC9yYo8LbjfcokUHW442TlFm5", "a string")
	macAddress := flag.String("macaddress", "44:38:39:ff:ef:57", "a string")
	flag.Parse()

	apiKeyParam := "apikey=" + *apiKey
	outputParam := "&output=json"
	searchParam := "&search=" + *macAddress

	urlParts := []string{baseURL, apiKeyParam, outputParam, searchParam}
	url := strings.Join(urlParts, "")

	fmt.Println("DEBUG: Making HTTP Get for ", url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: while querying {} there is an error:{}", baseURL, err.Error)
		return
	}
	defer resp.Body.Close()
	resB, _ := ioutil.ReadAll(resp.Body)
	s := string(resB)
	fmt.Printf("DEBUG: Response body=%v\n\n", s)
	parseJSONByteSlice(resB)
}*/

func parseJSONByteSlice(b []byte) {
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("ERROR: Got error while unmarshaling JSON:", err)
	}
	handleJSON(f)
}
func handleJSON(f interface{}) {
	m := f.(map[string]interface{})
	for k, v := range m {
		//fmt.Printf("---- call handleType k=%v \n v=%v \n type=%T \n", k, k, v, v)
		handleJSONTypes(k, v)
	}

}

func handleJSONTypes(k string, v interface{}) {

	//fmt.Printf("\n---- Inside handleType key = %v \n value=%v \n valueType=%T \n", k, v, v)
	switch vv := v.(type) {
	case []interface{}:
		//fmt.Println(k, "is an array:")
		for i, u := range vv {
			fmt.Printf("%v.%v=%v\n", k, i, u)
		}
	case map[string]interface{}:
		//fmt.Println("------k, is map[string]interface{}")
		m := v.(map[string]interface{})
		for k1, v1 := range m {
			//fmt.Printf("\n------k1=%v \n k1Type=%T\n v1=%v v1Type=%T\n", k1, k1, v1, v1)
			handleJSONTypes(k+"."+k1, v1)
		}
	default:
		fmt.Printf("%v=%v\n", k, vv)
	}
}
