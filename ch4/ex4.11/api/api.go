// Api package provides requests for create, read, update and close issues on
// GitHub.
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const APIURL = "https://api.github.com/" // By default, all requests receive the v3 of the REST API.

// get function send GET HTTP request to GitHub API and returns the response.
func get(url string) (*http.Response, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get failed: %s\n", resp.Status)
	}

	return resp, nil
}

// GetIssue gets one specific issue by calling get, decodes the body of response
// and returns a IssueList.
func GetIssue(owner, repo string, params map[string]string, issueN int) (*Issue, error) {
	url := APIURL + strings.Join([]string{"repos", owner, repo, "issues", strconv.Itoa(issueN)}, "/")
	url += "?"
	for k, v := range params {
		url += k + "=" + v
	}

	resp, err := get(url)
	if err != nil {
		return nil, err
	}

	var result Issue

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// GetIssues gets all issues of a GitHub repo, decodes the body of response and
// returns a IssueList.
func GetIssues(owner, repo string, params map[string]string) (*IssuesList, error) {
	url := APIURL + strings.Join([]string{"repos", owner, repo, "issues"}, "/")
	url += "?"
	for k, v := range params {
		url += k + "=" + v
	}

	resp, err := get(url)
	if err != nil {
		return nil, err
	}

	var result IssuesList

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func CreateIssue(owner, repo, title string) (*Issue, error) {
	url := APIURL + strings.Join([]string{"repos", owner, repo, "issues"}, "/")

	client := &http.Client{}
	fields := map[string]string{"title": title}
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(fields)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, err
	}

	if username == "" || password == "" {
		fmt.Print("Username: ")
		fmt.Scanf("%s", &username)
		fmt.Print("Password: ")
		fmt.Scanf("%s", &password)
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("fail to create issue: %s", resp.Status)

	}

	var result Issue

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil

}
