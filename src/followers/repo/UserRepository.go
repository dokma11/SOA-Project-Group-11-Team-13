package repo

import (
	"context"
	"followers/model"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

type UserRepository struct {
	driver neo4j.DriverWithContext
	logger *log.Logger
}

func New(logger *log.Logger) (*UserRepository, error) {
	//uri := os.Getenv("NEO4J_DB")
	//user := os.Getenv("NEO4J_USERNAME")
	//pass := os.Getenv("NEO4J_PASS")

	//uri := os.Getenv("neo4j://localhost:7687")
	//user := os.Getenv("neo4j")
	//pass := os.Getenv("password")

	auth := neo4j.BasicAuth("neo4j", "password", "")

	driver, err := neo4j.NewDriverWithContext("neo4j://localhost:7687", auth)
	if err != nil {
		logger.Panic(err)
		return nil, err
	}

	return &UserRepository{
		driver: driver,
		logger: logger,
	}, nil
}

func (ur *UserRepository) CheckConnection() {
	ctx := context.Background()
	err := ur.driver.VerifyConnectivity(ctx)
	if err != nil {
		ur.logger.Panic(err)
		return
	}
	ur.logger.Printf(`Neo4J server address: %s`, ur.driver.Target().Host)
}

func (ur *UserRepository) CloseDriverConnection(ctx context.Context) {
	ur.driver.Close(ctx)
}

func (ur *UserRepository) Create(user *model.User) error {
	ctx := context.Background()
	session := ur.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	savedMovie, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"CREATE (u:User) SET u.username = $username, u.password = $password, u.isActive = $isActive"+
					" RETURN u.username + ', from node ' + id(u)",
				map[string]any{"username": user.Username, "password": user.Password, "isActive": user.IsActive})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		ur.logger.Println("Error inserting User:", err)
		return err
	}
	ur.logger.Println(savedMovie.(string))
	return nil
}
