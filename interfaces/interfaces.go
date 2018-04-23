package interfaces

// IContact is an interface containing required methods for Contact
type IContact interface {
        ToSQLInsertQuery() string
        ToSQLUpdateQuery() string
}
