package repository

import (
	"L0/iternal/domain"
	"context"
	"database/sql"
	"encoding/json"
	"log"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (o *OrderRepo) Create(ctx context.Context, orderUid string, order domain.Order) error {

	jsonOrder, err := json.Marshal(order)
	if err != nil {
		log.Print("cant marshal order")
		return err
	}
	_, err = o.db.Exec("INSERT INTO orders (order_uid, order_info) values ($1, $2)", orderUid, jsonOrder)
	if err != nil {
		log.Print("cant create order in db")
		return err
	}
	return nil
}
