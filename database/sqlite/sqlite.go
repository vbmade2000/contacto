package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"github.com/vbmade2000/contacto/config"
	"github.com/vbmade2000/contacto/interfaces"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteManager struct {
	db *sql.DB
	Conf config.Config 
}

func (sman *SQLiteManager) Connect() error {
	err := errors.New("Error in connecting database")
	fmt.Println(sman.Conf.Host)
	fmt.Println("--------")
	fmt.Println(sman.Conf.DBProviderConf[0].Host)
	fmt.Println("--------------------------")
	sman.db, err = sql.Open(sman.Conf.DBProvider, sman.Conf.DBProviderConf[0].Host)
	return err
}

func (sman *SQLiteManager) CreateDatabase() error {
	fmt.Println("This is CreateDatabase() method")
	return nil
}

func (sman *SQLiteManager) CreateContact(c interfaces.IContact) error {
	fmt.Println("This is CreateContact() method")
	return nil
}

func (sman *SQLiteManager) UpdateContact(contactID int, updatedContact interfaces.IContact) error {
	fmt.Println("This is UpdateContact() method")
	return nil
}

func (sman *SQLiteManager) DeleteContact(contactID int) error {
	fmt.Println("This is DeleteContact() method")
	return nil
}

func (sman *SQLiteManager) GetContacts() error {
	return nil
}

func (sman *SQLiteManager) GetContact(contactID int) sql.Row{
	row := sman.db.QueryRow("Select * from contacts where contact_id=" + strconv.Itoa(contactID))
	return *row
}

