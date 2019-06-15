package proxy

// Proxy is the data structure that will handle information
type Proxy struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Domain   string `json:"domain" bson:"domain,omitempty"`
	Weigth   string `json:"weight" bson:"weight"`
	Priority string `json:"priority" bson:"priority,omitempty"`
}
