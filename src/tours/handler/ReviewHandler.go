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

func (handler *ReviewHandler) GetById(ctx context.Context, request *reviews.GetByIdRequest) (*reviews.GetByIdResponse, error) {
	review, _ := handler.ReviewService.GetById(request.ID)

	reviewResponse := reviews.Review{}
	reviewResponse.ID = review.ID
	reviewResponse.Rating = int32(review.Rating)
	reviewResponse.TourId = review.TourId
	reviewResponse.Comment = review.Comment
	reviewResponse.TouristId = int32(review.TouristId)
	reviewResponse.TourId = review.TourId
	reviewResponse.TourVisitDate = TimeToProtoTimestamp(review.TourVisitDate)
	reviewResponse.CommentDate = TimeToProtoTimestamp(review.CommentDate)
	reviewResponse.Images = review.Images

	ret := &reviews.GetByIdResponse{
		Review: &reviewResponse,
	}

	return ret, nil
}

func (handler *ReviewHandler) GetAll(ctx context.Context, request *reviews.GetAllRequest) (*reviews.GetAllResponse, error) {
	reviewList, _ := handler.ReviewService.GetAll()

	reviewsResponse := make([]*reviews.Review, len(*reviewList))

	if reviewList != nil && len(*reviewList) > 0 {
		for i, review := range *reviewList {
			reviewsResponse[i] = &reviews.Review{
				ID:            review.ID,
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

	ret := &reviews.GetAllResponse{
		Reviews: reviewsResponse,
	}

	return ret, nil
}

func (handler *ReviewHandler) Create(ctx context.Context, request *reviews.CreateRequest) (*reviews.CreateResponse, error) {
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

	return &reviews.CreateResponse{}, nil
}

func (handler *ReviewHandler) Delete(ctx context.Context, request *reviews.DeleteRequest) (*reviews.DeleteResponse, error) {
	handler.ReviewService.Delete(request.ID)
	return &reviews.DeleteResponse{}, nil
}

func (handler *ReviewHandler) Update(ctx context.Context, request *reviews.UpdateRequest) (*reviews.UpdateResponse, error) {
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

	return &reviews.UpdateResponse{}, nil
}
