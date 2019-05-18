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
	-- Definition for table product
	--
USE petshop;

    CREATE TABLE product (
	  pid VARCHAR(36) NOT NULL,
	  pname VARCHAR(128) NOT NULL,
	  oprice VARCHAR(8) NOT NULL,
	  pprice VARCHAR(8) NOT NULL,
	  onhand VARCHAR(8) NOT NULL,
	  pimg1 varchar(64) not null,
	  pimg2 varchar(64) not null,
	  pimg3 varchar(64) not null,
	  pimg4 varchar(64) not null,
	  pimg5 varchar(64) not null,
	  pimg6 varchar(64) not null,
	  PRIMARY KEY (pid)
	)
	ENGINE = INNODB
	AVG_ROW_LENGTH = 16384
	CHARACTER SET utf8
	COLLATE utf8_general_ci;
    ALTER TABLE product
        ADD UNIQUE INDEX pid(pid);
	INSERT INTO product VALUES
('001', '柯基幼犬活体 纯种威尔士柯基犬 宠物AA级 公','6500','7500','55', '../static/img/product/001A.jpg', '../static/img/product/001B.jpg', '../static/img/product/001C.jpg', '../static/img/product/001D.jpg', '../static/img/product/001E.jpg', '../static/img/product/001F.jpg'),
('002', '哈士奇幼犬活体 纯种西伯利亚雪橇犬 宠物级 公','2499','2999','122', '../static/img/product/002A.jpg', '../static/img/product/002B.jpg', '../static/img/product/002C.jpg', '../static/img/product/002D.jpg', '../static/img/product/002E.jpg', '../static/img/product/002F.jpg'),
('003', '泰迪幼犬活体 茶杯犬 纯种泰迪狗贵宾犬 宠物AA级 公','3999','4399','98', '../static/img/product/003A.jpg', '../static/img/product/003B.jpg', '../static/img/product/003C.jpg', '../static/img/product/003D.jpg', '../static/img/product/003E.jpg', '../static/img/product/003F.jpg'),
('004', '萨摩耶幼犬活体 纯种萨摩耶犬 限时特惠（高品质） 公','3199','3399','112', '../static/img/product/004A.jpg', '../static/img/product/004B.jpg', '../static/img/product/004C.jpg', '../static/img/product/004D.jpg', '../static/img/product/004E.jpg', '../static/img/product/004F.jpg'),
('005', '博美犬活体 茶杯犬 纯种博美俊介幼犬 宠物级（高品质） 公','4999','5099','69', '../static/img/product/005A.jpg', '../static/img/product/005B.jpg', '../static/img/product/005C.jpg', '../static/img/product/005D.jpg', '../static/img/product/005E.jpg', '../static/img/product/005F.jpg'),
('006', '纯种阿拉斯加幼犬活体 宠物A级 公','2499','2799','123', '../static/img/product/006A.jpg', '../static/img/product/006B.jpg', '../static/img/product/006C.jpg', '../static/img/product/006D.jpg', '../static/img/product/006E.jpg', '../static/img/product/006F.jpg'),
('007', '金毛犬活体 纯种金毛寻回犬 宠物级 公','1999','2199','134', '../static/img/product/007A.jpg', '../static/img/product/007B.jpg', '../static/img/product/007C.jpg', '../static/img/product/007D.jpg', '../static/img/product/007E.jpg', '../static/img/product/007F.jpg'),
('008', '佩妮6+1猫配餐7罐 幼猫罐头成猫猫粮猫湿粮猫零食','119','130','9888', '../static/img/product/008A.jpg', '../static/img/product/008B.jpg', '../static/img/product/008C.jpg', '../static/img/product/008D.jpg', '../static/img/product/008E.jpg', '../static/img/product/008F.jpg'),
('009', '柴犬幼犬活体 纯种日系柴犬 母','15999','16499','43', '../static/img/product/009A.jpg', '../static/img/product/009B.jpg', '../static/img/product/009C.jpg', '../static/img/product/009D.jpg', '../static/img/product/009E.jpg', '../static/img/product/009F.jpg'),
('010', '英短蓝猫 短毛蓝猫纯种英国短毛猫 宠物AA级 公','4999','5299','64', '../static/img/product/010A.jpg', '../static/img/product/010B.jpg', '../static/img/product/010C.jpg', '../static/img/product/010D.jpg', '../static/img/product/010E.jpg', '../static/img/product/010F.jpg'),
('011', '英短蓝白活体 短毛蓝白猫纯种英国短毛猫 宠物AA级 公','5499','5699','60', '../static/img/product/011A.jpg', '../static/img/product/011B.jpg', '../static/img/product/011C.jpg', '../static/img/product/011D.jpg', '../static/img/product/011E.jpg', '../static/img/product/011F.jpg'),
('012', '布偶猫 纯种布拉多尔猫 宠物AA级 母','19799','19999','35', '../static/img/product/012A.jpg', '../static/img/product/012B.jpg', '../static/img/product/012C.jpg', '../static/img/product/012D.jpg', '../static/img/product/012E.jpg', '../static/img/product/012F.jpg'),
('013', '短毛猫 孟加拉猫 宠物A级 母','6799','6999','51', '../static/img/product/013A.jpg', '../static/img/product/013B.jpg', '../static/img/product/013C.jpg', '../static/img/product/013D.jpg', '../static/img/product/013E.jpg', '../static/img/product/013F.jpg'),
('014', '英短蓝猫 短毛蓝猫纯种英国短毛猫 宠物A级 公','3999','4499','100', '../static/img/product/014A.jpg', '../static/img/product/014B.jpg', '../static/img/product/014C.jpg', '../static/img/product/014D.jpg', '../static/img/product/014E.jpg', '../static/img/product/014F.jpg'),
('015', '纯种边境牧羊犬犬活体 赛级 公','36999','37999','23', '../static/img/product/015A.jpg', '../static/img/product/015B.jpg', '../static/img/product/015C.jpg', '../static/img/product/015D.jpg', '../static/img/product/015E.jpg', '../static/img/product/015F.jpg'),
('016', '英短金渐层猫 短毛金渐层猫纯种英国短毛猫 血统级 母','11999','12999','35', '../static/img/product/016A.jpg', '../static/img/product/016B.jpg', '../static/img/product/016C.jpg', '../static/img/product/016D.jpg', '../static/img/product/016E.jpg', '../static/img/product/016F.jpg'),
('017', '英短金渐层猫 短毛金渐层猫纯种英国短毛猫 新上架幼猫（高品质） 母','16999','17399','33', '../static/img/product/017A.jpg', '../static/img/product/017B.jpg', '../static/img/product/017C.jpg', '../static/img/product/017D.jpg', '../static/img/product/017E.jpg', '../static/img/product/017F.jpg'),
('018', '耐威克狗粮 澳洲牛肉配方犬粮6kg','169','180','7889', '../static/img/product/018A.jpg', '../static/img/product/018B.jpg', '../static/img/product/018C.jpg', '../static/img/product/018D.jpg', '../static/img/product/018E.jpg', '../static/img/product/018F.jpg');

***************************************************************************/

type ProductRow struct {
	Pid    string
	Pname  string
	Oprice string
	Pprice string
	Onhand string
	Pimg1  string
	Pimg2  string
	Pimg3  string
	Pimg4  string
	Pimg5  string
	Pimg6  string
}

type ProductMapper struct {
	Tx       *sql.Tx
	Metadata TableMetadata
}

func (m *ProductMapper) fields(row *ProductRow) []interface{} {
	return []interface{}{
		&row.Pid,
		&row.Pname,
		&row.Oprice,
		&row.Pprice,
		&row.Onhand,
		&row.Pimg1,
		&row.Pimg2,
		&row.Pimg3,
		&row.Pimg4,
		&row.Pimg5,
		&row.Pimg6,
	}
}

func (m *ProductMapper) CreateProduct(pid, username, role, email, password string) (*ProductRow, error) {
	sqlInsert := fmt.Sprintf(`INSERT INTO %s(%s) VALUES (%s);`,
		m.Metadata.TableName(),
		m.Metadata.ColumnsString(),
		m.Metadata.QuestionMarkString())
	res, err := m.Tx.Exec(sqlInsert, pid, username, role, email, password)
	beego.Info(sqlInsert)
	beego.Info(pid, username, role, email, password)
	beego.Info(res)
	if err != nil {
		return nil, err
	}
	beego.Info(err)
	if num, err := res.RowsAffected(); err != nil || num != 1 {
		return nil, fmt.Errorf("RowsAffected: %d, %s", num, err.Error())
	}
	row := &ProductRow{}
	sqlSelect := fmt.Sprintf(`SELECT %s FROM %s WHERE pid = ?`, m.Metadata.ColumnsString(), m.Metadata.TableName())
	err = m.Tx.QueryRow(sqlSelect, pid).Scan(m.fields(row)...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (m *ProductMapper) ResetProduct(username, key, value string) error {
	sqlInsert := fmt.Sprintf(`Update %s SET %s=? WHERE username=?;`, m.Metadata.TableName(), key)
	res, err := m.Tx.Exec(sqlInsert, value, username)
	if err != nil {
		return err
	}
	if num, err := res.RowsAffected(); err != nil || num != 1 {
		return fmt.Errorf("RowsAffected: %d, %s", num, err.Error())
	}
	return nil
}

func (m *ProductMapper) UpdateDetails(pid, pnmae, oprice, onhand string) error {
	sqlInsert := fmt.Sprintf(`Update %s SET pname=?,oprice=?,onhand=? WHERE pid=?;`, m.Metadata.TableName())
	res, err := m.Tx.Exec(sqlInsert, pnmae, oprice, onhand, pid)
	if err != nil {
		beego.Info(err)
		return err
	}
	if num, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("RowsAffected: %d, %s", num, err.Error())
	}
	return nil
}

func (m *ProductMapper) FindProductByPid(pid string) (*ProductRow, error) {
	row := &ProductRow{}
	sqlSelect := fmt.Sprintf(`SELECT %s FROM %s WHERE pid = ?`, m.Metadata.ColumnsString(), m.Metadata.TableName())
	err := m.Tx.QueryRow(sqlSelect, pid).Scan(m.fields(row)...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (m *ProductMapper) FindProductByUserName(username string) (*ProductRow, error) {
	row := &ProductRow{}
	sqlSelect := fmt.Sprintf(`SELECT %s FROM %s WHERE username = ?`, m.Metadata.ColumnsString(), m.Metadata.TableName())
	err := m.Tx.QueryRow(sqlSelect, username).Scan(m.fields(row)...)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err

	}
	return row, nil
}

func (m *ProductMapper) DeleteProductByPid(pid string) error {
	sqlSelect := fmt.Sprintf(`DELETE FROM %s WHERE pid = ?`, m.Metadata.TableName())
	res, err := m.Tx.Exec(sqlSelect, pid)
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

func (m *ProductMapper) FindProductAll() ([]ProductRow, error) {
	sql := fmt.Sprintf(`SELECT %s FROM %s`, m.Metadata.ColumnsString(), m.Metadata.TableName())
	rows, err := m.Tx.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	slice := make([]ProductRow, 0)
	for rows.Next() {
		row := ProductRow{}
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

func (m *ProductMapper) FindProductCount() (int, error) {
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

func (m *ProductMapper) FindProductList(startIndex, count int32) ([]ProductRow, error) {
	sql := fmt.Sprintf(`SELECT %s FROM %s ORDER BY pid LIMIT ? , ?`,
		m.Metadata.ColumnsString(), m.Metadata.TableName())
	rows, err := m.Tx.Query(sql, startIndex, count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	slice := make([]ProductRow, 0)
	for rows.Next() {
		row := ProductRow{}
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

func NewProductMapper(tx *sql.Tx) *ProductMapper {
	mapper := &ProductMapper{
		Tx: tx,
		Metadata: TableMetadata{
			table: "product",
			columns: []string{
				"pid",
				"pname",
				"oprice",
				"pprice",
				"onhand",
				"pimg1",
				"pimg2",
				"pimg3",
				"pimg4",
				"pimg5",
				"pimg6",
			},
		},
	}
	return mapper
}
