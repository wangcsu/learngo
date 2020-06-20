package main

import (
	"fmt"
	"learngo/downloader/testing"
)

//func retrieve(url string) string {
//	resp, err := http.Get(url)
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//	bytes, _ := ioutil.ReadAll(resp.Body)
//	return string(bytes)
//}

func getRetriever() retriever {
	return testing.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
