package repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDatabse(t *testing.T) {
	confString := os.Getenv("DATABASE_URL")
	if confString == "" {
		t.Skip("no DATABASE_URL in env")
	}
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, confString)
	if err != nil {
		err = fmt.Errorf("error happened in sql.Open: %w", err)

		t.Fail()
	}
	defer db.Close()

	if err = db.Ping(context.Background()); err != nil {
		err = fmt.Errorf("error happened in db.Ping: %w", err)

		t.Fail()
	}

	productName := "macbook"
	searchRepo := NewSearchRepo(db)
	productsSlice, err := searchRepo.ReadProductsByName(ctx, productName)

	fmt.Printf("%+v", productsSlice)
	assert.Nil(t, err)
}
