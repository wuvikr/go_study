package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// func method() {
// 	response, err := http.Get("http://example.com")

// 	response2, err2 := http.Post("http://example.com/upload", "image/jpeg", &buf)

// 	response2, err2 := http.Post("http://example.com/upload", "image/jpeg", &buf)

// }

func get() {
	response, err := http.Get("https://www.liwenzhou.com/")
	if err != nil {
		fmt.Printf("get method exec failed, err: %v\n", err)
		return
	}

	// 程序使用完response后必须关闭回复的主体。
	defer response.Body.Close()

	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Printf("read from response.Body failed, err: %v\n", err2)
		return
	}
	fmt.Printf("string(body): %v\n", string(body))
}

func getWithParam() {
	apiUrl := "http:127.0.0.1:9090/get"

	data := url.Values{}
	data.Set("name", "zhangsan")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err: %v\n", err)
	}

	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())

	resp, err2 := http.Get((u.String()))
	if err2 != nil {
		fmt.Printf("post failed, err: %v\n", err2)
		return
	}

	defer resp.Body.Close()

	body, err3 := ioutil.ReadAll(resp.Body)

}

func main() {
	get()

}
