# goreq
[![GoDoc](https://godoc.org/github.com/smallnest/goreq?status.png)](http://godoc.org/github.com/smallnest/goreq) [![Drone Build Status](https://drone.io/github.com/smallnest/goreq/status.png)](https://drone.io/github.com/smallnest/goreq/latest) [![Go Report Card](http://goreportcard.com/badge/smallnest/goreq)](http://goreportcard.com/report/smallnest/goreq)

A Simplified Http Client. Its initial codes are cloned from [HttpRequest](https://github.com/parnurzeal/gorequest). I have refactored the codes and make it more friendly to programmers.  And some bugs are fixed and new features are added.

Major changes include:
- can send any string or bytes in body
- can set a shared client for concurrency
- won't clear setting actively
- ......

You can see the release notes for details.

## Installation

```sh
$ go get github.com/smallnest/goreq
```

## Usage
### why you should use goreq
[goreq](time.Now().Add(timeout)) comes from [gorequest]() but it added some new features and fixed some bugs.  The initial functions and major functions are from _gorequest_. Thanks to @parnurzeal and other contributors. Why have I not created pull requests to _gorequest_ and created a new repository? I want to refactor it a lot and add features quickly.

GoReq makes http thing more simple for you, using fluent styles to make http client more awesome. You can control headers, timeout, query parameters, binding response and others in one line:

Before

```go
client := &http.Client{
  CheckRedirect: redirectPolicyFunc,
}

req, err := http.NewRequest("GET", "http://example.com", nil)

req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
```

Using GoReq

```go
resp, body, errs := goreq.New().Get("http://example.com").
  RedirectPolicy(redirectPolicyFunc).
  SetHeader("If-None-Match", `W/"wyzzy"`).
  End()
```

### Http Methods
#### GET

```go
resp, body, err := goreq.New().Get("http://httpbin.org/get").End()
```

#### DELETE

```go
q := `{"Name":"Jerry"}`
resp, _, err := goreq.New().Delete("http://httpbin.org/delete").ContentType("json").SendMapString(q).End()
```

#### HEAD

```go
resp, body, err := goreq.New().Head("http://httpbin.org/headers").SendRawString("hello world").End()
```

#### POST

```go
resp, body, err := goreq.New().Post("http://httpbin.org/post").SendRawString("hello world").End()
```

#### PUT

```go
q := `{"Name":"Jerry"}`
resp, body, err := goreq.New().Put("http://httpbin.org/put").ContentType("json").SendMapString(q).End()
```

#### PATCH

```go
q := `{"Name":"Jerry"}`
resp, body, err := goreq.New().Patch("http://httpbin.org/patch").ContentType("json").SendMapString(q).End()
```

### Header
You can set one Header by:

```go

```

or set some headers by json:

```go

```

or set some headers by struct:

```go

```

### Proxy
In the case when you are behind proxy, GoRequest can handle it easily with Proxy func:

```go
request := goreq.New().Proxy("http://proxy:999")
resp, body, errs := request.Get("http://example-proxy.com").End()
```

_Socks5_ will be supported in future.

### Timeout
Timeout can be set in any time duration using time package:

```go
request := goreq.New().Timeout(2*time.Millisecond)
resp, body, errs:= request.Get("http://example.com").End()
```

Timeout func defines both dial + read/write timeout to the specified time parameter.

### SSL
### Basic Auth
To add a basic authentication header:

request := goreq.New().SetBasicAuth("username", "password") resp, body, errs := request.Get("[http://example-proxy.com").End(](http://example-proxy.com").End())

### Query Parameter
Query function accepts either json string or query strings which will form a query-string in url of GET method or body of POST method. For example, making "/search?query=bicycle&size=50x50&weight=20kg" using GET method:

```go
     goreq.New().
        Get("/search").
        Query(`{ "query": "bicycle" }`).
        Query(`{ "size": "50x50" }`).
        Query(`{ "weight": "20kg" }`).
        End()
```

It also support query string:

```go
      goreq.New().
        Get("/search").
        Query("query=bicycle&size=50x50").
        Query("weight=20kg").
        End()
```

even you can pass a struct:

```go
      qq := struct {
              Query1 string `json:"query1"`
              Query2 string `json:"query2"`
          }{
              Query1: "test1",
              Query2: "test2",
          }
      goreq.New().
        Get("/search").
        Query(qq).
        End()
```

`Param` can be used to set query value that contains ";" like _fields=f1;f2;f3_

### Request Body
For POST, PUT, PATCH, you can set content of request BODY.. It is convenient to BODY content.

#### JSON
You can use `SendMapString` or `SendStruct` to set JSON content. You should set content type to "application/json" by:

```go
goreq.New().Post("/user").ContentType("json")
```

or

```go
goreq.New().Post("/user").SetHeader("application/json")
```

GoReq will parse struct, json string or query string and  rebuild the json content:

```go
      type BrowserVersionSupport struct {
        Chrome string
        Firefox string
      }
      ver := BrowserVersionSupport{ Chrome: "37.0.2041.6", Firefox: "30.0" }
      goreq.New().
        Post("/update_version").
        SendStruct(ver).
        SendStruct(`{"Safari":"5.1.10"}`).
        End()
```

or

```go
      goreq.New().
        Post("/search").
        SendMapString("query=bicycle&size=50x50").
        SendMapString(`{ "wheel": "4"}`).
        End()
```

#### Form
If you set Content-Type as "application/x-www-form-urlencoded", GoReq rebuilds the below data into form style:

```go
      goreq.New().
        Post("/search").
        ContentType("form").
        SendMapString("query=bicycle&size=50x50").
        SendMapString(`{ "wheel": "4"}`).
        End()
```

#### Raw String
If you want upload XML or other plain text, you can use this method:

```go
        goreq.New().
        Post("/search").
        ContentType("text").
        SendRawString("hello world").
        End()
```

#### Raw Bytes
Even you can upload raw bytes:

```go
        goreq.New().
        Post("/search").
        ContentType("stream").
        SendRawBytes([]byte("hello world")).
        End()
```

### Bind Response Body
You can bind response body to a struct:

```go
    type Person struct {
        Name string
    }

    var friend Person
        _, _, err := goreq.New().Get(ts.URL).
        BindBody(&friend).
        End()
```

### Callback
GoReqalso supports callback function to handle response:

```go
func printStatus(resp goreq.Response, body string, errs []error){
  fmt.Println(resp.Status)
}
goreq.New().Get("http://example.com").End(printStatus)
```

### Debug
For deugging, GoReq leverages _httputil_ to dump details of every request/response. You can just use SetDebug to enable/disable debug mode and SetLogger to set your own choice of logger.

### Share Client
For concurrency, you can use a shared client in multiple GoReq instances:

```go
    sa := New().Get(ts.URL + case1_empty)
    sa.End()

    client := sa.Client;

    goreq.New().Get(ts.URL+case2_set_header).
    SetHeader("API-Key", "fookey").
    SetClient(client).
    End()
```

### Reset
You can reset GoReq and use it send another request. It only keep the client and reset other fields.

```go
goreq.New().Get("http://example.com").Reset()
```

### Retry
You can set a retry value and GoReq will retry until the value if it fails. So goreq sends request at most retry + 1 times.

```go
_, _, err := New().Get(ts.URL).
    Retry(3, 100, nil).
    End()
```

## License
goreq is MIT License.
