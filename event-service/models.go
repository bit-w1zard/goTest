package main

import "time"

type Event struct {
    ID             string    `json:"id"`
    CreatedAt      time.Time `json:"createdAt"`
    CreatedBy      string    `json:"createdBy"`
    IsDeleted      bool      `json:"isDeleted"`
    Type           string    `json:"type"`
    Contents       []Content `json:"contents"`
}

type Content struct {
    GTIN          string `json:"gtin"`
    Lot           string `json:"lot"`
    BestByDate    string `json:"bestByDate,omitempty"`
    ExpirationDate string `json:"expirationDate,omitempty"`
}
