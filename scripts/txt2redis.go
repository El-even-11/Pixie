package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
)

func main() {
	file, err := os.OpenFile("../data/words/word", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := bufio.NewReader(file)

	i := 0
	content := []string{}
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if i%2 == 0 {
			content = append(content, line)
		} else {
			content[i/2] = strings.Join([]string{content[i/2], line}, "\n")
		}

		if err == io.EOF {
			break
		}
		i++
	}

	ctx := context.Background()

	redisDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Println(len(content))

	t := make([]interface{}, len(content))
	for i, v := range content {
		t[i] = v
	}

	redisDB.SAdd(ctx, "背单词-t", t...)
}
