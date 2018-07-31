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
	Name      string
	Publisher string
	SaleDate  time.Time
	NewsURL   string
	ImgURL    string
	Category  string
}

// UpsertProduct save new product
func UpsertProduct(ctx context.Context, key string, product Product) (*datastore.Key, error) {
	return datastore.Put(ctx, datastore.NewKey(ctx, "Product", key, 0, nil), &product)
}

// MakeProduct create new Product object
func MakeProduct(name string, publisher string, saleDate time.Time, newsURL string, imgURL string, category string) (string, Product) {
	converted := sha256.Sum256([]byte(name))
	id := hex.EncodeToString((converted[:]))
	product := Product{Name: name, Publisher: publisher, SaleDate: saleDate, NewsURL: newsURL, ImgURL: imgURL, Category: category}
	return id, product
}

// FetchAllProduct fetch all product from datastore
func FetchAllProduct(ctx context.Context) ([]Product, error) {
	var products []Product
	_, err := datastore.NewQuery("Product").GetAll(ctx, &products)
	return products, err
}
