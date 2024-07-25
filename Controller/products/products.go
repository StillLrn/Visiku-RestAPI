package productcontroller

import (
	"encoding/json"
	"log"
	"net/http"
	productsmodel "visiku-restapi/Model/products"
)

func Product(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		// Get user query param from postman
		name := "%" + r.URL.Query().Get("name") + "%"
		desc := "%" + r.URL.Query().Get("description") + "%"
		catID := r.URL.Query().Get("product_category")
		page := r.URL.Query().Get("page")

		res, prod, now, all, count := productsmodel.QueryParam(name, desc, catID, page)
		log.Println(res)

		// Check such record exists or not
		if !res {
			Response := Msg{
				Message: "No record",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response)
			return
		}

		Paginate := Pagination{
			Page:      now,
			AllPage:   all,
			TotalData: count,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(prod)
		json.NewEncoder(w).Encode(Paginate)

	} else if r.Method == "POST" {
		if res := productsmodel.InsertProd(r); !res {
			Response := Msg{
				Message: "Error to process your data, please fill it correctly",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response)
			return
		}

		Response := Msg{
			Message: "Success, your data has been inserted",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response)
	}
}

type Msg struct {
	Message string `json:"message" gorm:"-"`
}

type Pagination struct {
	Page      int   `json:"page" gorm:"-"`
	AllPage   int   `json:"all_page" gorm:"-"`
	TotalData int64 `json:"total_data" gorm:"-"`
}
