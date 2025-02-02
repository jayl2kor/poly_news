package polymarket

import "time"

type PolyMarketResponse interface {
	EventListResponse | SummaryResponse
}

type EventListResponse struct {
	Data []struct {
		ID               string    `json:"id"`
		Ticker           string    `json:"ticker"`
		Slug             string    `json:"slug"`
		Title            string    `json:"title"`
		Description      string    `json:"description"`
		ResolutionSource string    `json:"resolutionSource,omitempty"`
		StartDate        time.Time `json:"startDate"`
		CreationDate     time.Time `json:"creationDate"`
		EndDate          time.Time `json:"endDate"`
		Image            string    `json:"image"`
		Icon             string    `json:"icon"`
		Active           bool      `json:"active"`
		Closed           bool      `json:"closed"`
		Archived         bool      `json:"archived"`
		New              bool      `json:"new"`
		Featured         bool      `json:"featured"`
		Restricted       bool      `json:"restricted"`
		Liquidity        float64   `json:"liquidity"`
		Volume           float64   `json:"volume"`
		OpenInterest     int       `json:"openInterest"`
		SortBy           string    `json:"sortBy,omitempty"`
		CreatedAt        time.Time `json:"createdAt"`
		UpdatedAt        time.Time `json:"updatedAt"`
		Competitive      float64   `json:"competitive"`
		Volume24Hr       float64   `json:"volume24hr"`
		EnableOrderBook  bool      `json:"enableOrderBook"`
		LiquidityClob    float64   `json:"liquidityClob"`
		NegRisk          bool      `json:"negRisk,omitempty"`
		NegRiskMarketID  string    `json:"negRiskMarketID,omitempty"`
		CommentCount     int       `json:"commentCount"`
		Markets          []struct {
			ID                           string    `json:"id"`
			Question                     string    `json:"question"`
			ConditionID                  string    `json:"conditionId"`
			Slug                         string    `json:"slug"`
			ResolutionSource             string    `json:"resolutionSource"`
			EndDate                      time.Time `json:"endDate"`
			Liquidity                    string    `json:"liquidity"`
			StartDate                    time.Time `json:"startDate"`
			Image                        string    `json:"image"`
			Icon                         string    `json:"icon"`
			Description                  string    `json:"description"`
			Outcomes                     string    `json:"outcomes"`
			OutcomePrices                string    `json:"outcomePrices"`
			Volume                       string    `json:"volume"`
			Active                       bool      `json:"active"`
			Closed                       bool      `json:"closed"`
			MarketMakerAddress           string    `json:"marketMakerAddress"`
			CreatedAt                    time.Time `json:"createdAt"`
			UpdatedAt                    time.Time `json:"updatedAt"`
			New                          bool      `json:"new"`
			Featured                     bool      `json:"featured"`
			SubmittedBy                  string    `json:"submitted_by"`
			Archived                     bool      `json:"archived"`
			ResolvedBy                   string    `json:"resolvedBy"`
			Restricted                   bool      `json:"restricted"`
			GroupItemTitle               string    `json:"groupItemTitle"`
			GroupItemThreshold           string    `json:"groupItemThreshold"`
			QuestionID                   string    `json:"questionID"`
			EnableOrderBook              bool      `json:"enableOrderBook"`
			OrderPriceMinTickSize        float64   `json:"orderPriceMinTickSize"`
			OrderMinSize                 int       `json:"orderMinSize"`
			VolumeNum                    float64   `json:"volumeNum"`
			LiquidityNum                 float64   `json:"liquidityNum"`
			EndDateIso                   string    `json:"endDateIso"`
			StartDateIso                 string    `json:"startDateIso"`
			HasReviewedDates             bool      `json:"hasReviewedDates"`
			Volume24Hr                   float64   `json:"volume24hr"`
			ClobTokenIds                 string    `json:"clobTokenIds"`
			UmaBond                      string    `json:"umaBond"`
			UmaReward                    string    `json:"umaReward"`
			Volume24HrClob               float64   `json:"volume24hrClob"`
			VolumeClob                   float64   `json:"volumeClob"`
			LiquidityClob                float64   `json:"liquidityClob"`
			AcceptingOrders              bool      `json:"acceptingOrders"`
			NegRisk                      bool      `json:"negRisk"`
			NegRiskMarketID              string    `json:"negRiskMarketID"`
			NegRiskRequestID             string    `json:"negRiskRequestID"`
			Ready                        bool      `json:"ready"`
			Funded                       bool      `json:"funded"`
			AcceptingOrdersTimestamp     time.Time `json:"acceptingOrdersTimestamp"`
			Cyom                         bool      `json:"cyom"`
			Competitive                  float64   `json:"competitive"`
			PagerDutyNotificationEnabled bool      `json:"pagerDutyNotificationEnabled"`
			Approved                     bool      `json:"approved"`
			ClobRewards                  []struct {
				ID               string `json:"id"`
				ConditionID      string `json:"conditionId"`
				AssetAddress     string `json:"assetAddress"`
				RewardsAmount    int    `json:"rewardsAmount"`
				RewardsDailyRate int    `json:"rewardsDailyRate"`
				StartDate        string `json:"startDate"`
				EndDate          string `json:"endDate"`
			} `json:"clobRewards"`
			RewardsMinSize      int     `json:"rewardsMinSize"`
			RewardsMaxSpread    float64 `json:"rewardsMaxSpread"`
			Spread              float64 `json:"spread"`
			LastTradePrice      float64 `json:"lastTradePrice"`
			BestBid             float64 `json:"bestBid"`
			BestAsk             float64 `json:"bestAsk"`
			AutomaticallyActive bool    `json:"automaticallyActive"`
			ClearBookOnStart    bool    `json:"clearBookOnStart"`
			SeriesColor         string  `json:"seriesColor"`
			ShowGmpSeries       bool    `json:"showGmpSeries"`
			ShowGmpOutcome      bool    `json:"showGmpOutcome"`
			ManualActivation    bool    `json:"manualActivation"`
			NegRiskOther        bool    `json:"negRiskOther"`
			OneDayPriceChange   float64 `json:"oneDayPriceChange,omitempty"`
		} `json:"markets"`
		Tags []struct {
			ID          string    `json:"id"`
			Label       string    `json:"label"`
			Slug        string    `json:"slug"`
			CreatedAt   time.Time `json:"createdAt"`
			ForceShow   bool      `json:"forceShow,omitempty"`
			PublishedAt string    `json:"publishedAt,omitempty"`
			CreatedBy   int       `json:"createdBy,omitempty"`
			UpdatedBy   int       `json:"updatedBy,omitempty"`
			UpdatedAt   time.Time `json:"updatedAt,omitempty"`
			ForceHide   bool      `json:"forceHide,omitempty"`
		} `json:"tags"`
		Cyom                bool      `json:"cyom"`
		ShowAllOutcomes     bool      `json:"showAllOutcomes"`
		ShowMarketImages    bool      `json:"showMarketImages"`
		EnableNegRisk       bool      `json:"enableNegRisk"`
		AutomaticallyActive bool      `json:"automaticallyActive,omitempty"`
		StartTime           time.Time `json:"startTime,omitempty"`
		GmpChartMode        string    `json:"gmpChartMode,omitempty"`
		NegRiskAugmented    bool      `json:"negRiskAugmented"`
		CountryName         string    `json:"countryName,omitempty"`
		ElectionType        string    `json:"electionType,omitempty"`
		Color               string    `json:"color,omitempty"`
		Live                bool      `json:"live,omitempty"`
		Ended               bool      `json:"ended,omitempty"`
		Series              []struct {
			ID              string    `json:"id"`
			Ticker          string    `json:"ticker"`
			Slug            string    `json:"slug"`
			Title           string    `json:"title"`
			SeriesType      string    `json:"seriesType"`
			Recurrence      string    `json:"recurrence"`
			Image           string    `json:"image"`
			Icon            string    `json:"icon"`
			Layout          string    `json:"layout"`
			Active          bool      `json:"active"`
			Closed          bool      `json:"closed"`
			Archived        bool      `json:"archived"`
			New             bool      `json:"new"`
			Featured        bool      `json:"featured"`
			Restricted      bool      `json:"restricted"`
			PublishedAt     string    `json:"publishedAt"`
			CreatedBy       string    `json:"createdBy"`
			UpdatedBy       string    `json:"updatedBy"`
			CreatedAt       time.Time `json:"createdAt"`
			UpdatedAt       time.Time `json:"updatedAt"`
			CommentsEnabled bool      `json:"commentsEnabled"`
			Competitive     string    `json:"competitive"`
			Volume24Hr      int       `json:"volume24hr"`
			Volume          float64   `json:"volume"`
			Liquidity       float64   `json:"liquidity"`
			StartDate       time.Time `json:"startDate"`
			CommentCount    int       `json:"commentCount"`
		} `json:"series,omitempty"`
		SeriesSlug string `json:"seriesSlug,omitempty"`
	} `json:"data"`
	Pagination struct {
		HasMore bool `json:"hasMore"`
	} `json:"pagination"`
}

type SummaryResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Created int    `json:"created"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Citations []string `json:"citations"`
	Object    string   `json:"object"`
	Choices   []struct {
		Index        int    `json:"index"`
		FinishReason string `json:"finish_reason"`
		Message      struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Delta struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"delta"`
	} `json:"choices"`
}
