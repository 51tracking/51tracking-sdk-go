package tracking51

import (
	"context"
	"net/http"
)

type Courier struct {
	CourierName            string      `json:"courier_name"`
	CourierCode            string      `json:"courier_code"`
	CourierCountryIso2     interface{} `json:"courier_country_iso2"`
	CourierUrl             string      `json:"courier_url"`
	CourierPhone           string      `json:"courier_phone"`
	CourierType            string      `json:"courier_type"`
	TrackingRequiredFields interface{} `json:"tracking_required_fields"`
	OptionalFields         interface{} `json:"optional_fields"`
	DefaultLanguage        string      `json:"default_language"`
	SupportLanguage        []string    `json:"support_language"`
	CourierLogo            string      `json:"courier_logo"`
}

func (client *Client) GetAllCouriers(ctx context.Context) (*Response, error) {
	var couriers []Courier
	response, err := client.sendApiRequest(ctx, http.MethodGet, "/couriers/all", nil, nil, &couriers)
	return response, err
}
