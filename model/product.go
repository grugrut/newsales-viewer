package model

import (
	"context"
	"google.golang.org/appengine/datastore"
)

// Product dao for datastore
type Product struct {
	Name string
}

// AddProduct save new product
func AddProduct(ctx context.Context, product Product) (*datastore.Key, error) {
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Product", nil), &product)
}
