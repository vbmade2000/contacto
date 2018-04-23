package database

import (
	"database/sql"
	"github.com/vbmade2000/contacto/config"
	"github.com/vbmade2000/contacto/interfaces"
	"github.com/vbmade2000/contacto/database/sqlite"
)

type DBManager interface {
	Connect() error
	CreateDatabase() error
	CreateContact(interfaces.IContact) error
	UpdateContact(contactID int, updatedContact interfaces.IContact) error
	DeleteContact(contactID int) error
	GetContacts() error
	GetContact(contactID int) sql.Row
}

func GetDatabaseManager(conf config.Config) (DBManager, error) {
	dbMan :=  sqlite.SQLiteManager{Conf:conf}
	err := dbMan.Connect()
	if err!=nil {
		return nil, err
	}
	return &dbMan, nil
}
