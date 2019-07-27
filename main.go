package main

import (
	"os"
	"strings"

	"github.com/ilhammhdd/kudaki-storefront-service/externals"

	"github.com/ilhammhdd/kudaki-storefront-service/externals/mysql"

	"github.com/ilhammhdd/go-toolkit/safekit"
)

func init() {
	if len(os.Args) > 1 {
		for _, val := range os.Args[1:] {
			f := strings.Split(val, " ")
			os.Setenv(string(f[1]), f[2])
		}
	}

	mysql.CommandDB = mysql.OpenDB(os.Getenv("DB_PATH"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	mysql.QueryDB = mysql.OpenDB(os.Getenv("QUERY_DB_PATH"), os.Getenv("QUERY_DB_USERNAME"), os.Getenv("QUERY_DB_PASSWORD"), os.Getenv("QUERY_DB_NAME"))
}

func main() {
	wp := safekit.NewWorkerPool()

	wp.Worker <- new(externals.AddStorefrontItem)
	wp.Worker <- new(externals.DeleteStorefrontItem)
	wp.Worker <- new(externals.UpdateStorefrontItem)
	wp.Worker <- new(externals.RetrieveStorefrontItems)
	wp.Worker <- new(externals.OwnerOrderApproved)
	wp.Worker <- new(externals.OwnerConfirmedReturnment)
	wp.Worker <- new(externals.UserVerificationEmailSent)
	wp.Worker <- new(externals.ItemsReviewed)

	wp.PoolWG.Wait()
}
