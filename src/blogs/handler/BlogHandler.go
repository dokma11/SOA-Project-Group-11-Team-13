package handler

import (
	"blogs/model"
	"blogs/proto/blogs"
	"blogs/proto/votes"
	"blogs/service"
	"context"
	"strconv"
	"strings"
)

type BlogHandler struct {
	BlogService *service.BlogService
	blogs.UnimplementedBlogsServiceServer
}

func (handler *BlogHandler) GetBlogById(ctx context.Context, request *blogs.GetBlogByIdRequest) (*blogs.GetBlogByIdResponse, error) {
	blog, _ := handler.BlogService.GetById(request.ID)

	commentList := make([]*blogs.BlogComment, len(blog.Comments))
	if blog.Comments != nil && len(blog.Comments) > 0 {
		for index, b := range blog.Comments {
			commentList[index] = &blogs.BlogComment{
				Id:        int32(b.ID),
				AuthorId:  int32(b.AuthorId),
				BlogId:    int32(b.BlogId),
				Text:      b.Text,
				CreatedAt: TimeToProtoTimestamp(b.CreatedAt),
				UpdatedAt: TimeToProtoTimestamp(b.UpdatedAt),
			}
		}
	}

	voteList := make([]*blogs.BlogVote, len(blog.Votes))
	if blog.Votes != nil && len(blog.Votes) > 0 {
		for index, b := range blog.Votes {
			voteList[index] = &blogs.BlogVote{
				Id:     int32(b.ID),
				UserId: int32(b.UserId),
				BlogId: int32(b.BlogId),
				Type:   blogs.BlogVote_VoteType(votes.Vote_VoteType(b.Type)),
			}
		}
	}

	recommendationList := make([]*blogs.BlogsRecommendation, len(blog.Recommendations))
	if blog.Recommendations != nil && len(blog.Recommendations) > 0 {
		for index, b := range blog.Recommendations {
			recommendationList[index] = &blogs.BlogsRecommendation{
				Id:                       int32(b.ID),
				BlogId:                   int32(b.BlogId),
				RecommenderId:            int32(b.RecommenderId),
				RecommendationReceiverId: int32(b.RecommendationReceiverId),
			}
		}
	}

	blogResponse := blogs.Blog{}
	blogResponse.Id = int32(blog.ID)
	blogResponse.Title = blog.Title
	blogResponse.Description = blog.Description
	blogResponse.Status = blogs.Blog_BlogStatus(blog.Status)
	blogResponse.AuthorId = int32(blog.AuthorId)
	blogResponse.Comments = commentList
	blogResponse.Votes = voteList
	blogResponse.Recommendations = recommendationList

	ret := &blogs.GetBlogByIdResponse{
		Blog: &blogResponse,
	}

	return ret, nil
}

func (handler *BlogHandler) GetAllBlogs(ctx context.Context, request *blogs.GetAllBlogsRequest) (*blogs.GetAllBlogsResponse, error) {
	blogList, _ := handler.BlogService.GetAll()

	blogsResponse := make([]*blogs.Blog, len(*blogList))

	if blogList != nil && len(*blogList) > 0 {
		for i, blog := range *blogList {

			commentList := make([]*blogs.BlogComment, len(blog.Comments))
			if blog.Comments != nil && len(blog.Comments) > 0 {
				for index, b := range blog.Comments {
					commentList[index] = &blogs.BlogComment{
						Id:        int32(b.ID),
						AuthorId:  int32(b.AuthorId),
						BlogId:    int32(b.BlogId),
						Text:      b.Text,
						CreatedAt: TimeToProtoTimestamp(b.CreatedAt),
						UpdatedAt: TimeToProtoTimestamp(b.UpdatedAt),
					}
				}
			}

			voteList := make([]*blogs.BlogVote, len(blog.Votes))
			if blog.Votes != nil && len(blog.Votes) > 0 {
				for index, b := range blog.Votes {
					voteList[index] = &blogs.BlogVote{
						Id:     int32(b.ID),
						UserId: int32(b.UserId),
						BlogId: int32(b.BlogId),
						Type:   blogs.BlogVote_VoteType(votes.Vote_VoteType(b.Type)),
					}
				}
			}

			recommendationList := make([]*blogs.BlogsRecommendation, len(blog.Recommendations))
			if blog.Recommendations != nil && len(blog.Recommendations) > 0 {
				for index, b := range blog.Recommendations {
					recommendationList[index] = &blogs.BlogsRecommendation{
						Id:                       int32(b.ID),
						BlogId:                   int32(b.BlogId),
						RecommenderId:            int32(b.RecommenderId),
						RecommendationReceiverId: int32(b.RecommendationReceiverId),
					}
				}
			}

			blogsResponse[i] = &blogs.Blog{
				Id:              int32(blog.ID),
				Title:           blog.Title,
				Description:     blog.Description,
				Status:          blogs.Blog_BlogStatus(blog.Status),
				AuthorId:        int32(blog.AuthorId),
				Comments:        commentList,
				Votes:           voteList,
				Recommendations: recommendationList,
			}
		}
	}

	ret := &blogs.GetAllBlogsResponse{
		Blogs: blogsResponse,
	}

	return ret, nil
}

func (handler *BlogHandler) CreateBlog(ctx context.Context, request *blogs.CreateBlogRequest) (*blogs.CreateBlogResponse, error) {
	blog := model.Blog{}

	if blog.Comments != nil && len(blog.Comments) > 0 {
		commentList := make([]model.Comment, len(blog.Comments))
		for index, b := range blog.Comments {
			commentList[index] = model.Comment{
				ID:        b.ID,
				AuthorId:  b.AuthorId,
				BlogId:    b.BlogId,
				Text:      b.Text,
				CreatedAt: b.CreatedAt,
				UpdatedAt: b.UpdatedAt,
			}
		}
		blog.Comments = commentList
	}

	if blog.Votes != nil && len(blog.Votes) > 0 {
		voteList := make([]model.Vote, len(blog.Votes))
		for index, b := range blog.Votes {
			voteList[index] = model.Vote{
				ID:     b.ID,
				UserId: b.UserId,
				BlogId: b.BlogId,
				Type:   b.Type,
			}
		}
		blog.Votes = voteList
	}

	if blog.Recommendations != nil && len(blog.Recommendations) > 0 {
		recommendationList := make([]model.BlogRecommendation, len(blog.Recommendations))
		for index, b := range blog.Recommendations {
			recommendationList[index] = model.BlogRecommendation{
				ID:                       b.ID,
				BlogId:                   b.BlogId,
				RecommenderId:            b.RecommenderId,
				RecommendationReceiverId: b.RecommendationReceiverId,
			}
		}
		blog.Recommendations = recommendationList
	}

	blog.ID = int(request.Blog.Id)
	blog.Title = request.Blog.Title
	blog.Description = request.Blog.Description
	blog.Status = model.BlogStatus(request.Blog.Status)
	blog.AuthorId = int(request.Blog.AuthorId)

	handler.BlogService.Create(&blog)

	return &blogs.CreateBlogResponse{}, nil
}

func (handler *BlogHandler) DeleteBlog(ctx context.Context, request *blogs.DeleteBlogRequest) (*blogs.DeleteBlogResponse, error) {
	handler.BlogService.Delete(request.ID)
	return &blogs.DeleteBlogResponse{}, nil
}

func (handler *BlogHandler) SearchBlogByName(ctx context.Context, request *blogs.SearchBlogByNameRequest) (*blogs.SearchBlogByNameResponse, error) {
	blogList, _ := handler.BlogService.SearchByName(request.Title)

	blogsResponse := make([]*blogs.Blog, len(*blogList))

	if blogList != nil && len(*blogList) > 0 {
		for i, blog := range *blogList {

			commentList := make([]*blogs.BlogComment, len(blog.Comments))
			if blog.Comments != nil && len(blog.Comments) > 0 {
				for index, b := range blog.Comments {
					commentList[index] = &blogs.BlogComment{
						Id:        int32(b.ID),
						AuthorId:  int32(b.AuthorId),
						BlogId:    int32(b.BlogId),
						Text:      b.Text,
						CreatedAt: TimeToProtoTimestamp(b.CreatedAt),
						UpdatedAt: TimeToProtoTimestamp(b.UpdatedAt),
					}
				}
			}

			voteList := make([]*blogs.BlogVote, len(blog.Votes))
			if blog.Votes != nil && len(blog.Votes) > 0 {
				for index, b := range blog.Votes {
					voteList[index] = &blogs.BlogVote{
						Id:     int32(b.ID),
						UserId: int32(b.UserId),
						BlogId: int32(b.BlogId),
						Type:   blogs.BlogVote_VoteType(votes.Vote_VoteType(b.Type)),
					}
				}
			}

			recommendationList := make([]*blogs.BlogsRecommendation, len(blog.Recommendations))
			if blog.Recommendations != nil && len(blog.Recommendations) > 0 {
				for index, b := range blog.Recommendations {
					recommendationList[index] = &blogs.BlogsRecommendation{
						Id:                       int32(b.ID),
						BlogId:                   int32(b.BlogId),
						RecommenderId:            int32(b.RecommenderId),
						RecommendationReceiverId: int32(b.RecommendationReceiverId),
					}
				}
			}

			blogsResponse[i] = &blogs.Blog{
				Id:              int32(blog.ID),
				Title:           blog.Title,
				Description:     blog.Description,
				Status:          blogs.Blog_BlogStatus(blog.Status),
				AuthorId:        int32(blog.AuthorId),
				Comments:        commentList,
				Votes:           voteList,
				Recommendations: recommendationList,
			}
		}
	}

	ret := &blogs.SearchBlogByNameResponse{
		Blogs: blogsResponse,
	}

	return ret, nil
}

func (handler *BlogHandler) PublishBlog(ctx context.Context, request *blogs.PublishBlogRequest) (*blogs.PublishBlogResponse, error) {
	handler.BlogService.Publish(strconv.FormatInt(int64(request.Blog.Id), 10))
	return &blogs.PublishBlogResponse{}, nil
}

func (handler *BlogHandler) GetBlogsByAuthorsId(ctx context.Context, request *blogs.GetBlogsByAuthorsIdRequest) (*blogs.GetBlogsByAuthorsIdResponse, error) {
	blogList, _ := handler.BlogService.GetByAuthorId(string(request.AuthorId))

	blogsResponse := make([]*blogs.Blog, len(*blogList))

	if blogList != nil && len(*blogList) > 0 {
		for i, blog := range *blogList {

			commentList := make([]*blogs.BlogComment, len(blog.Comments))
			if blog.Comments != nil && len(blog.Comments) > 0 {
				for index, b := range blog.Comments {
					commentList[index] = &blogs.BlogComment{
						Id:        int32(b.ID),
						AuthorId:  int32(b.AuthorId),
						BlogId:    int32(b.BlogId),
						Text:      b.Text,
						CreatedAt: TimeToProtoTimestamp(b.CreatedAt),
						UpdatedAt: TimeToProtoTimestamp(b.UpdatedAt),
					}
				}
			}

			voteList := make([]*blogs.BlogVote, len(blog.Votes))
			if blog.Votes != nil && len(blog.Votes) > 0 {
				for index, b := range blog.Votes {
					voteList[index] = &blogs.BlogVote{
						Id:     int32(b.ID),
						UserId: int32(b.UserId),
						BlogId: int32(b.BlogId),
						Type:   blogs.BlogVote_VoteType(votes.Vote_VoteType(b.Type)),
					}
				}
			}

			recommendationList := make([]*blogs.BlogsRecommendation, len(blog.Recommendations))
			if blog.Recommendations != nil && len(blog.Recommendations) > 0 {
				for index, b := range blog.Recommendations {
					recommendationList[index] = &blogs.BlogsRecommendation{
						Id:                       int32(b.ID),
						BlogId:                   int32(b.BlogId),
						RecommenderId:            int32(b.RecommenderId),
						RecommendationReceiverId: int32(b.RecommendationReceiverId),
					}
				}
			}

			blogsResponse[i] = &blogs.Blog{
				Id:              int32(blog.ID),
				Title:           blog.Title,
				Description:     blog.Description,
				Status:          blogs.Blog_BlogStatus(blog.Status),
				AuthorId:        int32(blog.AuthorId),
				Comments:        commentList,
				Votes:           voteList,
				Recommendations: recommendationList,
			}
		}
	}

	ret := &blogs.GetBlogsByAuthorsIdResponse{
		Blogs: blogsResponse,
	}

	return ret, nil
}

func (handler *BlogHandler) GetBlogsByAuthorsIds(ctx context.Context, request *blogs.GetBlogsByAuthorsIdsRequest) (*blogs.GetBlogsByAuthorsIdsResponse, error) {
	authorIds := strings.Split(request.AuthorsIds, ",")
	blogList, _ := handler.BlogService.GetByAuthorIds(authorIds)

	blogsResponse := make([]*blogs.Blog, len(*blogList))

	if blogList != nil && len(*blogList) > 0 {
		for i, blog := range *blogList {

			commentList := make([]*blogs.BlogComment, len(blog.Comments))
			if blog.Comments != nil && len(blog.Comments) > 0 {
				for index, b := range blog.Comments {
					commentList[index] = &blogs.BlogComment{
						Id:        int32(b.ID),
						AuthorId:  int32(b.AuthorId),
						BlogId:    int32(b.BlogId),
						Text:      b.Text,
						CreatedAt: TimeToProtoTimestamp(b.CreatedAt),
						UpdatedAt: TimeToProtoTimestamp(b.UpdatedAt),
					}
				}
			}

			voteList := make([]*blogs.BlogVote, len(blog.Votes))
			if blog.Votes != nil && len(blog.Votes) > 0 {
				for index, b := range blog.Votes {
					voteList[index] = &blogs.BlogVote{
						Id:     int32(b.ID),
						UserId: int32(b.UserId),
						BlogId: int32(b.BlogId),
						Type:   blogs.BlogVote_VoteType(votes.Vote_VoteType(b.Type)),
					}
				}
			}

			recommendationList := make([]*blogs.BlogsRecommendation, len(blog.Recommendations))
			if blog.Recommendations != nil && len(blog.Recommendations) > 0 {
				for index, b := range blog.Recommendations {
					recommendationList[index] = &blogs.BlogsRecommendation{
						Id:                       int32(b.ID),
						BlogId:                   int32(b.BlogId),
						RecommenderId:            int32(b.RecommenderId),
						RecommendationReceiverId: int32(b.RecommendationReceiverId),
					}
				}
			}

			blogsResponse[i] = &blogs.Blog{
				Id:              int32(blog.ID),
				Title:           blog.Title,
				Description:     blog.Description,
				Status:          blogs.Blog_BlogStatus(blog.Status),
				AuthorId:        int32(blog.AuthorId),
				Comments:        commentList,
				Votes:           voteList,
				Recommendations: recommendationList,
			}
		}
	}

	ret := &blogs.GetBlogsByAuthorsIdsResponse{
		Blogs: blogsResponse,
	}

	return ret, nil
}
