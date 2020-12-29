package parser

// DataType tells the parser what kind of type a given byte array should be
// parsed as.
type DataType int

const (
	// DTInt denotes an integer
	DTInt DataType = iota

	// DTFloat denotes a float
	DTFloat

	// DTString denotes a string
	DTString

	// DTObject denotes an object
	DTObject

	// DTArray denotes an array
	DTArray

	// DTBool denotes a boolean
	DTBool
)

func (dt *DataType) String() string {
	switch *dt {
	case DTInt:
		return "int"
	case DTFloat:
		return "float"
	case DTString:
		return "string"
	case DTObject:
		return "object"
	case DTArray:
		return "array"
	case DTBool:
		return "bool"
	}

	return "<unknown>"
}

// NewDataType will return a data type for a given type string.
func NewDataType(input string) DataType {
	switch input {
	case "int":
		return DTInt
	case "float":
		return DTFloat
	case "string":
		return DTString
	case "object":
		return DTObject
	case "array":
		return DTArray
	case "bool":
		return DTBool
	}

	return DTString
}
