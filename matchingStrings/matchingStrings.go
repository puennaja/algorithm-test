package matchingstrings

// strings => a database of the []string
// queries => a query to query string in strings
// output => the match of queries in straings

func MatchingStrings(strings []string, queries []string) []int32 {
	mapStrings := map[string]int32{}
	for _, s := range strings {
		mapStrings[s] = mapStrings[s] + 1
	}

	var out []int32
	for _, v := range queries {
		out = append(out, mapStrings[v])
	}
	return out
}
