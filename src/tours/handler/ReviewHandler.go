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

func (handler *ReviewHandler) GetById(ctx context.Context, request *reviews.GetReviewByIdRequest) (*reviews.GetReviewByIdResponse, error) {
	review, _ := handler.ReviewService.GetById(request.ID)

	var reviewResponse reviews.Review
	reviewResponse.ID = review.ID
	reviewResponse.Rating = int32(review.Rating)
	reviewResponse.TourId = review.TourId
	reviewResponse.Comment = review.Comment
	reviewResponse.TouristId = int32(review.TouristId)
	reviewResponse.TourId = review.TourId
	reviewResponse.TourVisitDate = TimeToProtoTimestamp(review.TourVisitDate)
	reviewResponse.CommentDate = TimeToProtoTimestamp(review.CommentDate)
	reviewResponse.Images = review.Images

	return &reviews.GetReviewByIdResponse{
		Review: &reviewResponse,
	}, nil
}

func (handler *ReviewHandler) GetAll(ctx context.Context, request *reviews.GetAllReviewsRequest) (*reviews.GetAllReviewsResponse, error) {
	reviewList, _ := handler.ReviewService.GetAll()

	var reviewsResponse []*reviews.Review

	for i, review := range *reviewList {
		reviewsResponse[i].ID = review.ID
		reviewsResponse[i].Rating = int32(review.Rating)
		reviewsResponse[i].TourId = review.TourId
		reviewsResponse[i].Comment = review.Comment
		reviewsResponse[i].TouristId = int32(review.TouristId)
		reviewsResponse[i].TourId = review.TourId
		reviewsResponse[i].TourVisitDate = TimeToProtoTimestamp(review.TourVisitDate)
		reviewsResponse[i].CommentDate = TimeToProtoTimestamp(review.CommentDate)
		reviewsResponse[i].Images = review.Images
	}

	return &reviews.GetAllReviewsResponse{
		Reviews: reviewsResponse,
	}, nil
}

func (handler *ReviewHandler) Create(ctx context.Context, request *reviews.CreateReviewRequest) (*reviews.CreateReviewResponse, error) {
	review := model.Review{}

	review.ID = request.Review.ID
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

func (handler *ReviewHandler) Delete(ctx context.Context, request *reviews.DeleteReviewRequest) (*reviews.DeleteReviewResponse, error) {
	handler.ReviewService.Delete(request.ID)
	return &reviews.DeleteReviewResponse{}, nil
}

func (handler *ReviewHandler) Update(ctx context.Context, request *reviews.UpdateReviewRequest) (*reviews.UpdateReviewResponse, error) {
	review := model.Review{}

	review.ID = request.Review.ID
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
