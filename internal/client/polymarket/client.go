package polymarket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Tag string

const (
	TagPolitics       Tag = "politics"
	TagCrypto         Tag = "crypto"
	TagGlobalElection Tag = "GlobalElection"
)

type PolyMarketClient interface {
	GetPoliticFromPolyMarket(ctx context.Context, tag Tag, limit, offset int)
	GetSummaryForEvent(ctx context.Context, ticker string) (string, error)
}
type Client struct {
	baseURL    string
	summaryURL string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) ListEventsFromPolyMarket(_ context.Context, tag Tag, limit, offset int) (EventListResponse, error) {
	queries := fmt.Sprintf("limit=%d&active=true&archived=false&tag_slug=%s&closed=false&order=volume24hr&ascending=false&offset=%d", limit, tag, offset)
	endpoint := fmt.Sprintf("%s%s?%s", c.baseURL, "/events/pagination", queries)
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return EventListResponse{}, err
	}
	return doHttpRequest[EventListResponse](c.httpClient, request)
}

func (c *Client) GetSummaryForEvent(_ context.Context, ticker string) (SummaryResponse, error) {
	queries := fmt.Sprintf("prompt=%s", ticker)
	endpoint := fmt.Sprintf("%s%s?%s", c.summaryURL, "/api/perplexity", queries)
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return SummaryResponse{}, err
	}

	return doHttpRequest[SummaryResponse](c.httpClient, request)
}

func doHttpRequest[T PolyMarketResponse](c *http.Client, req *http.Request) (T, error) {
	var data T

	resp, err := c.Do(req)
	if err != nil {
		return data, err
	}
	if resp.StatusCode != http.StatusOK {
		return data, errors.New("failed to get data from polymarket")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, errors.New("failed to unmarshal polymarket response")
	}

	if _, ok := any(data).(SummaryResponse); ok {
		lines := strings.Split(string(body), "data: ")
		body = []byte(lines[len(lines)-1])
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, errors.New("failed to unmarshal polymarket response")
	}

	return data, nil
}
