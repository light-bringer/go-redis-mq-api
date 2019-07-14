package main

import (
	"fmt"
	redis "github.com/go-redis/redis"
)
func RedisNewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "docker-server.cloudapp.net:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return client
}


// init seeds some ridiculous initial data
func init() {
	CreateUser(UserObject{
		Fname: "Debapriya",
		Lname: "Das",
		Age : 24,
	})
	CreateUser(UserObject{
		Fname: "Anuja",
		Lname: "Saha",
		Age : 21,
	})
}


func CreateUser(u UserObject) {

	c := RedisNewClient()
	defer c.Close()

	var user User = NewUser(u.Fname, u.Lname, u.Age)
	uuid := user.ID

	// HandleError(err)
	reply, err := c.HMSet(uuid, UserMap(&user)).Result()
	HandleError(err)

	fmt.Println("POST", reply, uuid)

}

