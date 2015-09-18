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
	"github.com/smallnest/goreq"
)

func ExampleSetClient() {
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


