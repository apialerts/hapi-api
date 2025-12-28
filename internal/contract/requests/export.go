package requests

type ExportRequest struct {
	Endpoints []EndpointRequest `json:"endpoints"`
}

type EndpointRequest struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	URL             string `json:"url"`
	Method          string `json:"method"`
	IsActive        bool   `json:"isActive"`
	SlowThresholdMs int64  `json:"slowThresholdMs"`
	CreatedAt       int64  `json:"createdAt"`
}
