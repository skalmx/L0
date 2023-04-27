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
		log.Println("cant marshal order")
		return err
	}
	_, err = o.db.Exec("INSERT INTO orders (order_uid, order_info) values ($1, $2)", orderUid, jsonOrder)
	if err != nil {
		log.Println("cant create order in db")
		return err
	}
	return nil
}

func (o *OrderRepo) GetById(ctx context.Context, orderUid string) (domain.Order, error) {
	var order domain.Order
	row := o.db.QueryRow("SELECT order_info FROM orders WHERE order_uid = $1", orderUid)
	err := row.Scan(&order)
	//err := o.db.QueryRow("SELECT order_info FROM orders WHERE order_uid = $1", orderUid).Scan(&order)
	if err == sql.ErrNoRows {
		log.Println("no order with this id")
		return domain.Order{}, err
	}
	if err != nil {
		log.Println("cant select order")
		return domain.Order{}, err
	}

	return order, nil
	
	// if err != nil {
	// 	log.Print("cant get order by id")
	// 	return domain.Order{}, err
	// }
	// err = json.Unmarshal(result, &order)
}
