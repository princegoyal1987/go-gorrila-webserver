package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	r "github.com/revel/revel"
	"log"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
	db, err := sql.Open("sqlite3", "testgorp.sqlite")
	if err != nil {
		log.Fatalln("failed to open db", err)
	}
	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTable(User{}).SetKeys(true, "UserId")
	setColumnSizes(t, map[string]int{
		"Name":           250,
		"HashedPassword": 250,
		"DeviceId":       250,
		"Email":          250,
	})

	t = Dbm.AddTable(UserCurrency{}).SetKeys(true, "UserId")
	setColumnSizes(t, map[string]int{
		"CurrencyId": 128,
	})

	Dbm.TraceOn("[gorp]", r.INFO)
	Dbm.CreateTables()

}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
