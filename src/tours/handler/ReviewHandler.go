package handler

import (
	"context"
	"tours/model"
	"tours/proto/reviews"
	"tours/service"
)

type ReviewHandler struct {
	ReviewService *service.ReviewService
	reviews.UnimplementedReviewsServiceServer
}

func (handler *ReviewHandler) GetReviewById(ctx context.Context, request *reviews.GetReviewByIdRequest) (*reviews.GetReviewByIdResponse, error) {
	review, _ := handler.ReviewService.GetById(request.ID)

	reviewResponse := reviews.Review{}
	reviewResponse.Id = review.ID
	reviewResponse.Rating = int32(review.Rating)
	reviewResponse.TourId = review.TourId
	reviewResponse.Comment = review.Comment
	reviewResponse.TouristId = int32(review.TouristId)
	reviewResponse.TourId = review.TourId
	reviewResponse.TourVisitDate = TimeToProtoTimestamp(review.TourVisitDate)
	reviewResponse.CommentDate = TimeToProtoTimestamp(review.CommentDate)
	reviewResponse.Images = review.Images

	ret := &reviews.GetReviewByIdResponse{
		Review: &reviewResponse,
	}

	return ret, nil
}

func (handler *ReviewHandler) GetAllReviews(ctx context.Context, request *reviews.GetAllReviewsRequest) (*reviews.GetAllReviewsResponse, error) {
	reviewList, _ := handler.ReviewService.GetAll()

	reviewsResponse := make([]*reviews.Review, len(*reviewList))

	if reviewList != nil && len(*reviewList) > 0 {
		for i, review := range *reviewList {
			reviewsResponse[i] = &reviews.Review{
				Id:            review.ID,
				Rating:        int32(review.Rating),
				TourId:        review.TourId,
				Comment:       review.Comment,
				TouristId:     int32(review.TouristId),
				TourVisitDate: TimeToProtoTimestamp(review.TourVisitDate),
				CommentDate:   TimeToProtoTimestamp(review.CommentDate),
				Images:        review.Images,
			}
		}
	}

	ret := &reviews.GetAllReviewsResponse{
		Reviews: reviewsResponse,
	}

	return ret, nil
}

func (handler *ReviewHandler) CreateReview(ctx context.Context, request *reviews.CreateReviewRequest) (*reviews.CreateReviewResponse, error) {
	review := model.Review{}

	review.ID = request.Review.Id
	review.Rating = int(request.Review.Rating)
	review.TourId = request.Review.TourId
	review.Comment = request.Review.Comment
	review.TouristId = int(request.Review.TouristId)
	review.TourId = request.Review.TourId
	review.TourVisitDate, _ = ProtoTimestampToTime(request.Review.TourVisitDate)
	review.CommentDate, _ = ProtoTimestampToTime(request.Review.CommentDate)
	review.Images = request.Review.Images

	handler.ReviewService.Create(&review)

	return &reviews.CreateReviewResponse{}, nil
}

func (handler *ReviewHandler) DeleteReview(ctx context.Context, request *reviews.DeleteReviewRequest) (*reviews.DeleteReviewResponse, error) {
	handler.ReviewService.Delete(request.ID)
	return &reviews.DeleteReviewResponse{}, nil
}

func (handler *ReviewHandler) UpdateReview(ctx context.Context, request *reviews.UpdateReviewRequest) (*reviews.UpdateReviewResponse, error) {
	review := model.Review{}

	review.ID = request.Review.Id
	review.Rating = int(request.Review.Rating)
	review.TourId = request.Review.TourId
	review.Comment = request.Review.Comment
	review.TouristId = int(request.Review.TouristId)
	review.TourId = request.Review.TourId
	review.TourVisitDate, _ = ProtoTimestampToTime(request.Review.TourVisitDate)
	review.CommentDate, _ = ProtoTimestampToTime(request.Review.CommentDate)
	review.Images = request.Review.Images

	handler.ReviewService.Update(&review)

	return &reviews.UpdateReviewResponse{}, nil
}
