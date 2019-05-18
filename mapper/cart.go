package mapper

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
)

/***************************************************************************
Product Mapper

SQL:
	--
	-- Definition for table cart
	--
USE petshop;

    CREATE TABLE cart (
	  uid VARCHAR(36) NOT NULL,
	  pid VARCHAR(36) NOT NULL,
	  pnum VARCHAR(128) NOT NULL,
	  PRIMARY KEY (uid,pid)
	)
    ENGINE = INNODB,
    AVG_ROW_LENGTH = 16384,
    CHARACTER SET utf8,
    COLLATE utf8_general_ci;

    ALTER TABLE cart
      ADD CONSTRAINT uid FOREIGN KEY (uid)
        REFERENCES account(uid);

    ALTER TABLE cart
      ADD CONSTRAINT pid FOREIGN KEY (pid)
        REFERENCES product(pid);

INSERT INTO cart VALUES
('001', '001', '1'),
('001', '008', '2'),
('001', '010', '3'),
('001', '018', '5'),
('002', '001', '3'),
('003', '002', '5'),
('003', '001', '3');
***************************************************************************/

type CartRow struct {
	Uid  string
	Pid  string
	Pnum string
}

type CartMapper struct {
	Tx       *sql.Tx
	Metadata TableMetadata
}

func (m *CartMapper) fields(row *CartRow) []interface{} {
	return []interface{}{
		&row.Uid,
		&row.Pid,
		&row.Pnum,
	}
}

func (m *CartMapper) CheckCart(uid, pid string) (*CartRow, error) {
	row := &CartRow{}
	sqlSelect := fmt.Sprintf(`SELECT %s FROM %s WHERE uid = ? and pid = ?`,
		m.Metadata.ColumnsString(),
		m.Metadata.TableName())
	err := m.Tx.QueryRow(sqlSelect, uid, pid).Scan(m.fields(row)...)
	return row, err
}

func (m *CartMapper) AddCart(uid, pid, pnum string) (*CartRow, error) {
	sqlInsert := fmt.Sprintf(`INSERT INTO %s(%s) VALUES (%s);`,
		m.Metadata.TableName(),
		m.Metadata.ColumnsString(),
		m.Metadata.QuestionMarkString())
	res, err := m.Tx.Exec(sqlInsert, uid, pid, pnum)
	beego.Info(sqlInsert)
	beego.Info(uid, pid, pnum)
	beego.Info(res)
	if err != nil {
		return nil, err
	}
	beego.Info(err)
	if num, err := res.RowsAffected(); err != nil || num != 1 {
		return nil, fmt.Errorf("RowsAffected: %d, %s", num, err.Error())
	}
	row := &CartRow{}
	sqlSelect := fmt.Sprintf(`SELECT %s FROM %s WHERE uid = ?`, m.Metadata.ColumnsString(), m.Metadata.TableName())
	err = m.Tx.QueryRow(sqlSelect, uid).Scan(m.fields(row)...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (m *CartMapper) ResetPnum(uid, pid, pnum string) error {
	sqlInsert := fmt.Sprintf(`Update %s SET pnum=? WHERE uid=? and pid=?;`, m.Metadata.TableName())
	res, err := m.Tx.Exec(sqlInsert, pnum, uid, pid)
	if err != nil {
		return err
	}
	if num, err := res.RowsAffected(); err != nil || num != 1 {
		return fmt.Errorf("RowsAffected: %d, %s", num, err.Error())
	}
	return nil
}

func (m *CartMapper) FindCartByUid(uid string) ([]CartRow, error) {
	sql := fmt.Sprintf(`SELECT %s FROM %s WHERE uid = ?`, m.Metadata.ColumnsString(), m.Metadata.TableName())
	rows, err := m.Tx.Query(sql, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	slice := make([]CartRow, 0)
	for rows.Next() {
		row := CartRow{}
		if err := rows.Scan(m.fields(&row)...); err != nil {
			return nil, err
		}
		slice = append(slice, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return slice, nil
}

func (m *CartMapper) FindCartAll() ([]CartRow, error) {
	sql := fmt.Sprintf(`SELECT %s FROM %s`, m.Metadata.ColumnsString(), m.Metadata.TableName())
	rows, err := m.Tx.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	slice := make([]CartRow, 0)
	for rows.Next() {
		row := CartRow{}
		if err := rows.Scan(m.fields(&row)...); err != nil {
			return nil, err
		}
		slice = append(slice, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return slice, nil
}

func (m *CartMapper) DeleteCart(uid, pid string) error {
	sqlSelect := fmt.Sprintf(`DELETE FROM %s WHERE uid= ? and pid = ?`, m.Metadata.TableName())
	_, err := m.Tx.Exec(sqlSelect, uid, pid)
	if err != nil {
		return err
	}
	return nil
}

func (m *CartMapper) DeleteCartByPid(pid string) error {
	sqlSelect := fmt.Sprintf(`DELETE FROM %s WHERE pid = ?`, m.Metadata.TableName())
	_, err := m.Tx.Exec(sqlSelect, pid)
	if err != nil {
		return err
	}
	return nil
}

func (m *CartMapper) DeleteCartByUid(uid string) error {
	sqlSelect := fmt.Sprintf(`DELETE FROM %s WHERE account_id = ?`, m.Metadata.TableName())
	_, err := m.Tx.Exec(sqlSelect, uid)
	if err != nil {
		return err
	}
	return nil
}

func NewCartMapper(tx *sql.Tx) *CartMapper {
	mapper := &CartMapper{
		Tx: tx,
		Metadata: TableMetadata{
			table: "cart",
			columns: []string{
				"uid",
				"pid",
				"pnum",
			},
		},
	}
	return mapper
}
