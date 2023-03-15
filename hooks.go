package main

import (
	"context"
	"log"

	"github.com/fasibio/autogql_example/graph/db"
	"github.com/fasibio/autogql_example/graph/model"
	"gorm.io/gorm"
)

type DeleteCompanyHook struct {
	db.DefaultDeleteHook[model.CompanyFiltersInput, model.DeleteCompanyPayload]
}

func (h DeleteCompanyHook) BeforeReturn(ctx context.Context, db *gorm.DB, res *model.DeleteCompanyPayload) (*model.DeleteCompanyPayload, error) {
	log.Println("BeforeReturn")
	return res, nil
}

type UpdateCompanyHook struct {
	db.DefaultUpdateHook[model.UpdateCompanyInput, model.UpdateCompanyPayload]
}

func (h UpdateCompanyHook) Received(ctx context.Context, dbHelper *db.AutoGqlDB, input *model.UpdateCompanyInput) (*gorm.DB, model.UpdateCompanyInput, error) {
	log.Println("Received")
	return dbHelper.Db, *input, nil
}

type AddCompanyHook struct {
	db.DefaultAddHook[model.Company, model.CompanyInput, model.AddCompanyPayload]
}

func (h AddCompanyHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data []model.Company) (*gorm.DB, []model.Company, error) {
	log.Println("BeforeCallDb")
	return db, data, nil
}

type GetUserHook struct {
	db.DefaultGetHook[model.User, int]
}

func (h GetUserHook) AfterCallDb(ctx context.Context, data *model.User) (*model.User, error) {
	log.Println("AfterCallDb")
	return data, nil
}
func (h GetUserHook) BeforeReturn(ctx context.Context, data *model.User, db *gorm.DB) (*model.User, error) {
	log.Println("BeforeReturn")
	return data, nil
}

type QueryCompanyHook struct {
	db.DefaultQueryHook[model.Company, model.CompanyFiltersInput, model.CompanyOrder]
}

func (h QueryCompanyHook) Received(ctx context.Context, dbHelper *db.AutoGqlDB, filter *model.CompanyFiltersInput, order *model.CompanyOrder, first, offset *int) (*gorm.DB, *model.CompanyFiltersInput, *model.CompanyOrder, *int, *int, error) {
	log.Println("Received")
	return dbHelper.Db, filter, order, first, offset, nil
}
