package repository

import (
	"testing"

	"github.com/AlehBelski/go-card-api/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestStorage_Create(t *testing.T) {
	assertion := assert.New(t)

	expectedCart := model.Cart{
		ID:    123,
		Items: []model.CartItem{},
	}

	expectedQuery := "INSERT INTO cart VALUES(DEFAULT)*"

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	sqlMock.ExpectQuery(expectedQuery).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			FromCSVString("123"))

	storage := NewStorage(db)

	actualCart, err := storage.Create()

	if err != nil {
		t.Fatal(err)
	}

	assertion.EqualValues(expectedCart, actualCart)
}

func TestStorage_Read(t *testing.T) {
	assertion := assert.New(t)

	item := model.CartItem{
		ID:       1,
		CartID:   123,
		Product:  "Shoes",
		Quantity: 10,
	}

	expectedCart := model.Cart{
		ID:    123,
		Items: []model.CartItem{item},
	}

	expectedQuery := "SELECT \\* FROM cart_item WHERE *"

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"id", "product", "quantity", "fk_cart_id"}).
		AddRow(1, "Shoes", 10, 123)

	sqlMock.ExpectQuery(expectedQuery).
		WithArgs(123).
		WillReturnRows(rows)

	storage := NewStorage(db)

	actualCart, err := storage.Read(123)

	if err != nil {
		t.Fatal(err)
	}

	assertion.EqualValues(expectedCart, actualCart)
}

func TestStorage_Update(t *testing.T) {
	assertion := assert.New(t)

	expectedQuery := "INSERT INTO cart_item(.+) VALUES (.+)*"

	itemToUpdate := model.CartItem{
		Product:  "Shoes",
		Quantity: 10,
	}

	expectedCartItem := model.CartItem{
		ID:       1,
		CartID:   123,
		Product:  "Shoes",
		Quantity: 10,
	}

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"id", "fk_cart_id"}).
		AddRow(1, 123)

	sqlMock.ExpectQuery(expectedQuery).
		WithArgs("Shoes", 10, 123).
		WillReturnRows(rows)

	storage := NewStorage(db)

	actualCartItem, err := storage.Update(123, itemToUpdate)

	if err != nil {
		t.Fatal(err)
	}

	assertion.EqualValues(expectedCartItem, actualCartItem)
}

func TestStorage_Delete(t *testing.T) {
	assertion := assert.New(t)

	expectedQuery := "DELETE FROM cart_item WHERE *"

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	sqlMock.ExpectQuery(expectedQuery).
		WithArgs(123, 1).
		WillReturnRows(sqlMock.NewRows([]string{}))

	storage := NewStorage(db)

	err = storage.Delete(123, 1)

	if err != nil {
		t.Fatal(err)
	}

	assertion.Nil(err)
}

func TestStorage_IsCartExists(t *testing.T) {
	assertion := assert.New(t)

	expectedQuery := "SELECT 1 FROM cart WHERE *"

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	sqlMock.ExpectQuery(expectedQuery).
		WithArgs(1).
		WillReturnRows(sqlMock.NewRows([]string{""}).FromCSVString("true"))

	storage := NewStorage(db)

	isExists, err := storage.IsCartExists(1)

	if err != nil {
		t.Fatal(err)
	}

	assertion.True(isExists)
}

func TestStorage_IsCartItemExists(t *testing.T) {
	assertion := assert.New(t)

	expectedQuery := "SELECT 1 FROM cart_item WHERE *"

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	sqlMock.ExpectQuery(expectedQuery).
		WithArgs(1, 123).
		WillReturnRows(sqlMock.NewRows([]string{""}).FromCSVString("true"))

	storage := NewStorage(db)

	isExists, err := storage.IsCartItemExists(1, 123)

	if err != nil {
		t.Error(err)
	}

	assertion.True(isExists)
}
