package services

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/Reckner/rad_product/src/database"
	"github.com/Reckner/rad_product/src/models"
)

type itemService struct{}

var (
	ItemService itemService
)

func (i *itemService) GetItems() ([]models.Item, error) {
	db, err := database.CreateConnection()
	if err != nil {
		return []models.Item{}, err
	}

	query, _, err := sq.Select("*").From("item").ToSql()
	if err != nil {
		return []models.Item{}, err
	}

	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	items := scanItems(result)
	db.Close()
	defer result.Close()
	return items, err
}

func scanItems(result *sql.Rows) []models.Item {
	var item models.Item
	var items []models.Item
	for result.Next() {
		err := result.Scan(&item.ID, &item.Name, &item.Price, &item.CategoryID, &item.InSale, &item.AddedTime, &item.LastUpdated, &item.RemovedTime)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, item)
	}

	return items
}
