package accountrepo

import (
	"aspect_apps/internal/services/account"
	"aspect_apps/internal/shared"
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

type GormAccountRepository struct {
	db *gorm.DB
}

func NewGormAccountRepository(db *gorm.DB) *GormAccountRepository {
	db.AutoMigrate(&gormAccount{})
	return &GormAccountRepository{
		db: db,
	}
}

func (repo *GormAccountRepository) FindByID(id account.AccountID) (*account.Account, shared.AggregateVersion, error) {
	var accountGorm gormAccount
	if err := repo.db.First(&accountGorm, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &account.Account{}, 0, nil
		}
		log.Println(err)
		return &account.Account{}, -1, shared.NewInfrastructureError(err)
	}
	return accountGorm.ToAccountModel(), accountGorm.Version, nil
}

func (repo *GormAccountRepository) FindByEmail(email shared.Email) (*account.Account, shared.AggregateVersion, error) {
	var accountGorm gormAccount
	if err := repo.db.First(&accountGorm, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &account.Account{}, 0, nil
		}
		log.Println(err)
		return &account.Account{}, -1, shared.NewInfrastructureError(err)
	}
	return accountGorm.ToAccountModel(), accountGorm.Version, nil
}

func (repo *GormAccountRepository) Save(acc *account.Account, currentVersion shared.AggregateVersion) (*account.Account, error) {

	gormAcc := newGormAccount(acc)

	if currentVersion == shared.NewAggregate {
		gormAcc.Version = 1
		if err := repo.db.Save(&gormAcc).Error; err != nil {
			log.Println(err)
			return &account.Account{}, shared.NewInfrastructureError(err)
		}
		return gormAcc.ToAccountModel(), nil
	}

	gormAcc.Version = currentVersion + 1

	res := repo.db.Where("version = ?", currentVersion).Save(gormAcc)

	if err := res.Error; err != nil {
		log.Print(err)
		return &account.Account{}, shared.NewInfrastructureError(err)
	}

	if res.RowsAffected <= 0 {
		return &account.Account{}, shared.AggregateVersionOutdatedError
	}

	return &account.Account{}, nil
}
