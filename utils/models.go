package utils

//Map a map types
type Map map[string]interface{}

//DatabaseTable the type for database tables
type DatabaseTable string

//String converts DatabaseTable To String
func (d DatabaseTable) String() string {
	return string(d)
}
