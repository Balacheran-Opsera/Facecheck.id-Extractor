package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// InfoResponse represents the InfoResponse schema from the OpenAPI specification
type InfoResponse struct {
	Has_credits_to_search bool `json:"has_credits_to_search,omitempty"`
	Is_online bool `json:"is_online,omitempty"`
	Remaining_credits int `json:"remaining_credits,omitempty"`
	Faces int `json:"faces,omitempty"`
}

// InputImage represents the InputImage schema from the OpenAPI specification
type InputImage struct {
	Url_source string `json:"url_source,omitempty"`
	Base64 string `json:"base64,omitempty"`
	Id_pic string `json:"id_pic,omitempty"`
	Svg_anim string `json:"svg_anim,omitempty"`
}

// SearchData represents the SearchData schema from the OpenAPI specification
type SearchData struct {
	Id_captcha string `json:"id_captcha,omitempty"` // captcha is not used
	Id_search string `json:"id_search,omitempty"`
	Status_only bool `json:"status_only,omitempty"` // true = don't submit a new search
	With_progress bool `json:"with_progress,omitempty"` // true = return imediately with a progress. False waits until search is finished.
	Demo bool `json:"demo,omitempty"` // true = searches only the first 100,000 faces, good for testing/debugging
}

// SearchItem represents the SearchItem schema from the OpenAPI specification
type SearchItem struct {
	Guid string `json:"guid,omitempty"`
	Index int `json:"index,omitempty"`
	Indexdb int64 `json:"indexDB,omitempty"`
	Score int `json:"score,omitempty"`
	Seen int64 `json:"seen,omitempty"`
	Url string `json:"url,omitempty"`
	Base64 string `json:"base64,omitempty"`
	Group int `json:"group,omitempty"`
}

// SearchResults represents the SearchResults schema from the OpenAPI specification
type SearchResults struct {
	Images_in_bundle int `json:"images_in_bundle,omitempty"`
	Performance string `json:"performance,omitempty"`
	Searchedfaces int `json:"searchedFaces,omitempty"`
	Demo bool `json:"demo,omitempty"`
	Tooksecondsdownload float64 `json:"tookSecondsDownload,omitempty"`
	Scaned_till_index int `json:"scaned_till_index,omitempty"`
	Tookseconds float64 `json:"tookSeconds,omitempty"`
	Freeram float64 `json:"freeRam,omitempty"`
	Items []SearchItem `json:"items,omitempty"`
	Max_score int `json:"max_score,omitempty"`
	Tooksecondsqueue float64 `json:"tookSecondsQueue,omitempty"`
	Face_per_sec int `json:"face_per_sec,omitempty"`
}

// BrowserJsonResponse represents the BrowserJsonResponse schema from the OpenAPI specification
type BrowserJsonResponse struct {
	Output SearchResults `json:"output,omitempty"`
	Progress int `json:"progress,omitempty"`
	Code string `json:"code,omitempty"`
	ErrorField string `json:"error,omitempty"`
	Hasemptyimages bool `json:"hasEmptyImages,omitempty"`
	Id_search string `json:"id_search,omitempty"`
	Input []InputImage `json:"input,omitempty"`
	Message string `json:"message,omitempty"`
}
