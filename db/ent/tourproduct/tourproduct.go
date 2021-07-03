// Code generated by entc, DO NOT EDIT.

package tourproduct

const (
	// Label holds the string label denoting the tourproduct type in the database.
	Label = "tour_product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldForSale holds the string denoting the forsale field in the database.
	FieldForSale = "for_sale"
	// EdgeManager holds the string denoting the manager edge name in mutations.
	EdgeManager = "manager"
	// Table holds the table name of the tourproduct in the database.
	Table = "tour_products"
	// ManagerTable is the table the holds the manager relation/edge.
	ManagerTable = "tour_products"
	// ManagerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ManagerInverseTable = "users"
	// ManagerColumn is the table column denoting the manager relation/edge.
	ManagerColumn = "user_products"
)

// Columns holds all SQL columns for tourproduct fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPrice,
	FieldForSale,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tour_products"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_products",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultForSale holds the default value on creation for the "forSale" field.
	DefaultForSale bool
)
