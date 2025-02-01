package polymarket

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestClient_ListEventsFromPolyMarket(t *testing.T) {
	type fields struct {
		baseURL    string
		summaryURL string
		httpClient *http.Client
	}
	type args struct {
		ctx    context.Context
		tag    Tag
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    EventListResponse
		wantErr bool
	}{
		{
			name: "#1 Get politics events successfully",
			fields: fields{
				baseURL:    "https://gamma-api.polymarket.com",
				summaryURL: "https://polymarket.com",
				httpClient: &http.Client{
					Timeout: 10 * time.Second,
				},
			},
			args: args{
				ctx:    context.Background(),
				tag:    TagPolitics,
				limit:  20,
				offset: 0,
			},
			wantErr: false,
		},
		{
			name: "#2 Get crypto events successfully",
			fields: fields{
				baseURL:    "https://gamma-api.polymarket.com",
				summaryURL: "https://polymarket.com",
				httpClient: &http.Client{
					Timeout: 10 * time.Second,
				},
			},
			args: args{
				ctx:    context.Background(),
				tag:    TagCrypto,
				limit:  20,
				offset: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
			}
			_, err := c.ListEventsFromPolyMarket(tt.args.ctx, tt.args.tag, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListEventsFromPolyMarket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_GetSummaryForEvent(t *testing.T) {
	type fields struct {
		baseURL    string
		summaryURL string
		httpClient *http.Client
	}
	type args struct {
		ctx    context.Context
		ticker string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    SummaryResponse
		wantErr bool
	}{
		{
			name: "Get summary successfully",
			fields: fields{
				baseURL:    "https://gamma-api.polymarket.com",
				summaryURL: "https://polymarket.com",
				httpClient: &http.Client{},
			},
			args: args{
				ctx:    context.Background(),
				ticker: "germany-parliamentary-election",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				summaryURL: tt.fields.summaryURL,
				httpClient: tt.fields.httpClient,
			}
			got, err := c.GetSummaryForEvent(tt.args.ctx, tt.args.ticker)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSummaryForEvent() error = %v, wantErr %v", err, tt.wantErr)
				return

			}

			log.Println(got)
		})
	}
}
