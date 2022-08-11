package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (clnt *Client) SendRequest(method, url string) error {
	client := clnt.client
	requestUrl := "https://" + url

	req, err := http.NewRequest(method, requestUrl, nil)
	if err != nil {
		fmt.Printf("Error %s", err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(err)
	if err != nil {
		fmt.Printf("Error %s", err)
		return err
	}

	defer resp.Body.Close()
	fmt.Printf("Body : %s \n ", resp.Body)
	fmt.Printf("Response status : %s \n", resp.Status)
	return nil
}
