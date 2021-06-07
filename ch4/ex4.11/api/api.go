// Api package provides requests for create, read, update and close issues on
// GitHub.
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const APIURL = "https://api.github.com/" // By default, all requests receive the v3 of the REST API.

// get function send GET HTTP request to GitHub API and returns the response.
func get(url string) (*http.Response, error) {

	resp, err := http.Get(APIURL + url)
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
func GetIssue(owner, repo string, issueN int) (*Issue, error) {
	url := strings.Join([]string{"repos", owner, repo, "issues", strconv.Itoa(issueN)}, "/")

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
func GetIssues(owner, repo string) (*IssuesList, error) {
	url := strings.Join([]string{"repos", owner, repo, "issues"}, "/")

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
