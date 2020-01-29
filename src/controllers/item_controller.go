package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Reckner/rad_product/src/services"
)

func GetItems(res http.ResponseWriter, req *http.Request) {
	items, err := services.ItemService.GetItems()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(res).Encode(items)
}
