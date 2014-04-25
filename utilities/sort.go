package utilities

// Type for sorted list of comservers by name
type StringsByName []string

// Returns number of strings in slice (needed for sort.Sort)
func (n StringsByName) Len() int { return len(n) }

// Swaps two string in slice (needed for sort.Sort)
func (n StringsByName) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

// checks which string is lesser than the other one (needed for sort.Sort)
func (n StringsByName) Less(i, j int) bool { return n[i] < n[j] }
