package request

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	request  *http.Request
	client   *http.Client
	response *http.Response
}

//CreateRequest create Request
//  method request method Get/Post/PUT...
//  url request url
//  params request params
func CreateRequest(method, url, params string) (*Request, error) {
	re, err := http.NewRequest(method, url, strings.NewReader(params))
	if err != nil {
		return nil, err
	}
	request := new(Request)
	request.request = re
	return request, nil
}

//Header set header
func (r *Request) Header(header http.Header) {
	r.request.Header = header
}

//SetHeader set header from key/value
func (r *Request) SetHeader(k, v string) {
	r.request.Header.Set(k, v)
}

//SetHeaders set header from map
func (r *Request) SetHeaders(headers map[string]string) {
	for k, v := range headers {
		r.SetHeader(k, v)
	}
}

//AddHeader add header from key/value
func (r *Request) AddHeader(k, v string) {
	r.request.Header.Add(k, v)
}

//AddHeaders add header from map
func (r *Request) AddHeaders(headers map[string]string) {
	for k, v := range headers {
		r.AddHeader(k, v)
	}
}

// SetCookies set Cookie
func (r *Request) SetCookies(cookies string) {
	r.request.Header.Set("Cookie", cookies)
}

// AddCookie Add Cookie
func (r *Request) AddCookie(cookie *http.Cookie) {
	r.request.AddCookie(cookie)
}

// AddCookies Add Cookie
func (r *Request) AddCookies(cookies []*http.Cookie) {
	for _, k := range cookies {
		r.AddCookie(k)
	}
}

//AddCookieFromNameValue Cookie from name/value Add
func (r *Request) AddCookieFromNameValue(name, value string) {
	c := &http.Cookie{
		Name:  name,
		Value: value,
	}
	r.AddCookie(c)
}

//AddCookiesFromMap Cookie from Map Add
func (r *Request) AddCookiesFromMap(c map[string]string) {
	cookies := make([]*http.Cookie, 0)
	for k, v := range c {
		cookie := &http.Cookie{
			Name:  k,
			Value: v,
		}
		cookies = append(cookies, cookie)
	}
	r.AddCookies(cookies)
}

//GetRequest get http request
func (r *Request) GetRequest() *http.Request {
	return r.request
}

//SetClient set http client
func (r *Request) SetClient(client *http.Client) {
	r.client = client
}

// Do request
func (r *Request) Do() error {
	if r.client == nil {
		r.client = http.DefaultClient
	}
	response, err := r.client.Do(r.request)
	if err != nil {
		return err
	}
	r.response = response
	return nil
}

//GetResponse get response
func (r *Request) GetResponse() *http.Response {
	return r.response
}

// GetBody get response body
func (r *Request) GetBody() ([]byte, error) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.response.Body)
	var reader io.Reader
	var err error
	if r.response.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(r.response.Body)
		if err != nil {
			return nil, err
		}
	} else {
		reader = r.response.Body
	}
	return ioutil.ReadAll(reader)
}

// ConvertHeader 将 map[string]string 转成 http.header
//   headers 可以为空
func ConvertHeader(header http.Header, headers map[string]string) http.Header {
	if header == nil {
		header = http.Header{}
	}
	for k, v := range headers {
		header[k] = []string{v}
	}
	return header
}

// SetGetParams Set Get Request Params
func SetGetParams(req *Request, params map[string]string) {
	query := req.GetRequest().URL.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	req.GetRequest().URL.RawQuery = query.Encode()
}
