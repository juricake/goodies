package split

// Split a given slice by the test condition. The result are two slices,
// where the first contains values passing the condition, and the second contains values failing it.
// Nil, Nil will be returned on an empty slice.
func Split[V any](values []V, test func(value V) bool) ([]V, []V) {
	if len(values) == 0 {
		return nil, nil
	}

	pass := make([]V, 0)
	fail := make([]V, 0)
	for _, value := range values {
		if test(value) {
			pass = append(pass, value)
		} else {
			fail = append(fail, value)
		}
	}

	return pass, fail
}
