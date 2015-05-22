// Library to shorten URIs using http://is.gd
package goisgd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Shortens the URI using the "API" listed here: http://is.gd/api_info.php
func Shorten(uri string) (string, error) {
	// नये अनुप्रयोग क्रमादेशन अंतर्फलक का प्रयोग किया ।
	// @see http://is.gd/apishorteningreference.php
	// Used the new API.
	// @see http://is.gd/apishorteningreference.php
	u := "http://is.gd/create.php?format=simple&url=" + url.QueryEscape(uri)

	response, err := http.Get(u)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	shortUri := string(body)

	// Make sure we get a 200 response code, otherwise,
	// return the error message returned by is.gd
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf(shortUri)
	}

	return shortUri, nil
}
