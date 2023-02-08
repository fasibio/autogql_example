package main

import (
	"context"
	"log"

	"github.com/fasibio/autogql_example/graph/db"
	"github.com/fasibio/autogql_example/graph/model"
	"gorm.io/gorm"
)

type DeleteCompanyHook struct{}

func (h DeleteCompanyHook) Received(ctx context.Context, dbHelper *db.AutoGqlDB, input *model.CompanyFiltersInput) (*gorm.DB, model.CompanyFiltersInput, error) {
	log.Println("Received")
	return dbHelper.Db, *input, nil
}
func (h DeleteCompanyHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	log.Println("BeforeCallDb")
	return db, nil
}
func (h DeleteCompanyHook) BeforeReturn(ctx context.Context, db *gorm.DB, res *model.DeleteCompanyPayload) (*model.DeleteCompanyPayload, error) {
	log.Println("BeforeReturn")
	return res, nil
}

type UpdateCompanyHook struct{}

func (h UpdateCompanyHook) Received(ctx context.Context, dbHelper *db.AutoGqlDB, input *model.UpdateCompanyInput) (*gorm.DB, model.UpdateCompanyInput, error) {
	log.Println("Received")
	return dbHelper.Db, *input, nil
}
func (h UpdateCompanyHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data map[string]interface{}) (*gorm.DB, map[string]interface{}, error) {
	log.Println("BeforeCallDb")
	return db, data, nil
}
func (h UpdateCompanyHook) BeforeReturn(ctx context.Context, db *gorm.DB, res *model.UpdateCompanyPayload) (*model.UpdateCompanyPayload, error) {
	log.Println("BeforeReturn")
	return res, nil
}

type AddCompanyHook struct{}

func (h AddCompanyHook) Received(ctx context.Context, dbHelper *db.AutoGqlDB, input []*model.CompanyInput) (*gorm.DB, []*model.CompanyInput, error) {
	log.Println("Received")
	return dbHelper.Db, input, nil
}
func (h AddCompanyHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data []model.Company) (*gorm.DB, []model.Company, error) {
	log.Println("BeforeCallDb")
	return db, data, nil
}
func (h AddCompanyHook) BeforeReturn(ctx context.Context, db *gorm.DB, res *model.AddCompanyPayload) (*model.AddCompanyPayload, error) {
	log.Println("BeforeReturn")
	return res, nil
}

type GetUserHook struct{}

func (h GetUserHook) Received(ctx context.Context, dbHelper *db.AutoGqlDB, id int) (*gorm.DB, error) {
	log.Println("Received")
	return dbHelper.Db, nil
}
func (h GetUserHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	log.Println("BeforeCallDb")
	return db, nil
}
func (h GetUserHook) AfterCallDb(ctx context.Context, data *model.User) (*model.User, error) {
	log.Println("AfterCallDb")
	return data, nil
}
func (h GetUserHook) BeforeReturn(ctx context.Context, data *model.User, db *gorm.DB) (*model.User, error) {
	log.Println("BeforeReturn")
	return data, nil
}

type QueryCompanyHook struct{}

func (h QueryCompanyHook) Received(ctx context.Context, dbHelper *db.AutoGqlDB, filter *model.CompanyFiltersInput, order *model.CompanyOrder, first, offset *int) (*gorm.DB, *model.CompanyFiltersInput, *model.CompanyOrder, *int, *int, error) {
	log.Println("Received")
	return dbHelper.Db, filter, order, first, offset, nil
}
func (h QueryCompanyHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	log.Println("BeforeCallDb")
	return db, nil
}
func (h QueryCompanyHook) AfterCallDb(ctx context.Context, data []*model.Company) ([]*model.Company, error) {
	log.Println("AfterCallDb")
	return data, nil
}
func (h QueryCompanyHook) BeforeReturn(ctx context.Context, data []*model.Company, db *gorm.DB) ([]*model.Company, error) {
	log.Println("BeforeReturn")
	return data, nil
}
