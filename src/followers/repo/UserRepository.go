package repo

import (
	"context"
	"followers/model"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type UserRepository struct {
	driver neo4j.DriverWithContext
	logger *log.Logger
}

func New(logger *log.Logger) (*UserRepository, error) {
	uri := os.Getenv("NEO4J_DB")
	user := os.Getenv("NEO4J_USERNAME")
	pass := os.Getenv("NEO4J_PASS")

	//uri := os.Getenv("neo4j://localhost:7687")
	//user := os.Getenv("neo4j")
	//pass := os.Getenv("password")

	auth := neo4j.BasicAuth(user, pass, "")

	driver, err := neo4j.NewDriverWithContext(uri, auth)
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

	savedUser, err := session.ExecuteWrite(ctx,
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
	ur.logger.Println(savedUser.(string))
	return nil
}

func (ur *UserRepository) CreateFollowConnectionBetweenUsers(user1 *model.User, user2 *model.User) error {
	ctx := context.Background()
	session := ur.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	savedUser, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"MATCH (u1:User), (u2:User) WHERE u1.username = $username1 AND u2.username = $username2 "+
					"CREATE (u1)-[r:FOLLOWS]->(u2) RETURN type(r)",
				map[string]any{"username1": user1.Username, "username2": user2.Username})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		ur.logger.Println("Error while creating a follow branch", err)
		return err
	}
	ur.logger.Println(savedUser.(string))
	return nil
}
