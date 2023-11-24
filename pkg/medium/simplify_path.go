package medium

import "strings"

func simplifyPath(path string) string {
	items := strings.Split(path, "/")

	newItems := []string{}

	for _, i := range items {
		if i == ".." {
			if len(newItems) > 0 {
				newItems = newItems[:len(newItems)-1]
			} else {
				continue
			}
		} else if i != "." && i != "" {
			newItems = append(newItems, i)
		}
	}

	return "/" + strings.Join(newItems, "/")
}
