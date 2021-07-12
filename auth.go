package hashicups

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// SignIn - Get a new token for user
func (c *Client) SignIn(auth AuthStruct) (*AuthResponse, error) {
	rb, err := json.Marshal(auth)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signin", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

// SignOut - Revoke the token for a user
func (c *Client) SignOut() error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signout", c.HostURL), strings.NewReader(string("")))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "Signed out user" {
		return errors.New(string(body))
	}

	return nil
}
