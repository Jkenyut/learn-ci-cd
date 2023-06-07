package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Users []User

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)
	user := Users{
		User{
			ID:   1,
			Name: "kenyut",
		},
		User{
			ID:   2,
			Name: "halo",
		},
	}
	fmt.Println(user)

	//userValue := reflect.ValueOf(user)
	//userType := userValue.Type()
	//
	//for i := 0; i < userValue.NumField(); i++ {
	//	field := userValue.Field(i)
	//	fieldName := userType.Field(i).Name
	//	fieldValue := field.Interface()
	//use := User{
	//	ID:   1,
	//	Name: "kenyut",
	//}

	bs, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	err = client.HSet(ctx, "user-session:123", "data", bs).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Data JSON berhasil disimpan di Redis")

	// Mengambil data JSON dari Redis
	val, err := client.HGet(ctx, "user-session:123", "data").Result()
	if err != nil {
		panic(err)
	}
	var data Users
	fmt.Println("val", val)
	err = json.Unmarshal([]byte(val), &data)
	fmt.Println(user)
	fmt.Println(data)

	fmt.Println("ata dari redis", val)
}
