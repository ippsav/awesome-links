package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ippsav/awesome-links/api/cmd/api/config"
	"github.com/ippsav/awesome-links/api/ent"
	"github.com/ippsav/awesome-links/api/ent/migrate"
	_ "github.com/lib/pq"
)

func main() {
	config := &config.Config{}
	if err := config.Parse(); err != nil {
		fmt.Printf("could not parse config, err: %s ", err.Error())
		os.Exit(1)
	}
	// ent client
	c, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", config.Database.Host, config.Database.Port, config.Database.User, config.Database.DBName, config.Database.Password, config.Database.SSLMode))
	if err != nil {
		fmt.Printf("could not open connection , err: %s ", err.Error())
	}
	defer c.Close()

	ctx := context.Background()
	// running the migrations
	if err := c.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		fmt.Printf("could run the migrations, err: %s ", err.Error())
	}
}
