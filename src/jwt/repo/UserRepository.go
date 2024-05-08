package repo

import (
	"context"
	"fmt"
	"jwt/model"
	"log"
	"os"
	"strconv"

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
				`CREATE (u:User) SET u.id = $id, u.username = $username, u.password = $password,
 					    u.isActive = $isActive, u.profilePicture = $profilePicture, u.role = $role 
						RETURN u.username + ', from node ' + id(u)`,
				map[string]any{"id": user.ID, "username": user.Username, "password": user.Password,
					"isActive": user.IsActive, "profilePicture": user.ProfilePicture, "role": user.Role})
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

func (ur *UserRepository) CreateFollowConnectionBetweenUsers(userId string, followedById string) error {
	ctx := context.Background()
	session := ur.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	savedUser, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"MATCH (u1:User {id:"+followedById+"}), (u2:User {id:"+userId+"})"+
					"CREATE (u1)-[r:FOLLOWS]->(u2) RETURN type(r)",
				map[string]any{})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		ur.logger.Println("Error while creating a follow relationship", err)
		return err
	}
	ur.logger.Println(savedUser.(string))
	return nil
}

func (ur *UserRepository) DeleteFollowConnectionBetweenUsers(userId string, followingId string) error {
	ctx := context.Background()
	session := ur.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (interface{}, error) {
			result, err := transaction.Run(ctx,
				"MATCH (node1:User {id:"+userId+"})-[f:FOLLOWS]->(node2:User {id:"+followingId+"}) DELETE f",
				map[string]any{})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		ur.logger.Println("Error while deleting a follow relationship", err)
		return err
	}
	return nil
}

func (ur *UserRepository) GetFollowers(userId string) ([]model.User, error) {
	ctx := context.Background()
	session := ur.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	var followers []model.User

	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (interface{}, error) {
			result, err := transaction.Run(ctx,
				"MATCH (user:User {id:"+userId+"}) MATCH (follower)-[:FOLLOWS]->(user) RETURN follower",
				map[string]any{"user_id": userId})
			if err != nil {
				return nil, err
			}

			for result.Next(ctx) {
				node, ok := result.Record().Get("follower")
				if !ok {
					return nil, fmt.Errorf("follower node not found")
				}

				userNode, ok := node.(neo4j.Node)
				if !ok {
					return nil, fmt.Errorf("follower node not found or not of expected type")
				}

				id, _ := userNode.Props["id"].(int64)
				username, _ := userNode.Props["username"].(string)
				password, _ := userNode.Props["password"].(string)
				role, _ := userNode.Props["role"].(model.UserRole)
				profilePicture, _ := userNode.Props["profilePicture"].(string)
				isActive, _ := userNode.Props["isActive"].(bool)

				follower := model.User{
					ID:             id,
					Username:       username,
					Password:       password,
					Role:           role,
					ProfilePicture: profilePicture,
					IsActive:       isActive,
				}
				followers = append(followers, follower)
			}
			if err := result.Err(); err != nil {
				return nil, err
			}

			return nil, nil
		})
	if err != nil {
		ur.logger.Println("Error while retrieving users followers", err)
		return nil, err
	}

	return followers, nil
}

func (ur *UserRepository) GetFollowings(userId string) ([]model.User, error) {
	ctx := context.Background()
	session := ur.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	var followedUsers []model.User

	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (interface{}, error) {
			cypherQuery := "MATCH (user:User {id: $userId})-[:FOLLOWS]->(following:User) RETURN following"
			intUserId, err := strconv.Atoi(userId)

			if err != nil {
				log.Printf("Error converting userId to integer: %v", err)
				return nil, err
			}

			result, err := transaction.Run(ctx, cypherQuery, map[string]interface{}{"userId": intUserId})

			if err != nil {
				return nil, err
			}

			for result.Next(ctx) {
				node, ok := result.Record().Get("following")
				if !ok {
					return nil, fmt.Errorf("following node not found")
				}

				userNode, ok := node.(neo4j.Node)
				if !ok {
					return nil, fmt.Errorf("following node not found or not of expected type")
				}

				id, _ := userNode.Props["id"].(int64)
				username, _ := userNode.Props["username"].(string)
				password, _ := userNode.Props["password"].(string)
				role, _ := userNode.Props["role"].(model.UserRole)
				profilePicture, _ := userNode.Props["profilePicture"].(string)
				isActive, _ := userNode.Props["isActive"].(bool)

				following := model.User{
					ID:             id,
					Username:       username,
					Password:       password,
					Role:           role,
					ProfilePicture: profilePicture,
					IsActive:       isActive,
				}
				followedUsers = append(followedUsers, following)
			}

			if err := result.Err(); err != nil {
				ur.logger.Println("Error while retrieving users' followings", err)
				return nil, err
			}

			return nil, nil
		})

	if err != nil {
		return nil, err
	}

	return followedUsers, nil
}

func (ur *UserRepository) GetByUsername(username string) (model.User, error) {
	ctx := context.Background()
	session := ur.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	var savedUser model.User

	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (interface{}, error) {
			result, err := transaction.Run(ctx,
				`MATCH (user:User {username: $user_username}) RETURN user`,
				map[string]interface{}{"user_username": username})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				node, ok := result.Record().Get("user")
				if !ok {
					return nil, fmt.Errorf("user not found")
				}

				userNode, ok := node.(neo4j.Node)
				if !ok {
					return nil, fmt.Errorf("unexpected type for user node")
				}

				id, _ := userNode.Props["id"].(int64)
				username, _ := userNode.Props["username"].(string)
				password, _ := userNode.Props["password"].(string)
				role, _ := userNode.Props["role"].(model.UserRole)
				profilePicture, _ := userNode.Props["profilePicture"].(string)
				isActive, _ := userNode.Props["isActive"].(bool)

				savedUser = model.User{
					ID:             id,
					Username:       username,
					Password:       password,
					Role:           role,
					ProfilePicture: profilePicture,
					IsActive:       isActive,
				}
				return savedUser, nil
			}

			return nil, result.Err()
		})

	if err != nil {
		ur.logger.Println("Error while finding user by his/her username", err)
		return model.User{}, err
	}

	return savedUser, nil
}

func (ur *UserRepository) GetRecommendedUsers(userId string) ([]model.User, error) {
	ctx := context.Background()
	session := ur.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	var recommendedUsers []model.User

	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (interface{}, error) {
			result, err := transaction.Run(ctx,
				"MATCH (me:User {id:"+userId+"})-[:FOLLOWS]->(myFollowing:User)-[:FOLLOWS]->(recommended:User) "+
					"WHERE NOT (me)-[:FOLLOWS]->(recommended) AND recommended <> me "+
					"RETURN recommended LIMIT 10",
				map[string]interface{}{})
			if err != nil {
				return nil, err
			}

			for result.Next(ctx) {
				node, ok := result.Record().Get("recommended")
				if !ok {
					return nil, fmt.Errorf("users not found")
				}

				userNode, ok := node.(neo4j.Node)
				if !ok {
					return nil, fmt.Errorf("unexpected type for user node")
				}

				id, _ := userNode.Props["id"].(int64)
				username, _ := userNode.Props["username"].(string)
				password, _ := userNode.Props["password"].(string)
				role, _ := userNode.Props["role"].(model.UserRole)
				profilePicture, _ := userNode.Props["profilePicture"].(string)
				isActive, _ := userNode.Props["isActive"].(bool)

				recommended := model.User{
					ID:             id,
					Username:       username,
					Password:       password,
					Role:           role,
					ProfilePicture: profilePicture,
					IsActive:       isActive,
				}
				recommendedUsers = append(recommendedUsers, recommended)
			}
			if err := result.Err(); err != nil {
				return nil, err
			}

			return nil, nil
		})

	if err != nil {
		ur.logger.Println("Error while finding recommended users", err)
		return nil, err
	}

	return recommendedUsers, nil
}
