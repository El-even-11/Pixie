package log

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func Log(format string, a ...interface{}) {
	ymd := time.Now().Format("2006-01-02")
	ymdhms := time.Now().Format("2006-01-02 15:04:05")
	file, err := os.OpenFile("../logs/"+ymd, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		log.Printf("open file failed")
		return
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(fmt.Sprintf(ymdhms+" "+format+"\n", a...))
	writer.Flush()
}
