package engine

import (
	"encoding/json"
	"fmt"
	"github.com/oliveagle/jsonpath"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type httpNode struct {
	method      string                 // http method
	rawQuery    string                 // raw query
	host        string                 // host
	path        string                 // path
	scheme      string                 // scheme
	headers     string                 // header {"k":"v"}
	contentType string                 // Content-Type
	root        string                 // root source
	endpoint    string                 // resource endpoint
	postBody    string                 // post body
	args        map[string]interface{} // args
}

func (h *httpNode) do() (interface{}, error) {
	r, err := h.httpRequest(h.method)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func jsonPathWithRoot(rootPath string, body []byte) (interface{}, error) {
	var jsonData interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return nil, err
	}
	res, err := jsonpath.JsonPathLookup(jsonData, rootPath)
	if err != nil {
		fmt.Printf("jsonpath JsonPathLookup error = %v", err)
		return nil, err
	}
	return res, nil
}

func (h *httpNode) httpRequest(method string) (interface{}, error) {
	if method == "POST" {
		return h.postRequest(method)
	}
	if method == "GET" {
		return h.getRequest(method)
	}
	return nil, nil
}

func (h *httpNode) getRequest(method string) (interface{}, error) {
	client := &http.Client{}
	u, err := url.ParseRequestURI(h.endpoint)
	if err != nil {
		return nil, err
	}
	// raw query is len , use args
	if len(u.RawQuery) == 0 {
		var rawQuery []string
		for k, v := range h.args {
			rawQuery = append(rawQuery, fmt.Sprintf("%s=%s", k, v))
		}
		u.RawQuery = strings.Join(rawQuery, `&`)
	}
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		log.Printf("http new request error = %v", err)
		return nil, err
	}
	var headers map[string]string
	if err := json.Unmarshal([]byte(h.headers), &headers); err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Content-Type", "application/json")
	if len(h.contentType) != 0 {
		req.Header.Add("Content-Type", h.contentType)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("http get error = %v \n", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil ReadAll error = %v \n", err)
		return nil, err
	}
	return jsonPathWithRoot(h.root, body)
}

func (h *httpNode) postRequest(method string) (interface{}, error) {
	client := &http.Client{}
	u, err := url.ParseRequestURI(h.endpoint)
	if err != nil {
		return nil, err
	}
	postBody := "{}"
	if h.postBody != "" {
		postBody = h.postBody
	} else {
		postBodyByte, err := json.Marshal(h.args)
		if err != nil {
			return nil, err
		}
		postBody = string(postBodyByte)
	}
	payload := strings.NewReader(postBody)
	req, err := http.NewRequest(method, u.String(), payload)
	if err != nil {
		log.Printf("http new request error = %v", err)
		return nil, err
	}
	var headers map[string]string
	if err := json.Unmarshal([]byte(h.headers), &headers); err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Content-Type", "application/json")
	if len(h.contentType) != 0 {
		req.Header.Add("Content-Type", h.contentType)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("http get error = %v \n", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil ReadAll error = %v \n", err)
		return nil, err
	}
	return jsonPathWithRoot(h.root, body)
}
