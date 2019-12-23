package repository

import (
	"fmt"
	"github.com/boltdb/bolt"
	"main/cart-service/structs"
	"strconv"
)

const BucketKey = "user-cart:%d"

type IRepository interface {
	AddToCart(userId int64, items []structs.CartItem) (err error)
	GetCart(userId int64) (items []structs.CartItem, err error)
	ClearCart(userId int64) (err error)
}

type Repository struct {
	Db *bolt.DB
}

// Init initialize repository
func (r *Repository) Init(Db *bolt.DB) {
	r.Db = Db
}

// AddToCart adding new items to cart
func (r *Repository) AddToCart(userId int64, items []structs.CartItem) (err error) {
	err = r.Db.Update(func(tx *bolt.Tx) (err error) {
		b, err := tx.CreateBucketIfNotExists([]byte(fmt.Sprintf(BucketKey, userId)))
		if err != nil {
			return
		}

		for _, item := range items {
			bItemId := []byte(strconv.FormatInt(item.ItemId, 10))
			bQuantity := []byte(strconv.FormatInt(item.Quantity, 10))

			err = b.Put(bItemId, bQuantity)
			if err != nil {
				return
			}
		}

		return
	})

	if err != nil {
		return
	}

	return
}

// GetCart get cart items
func (r *Repository) GetCart(userId int64) (items []structs.CartItem, err error) {
	err = r.Db.View(func(tx *bolt.Tx) (err error) {
		b := tx.Bucket([]byte(fmt.Sprintf(BucketKey, userId)))
		if b == nil {
			return
		}

		items = make([]structs.CartItem, b.Stats().KeyN)
		i := 0

		err = b.ForEach(func(k, v []byte) (err error) {
			items[i].ItemId, err = strconv.ParseInt(string(k), 10, 64)
			if err != nil {
				return
			}

			items[i].Quantity, err = strconv.ParseInt(string(v), 10, 64)
			if err != nil {
				return
			}

			i++
			return
		})

		return
	})
	return
}

// ClearCart clear all items in the cart
func (r *Repository) ClearCart(userId int64) (err error) {
	err = r.Db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(fmt.Sprintf(BucketKey, userId)))
	})

	return
}
