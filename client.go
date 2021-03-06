package gofire

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Client struct {
	intClient http.Client
	ssl       bool
	GeodeUrl  string
	Region    string
}

// Builds a new client for Geode.
// target: https://localhost:7070 (no trailing /)
// insecure: true to skip SSL check.
func NewClient(target string, insecure bool) (*Client, error) {
	client := &Client{
		GeodeUrl: target,
	}
	if strings.Contains(target, "https") && insecure {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.intClient.Transport = tr
	}
	statusCode, err, ok := client.Ping()
	if !ok {
		return &Client{}, errors.New(fmt.Sprintf("error testing connectivity to Geode. status code: %d, error: %s", statusCode, err))
	}
	return client, nil
}

// TODO (mxplusb): could prolly concatenate these into one. but I am suuuuper lazy right now yo.

func (cl Client) getRequestBuilder(finalPath string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s/%s",
		cl.GeodeUrl, BaseURIPath, CurrentAPIVersion, finalPath), nil)
	if err != nil {
		return &http.Request{}, err
	}
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("User-Agent", UserAgent)
	return req, nil
}

func (cl Client) headRequestBuilder(finalPath string) (*http.Request, error) {
	req, err := http.NewRequest("HEAD", fmt.Sprintf("%s%s%s/%s",
		cl.GeodeUrl, BaseURIPath, CurrentAPIVersion, finalPath), nil)
	if err != nil {
		return &http.Request{}, err
	}
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("User-Agent", UserAgent)
	return req, nil
}

func (cl Client) putRequestBuilder(finalPath string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s%s%s/%s",
		cl.GeodeUrl, BaseURIPath, CurrentAPIVersion, finalPath), bytes.NewBuffer(data))
	if err != nil {
		return &http.Request{}, err
	}
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("User-Agent", UserAgent)
	return req, nil
}

func (cl Client) deleteRequestBuilder(finalPath string) (*http.Request, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s%s/%s",
		cl.GeodeUrl, BaseURIPath, CurrentAPIVersion, finalPath), nil)
	if err != nil {
		return &http.Request{}, err
	}
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("User-Agent", UserAgent)
	return req, nil
}