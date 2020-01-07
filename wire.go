//+build wireinject

package main

import (
	"github.com/garaekz/goshort/url"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initURLAPI(db *gorm.DB) url.API {
	wire.Build(url.ProvideRepository, url.ProvideService, url.ProvideAPI)

	return url.API{}
}
