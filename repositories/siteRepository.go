package repositories

import (
	"context"
	"github.com/chunganhbk/chat-golang/database"
	"github.com/chunganhbk/chat-golang/models"
)


type SiteRepository interface{
	Find(ctx context.Context, filters interface{}) ([]models.SiteSchema, error)
}

type siteRepository struct {
	store      *database.MongoDataStore
}

