package repository

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/AlehBelski/go-card-api/model"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

//todo redone
func newDB(userName, userPassword, host, dbName string) (CartRepositoryImpl, error) {
	storage := CartRepositoryImpl{}
	dataSource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", userName, userPassword, host, dbName)
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return storage, err
	}
	if err = db.Ping(); err != nil {
		return storage, err
	}
	storage = NewStorage(db)
	return storage, nil
}

//fixme host not found?
func TestStorage_CreateIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip TestStorage_CreateIntegration")
	}
	assertion := assert.New(t)

	expectedCart := model.NewCart(1, []model.CartItem{})

	storage, err := newDB("postgres", "postgres", "pstgr", "postgres")

	if err != nil {
		t.Fatal(err)
	}

	actualCart, err := storage.Create()

	if err != nil {
		t.Fatal(err)
	}

	assertion.EqualValues(expectedCart, actualCart)
}

func TestStorage_ReadIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip TestStorage_CreateIntegration")
	}
	assertion := assert.New(t)

	item := model.NewCartItem(1, 123, "Shoes", 10)

	expectedCart := model.NewCart(123, []model.CartItem{item})

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

func TestStorage_UpdateIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip TestStorage_CreateIntegration")
	}
	assertion := assert.New(t)

	expectedQuery := "INSERT INTO cart_item(.+) VALUES (.+)*"

	itemToUpdate := model.CartItem{}

	itemToUpdate.SetProduct("Shoes")
	itemToUpdate.SetQuantity(10)

	expectedCartItem := model.NewCartItem(1, 123, "Shoes", 10)

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

func TestStorage_DeleteIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip TestStorage_CreateIntegration")
	}
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

func TestStorage_IsCartExistsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip TestStorage_CreateIntegration")
	}
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

func TestStorage_IsCartItemExistsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip TestStorage_CreateIntegration")
	}
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
