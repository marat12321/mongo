package collection

import "go.mongodb.org/mongo-driver/mongo"

//Transaction структура канала для создания транзакций
type Transaction struct {
	Context mongo.SessionContext
	Error   error
}
