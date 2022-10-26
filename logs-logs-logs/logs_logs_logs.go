package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	returnedString := "default"

	for _, c := range log {

		if c == '❗' {
			returnedString = "recommendation"
			break
		} else if c == '🔍' {
			returnedString = "search"
			break
		} else if c == '☀' {
			returnedString = "weather"
			break
		}
	}
	return returnedString
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var returnedRuneString string
	sliceLog := []rune(log)

	for i := 0; i < len(sliceLog); i++ {
		if sliceLog[i] == oldRune {
			sliceLog[i] = newRune
		}
	}
	returnedRuneString = string(sliceLog)
	return returnedRuneString
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if len([]rune(log)) <= limit {
		return true
	}
	return false
}
