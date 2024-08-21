package problems

type Key [26]byte

func GroupAnagram(list []string) [][]string {
	groups := make(map[Key][]string)

	for _, values := range list {

		key := strKey(values)

		groups[key] = append(groups[key], values)
	}

	result := make([][]string, 0, len(groups))

	for _, v := range groups {
		result = append(result, v)
	}

	return result
}

func strKey(str string) Key {

	var key Key

	for i := range str {
		key[str[i]-'a']++
	}

	return key
}
