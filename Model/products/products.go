package productsmodel

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	configuration "visiku-restapi/Configuration"
	productsentity "visiku-restapi/Entity/products"

	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
)

func QueryParam(name string, desc string, ctgID string, page string) (bool, []productsentity.Products, int, int, int64) {

	var prod []productsentity.Products
	var res *gorm.DB
	var count, start, end int64

	// this action will prevent string "%" from logic "%" + ... + "%" categorized as a string instead of empty value
	sliceName := strings.Split(name, "%")
	sliceDesc := strings.Split(desc, "%")

	log.Printf("Name: %s -> Desc: %s -> CtgID: %s", sliceName[1], sliceDesc[1], ctgID)

	if sliceName[1] != "" && sliceDesc[1] != "" && ctgID != "" {
		res = configuration.DB.Where("name LIKE ? AND description LIKE ? AND category_id = ?", name, desc, ctgID).Order("id DESC").Find(&prod)
		res.Count(&count)
	} else if sliceName[1] != "" && sliceDesc[1] != "" {
		res = configuration.DB.Where("name LIKE ? AND description LIKE ?", name, desc).Order("id DESC").Find(&prod)
		res.Count(&count)
	} else if sliceName[1] != "" {
		res = configuration.DB.Where("name LIKE ?", name).Order("id DESC").Find(&prod)
		res.Count(&count)
	} else if sliceDesc[1] != "" {
		res = configuration.DB.Where("description LIKE ?", desc).Order("id DESC").Find(&prod)
		res.Count(&count)
	} else if ctgID != "" {
		res = configuration.DB.Where("category_id = ?", ctgID).Order("id DESC").Find(&prod)
		res.Count(&count)
	} else {
		res = configuration.DB.Order("id DESC").Find(&prod)
		res.Count(&count)
	}

	// limit 1 page = 5

	// count total page in db
	totalPage := int(math.Ceil(float64(count) / 5))

	// logic for pagination
	currPage, _ := strconv.Atoi(page)
	start = int64((currPage * 5) - 5)
	end = start + 5

	// If the end array is more than db count and page is filled
	if page != "" && end > count {
		end = count
		nowPage, _ := strconv.ParseInt(page, 0, 32)
		return res.RowsAffected > 0, prod[start:end], int(nowPage), totalPage, count
	} else if page != "" && start < count {
		// If data in the page still smaller than count and the page is filled
		nowPage, _ := strconv.ParseInt(page, 0, 32)
		return res.RowsAffected > 0, prod[start:end], int(nowPage), totalPage, count
	} else if page == "" && count < 5 {
		// If data in the page smaller than limit (5) and page is blank
		end = count
		return res.RowsAffected > 0, prod[0:end], 1, totalPage, count
	}

	// none of those if's statement
	// generally when displaying all without filling page params
	return res.RowsAffected > 0, prod[0:5], 1, totalPage, count
}

func InsertProd(r *http.Request) bool {

	var prod productsentity.Products
	var err error

	// using IO lib to read the body request from Postman
	RequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	// Read JSON data and change it into struct like data
	err = json.Unmarshal(RequestBody, &prod)
	if err != nil {
		log.Println(err)
	}

	// Sanitize from XSS case scenario
	sanitize := bluemonday.UGCPolicy()
	stzName := sanitize.Sanitize(prod.Name)
	stzDsc := sanitize.Sanitize(prod.Description)
	stzCtg := sanitize.Sanitize(strconv.Itoa(int(prod.Category_ID)))
	stzPrsCtg, _ := strconv.ParseUint(stzCtg, 0, 64)

	sanitized := productsentity.Products{
		Name:        stzName,
		Description: stzDsc,
		Category_ID: uint(stzPrsCtg),
	}

	res := configuration.DB.Create(&sanitized)

	return res.RowsAffected > 0
}
