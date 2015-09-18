package goreq

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
)

func ExampleSetClient() {
	client := &http.Client{}
	resp, body, err := New().SetClient(client).
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

