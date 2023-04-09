package retro

type GetDataResponse struct {
	Data       map[string]interface{} `json:"data"`
	Schema     map[string]interface{} `json:"schema"`
	StatusCode string                 `json:"status_code"`
}
