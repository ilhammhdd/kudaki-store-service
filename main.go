package main

import (
	"os"
	"strings"

	"github.com/ilhammhdd/kudaki-store-service/externals/mysql"

	"github.com/ilhammhdd/go-toolkit/safekit"
	"github.com/ilhammhdd/kudaki-store-service/externals/event_driven"
)

func init() {
	for _, val := range os.Args[1:] {
		f := strings.Split(val, " ")
		os.Setenv(string(f[1]), f[2])
	}

	mysql.OpenDB(os.Getenv("DB_PATH"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}

func main() {
	wp := safekit.NewWorkerPool()

	wp.Work <- event_driven.AddStorefrontItem

	wp.PoolWG.Wait()
}
