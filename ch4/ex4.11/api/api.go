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

	"github.com/toutane/gopl/ch4/ex4.11/auth"
	"github.com/toutane/gopl/ch4/ex4.11/cli"
	"github.com/toutane/gopl/ch4/ex4.11/util"
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

func CreateIssue(repo, title string) (*Issue, error) {

	if !auth.IsLogged() {
		return nil, fmt.Errorf("\nYou must be logged in to create an issue (run auth login).")

	}

	hosts, err := util.LoadHosts()
	if err != nil {
		return nil, err
	}

	username := hosts.GitHubUser
	token := hosts.GitHubAccessToken

	fmt.Printf("\nCreating issue in %s/%s", username, repo)

	body, err := cli.CreateBody()
	if err != nil {
		return nil, err
	}

	fields := map[string]string{"title": title, "body": body}

	url := APIURL + strings.Join([]string{"repos", username, repo, "issues"}, "/")

	client := &http.Client{}

	buf := &bytes.Buffer{}

	encoder := json.NewEncoder(buf)
	err = encoder.Encode(fields)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", "token "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("\n\nFail to create issue: %s", resp.Status)

	}

	var result Issue

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil

}

func UpdateIssue(repo string, number int) (*Issue, error) {

	if !auth.IsLogged() {
		return nil, fmt.Errorf("\nYou must be logged in to update an issue (run auth login).")

	}

	hosts, err := util.LoadHosts()
	if err != nil {
		return nil, err
	}

	username := hosts.GitHubUser
	token := hosts.GitHubAccessToken

	fmt.Printf("\nUpdating issue #%d in %s/%s", number, username, repo)

	issue, err := GetIssue(username, repo, map[string]string{"state": "open"}, number)
	if err != nil {
		return nil, err
	}

	title, err := cli.UpdateContent(issue.Title, "title")
	if err != nil {
		return nil, err
	}

	body, err := cli.UpdateContent(issue.Body, "body")
	if err != nil {
		return nil, err
	}

	url := APIURL + strings.Join([]string{"repos", username, repo, "issues", strconv.Itoa(number)}, "/")
	url = url[:len(url)]

	client := &http.Client{}

	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err = encoder.Encode(map[string]string{"title": title, "body": body})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", "token "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("\n\nFailed to update issue: %s", resp.Status)

	}

	var result Issue

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func ClosedIssue(repo string, number int) (*Issue, error) {

	if !auth.IsLogged() {
		return nil, fmt.Errorf("\nYou must be logged in to close an issue (run auth login).")

	}

	hosts, err := util.LoadHosts()
	if err != nil {
		return nil, err
	}

	username := hosts.GitHubUser
	token := hosts.GitHubAccessToken

	fmt.Printf("\nClosing issue #%d in %s/%s", number, username, repo)

	url := APIURL + strings.Join([]string{"repos", username, repo, "issues", strconv.Itoa(number)}, "/")
	url = url[:len(url)]

	client := &http.Client{}

	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err = encoder.Encode(map[string]string{"state": "closed"})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", "token "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("\n\nFailed to close issue: %s", resp.Status)

	}

	var result Issue

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
