package client

func QueryOptionsFiller(pageSize string, dataType string, requireAllWords string) map[string]string {
	m := make(map[string]string)
	m["pageSize"] = pageSize
	m["dataType"] = dataType
	m["requireAllWords"] = requireAllWords
	return m
}
