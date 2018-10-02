package fulfillment

type WebhookResponse struct {
	FulfillmentText string `json:"fulfillmentText,omitempty"`
}

type WebhookRequest struct {
	QueryResult QueryResult `json:"queryResult"`
}
type QueryResult struct {
	Action     string            `json:"action"`
	Parameters map[string]string `json:"parameters"`
}
