package repositories

import "gorm.io/gorm"

// แท็ก
type HotelIssues struct {
	gorm.Model
	HotelID     uint
	IssueCode   string `json:"issueCode"`
	IssueType   string `json:"issueType"`
	DateFrom    string `json:"dateFrom"`
	DateTo      string `json:"dateTo"`
	Order       int    `json:"order"`
	Alternative bool   `json:"alternative"`
}
