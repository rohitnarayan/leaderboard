package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rohitnarayan/leaderboard/internal/config"
)

type employee struct {
	Id     int     `db:"id"`
	Name   string  `db:"name"`
	Salary float64 `db:"salary"`
}

func TestDBMethods(t *testing.T) {
	ctx := context.Background()
	query := []*Query{
		{
			Query: "SELECT * FROM employee WHERE id=$1",
			Args:  []interface{}{3},
		},
		{
			Query: "INSERT INTO employee VALUES($1, $2, $3)",
			Args:  []interface{}{1000, "Byomkesh Bakshi", 37500},
		},
		{
			Query: "SELECT * FROM employee WHERE id=$1",
			Args:  []interface{}{1000},
		},
		{
			Query: "DELETE FROM employee WHERE id=$1",
			Args:  []interface{}{1000},
		},
	}

	config.InitTestConfig()
	db, _ := NewDB(config.App.Database.Postgres)

	var employees []*employee
	err := Get(ctx, db, config.App.Database.Postgres.ReadTimeout, &employees, query[0])
	assert.Equal(t, err, nil)
	assert.Equal(t, "Employee 3", employees[0].Name)

	err = Exec(ctx, db, config.App.Database.Postgres.WriteTimeout, query[3])
	assert.Equal(t, nil, err)

	err = Exec(ctx, db, config.App.Database.Postgres.WriteTimeout, query[1])
	assert.Equal(t, nil, err)

	err = Get(ctx, db, config.App.Database.Postgres.ReadTimeout, &employees, query[2])
	assert.Equal(t, err, nil)
	assert.Equal(t, "Byomkesh Bakshi", employees[0].Name)
}
