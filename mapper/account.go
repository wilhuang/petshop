package mapper

import (
	"database/sql"
	"fmt"
)

/***************************************************************************
Account Mapper

SQL:
	--
	-- Definition for table account
	--
USE accounts;

    CREATE TABLE account (
	  uid VARCHAR(36) NOT NULL,
	  username VARCHAR(128) NOT NULL,
	  role VARCHAR(8) NOT NULL,
	  email VARCHAR(128) NOT NULL,
	  password varchar(36) not null,
	  PRIMARY KEY (uid)
	)
	ENGINE = INNODB
	AVG_ROW_LENGTH = 16384
	CHARACTER SET utf8
	COLLATE utf8_general_ci;

INSERT INTO account VALUES
('001', 'Test1','1', 'Test1@qq.com', '123456'),
('002', 'Test2','0', 'Test2@qq.com', '123456'),
('003', 'Test3','0', 'Test3@qq.com', '123456'),
('004', 'Test4','0', 'Test4@qq.com', '123456'),
('005', 'Test5','0', 'Test5@qq.com', '123456'),
('006', 'Test6','0', 'Test6@qq.com', '123456'),
('007', 'Test7','0', 'Test7@qq.com', '123456'),
('008', 'Test8','0', 'Test8@qq.com', '123456');
***************************************************************************/

type AccountRow struct {
	Uid      string
	UserName string
	Role     string
	Email    string
	password string
}

type AccountMapper struct {
	Tx       *sql.Tx
	Metadata TableMetadata
}

func (m *AccountMapper) fields(row *AccountRow) []interface{} {
	return []interface{}{
		&row.Uid,
		&row.UserName,
		&row.Role,
		&row.Email,
		&row.password,
	}
}

func (m *AccountMapper) CreateAccount(uid, username, role, email, password string) (*AccountRow, error) {
	sqlInsert := fmt.Sprintf(`INSERT INTO %s(%s) VALUES (%s, ?);`,
		m.Metadata.TableName(),
		m.Metadata.ColumnsString(),
		m.Metadata.QuestionMarkString())
	res, err := m.Tx.Exec(sqlInsert, uid, username, role, email, password)
	if err != nil {
		return nil, err
	}
	if num, err := res.RowsAffected(); err != nil || num != 1 {
		return nil, fmt.Errorf("RowsAffected: %d, %s", num, err.Error())
	}
	row := &AccountRow{}
	sqlSelect := fmt.Sprintf(`SELECT %s FROM %s WHERE uid = ?`, m.Metadata.ColumnsString(), m.Metadata.TableName())
	err = m.Tx.QueryRow(sqlSelect, uid).Scan(m.fields(row)...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (m *AccountMapper) ResetAccount(username, key, value string) error {
	sqlInsert := fmt.Sprintf(`Update %s SET %s=? VALUES username=?;`, m.Metadata.TableName(), key)
	res, err := m.Tx.Exec(sqlInsert, value, username)
	if err != nil {
		return err
	}
	if num, err := res.RowsAffected(); err != nil || num != 1 {
		return fmt.Errorf("RowsAffected: %d, %s", num, err.Error())
	}
	return nil
}

func (m *AccountMapper) FindAccountByUid(uid string) (*AccountRow, error) {
	row := &AccountRow{}
	sqlSelect := fmt.Sprintf(`SELECT %s FROM %s WHERE uid = ?`, m.Metadata.ColumnsString(), m.Metadata.TableName())
	err := m.Tx.QueryRow(sqlSelect, uid).Scan(m.fields(row)...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (m *AccountMapper) FindAccountByUserName(username string) (*AccountRow, error) {
	row := &AccountRow{}
	sqlSelect := fmt.Sprintf(`SELECT %s FROM %s WHERE username = ?`, m.Metadata.ColumnsString(), m.Metadata.TableName())
	err := m.Tx.QueryRow(sqlSelect, username).Scan(m.fields(row)...)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err

	}
	return row, nil
}

func (m *AccountMapper) CheckLoginInfo(username, password string) error {
	row := &AccountRow{}
	sqlSelect := fmt.Sprintf(`SELECT %s FROM %s WHERE username = ? and password = ?`,
		m.Metadata.ColumnsString(),
		m.Metadata.TableName())
	err := m.Tx.QueryRow(sqlSelect, username, password).Scan(m.fields(row)...)
	return err
}

func (m *AccountMapper) DeleteAccountByUserName(username string) error {
	sqlSelect := fmt.Sprintf(`DELETE FROM %s WHERE username = ?`, m.Metadata.TableName())
	res, err := m.Tx.Exec(sqlSelect, username)
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num != 1 {
		return fmt.Errorf("AffectedRows:%d, now rows found", num)
	}
	return nil
}

func (m *AccountMapper) FindAccountAll() ([]AccountRow, error) {
	sql := fmt.Sprintf(`SELECT %s FROM %s`, m.Metadata.ColumnsString(), m.Metadata.TableName())
	rows, err := m.Tx.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	slice := make([]AccountRow, 0)
	for rows.Next() {
		row := AccountRow{}
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

func (m *AccountMapper) FindAccountCount() (int, error) {
	sql := fmt.Sprintf(`SELECT COUNT(*) FROM  %s`, m.Metadata.TableName())
	rows, err := m.Tx.Query(sql)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	if rows.Next() {
		count := -1
		err = rows.Scan(&count)
		if err != nil {
			return -1, err
		}
		return count, nil
	}
	return -1, fmt.Errorf("invalid result")
}

func (m *AccountMapper) FindAccountList(startIndex, count int32) ([]AccountRow, error) {
	sql := fmt.Sprintf(`SELECT %s FROM %s ORDER BY uid LIMIT ? , ?`,
		m.Metadata.ColumnsString(), m.Metadata.TableName())
	rows, err := m.Tx.Query(sql, startIndex, count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	slice := make([]AccountRow, 0)
	for rows.Next() {
		row := AccountRow{}
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

func NewAccountMapper(tx *sql.Tx) *AccountMapper {
	mapper := &AccountMapper{
		Tx: tx,
		Metadata: TableMetadata{
			table: "account",
			columns: []string{
				"uid",
				"username",
				"role",
				"email",
				"password",
			},
		},
	}

	return mapper
}
