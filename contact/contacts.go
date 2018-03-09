package contact 

import (
	"errors"
)

// IContact is an interface containing required methods for Contact
type IContact interface {

	// Getter methods
	GetId() int
	GetName() string
	GetMobile() string
	GetLandline() string
	GetAddress() string
	GetOccupation() string
	ToJSON() string

	// Setter methods
        SetName(name string) error
        SetMobile(mobile string) error
        SetLandline(landline string) error
        SetAddress(addresss string) error
        SetOccupation(occupation string) error
        FromJSON(json string) IContact
}

// Struct to hold contact detail
type Contact struct {
        id int
        name string
        mobile string
        landline string
        address string
        occupation string
}

func (con *Contact)SetName(name string) error {
	if name == "" {
		return errors.New("Contact name is blank") 
	}
	con.name = name
	return nil
}

// New returns a new instance of Contact with all the fields blank 
func NewContact() Contact {
	return Contact{}	
}
