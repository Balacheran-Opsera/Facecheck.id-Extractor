package main

import (
	"github.com/facial-recognition-reverse-image-face-search-api/mcp-server/config"
	"github.com/facial-recognition-reverse-image-face-search-api/mcp-server/models"
	tools_facecheckapi "github.com/facial-recognition-reverse-image-face-search-api/mcp-server/tools/facecheckapi"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_facecheckapi.CreatePost_api_delete_picTool(cfg),
		tools_facecheckapi.CreatePost_api_infoTool(cfg),
		tools_facecheckapi.CreatePost_api_searchTool(cfg),
	}
}
