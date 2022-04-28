package store

type TransactionStack interface {
	Start()
	Quit()
	Commit()
	Abort()
}

type Store interface {
	TransactionStack

	Get(key string) string
	Set(key, value string)
	Delete(key string)
}
