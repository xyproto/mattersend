// Package for sending messages to a Mattermost host
package mattersend

import (
	"bytes"
	"github.com/xyproto/jpath"
	"net/http"
)

// Send a dictionary to Mattermost. Returns the HTTP status and err.
func Send(hookURL string, node *jpath.Node) (string, error) {
	// Render JSON
	jsonData, err := node.JSON()
	if err != nil {
		return "", err
	}

	// Build request
	client := &http.Client{}
	r, _ := http.NewRequest("POST", hookURL, bytes.NewReader(jsonData))
	//r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	r.Header.Add("Content-Type", "application/json")

	// Send request and return result
	resp, err := client.Do(r)
	return resp.Status, err
}
