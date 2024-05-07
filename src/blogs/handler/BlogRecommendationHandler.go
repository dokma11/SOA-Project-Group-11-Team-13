package handler

import (
	"blogs/model"
	"blogs/proto/blog_recommendations"
	"blogs/service"
	"context"
)

type BlogRecommendationHandler struct {
	BlogRecommendationService *service.BlogRecommendationService
	blog_recommendations.UnimplementedBlogRecommendationServiceServer
}

func (handler *BlogRecommendationHandler) GetBlogRecommendationById(ctx context.Context, request *blog_recommendations.GetBlogRecommendationByIdRequest) (*blog_recommendations.GetBlogRecommendationByIdResponse, error) {
	recommendation, _ := handler.BlogRecommendationService.GetById(request.ID)

	recommendationResponse := blog_recommendations.BlogRecommendation{}
	recommendationResponse.ID = int32(recommendation.ID)
	recommendationResponse.BlogId = int32(recommendation.BlogId)
	recommendationResponse.RecommenderId = int32(recommendation.RecommenderId)
	recommendationResponse.RecommendationReceiverId = int32(recommendation.RecommendationReceiverId)
	recommendationResponse.Blog = &blog_recommendations.RecommenderBlog{
		ID:          int32(recommendation.Blog.ID),
		Title:       recommendation.Blog.Title,
		Description: recommendation.Blog.Description,
		Status:      blog_recommendations.RecommenderBlog_BlogStatus(recommendation.Blog.Status),
		AuthorId:    int32(recommendation.Blog.AuthorId),
		//Comments : recommendation.Blog.Comments,
		//Votes : recommendation.Blog.Votes,
		//Recommendations : recommendation.Blog.Recommendations,
	}

	ret := &blog_recommendations.GetBlogRecommendationByIdResponse{
		Recommendation: &recommendationResponse,
	}

	return ret, nil
}

func (handler *BlogRecommendationHandler) GetAllBlogRecommendations(ctx context.Context, request *blog_recommendations.GetAllBlogRecommendationsRequest) (*blog_recommendations.GetAllBlogRecommendationsResponse, error) {
	recommendationList, _ := handler.BlogRecommendationService.GetAll()

	recommendationResponse := make([]*blog_recommendations.BlogRecommendation, len(*recommendationList))

	if recommendationList != nil && len(*recommendationList) > 0 {
		for i, recommendation := range *recommendationList {
			recommendationResponse[i] = &blog_recommendations.BlogRecommendation{
				ID:                       int32(recommendation.ID),
				BlogId:                   int32(recommendation.BlogId),
				RecommenderId:            int32(recommendation.RecommenderId),
				RecommendationReceiverId: int32(recommendation.RecommendationReceiverId),
				Blog: &blog_recommendations.RecommenderBlog{
					ID:          int32(recommendation.Blog.ID),
					Title:       recommendation.Blog.Title,
					Description: recommendation.Blog.Description,
					Status:      blog_recommendations.RecommenderBlog_BlogStatus(recommendation.Blog.Status),
					AuthorId:    int32(recommendation.Blog.AuthorId),
					//Comments : recommendation.Blog.Comments,
					//Votes : recommendation.Blog.Votes,
					//Recommendations : recommendation.Blog.Recommendations,
				},
			}
		}
	}

	ret := &blog_recommendations.GetAllBlogRecommendationsResponse{
		Recommendations: recommendationResponse,
	}

	return ret, nil
}

func (handler *BlogRecommendationHandler) CreateBlogRecommendation(ctx context.Context, request *blog_recommendations.CreateBlogRecommendationRequest) (*blog_recommendations.CreateBlogRecommendationResponse, error) {
	recommendation := model.BlogRecommendation{}

	recommendation.ID = int(request.Recommendation.ID)
	recommendation.BlogId = int(request.Recommendation.BlogId)
	recommendation.RecommenderId = int(request.Recommendation.RecommenderId)
	recommendation.RecommendationReceiverId = int(request.Recommendation.RecommendationReceiverId)
	recommendation.Blog = model.Blog{
		ID:          int(request.Recommendation.Blog.ID),
		Title:       request.Recommendation.Blog.Title,
		Description: request.Recommendation.Blog.Description,
		Status:      model.BlogStatus(request.Recommendation.Blog.Status),
		AuthorId:    int(request.Recommendation.Blog.AuthorId),
		//Comments : recommendation.Blog.Comments,
		//Votes : recommendation.Blog.Votes,
		//Recommendations : recommendation.Blog.Recommendations,
	}
	handler.BlogRecommendationService.Create(&recommendation)

	return &blog_recommendations.CreateBlogRecommendationResponse{}, nil
}

func (handler *BlogRecommendationHandler) GetBlogRecommendationByReceiverId(ctx context.Context, request *blog_recommendations.GetBlogRecommendationByReceiverIdRequest) (*blog_recommendations.GetBlogRecommendationByReceiverIdResponse, error) {
	recommendationList, _ := handler.BlogRecommendationService.GetByReceiverId(int(request.ReceiverId))

	recommendationResponse := make([]*blog_recommendations.BlogRecommendation, len(*recommendationList))

	if recommendationList != nil && len(*recommendationList) > 0 {
		for i, recommendation := range *recommendationList {
			recommendationResponse[i] = &blog_recommendations.BlogRecommendation{
				ID:                       int32(recommendation.ID),
				BlogId:                   int32(recommendation.BlogId),
				RecommenderId:            int32(recommendation.RecommenderId),
				RecommendationReceiverId: int32(recommendation.RecommendationReceiverId),
				Blog: &blog_recommendations.RecommenderBlog{
					ID:          int32(recommendation.Blog.ID),
					Title:       recommendation.Blog.Title,
					Description: recommendation.Blog.Description,
					Status:      blog_recommendations.RecommenderBlog_BlogStatus(recommendation.Blog.Status),
					AuthorId:    int32(recommendation.Blog.AuthorId),
					//Comments : recommendation.Blog.Comments,
					//Votes : recommendation.Blog.Votes,
					//Recommendations : recommendation.Blog.Recommendations,
				},
			}
		}
	}

	ret := &blog_recommendations.GetBlogRecommendationByReceiverIdResponse{
		Recommendations: recommendationResponse,
	}

	return ret, nil
}
