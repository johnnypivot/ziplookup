package zippopotam

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string, httpClient *http.Client) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

func (z *Client) Lookup(zip string) (*Place, error) {
	req, err := http.NewRequest(http.MethodGet, z.baseURL+zip, nil)
	if err != nil {
		return nil, err
	}

	res, err := z.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resp Response
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Places) == 0 {
		return nil, ErrNoResults{Zip: zip}
	}

	return &resp.Places[0], nil
}
