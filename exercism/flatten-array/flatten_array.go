package flatten

//Flatten flattens the input arbitrarily deep list of integer values.
func Flatten(inp interface{}) []interface{} {
	returned := []interface{}{}
	lis, ok := inp.([]interface{})
	if !ok {
		return nil
	}
	for _, v := range lis {
		if v == nil {
			continue
		}
		switch v.(type) {
		case int:
			returned = append(returned, v)

		case []interface{}:
			returned = append(returned, Flatten(v)...)
		default:
			panic("You lied to me!")
		}
	}
	return returned
}
