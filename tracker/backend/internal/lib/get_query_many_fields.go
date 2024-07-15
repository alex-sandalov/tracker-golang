package lib

import "fmt"

// GetQueryManyFields generates the SQL set statements and arguments for updating multiple fields.
//
// Parameters:
// - fields: A map of field names to their new values.
//
// Returns:
// - setStatements: A slice of SQL set statements.
// - args: A slice of values to be interpolated into the set statements.
// - idx: The index of the next argument to be used.
func GetQueryManyFields(fields map[string]string) (setStatements []string, args []interface{}, idx int) {
	for field, value := range fields {
		setStatement := fmt.Sprintf("%s = $%d", field, idx)
		setStatements = append(setStatements, setStatement)
		args = append(args, value)
		idx++
	}

	return setStatements, args, idx
}
