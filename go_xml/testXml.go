package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func ErrHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func XmlTest() {
	content, err := ioutil.ReadFile("test.xml")
	ErrHandler(err)

	var result Reusult
	err2 := xml.Unmarshal(content, &result)
	ErrHandler(err2)

	fmt.Println("xml 解析的内容：")
	fmt.Println(result)

	data, err3 := json.Marshal(&result)
	ErrHandler(err3)
	strJson := string(data)
	fmt.Println("xml to json :")
	fmt.Println(strJson)
}

func main() {
	XmlTest()
}
