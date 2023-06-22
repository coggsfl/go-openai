package openai

import (
	"context"
	"net/http"
)

// CreateImage - API call to create an image. This is the main endpoint of the DALL-E API.
func (c *AzureClient) CreateAzureImage(ctx context.Context, request ImageRequest) (response ImageResponse, err error) {
	urlSuffix := "/images/generations"
	req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.fullAzureURL(urlSuffix), request)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
