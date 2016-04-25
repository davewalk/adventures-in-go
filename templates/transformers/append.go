package transformers

func Append(s string, args []map[string]interface{}) string {
	for _, arg := range args {
		appendable := arg["value"]
		if appendable == nil {
			return s
		}

		appendableString, ok := appendable.(string)
		if !ok {
			return s
		}
		s += appendableString
	}

	return s
}
