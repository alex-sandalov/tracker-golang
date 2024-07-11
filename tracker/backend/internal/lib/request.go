package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"tracker-app/backend/internal/http-server/response"
)

func RequestFormat(baseURL string, params url.Values) string {
	return baseURL + "?" + params.Encode()
}

// GET sends a GET request to the specified URL and returns the response body as a GetInfoRequest.
//
// Parameters:
// - ctx: The context for the HTTP request.
// - url: The URL to send the GET request to.
//
// Returns:
// - The response body as a GetInfoRequest.
// - An error if the request fails.
func GET(ctx context.Context, url string) (response.GetInfoResponse, error) {
	// Create a new HTTP request.
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	var info response.GetInfoResponse

	if err != nil {
		return info, err
	}

	// Send the HTTP request and get the response.
	res, err := http.DefaultClient.Do(req)

	if res.StatusCode != 200 {
		return info, fmt.Errorf("failed to get user info: %s", res.Status)
	}

	if err != nil {
		return info, err
	}

	// Close the response body after reading it.
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return info, err
	}

	// Unmarshal the response body into the GetInfoRequest struct.
	err = json.Unmarshal(body, &info)
	if err != nil {
		return info, err
	}

	return info, nil
}
