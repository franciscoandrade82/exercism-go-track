package twofer

// ShareWith will return the string "One for X, one for me." where X is the name passed in.
// If no name is passed in, it will default to "you".
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
