package schema

// SQLSchema represents the structure of a parsed SQL schema
type SQLSchema struct {
	Tables []SQLTable
}

// SQLTable represents a table in the SQL schema
type SQLTable struct {
	Name    string
	Columns []SQLColumn
}

// SQLColumn represents a column in an SQL table
type SQLColumn struct {
	Name        string
	DataType    string
	Constraints []string
}
