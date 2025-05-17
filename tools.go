package main

func compareSlices(s1, s2 []string) bool {
	if len(s1) == len(s2) {
		for i := range s1 {
			if s1[i] != s2[i] {
				return false
			}
		}
		return true
	}
	return false
}
