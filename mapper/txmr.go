package mapper

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// GetTx to link to mysql return a tx obj
func GetTx() (*sql.Tx, error) {
	// TODO. read configuration in petshop
	db, err := sql.Open("mysql", "root:Qq094130@tcp(127.0.0.1:3306)/petshop?charset=utf8")
	if err != nil {
		return nil, err
	}
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

//CloseTx a tx link if it face error we have rollback
func CloseTx(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		// TODO. rollback
		tx.Rollback()
		return nil
	}
	return nil
}
