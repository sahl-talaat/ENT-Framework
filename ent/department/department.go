// Code generated by ent, DO NOT EDIT.

package department

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the department type in the database.
	Label = "department"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeEmployees holds the string denoting the employees edge name in mutations.
	EdgeEmployees = "employees"
	// Table holds the table name of the department in the database.
	Table = "departments"
	// EmployeesTable is the table that holds the employees relation/edge.
	EmployeesTable = "employees"
	// EmployeesInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeesInverseTable = "employees"
	// EmployeesColumn is the table column denoting the employees relation/edge.
	EmployeesColumn = "department_employees"
)

// Columns holds all SQL columns for department fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Department queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByEmployeesCount orders the results by employees count.
func ByEmployeesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEmployeesStep(), opts...)
	}
}

// ByEmployees orders the results by employees terms.
func ByEmployees(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEmployeesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEmployeesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EmployeesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EmployeesTable, EmployeesColumn),
	)
}
