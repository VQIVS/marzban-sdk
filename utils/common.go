package utils

import "net/url"

// Convert string to net/url URL

func StringToURL(rawURL string) (*url.URL, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return parsedURL, nil
}
