package repo

import (
	"blogs/model"
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepository struct {
	MongoConnection *mongo.Client
}

func (repo *CommentRepository) GetById(id string) (model.Comment, error) {
	var comment model.Comment
	converted, _ := strconv.Atoi(id)
	filter := bson.D{{Key: "id", Value: converted}}
	commentsCollection := repo.getCollection()
	err := commentsCollection.FindOne(context.Background(), filter).Decode(&comment)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (repo *CommentRepository) GetByBlogId(id string) ([]model.Comment, int, error) {
	var comments = make([]model.Comment, 0)
	var totalCount int64

	converted, _ := strconv.Atoi(id)
	filter := bson.D{{Key: "blogid", Value: converted}}
	commentsCollection := repo.getCollection()
	cur, err := commentsCollection.Find(context.Background(), filter)

	if err != nil {
		return nil, 0, err
	}

	for cur.Next(context.Background()) {
		var comment model.Comment
		err := cur.Decode(&comment)
		if err != nil {
			return nil, 0, err
		}
		comments = append(comments, comment)
	}
	if err := cur.Err(); err != nil {
		return nil, 0, err
	}

	return comments, int(totalCount), nil
}

func (repo *CommentRepository) GetAll() ([]model.Comment, error) {
	var comments = make([]model.Comment, 0)

	commentsCollection := repo.getCollection()
	cur, err := commentsCollection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		var comment model.Comment
		err := cur.Decode(&comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (repo *CommentRepository) Create(comment *model.Comment) error {

	commentsCollection := repo.getCollection()
	comment.ID = repo.nextId()
	_, err := commentsCollection.InsertOne(context.Background(), comment)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CommentRepository) Delete(id string) error {

	commentsCollection := repo.getCollection()
	converted, _ := strconv.Atoi(id)
	filter := bson.M{"id": converted}

	// Find the comment by ID
	var comment model.Comment
	err := commentsCollection.FindOne(context.Background(), filter).Decode(&comment)
	if err != nil {
		return err
	}

	// Delete the blog
	_, err = commentsCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (repo *CommentRepository) Update(comment *model.Comment) error {
	commentsCollection := repo.getCollection()
	filter := bson.M{"id": comment.ID}
	update := bson.M{"$set": bson.M{"text": comment.Text}}
	commentsCollection.UpdateOne(context.Background(), filter, update)
	var updatedComment model.Comment
	err := commentsCollection.FindOneAndUpdate(context.Background(), filter, update).Decode(&updatedComment)

	if err != nil {
		return err
	}
	return nil
}

func (repo *CommentRepository) getCollection() *mongo.Collection {
	patientDatabase := repo.MongoConnection.Database("soa")
	patientsCollection := patientDatabase.Collection("comments")
	return patientsCollection
}

func (repo *CommentRepository) nextId() int {
	blogs, _ := repo.GetAll()

	maxId := 0
	for _, blog := range blogs {
		if blog.ID > maxId {
			maxId = blog.ID
		}
	}

	return maxId + 1
}
