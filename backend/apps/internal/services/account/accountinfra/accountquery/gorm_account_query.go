package accountquery

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
)

type GormAccountQuery struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewGormAccountQuery(db *gorm.DB) *GormAccountQuery {
	return &GormAccountQuery{
		db:    db,
		cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}
