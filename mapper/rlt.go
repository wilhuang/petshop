package mapper

type TableMetadata struct {
	table   string
	columns []string
}

func (d *TableMetadata) TableName() string {
	return d.table
}

func (d *TableMetadata) Columns() []string {
	return d.columns
}

//"a,b,c,d"
func (d *TableMetadata) ColumnsString() string {
	str := ""
	for i, col := range d.columns {
		if i != 0 {
			str += ", "
		}
		str += col
	}
	return str
}

//a=?,
func (d *TableMetadata) ColumnsEqualString() string {
	str := ""
	for i, col := range d.columns {
		if i != 0 {
			str += ", "
		}
		str += (col + " = ?")
	}
	return str
}

func (d *TableMetadata) ColumnsNum() int {
	return len(d.columns)
}

//?,?,?..
func (d *TableMetadata) QuestionMarkString() string {
	str := ""
	for i, _ := range d.columns {
		if i != 0 {
			str += ", "
		}
		str += "?"
	}
	return str
}
