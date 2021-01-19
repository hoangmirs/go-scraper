package bootstrap

import (
	_ "github.com/hoangmirs/go-scraper/models"  // Models
	_ "github.com/hoangmirs/go-scraper/routers" // Routers
)

func init() {
	SetUpDB()
}
