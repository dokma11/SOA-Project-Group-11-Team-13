// package repo

// import (
// 	"blogs/model"

// 	"gorm.io/gorm"
// )

// type BlogRepository struct {
// 	DatabaseConnection *gorm.DB
// }

// func (repo *BlogRepository) GetById(id string) (model.Blog, error) {
// 	var blog model.Blog
// 	dbResult := repo.DatabaseConnection.Preload("Comments").Where("id = ?", id).First(&blog)
// 	if dbResult.Error != nil {
// 		return blog, dbResult.Error
// 	}
// 	return blog, nil
// }

// func (repo *BlogRepository) GetAll() ([]model.Blog, error) {
// 	var blogs []model.Blog
// 	dbResult := repo.DatabaseConnection.Find(&blogs)
// 	if dbResult.Error != nil {
// 		return nil, dbResult.Error
// 	}
// 	return blogs, nil
// }

// func (repo *BlogRepository) Save(blog *model.Blog) error {
// 	dbResult := repo.DatabaseConnection.Create(blog)
// 	if dbResult.Error != nil {
// 		return dbResult.Error
// 	}
// 	println("Rows affected: ", dbResult.RowsAffected)
// 	return nil
// }

// func (repo *BlogRepository) UpdateStatus(id string, status model.BlogStatus) (model.Blog, error) {
// 	var blog model.Blog
// 	dbResult := repo.DatabaseConnection.Where("id = ?", id).First(&blog)
// 	if dbResult.Error != nil {
// 		return blog, dbResult.Error
// 	}
// 	blog.Status = status
// 	updateDbResult := repo.DatabaseConnection.Model(&model.Blog{}).Where("id = ?", id).Updates(blog)
// 	if updateDbResult.Error != nil {
// 		return blog, dbResult.Error
// 	}
// 	return blog, nil
// }

package repo

import (
	"blogs/model"
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository struct {
	DatabaseConnection *mongo.Client
}

func (repo *BlogRepository) GetById(id string) (model.Blog, error) {
	var blog model.Blog
	converted, _ := strconv.Atoi(id)
	filter := bson.D{{Key: "id", Value: converted}}
	blogsCollection := repo.getCollection();
	err := blogsCollection.FindOne(context.Background(), filter).Decode(&blog)
	if err != nil {
		return blog, err
	}
	return blog, nil
}

func (repo *BlogRepository) GetAll() ([]model.Blog, error) {
	var blogs = make([]model.Blog, 0)
	blogsCollection := repo.getCollection();
	cur, err := blogsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var blog model.Blog
		err := cur.Decode(&blog)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (repo *BlogRepository) GetByAuthorId(authorId string) ([]model.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var blogs []model.Blog

	converted, _ := strconv.Atoi(authorId)
	filter := bson.D{{Key: "authorid", Value: converted}}

	blogsCollection := repo.getCollection();

	cur, err := blogsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var blog model.Blog
		err := cur.Decode(&blog)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (repo *BlogRepository) GetByAuthorIds(authorIds []string) ([]model.Blog, error) {

	var blogs []model.Blog
	for _, authorId := range authorIds {
		tempBlogs, _ := repo.GetByAuthorId(authorId)
		blogs = append(blogs, tempBlogs...)
	}

	return blogs, nil
}

func (repo *BlogRepository) Save(blog *model.Blog) error {
	blogsCollection := repo.getCollection();
	blog.ID = repo.nextId()
	_, err := blogsCollection.InsertOne(context.Background(), blog)
	if err != nil {
		return err
	}
	return nil
}

func (repo *BlogRepository) UpdateStatus(id string, status model.BlogStatus) (model.Blog, error) {
	blogsCollection := repo.getCollection();
	converted, _ := strconv.Atoi(id)
	filter := bson.M{"id": converted}
	update := bson.M{"$set": bson.M{"status": status}}
	var updatedBlog model.Blog
	err := blogsCollection.FindOneAndUpdate(context.Background(), filter, update).Decode(&updatedBlog)
	if err != nil {
		return updatedBlog, err
	}
	return updatedBlog, nil
}

func (repo *BlogRepository) getCollection() *mongo.Collection {
	patientDatabase := repo.DatabaseConnection.Database("soa")
	patientsCollection := patientDatabase.Collection("blogs")
	return patientsCollection
}

func (repo *BlogRepository) nextId() int {
	blogs, _ := repo.GetAll();

	maxId := 0;
	for _, blog := range blogs {
		if blog.ID > maxId {
			maxId = blog.ID
		}
	}

	return maxId + 1
}

func (repo *BlogRepository) Delete(id string) error {
	blogsCollection := repo.getCollection()
	converted, _ := strconv.Atoi(id)
	filter := bson.M{"id": converted}

	// Find the blog by ID
	var blog model.Blog
	err := blogsCollection.FindOne(context.Background(), filter).Decode(&blog)
	if err != nil {
		return err
	}

	// Delete the blog
	_, err = blogsCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
