package types

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const IP_API_URL = "https://api.ipify.org?format=text"

type IPublicAddres interface {
	Get() error
}

type PublicIp struct {
	Ip string
}

func (ip *PublicIp) Get() error {
	public_url := os.Getenv("APP_PUBLIC_IP_API_URL")
	if public_url == "" {
		public_url = IP_API_URL
	}

	// parse url
	_, err := url.Parse(public_url)
	if err != nil {
		return fmt.Errorf("invalid url: %s", public_url)
	}

	// Make the request
	res, err := http.Get(public_url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check the status code
	if res.StatusCode != 200 {
		return fmt.Errorf("status code error: code=%d status=%s", res.StatusCode, res.Status)
	}

	// Read the response
	ret, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Return the IP
	ip.Ip = string(ret)
	return nil
}
