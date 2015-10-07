package goreq_test

import (
	//	"bytes"
	//	"encoding/base64"
	"fmt"
	//	"io/ioutil"
	"net/http"
	//	"net/http/httptest"
	//	"net/url"
	//	"strings"
	//	"testing"
	//	"time"
	//
	//	"github.com/elazarl/goproxy"
	//	"encoding/json"
	"../goreq"
	"strings"
)

func ExampleGoReq_SetClient() {
	client := &http.Client{}
	resp, body, err := goreq.New().SetClient(client).
		Get("http://httpbin.org/get").
		End()

	fmt.Println(resp.StatusCode == 200)
	fmt.Println(len(err) == 0)
	fmt.Println(body != "")
	// Output:
	// true
	// true
	// true
}

func ExampleGoReq_Reset() {

	gr := goreq.New()
	gr.Get("http://httpbin.org/get").
		End()

	resp, body, err := gr.Reset().Get("http://httpbin.org/").
		End()

	fmt.Println(resp.StatusCode == 200)
	fmt.Println(len(err) == 0)
	fmt.Println(body != "")
	// Output:
	// true
	// true
	// true
}

func ExampleGoReq_Get() {
	resp, body, err := goreq.New().
		Get("http://httpbin.org/get").
		End()

	fmt.Println(resp.StatusCode == 200)
	fmt.Println(len(err) == 0)
	fmt.Println(body != "")
	// Output:
	// true
	// true
	// true
}

func ExampleGoReq_Post() {
	resp, body, err := goreq.New().
		Post("http://httpbin.org/post").
		SendRawString("hello world").
		End()

	fmt.Println(resp.StatusCode == 200)
	fmt.Println(len(err) == 0)
	fmt.Println(strings.Contains(body, "hello world"))
	// Output:
	// true
	// true
	// true
}

func ExampleGoReq_Head() {
	resp, body, err := goreq.New().
		Head("http://httpbin.org/headers").
		SendRawString("hello world").
		End()

	fmt.Println(resp.StatusCode == 200)
	fmt.Println(len(err) == 0)
	fmt.Println(body == "")
	// Output:
	// true
	// true
	// true
}

func ExampleGoReq_Put() {
	q := `{"Name":"Jerry"}`
	resp, body, err := goreq.New().
		Put("http://httpbin.org/put").
		ContentType("json").
		SendMapString(q).
		End()

	fmt.Println(resp.StatusCode == 200)
	fmt.Println(len(err) == 0)
	fmt.Println(strings.Contains(body, "Jerry"))
	// Output:
	// true
	// true
	// true
}

func ExampleGoReq_Delete() {
	q := `{"Name":"Jerry"}`
	resp, _, err := goreq.New().
		Delete("http://httpbin.org/delete").
		ContentType("json").
		SendMapString(q).
		End()

	fmt.Println(resp.StatusCode == 200)
	fmt.Println(len(err) == 0)
	// Output:
	// true
	// true
}

func ExampleGoReq_Patch() {
	q := `{"Name":"Jerry"}`
	resp, body, err := goreq.New().
		Patch("http://httpbin.org/patch").
		ContentType("json").
		SendMapString(q).
		End()

	fmt.Println(resp.StatusCode == 200)
	fmt.Println(len(err) == 0)
	fmt.Println(strings.Contains(body, "Jerry"))
	// Output:
	// true
	// true
	// true
}
