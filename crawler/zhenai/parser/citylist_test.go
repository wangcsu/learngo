package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://localhost:8080/mock/www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", contents)
	result := ParseCityList(contents)
	//fmt.Println(result.Items, result.Requests)
	// Verify result
	const resultSize = 470
	expectedURL := []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City 阿坝",
		"City 阿克苏",
		"City 阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedURL {
		if result.Requests[i].URL != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].URL)
		}
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d items, but had %d", resultSize, len(result.Items))
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but was %s", i, city, result.Items[i].(string))
		}
	}
}
