package main


import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	goredis "github.com/gomodule/redigo/redis"
	"github.com/rueian/rueidis"
)

func main(){
	goredisClient()
	redigoClient()
	rueidisClient()
}


//example for goredis client
func goredisClient() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	ctx := context.Background()

	//ping redis
	err := client.Ping(context.Background()).Err()

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	log.Println("using go-redis client")

	//set value for a key
	client.Set(ctx, "go-redis", "github.com/go-redis/redis", 0)
	log.Println("executed SET")

	//get value for the key
	r := client.Get(ctx, "go-redis").Val()
	log.Println("executed GET")
	log.Println("value for go-redis", r)

}

//example for redigo client
func redigoClient() {

	//establish connection to redis
	c, err := goredis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)

	}
	defer c.Close()

	log.Println("using redigo client")

	//set value for a key
	c.Do("SET", "redigo", "github.com/gomodule/redigo/redis")
	log.Println("executed SET")

	//get value for the key
	s, _ := goredis.String(c.Do("GET", "redigo"))
	log.Println("executed GET")

	log.Println("value for redigo =", s)
}

//example for rueids client
func rueidsClient() {

	//create new redis client instance
	c, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"localhost:6379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	ctx := context.Background()

	log.Println("using rueidis client")

	//set value for a key
	c.Do(ctx, c.B().Set().Key("rueids").Value("github.com/rueian/rueidis").Nx().Build()).Error()
	log.Println("executed SET")

	//get value for the key
	r, _ := c.Do(ctx, c.B().Get().Key("rueids").Build()).ToString()
	log.Println("executed GET")

	log.Println("value for rueidis =", r)

}