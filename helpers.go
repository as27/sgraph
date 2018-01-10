package sgraph

func inStringSlice(sl []string, search string) bool {
	for _, v := range sl {
		if v == search {
			return true
		}
	}
	return false
}
