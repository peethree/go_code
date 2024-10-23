package main

func getCounts(messagedUsers []string, validUsers map[string]int) {
	// for every name in messagedUsers
	for _, name := range messagedUsers {
		// if that name is a key in the validUsers map
		if _, ok := validUsers[name]; ok {
			// increment the count by 1
			validUsers[name] += 1
		}
	}
}
