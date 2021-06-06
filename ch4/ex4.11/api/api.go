// Api package provides requests for create, read, update and close issues on
// GitHub.
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const URL = "https://api.github.com" // By default, all requests receive the v3 of the REST API.

func Get(owner, repo string, issueNumber int) (*IssuesList, error) {
	var url = "/repos/" + owner + "/" + repo + "/issues"
	if issueNumber != 0 {
		url += "/" + strconv.Itoa(issueNumber)
	}
	resp, err := http.Get(URL + url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Get issues failed: %s", resp.Status)
	}

	var result IssuesList
	decoder := json.NewDecoder(resp.Body)

	if issueNumber != 0 {
		var singleResult Issue
		if err := decoder.Decode(&singleResult); err != nil {
			resp.Body.Close()
			return nil, err
		}
		result = append(result, singleResult)
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return nil, err
		}
	}

	resp.Body.Close()
	return &result, nil
}
