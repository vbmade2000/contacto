package contact 

import (
	"database/sql"
	"fmt"
	"github.com/vbmade2000/contacto/database"
	"github.com/vbmade2000/contacto/config"
	"github.com/vbmade2000/contacto/interfaces"
)


type ContactManager struct {
	DbMan database.DBManager
}

var cm ContactManager

// Struct to hold contact detail
type Contact struct {
        Id int
        Name string
        Address string
        Email string
	Phone string
        Landline string
        Occupation string
}

// ToSQLInsertQuery returns update sql query for Contact
func (c Contact) ToSQLInsertQuery() (string) {
	return "This method is not implemented yet"
}

// ToSQLUpdateQuery returns update sql query for Contact
func (c Contact) ToSQLUpdateQuery() (string) {
	return "This method is not implemented yet"
}

// New returns a new instance of Contact with all the fields blank 
func New() Contact {
	return Contact{}	
}
func InitDB(conf config.Config) error{
	dbManager, err := database.GetDatabaseManager(conf)
	if err!=nil {
		return err
	}
	cm = ContactManager{}
	cm.DbMan = dbManager
	return nil
}
func (conMan ContactManager) GetContact(contactID int) (interfaces.IContact, error) {
	row := cm.DbMan.GetContact(contactID)
	c := Contact{}
	err := row.Scan(&c.Id, &c.Name, &c.Address, &c.Email, &c.Phone, &c.Landline, &c.Occupation)
	if err == sql.ErrNoRows {
		err = fmt.Errorf("Contact not found")
	}
	return c, err	
}
