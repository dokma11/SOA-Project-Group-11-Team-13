package repo

import (
	"blogs/model"
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRecommendationRepository struct {
	DatabaseConnection *mongo.Client
}

func (repo *BlogRecommendationRepository) GetById(id string) (model.BlogRecommendation, error) {
	var blogRecommendation model.BlogRecommendation
	converted, _ := strconv.Atoi(id)
	filter := bson.D{{Key: "id", Value: converted}}
	blogRecommendationsCollection := repo.getCollection();
	err := blogRecommendationsCollection.FindOne(context.Background(), filter).Decode(&blogRecommendation)
	if err != nil {
		return blogRecommendation, err
	}
	return blogRecommendation, nil
}

func (repo *BlogRecommendationRepository) GetAll() ([]model.BlogRecommendation, error) {
	var blogRecommendations []model.BlogRecommendation
	blogRecommendationsCollection := repo.getCollection();
	cur, err := blogRecommendationsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var blog model.BlogRecommendation
		err := cur.Decode(&blog)
		if err != nil {
			return nil, err
		}
		blogRecommendations = append(blogRecommendations, blog)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return blogRecommendations, nil
}

func (repo *BlogRecommendationRepository) Save(blogRecommendation *model.BlogRecommendation) error {
	blogRecommendationsCollection := repo.getCollection();
	blogRecommendation.ID = repo.nextId()
	_, err := blogRecommendationsCollection.InsertOne(context.Background(), blogRecommendation)
	if err != nil {
		return err
	}
	return nil
}

func (repo *BlogRecommendationRepository) GetByReceiverId(receiverId int) ([]model.BlogRecommendation, error) {
	blogRecommendations, _ := repo.GetAll();
	var filteredBlogRecommendations []model.BlogRecommendation
	for _, blogRecommendation := range blogRecommendations {
		if blogRecommendation.RecommendationReceiverId == receiverId {
			filteredBlogRecommendations = append(filteredBlogRecommendations, blogRecommendation)
		}
	}
	return filteredBlogRecommendations, nil
}

func (repo *BlogRecommendationRepository) getCollection() *mongo.Collection {
	patientDatabase := repo.DatabaseConnection.Database("soa")
	patientsCollection := patientDatabase.Collection("blog-recommendations")
	return patientsCollection
}

func (repo *BlogRecommendationRepository) nextId() int {
	blogRecommendations, _ := repo.GetAll();

	maxId := 0;
	for _, blog := range blogRecommendations {
		if blog.ID > maxId {
			maxId = blog.ID
		}
	}

	return maxId + 1
}