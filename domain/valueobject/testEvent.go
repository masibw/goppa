package valueobject

type TestEvent struct {
	Name    string
	Elapsed float64 // Elapsed is the time taken for the test, expressed in seconds.
}

func (t *TestEvent) IsSlowerThan(prev float64, border float64) bool {
	return t.Elapsed > prev*border
}
