package handler

import (
	"blogs/model"
	"blogs/proto/blogs"
	"blogs/service"
	"context"
	"strings"
)

type BlogHandler struct {
	BlogService *service.BlogService
	blogs.UnimplementedBlogsServiceServer
}

func (handler *BlogHandler) GetById(ctx context.Context, request *blogs.GetByIdRequest) (*blogs.GetByIdResponse, error) {
	blog, _ := handler.BlogService.GetById(request.ID)

	blogResponse := blogs.Blog{}
	blogResponse.ID = int32(blog.ID)
	blogResponse.Title = blog.Title
	blogResponse.Description = blog.Description
	blogResponse.Status = blogs.Blog_BlogStatus(blog.Status)
	blogResponse.AuthorId = int32(blog.AuthorId)
	//blogResponse.Comments = blog.Comments
	//blogResponse.Votes = blog.Votes
	//blogResponse.Recommendations = blog.Recommendations

	ret := &blogs.GetByIdResponse{
		Blog: &blogResponse,
	}

	return ret, nil
}

func (handler *BlogHandler) GetAll(ctx context.Context, request *blogs.GetAllRequest) (*blogs.GetAllResponse, error) {
	blogList, _ := handler.BlogService.GetAll()

	blogsResponse := make([]*blogs.Blog, len(*blogList))

	if blogList != nil && len(*blogList) > 0 {
		for i, blog := range *blogList {
			blogsResponse[i] = &blogs.Blog{
				ID:          int32(blog.ID),
				Title:       blog.Title,
				Description: blog.Description,
				Status:      blogs.Blog_BlogStatus(blog.Status),
				AuthorId:    int32(blog.AuthorId),
				//Comments : blog.Comments,
				//Votes : blog.Votes,
				//Recommendations : blog.Recommendations,
			}
		}
	}

	ret := &blogs.GetAllResponse{
		Blogs: blogsResponse,
	}

	return ret, nil
}

func (handler *BlogHandler) Create(ctx context.Context, request *blogs.CreateRequest) (*blogs.CreateResponse, error) {
	blog := model.Blog{}

	blog.ID = int(request.Blog.ID)
	blog.Title = request.Blog.Title
	blog.Description = request.Blog.Description
	blog.Status = model.BlogStatus(request.Blog.Status)
	blog.AuthorId = int(request.Blog.AuthorId)
	//blogResponse.Comments = blog.Comments
	//blogResponse.Votes = blog.Votes
	//blogResponse.Recommendations = blog.Recommendations

	handler.BlogService.Create(&blog)

	return &blogs.CreateResponse{}, nil
}

func (handler *BlogHandler) Delete(ctx context.Context, request *blogs.DeleteRequest) (*blogs.DeleteResponse, error) {
	handler.BlogService.Delete(request.ID)
	return &blogs.DeleteResponse{}, nil
}

func (handler *BlogHandler) SearchByName(ctx context.Context, request *blogs.SearchByNameRequest) (*blogs.SearchByNameResponse, error) {
	blogList, _ := handler.BlogService.SearchByName(request.Title)

	blogsResponse := make([]*blogs.Blog, len(*blogList))

	if blogList != nil && len(*blogList) > 0 {
		for i, blog := range *blogList {
			blogsResponse[i] = &blogs.Blog{
				ID:          int32(blog.ID),
				Title:       blog.Title,
				Description: blog.Description,
				Status:      blogs.Blog_BlogStatus(blog.Status),
				AuthorId:    int32(blog.AuthorId),
				//Comments : blog.Comments,
				//Votes : blog.Votes,
				//Recommendations : blog.Recommendations,
			}
		}
	}

	ret := &blogs.SearchByNameResponse{
		Blogs: blogsResponse,
	}

	return ret, nil
}

func (handler *BlogHandler) Publish(ctx context.Context, request *blogs.PublishRequest) (*blogs.PublishResponse, error) {
	handler.BlogService.Publish(request.ID)
	return &blogs.PublishResponse{}, nil
}

func (handler *BlogHandler) GetByAuthorsId(ctx context.Context, request *blogs.GetByAuthorsIdRequest) (*blogs.GetByAuthorsIdResponse, error) {
	blogList, _ := handler.BlogService.GetByAuthorId(string(request.AuthorId))

	blogsResponse := make([]*blogs.Blog, len(*blogList))

	if blogList != nil && len(*blogList) > 0 {
		for i, blog := range *blogList {
			blogsResponse[i] = &blogs.Blog{
				ID:          int32(blog.ID),
				Title:       blog.Title,
				Description: blog.Description,
				Status:      blogs.Blog_BlogStatus(blog.Status),
				AuthorId:    int32(blog.AuthorId),
				//Comments : blog.Comments,
				//Votes : blog.Votes,
				//Recommendations : blog.Recommendations,
			}
		}
	}

	ret := &blogs.GetByAuthorsIdResponse{
		Blogs: blogsResponse,
	}

	return ret, nil
}

func (handler *BlogHandler) GetByAuthorsIds(ctx context.Context, request *blogs.GetByAuthorsIdsRequest) (*blogs.GetByAuthorsIdsResponse, error) {
	authorIds := strings.Split(request.AuthorsIds, ",")
	blogList, _ := handler.BlogService.GetByAuthorIds(authorIds)

	blogsResponse := make([]*blogs.Blog, len(*blogList))

	if blogList != nil && len(*blogList) > 0 {
		for i, blog := range *blogList {
			blogsResponse[i] = &blogs.Blog{
				ID:          int32(blog.ID),
				Title:       blog.Title,
				Description: blog.Description,
				Status:      blogs.Blog_BlogStatus(blog.Status),
				AuthorId:    int32(blog.AuthorId),
				//Comments : blog.Comments,
				//Votes : blog.Votes,
				//Recommendations : blog.Recommendations,
			}
		}
	}

	ret := &blogs.GetByAuthorsIdsResponse{
		Blogs: blogsResponse,
	}

	return ret, nil
}
