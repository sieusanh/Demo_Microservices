package repoimpl

import (
	"fmt"
	"database/sql"
	repo "go-module/repository"
	models "go-module/model"
)

// Implementing OrderRepo's interface

type OrderRepoImpl struct {
	Db* sql.DB
}

func NewOrderRepo(db *sql.DB) repo.OrderRepo {
	return &OrderRepoImpl {
		Db: db,
	}
}																																							

func (o *OrderRepoImpl) Select(field, 
	value string) ([]models.Order, error) {
	orders := make([]models.Order, 0)
	generalQuery := "SELECT * FROM Orders"
	queryString := generalQuery
	fmt.Printf("field=%s, value=%s", field, value)
	// Select with specified field and value
	if (field != "" && value != "") {
		generalQuery += "\nWHERE %s = %s"
		queryString = fmt.Sprintf(generalQuery, field, value)
	}
	fmt.Println("q: ", queryString)
	rows, err := o.Db.Query(queryString)

	if (err != nil) {
		return orders, err
	}

	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.Id, &order.TotalPrice, &order.UserId, &order.ProductList)
		if err != nil {
			break
		}

		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (o *OrderRepoImpl) Insert(order models.Order) (
	error) {
	insertStatement := `
	INSERT INTO Orders (id, totalPrice, userId, productList)
	VALUES(?, ?, ?, ?)
	`
	_, err := o.Db.Exec(insertStatement, 
		order.Id, 
		order.TotalPrice, 
		order.UserId,
		order.ProductList,
	)

	return err
}

func (o *OrderRepoImpl) UpdateById(id string, 
	order models.Order) error {

	updateStatement := `
	UPDATE Orders
	SET totalPrice = ?, productList = ?
	WHERE id = ?
	`
	_, err := o.Db.Exec(updateStatement, 
		order.TotalPrice, 
		// order.UserId,
		order.ProductList,
		id, 
	)
	return err
}

func (o *OrderRepoImpl) DeleteById(id string) error {
		
	deleteStatement := `
	DELETE FROM Orders
	WHERE id = ?
	`
	_, err := o.Db.Exec(deleteStatement, id)
	return err
}

func (o *OrderRepoImpl) CreateTable(TableName string) (
	error) {
	insertStatement := `
	CREATE TABLE IF NOT EXISTS Orders(
		id int primary key not null auto_increment, 
		totalPrice float, 
		userId text,
		productList JSON
	)  
	`
	_, err := o.Db.Exec(insertStatement)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepoImpl) DropTable(TableName string) (
	error) {
	insertStatement := `
		DROP TABLE Orders;
	`
	_, err := o.Db.Exec(insertStatement)
	if err != nil {
		return err
	}

	return nil
}	