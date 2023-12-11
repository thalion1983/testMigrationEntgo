package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testMigrationEntgo/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	pgDriver = "pgx"
	seedInfo = []struct {
		Name  string
		Email string
		Title string
		Blogs int
	}{
		{
			Name:  "User1",
			Email: "user1@gmail.com",
			Title: "User 1 title",
			Blogs: 2,
		},
		{
			Name:  "User2",
			Email: "user2@hotmail.com",
			Blogs: 1,
		},
		{
			Name:  "User3",
			Email: "user3@yahoo.com",
			Title: "User 3 title",
			Blogs: 2,
		},
	}
)

// Gets a new entgo client to a database
func getClient(connStr string) (*ent.Client, error) {
	// Open Database
	db, err := sql.Open(pgDriver, connStr)
	if err != nil {
		return nil, err
	}

	// Create driver and return
	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver)), nil
}

// seed seeds initial info into database
func seed(ctx context.Context, cli *ent.Client) error {
	for _, user := range seedInfo {
		u, err := cli.User.Create().SetName(user.Name).SetEmail(user.Email).Save(ctx)
		if err != nil {
			return fmt.Errorf("while creating user %s: %w", user.Name, err)
		}

		if user.Title != "" {
			u, err = u.Update().SetTitle(user.Title).Save(ctx)
			if err != nil {
				return fmt.Errorf("while setting title for user %s: %w", user.Name, err)
			}
		}

		for i := 0; i < user.Blogs; i++ {
			title := fmt.Sprintf("%s blog %d", user.Name, i+1)
			content := fmt.Sprintf("%s blog %d body", user.Name, i+1)
			blog, err := cli.Blog.Create().SetTitle(title).SetBody(content).Save(ctx)
			if err != nil {
				return fmt.Errorf("while creating blog %d for user %s: %w", i+1, user.Name, err)
			}
			u, err = u.Update().AddBlogPosts(blog).Save(ctx)
			if err != nil {
				return fmt.Errorf("while adding blog %d for user %s: %w", i+1, user.Name, err)
			}
		}
	}
	return nil
}

func main() {
	client, err := getClient("host=localhost port=5432 user=testuser dbname=test_migration password=testpswd")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating ORM resources: %v", err)
	}

	if err := seed(ctx, client); err != nil {
		log.Fatalf("failed seeding data: %v", err)
	}
}
