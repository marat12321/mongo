package mongo

//import (
//	connect "github.com/marat12321/mongo/connectOptions"
//	mongodb "github.com/marat12321/mongo/mongo"
//	"fmt"
//	"os"
//	"go.mongodb.org/mongo-driver/bson"
//	"context"
//)
//
//func main() {
//	var db2 = struct {
//		Tokens mongodb.Collection
//		Users  mongodb.Collection
//	}{}
//
//	conOpts := connect.Options("test")
//	conOpts.Hosts([]string{"172.24.18.25:27017", "172.24.18.25:27018", "172.24.18.25:27019"})
//	conOpts.Replica("myapp")
//
//	db, err := mongodb.ConnectDB(conOpts, "ReforceID")
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	raw := db.Collection("test").FindOne(context.TODO(), bson.M{})
//	fmt.Println(raw.Err())
//	fmt.Println(raw.DecodeBytes())
//	db2.Tokens = db.Collection("tokens")
//	db2.Users = db.Collection("users")
//
//}
