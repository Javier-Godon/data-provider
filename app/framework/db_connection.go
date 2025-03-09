package framework

//_ "github.com/jackc/pgx/v5"
import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDatabase() {
	log.Println("Initializing QuestDB")
	dbURI := AppConfig.PostgresUrl.URI

	config, err := pgxpool.ParseConfig(dbURI)
	if err != nil {
		log.Fatalf("Invalid database URI: %v", err)
	}

	// Tune connection pool settings (adjust based on your needs)
	config.MaxConns = 10                      // Limit max connections
	config.MinConns = 2                       // Keep a minimum of 2 connections
	config.MaxConnLifetime = time.Hour        // Close connections after 1 hour
	config.MaxConnIdleTime = 30 * time.Minute // Idle connections are closed after 30 minutes
	config.HealthCheckPeriod = time.Minute    // Periodically check health

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Cannot connect to QuestDB:", err)
	}

	// **Explicitly test the connection**
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = DB.Ping(ctx)
	if err != nil {
		log.Fatal("Database connection test failed:", err)
	}

	log.Println("Connected to QuestDB successfully")
}
