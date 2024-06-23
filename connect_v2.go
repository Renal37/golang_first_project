package main
import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)
type Product struct{
    Id bson.ObjectId `bson:"_id"`
    Model string `bson:"model"`
    Company string `bson:"company"`
    Price int `bson:"price"`
}
func main() { 
 
    // открываем соединение
    session, err := mgo.Dial("mongodb://localhost:27017/golang")
    if err != nil {
        panic(err)
    }
    defer session.Close()
     
    // получаем коллекцию
    productCollection := session.DB("productdb").C("products")
     
    p1 := &Product{Id:bson.NewObjectId(), Model:"iPhone 8", Company:"Apple", Price:64567}
    // добавляем один объект
    err = productCollection.Insert(p1)
    if err != nil{
        fmt.Println(err)
    }
     
    p2 := &Product{Id:bson.NewObjectId(), Model:"Pixel 2", Company:"Google", Price:58000}
    p3 := &Product{Id:bson.NewObjectId(), Model:"Xplay7", Company:"Vivo", Price:49560}
    // добавляем два объекта
    err = productCollection.Insert(p2, p3)
    if err != nil{
        fmt.Println(err)
    }
}