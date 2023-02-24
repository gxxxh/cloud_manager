package utils

// golang type convert to java type
func TypeConvert(typeName string) string {
	switch typeName {
	case "string":
		return "String"
	case "bool":
		return "Boolean"
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return "Integer"
	default:
		return typeName

	}
	return typeName
}
