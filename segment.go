package awsxray

import (
	"sync"
)

type Segment struct {
	// user expected to protect changes
	sync.RWMutex
	Name      string  `json:"name,omitempty"`
	Type      string  `json:"type,omitempty"`
	Id        string  `json:"id,omitempty"`
	TraceId   string  `json:"trace_id,omitempty"`
	ParentId  string  `json:"parent_id,omitempty"`
	StartTime float64 `json:"start_time,omitempty"`
	EndTime   float64 `json:"end_time,omitempty"`
	HTTP      *HTTP   `json:"http,omitempty"`
	Error     bool    `json:"error,omitempty"`
	Fault     bool    `json:"fault,omitempty"`
}

type HTTP struct {
	Request  *Request  `json:"request,omitempty"`
	Response *Response `json:"response,omitempty"`
}

type Request struct {
	Method    string `json:"method,omitempty"`
	URL       string `json:"url,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
	ClientIP  string `json:"client_ip,omitempty"`
}

type Response struct {
	Status int `json:"status,omitempty"`
}
