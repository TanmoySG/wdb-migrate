package retro

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TanmoySG/wdb-migrate/internal/requests"
)

const (
	GetDataAction    = "view-data"
	GetAllDataAction = "get-data"
	ContentType      = "application/json"
)

type Data map[string]interface{}

type Payload struct {
	Database   string  `json:"database"`
	Collection string  `json:"collection"`
	Marker     *string `json:"marker,omitempty"`
	Data       *Data   `json:"data,omitempty"`
}

type RequestBody struct {
	Action  string  `json:"action"`
	Payload Payload `json:"payload"`
}

type WdbRetroClient struct {
	httpClient    http.Client
	connectionURL string
}

func NewClient(baseURL, cluster, token string) WdbRetroClient {
	return WdbRetroClient{
		httpClient:    *http.DefaultClient,
		connectionURL: fmt.Sprintf("%s/connect?cluster=%s&token=%s", baseURL, cluster, token),
	}
}

func getError(responseBytes []byte) error {

	var resp map[string]interface{}
	err := json.Unmarshal(responseBytes, &resp)
	if err != nil {
		return nil
	}

	switch resp["status_code"].(string) {
	case "0":
		return fmt.Errorf(resp["response"].(string))
	case "1":
		return nil
	default:
		return nil
	}
}

func getMarker(key string, value string) string {
	return fmt.Sprintf("%s : %s", key, value)
}

func (w WdbRetroClient) GetData(database, collection string, marker ...string) (*GetDataResponse, error) {
	var markerString *string
	var action string = GetAllDataAction

	// filtered data feature
	// if len(marker) == 2 {
	// 	formattedMarkerString := getMarker(marker[0], marker[1])
	// 	markerString = &formattedMarkerString
	// 	action = GetDataAction
	// }

	requestBody := RequestBody{
		Action: action,
		Payload: Payload{
			Database:   database,
			Collection: collection,
			Marker:     markerString,
		},
	}

	responseBytes, err := requests.Query(w.httpClient, http.MethodPost, w.connectionURL, requestBody)
	if err != nil {
		return nil, err
	}

	err = getError(responseBytes)
	if err != nil {
		return nil, err
	}

	var getDataResponse GetDataResponse
	err = json.Unmarshal(responseBytes, &getDataResponse)
	if err != nil {
		return nil, err
	}

	return &getDataResponse, nil
}
