package rally

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var baseURL = "https://rally1.rallydev.com/slm/webservice/v2.0"

func (r *Client) getRequest(path string, params url.Values, v interface{}) (*http.Response, error) {
	theURL := fmt.Sprintf("%s/%s", baseURL, path)
	if params != nil {
		theURL += "?" + params.Encode()
	}
	fmt.Println("GETing", theURL)
	req, err := http.NewRequest("GET", theURL, nil)
	if err != nil {
		return nil, err
	}

	if r.apikey != "" {
	    req.Header.Add("ZSESSIONID", r.apikey)
    } else {        
        req.SetBasicAuth(r.username, r.password)     
    }

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return resp, err
	}

	cnt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}
	resp.Body.Close()

	if Debug {
		fmt.Println("Status code", resp.Status)
		fmt.Println("Headers", resp.Header)
		fmt.Printf("Debug content: \n%s\n\n", string(cnt))
	}

	var operationResult struct {
		OperationResult struct {
			Errors   []string
			Warnings []string
		}
	}
	json.Unmarshal(cnt, &operationResult)
	if len(operationResult.OperationResult.Errors) > 0 {
		return resp, fmt.Errorf("OperationResult error: %s", strings.Join(operationResult.OperationResult.Errors, ", "))
	}

	if Debug {
		fmt.Println("Content to queryContent.json")
		ioutil.WriteFile("queryContent.json", cnt, 0644)
	}

	err = json.Unmarshal(cnt, v)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
