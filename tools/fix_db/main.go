package main

import (
	"context"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Type: "pgsql",
				Link: "pgsql:postgres:postgres@tcp(127.0.0.1:5432)/cex",
			},
		},
	})

	db := g.DB()
	ctx := context.Background()
	var err error

	tables := []string{
		"t_app_recharge",
		"t_second_contract_order",
	}

	for _, table := range tables {
		seqName := table + "_id_seq"
		_, err = db.Exec(ctx, fmt.Sprintf(`
			CREATE SEQUENCE IF NOT EXISTS %s;
			ALTER TABLE %s ALTER COLUMN id SET DEFAULT nextval('%s');
			ALTER SEQUENCE %s OWNED BY %s.id;
		`, seqName, table, seqName, seqName, table))
		if err != nil {
			fmt.Printf("Error fixing sequence for %s: %v\n", table, err)
		} else {
			_, _ = db.Exec(ctx, fmt.Sprintf(`SELECT setval('%s', COALESCE((SELECT MAX(id)+1 FROM %s), 1), false);`, seqName, table))
			fmt.Printf("Fixed sequence for %s\n", table)
		}
	}
}
