package main

import (
	"os"
	"strings"

	"github.com/ilhammhdd/kudaki-store-service/externals/eventdriven"

	"github.com/ilhammhdd/kudaki-externals/mysql"

	"github.com/ilhammhdd/go-toolkit/safekit"
)

func init() {
	if len(os.Args) > 1 {
		for _, val := range os.Args[1:] {
			f := strings.Split(val, " ")
			os.Setenv(string(f[1]), f[2])
		}
	}

	mysql.OpenDB(os.Getenv("DB_PATH"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}

func main() {
	wp := safekit.NewWorkerPool()

	wp.Worker <- new(eventdriven.AddStorefrontItem)
	wp.Worker <- new(eventdriven.UpdateStorefrontItem)
	wp.Worker <- new(eventdriven.DeleteStorefrontItem)
	wp.Worker <- new(eventdriven.CartItemAdded)

	wp.PoolWG.Wait()
}
