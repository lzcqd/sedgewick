package symboltable

type StringToIntST interface {
	Put(key string, val int)
	Get(key string) (int, bool)
	Delete(key string)
}
//changed
