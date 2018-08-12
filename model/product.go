package model

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"google.golang.org/appengine/datastore"
	"time"
)

// Product dao for datastore
type Product struct {
	ID        string
	Name      string
	Publisher string
	SaleDate  time.Time
	NewsURL   string
	ImgURL    string
	Category  string
}

// ProductDateEntry is entry for json
type ProductDateEntry struct {
	SaleDate time.Time
	Products []Product
}

// ProductDateList is list of ProductDateEntry
type ProductDateList []ProductDateEntry

func (l ProductDateList) Len() int {
	return len(l)
}

func (l ProductDateList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ProductDateList) Less(i, j int) bool {
	return l[i].SaleDate.After(l[j].SaleDate)
}

// UpsertProduct save new product
func UpsertProduct(ctx context.Context, key string, product Product) (*datastore.Key, error) {
	return datastore.Put(ctx, datastore.NewKey(ctx, "Product", key, 0, nil), &product)
}

// MakeProduct create new Product object
func MakeProduct(name string, publisher string, saleDate time.Time, newsURL string, imgURL string, category string) (string, Product) {
	converted := sha256.Sum256([]byte(name))
	id := hex.EncodeToString((converted[:]))
	product := Product{ID: id, Name: name, Publisher: publisher, SaleDate: saleDate, NewsURL: newsURL, ImgURL: imgURL, Category: category}
	return id, product
}

// FetchAllProduct fetch all product from datastore
func FetchAllProduct(ctx context.Context) ([]Product, error) {
	var products []Product
	_, err := datastore.NewQuery("Product").GetAll(ctx, &products)
	return products, err
}
