package service

import (
	"blogs/model"
	"blogs/repo"
	"fmt"
	"strings"
)

type BlogService struct {
	BlogRepository *repo.BlogRepository
	BlogRecommendationRepository *repo.BlogRecommendationRepository
}

func (service *BlogService) GetById(id string) (*model.Blog, error) {
	blog, err := service.BlogRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("blog with id %s not found", id))
	}
	return &blog, nil
}

func (service *BlogService) GetAll() (*[]model.Blog, error) {
	blogs, err := service.BlogRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("no blogs were found")
	}
	return &blogs, nil
}

func (service *BlogService) Create(blog *model.Blog) error {
	err := service.BlogRepository.Save(blog)
	if err != nil {
		return err
	}
	return nil
}

func (service *BlogService) SearchByName(name string) (*[]model.Blog, error) {
	blogs, err := service.BlogRepository.GetAll()
	var filteredBlogs []model.Blog
	for _, blog := range blogs {
		if strings.Contains(strings.ToLower(blog.Title), strings.ToLower(name)) {
			filteredBlogs = append(filteredBlogs, blog)
		}
	}
	if err != nil {
		return nil, fmt.Errorf("no blogs were found")
	}
	return &filteredBlogs, nil
}

func (service *BlogService) Publish(id string) (model.Blog, error) {
	blog, err := service.BlogRepository.UpdateStatus(id, model.BlogStatus(model.Published))
	if err != nil {
		return blog, err
	}
	return blog, nil
}

func (service *BlogService) GetByAuthorId(id string) (*[]model.Blog, error) {
	blogs, err := service.BlogRepository.GetByAuthorId(id)
	if err != nil {
		return nil, fmt.Errorf("no blogs were found")
	}
	return &blogs, nil
}

func (service *BlogService) GetByAuthorIds(authorIds []string) (*[]model.Blog, error) {
	blogs, err := service.BlogRepository.GetByAuthorIds(authorIds)
	if err != nil {
		return nil, fmt.Errorf("no blogs were found")
	}
	return &blogs, nil
}

func (service *BlogService) GetRecommendationsByReceiverId(receiverId int) (*[]model.BlogRecommendation, error) {
	blogs, _ := service.BlogRepository.GetAll()
	var brs []model.BlogRecommendation
	for _, blog := range blogs {
		for _, br := range blog.Recommendations {
			if br.RecommendationReceiverId == receiverId {
				brs = append(brs, br)
			}
		}
	}

	return &brs, nil
}