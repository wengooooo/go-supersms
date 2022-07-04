package go_supersms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	// FivesimAPIEndpoint is the basepoint for 5sim API
	SupersmsAPIEndpoint = "https://www.supersms.ml/api"

	// ANY represents an "any" parameter
	ANY = "any"

	// VERSION is this API wrapper version
	VERSION = "1.0"
)

// Client will perform all the API-related tasks
type Client struct {
	APIKey   string
	Referral string
}

//NewClient get a new Client with a given APIKey
func NewClient(APIKey string) *Client {
	return &Client{APIKey: APIKey}
}

// makeGetRequest performs a simple get request with custom header and query values
func (c *Client) makeGetRequest(url string, queryValues *url.Values) (*http.Response, error) {
	// Creates a client
	client := &http.Client{}
	// Creates a request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Response{}, err
	}

	queryValues.Set("secret_key", c.APIKey)
	// Encode the query values (if any)
	req.URL.RawQuery = queryValues.Encode()

	return client.Do(req)
}

// GetUserInfo returns ID, Email, Balance and rating of the user in a single request
func (c *Client) GetCode(taskid int) (*CodeDetail, error) {

	queryValues := url.Values{}
	queryValues.Add("taskid", strconv.Itoa(taskid))

	// Make request
	resp, err := c.makeGetRequest(
		fmt.Sprintf("%s/getcode", SupersmsAPIEndpoint),
		&queryValues,
	)

	if err != nil {
		return &CodeDetail{}, err
	}

	// Check status code
	if resp.StatusCode != 200 {
		return &CodeDetail{}, fmt.Errorf("%s", resp.Status)
	}

	// Read request body
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return &CodeDetail{}, err
	}
	resp.Body.Close()

	// Unmarshal the body into a struct
	var info CodeDetail
	err = json.Unmarshal(r, &info)
	if err != nil {
		return &CodeDetail{}, err
	}

	return &info, nil
}

// GetEmail returns user's email
func (c *Client) ReleaseNumber(phone string) (*ReleaseDetail, error) {
	queryValues := url.Values{}
	queryValues.Add("phone", phone)

	// Make request
	resp, err := c.makeGetRequest(
		fmt.Sprintf("%s/getcode", SupersmsAPIEndpoint),
		&queryValues,
	)

	if err != nil {
		return &ReleaseDetail{}, err
	}

	// Check status code
	if resp.StatusCode != 200 {
		return &ReleaseDetail{}, fmt.Errorf("%s", resp.Status)
	}

	// Read request body
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return &ReleaseDetail{}, err
	}
	resp.Body.Close()

	// Unmarshal the body into a struct
	var info ReleaseDetail
	err = json.Unmarshal(r, &info)
	if err != nil {
		return &ReleaseDetail{}, err
	}

	return &info, nil
}

// BuyActivationNumber performs a "buy activation number" operation by selecting country, operator and product name
// and returns the operation information
func (c *Client) GetNumber(channel, country, pid string) (*NumberDetail, error) {

	// Check if any additional query values could be encapsulated
	queryValues := url.Values{}

	queryValues.Add("channel", channel)
	queryValues.Add("country", country)
	queryValues.Add("pid", pid)

	// Make request
	resp, err := c.makeGetRequest(
		fmt.Sprintf("%s/getnumber", SupersmsAPIEndpoint),
		&queryValues,
	)

	if err != nil {
		return &NumberDetail{}, err
	}

	// Check status code
	if resp.StatusCode != 200 {
		return &NumberDetail{}, fmt.Errorf("%s", resp.Status)
	}

	// Read request body
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return &NumberDetail{}, err
	}
	resp.Body.Close()

	// Unmarshal the body into a struct
	var numberDetail NumberDetail
	err = json.Unmarshal(r, &numberDetail)
	if err != nil {
		return &NumberDetail{}, err
	}

	return &numberDetail, nil
}
