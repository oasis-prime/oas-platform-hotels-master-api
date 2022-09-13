package googledm

import (
	"encoding/json"
	"errors"
)

type QueryautocompleteRequest struct {
	Input string `url:"input"`
	Key   string `url:"key"`
}

type QueryautocompleteResponse struct {
	Predictions []Predictions `json:"predictions"`
	Status      string        `json:"status"`
}

type Predictions struct {
	Description       string `json:"description"`
	MatchedSubstrings []struct {
		Length int `json:"length"`
		Offset int `json:"offset"`
	} `json:"matched_substrings"`
	PlaceID              string `json:"place_id"`
	Reference            string `json:"reference"`
	StructuredFormatting struct {
		MainText                  string `json:"main_text"`
		MainTextMatchedSubstrings []struct {
			Length int `json:"length"`
			Offset int `json:"offset"`
		} `json:"main_text_matched_substrings"`
		SecondaryText string `json:"secondary_text"`
	} `json:"structured_formatting"`
	Terms []struct {
		Offset int    `json:"offset"`
		Value  string `json:"value"`
	} `json:"terms"`
	Types []string `json:"types"`
}

func (r *QueryautocompleteResponse) BodyParser(inf interface{}) error {
	if r.Predictions == nil {
		return errors.New("cannot parse nil interface data")
	}
	b, err := json.Marshal(r.Predictions)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, inf)
	if err != nil {
		return err
	}

	return nil
}

type PlaceDetailsRequest struct {
	PlaceId string `url:"place_id"`
	Key     string `url:"key"`
	Fields  string `url:"fields"`
}

type PlaceDetailsResponse struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	Result           struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			Viewport struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
	} `json:"result"`
	Status string `json:"status"`
}
