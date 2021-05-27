package cmd

import (
	"aspect_apps/config"
	"aspect_apps/internal/http/rest"
	"aspect_apps/internal/services/account/accountapp"
	"aspect_apps/internal/services/account/accountinfra/accountrepo"
	"aspect_apps/internal/shared"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	conf config.Config
	err  error
)

func init() {
	conf, err = config.DefaultConfig()
	if err != nil {
		panic(err)
	}
	shared.BeautyCli(conf.ENV)
}

func StartAPI() {

	var service rest.Service
	domainEvent := shared.NewDomainEventHub()

	// ! ======= Setup Database ======= !
	gormDB, err := gorm.Open(conf.GormConfig.Dialect, conf.PostgresConfig.Url)
	if conf.ENV == "dev" {
		gormDB.LogMode(true)
	}
	if err != nil {
		log.Println(err)
		return
	}
	defer gormDB.Close()

	if err := gormDB.DB().Ping(); err != nil {
		log.Println(err)
		return
	}
	// ! ===== Setup Database End ===== !

	// ? ======== Utils ======== ?

	// bcryptHasher := accountinfra.NewBcryptHasher()
	// domainEventHub := shared.NewDomainEventHub()
	// restJwtUserIdentifier := rest.NewJwtUserIdentifier([]byte(conf.App.JWT.Salt))

	// ? ====== Utils End ====== ?

	// ! ======== Define all service ======== !

	// ? Account Service Start
	accRepo := accountrepo.NewGormAccountRepository(gormDB)
	accAppSvc := accountapp.NewAccountApplicationService(accRepo)
	accSvc := rest.NewAccountMainService(accAppSvc)
	// ? Account Service End

	// ! === End of defining each service === !

	// ? === injecting app service Start === ?
	service.
		NewInjectAccountService(accSvc).
		NewInjectSService(accSvc)
	// ? ====== injecting service End ====== ?

	// ! ==== Event Driven Service Start ==== !

	domainEvent.Start() // Start the Event service
	// ! ===== Event Driven Service End ===== !

	// Run the apps

	rest := rest.NewApp(&service)
	rest.StartApp(conf.App.ApiKey, fmt.Sprintf(":%s", conf.App.Port))
}
