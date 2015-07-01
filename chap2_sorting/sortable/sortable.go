package sortable

type Interface interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}
