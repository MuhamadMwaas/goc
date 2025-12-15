package domain

import (
	"context"
	"encoding/json"
)

type Location struct {
	ID                   string      `json:"id"`
	Type                 string      `json:"type"`
	Names                Names       `json:"names"`
	Coordinates          Coordinates `json:"coordinates"`
	ParentID             string      `json:"parentId"`
	IsGovernorateCapital bool        `json:"isGovernorateCapital"`
	IsSubdistrictCenter  bool        `json:"isSubdistrictCenter"`
}

type Names struct {
	En string `json:"en"`
	Ar string `json:"ar"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (l *Location) ToJson() ([]byte, error) {
	return json.Marshal(l)
}

func FromJson(data []byte) (*Location, error) {
	var l Location
	err := json.Unmarshal(data, &l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

type LocationRepository interface {
	GetByID(ctx context.Context, id string) (*Location, error)
}
