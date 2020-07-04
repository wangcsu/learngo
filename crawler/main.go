package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		URL:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
