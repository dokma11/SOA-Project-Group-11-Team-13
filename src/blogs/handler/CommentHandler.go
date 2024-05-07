package handler

import (
	"blogs/model"
	"blogs/proto/comments"
	"blogs/service"
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

type CommentHandler struct {
	CommentService *service.CommentService
	comments.UnimplementedCommentsServiceServer
}

func (handler *CommentHandler) GetById(ctx context.Context, request *comments.GetByIdRequest) (*comments.GetByIdResponse, error) {
	comment, _ := handler.CommentService.GetById(request.ID)

	commentResponse := comments.Comment{}
	commentResponse.ID = int32(comment.ID)
	commentResponse.AuthorId = int32(comment.AuthorId)
	commentResponse.BlogId = int32(comment.BlogId)
	commentResponse.Text = comment.Text
	commentResponse.CreatedAt = TimeToProtoTimestamp(comment.CreatedAt)
	commentResponse.UpdatedAt = TimeToProtoTimestamp(comment.UpdatedAt)

	ret := &comments.GetByIdResponse{
		Comment: &commentResponse,
	}

	return ret, nil
}

func (handler *CommentHandler) GetByBlogId(ctx context.Context, request *comments.GetByBlogIdRequest) (*comments.GetByBlogIdResponse, error) {
	commentList, _, _ := handler.CommentService.GetByBlogId(request.BlogId, int(request.Page), int(request.PageSize))

	commentResponse := make([]*comments.Comment, len(commentList))

	if commentList != nil && len(commentList) > 0 {
		for i, comment := range commentList {
			commentResponse[i] = &comments.Comment{
				ID:        int32(comment.ID),
				AuthorId:  int32(comment.AuthorId),
				BlogId:    int32(comment.BlogId),
				Text:      comment.Text,
				CreatedAt: TimeToProtoTimestamp(comment.CreatedAt),
				UpdatedAt: TimeToProtoTimestamp(comment.UpdatedAt),
			}
		}
	}

	ret := &comments.GetByBlogIdResponse{
		Comments: commentResponse,
	}

	return ret, nil
}

func (handler *CommentHandler) GetAll(ctx context.Context, request *comments.GetAllRequest) (*comments.GetAllResponse, error) {
	commentList, _ := handler.CommentService.GetAll()

	commentResponse := make([]*comments.Comment, len(*commentList))

	if commentList != nil && len(*commentList) > 0 {
		for i, comment := range *commentList {
			commentResponse[i] = &comments.Comment{
				ID:        int32(comment.ID),
				AuthorId:  int32(comment.AuthorId),
				BlogId:    int32(comment.BlogId),
				Text:      comment.Text,
				CreatedAt: TimeToProtoTimestamp(comment.CreatedAt),
				UpdatedAt: TimeToProtoTimestamp(comment.UpdatedAt),
			}
		}
	}

	ret := &comments.GetAllResponse{
		Comments: commentResponse,
	}

	return ret, nil
}

func (handler *CommentHandler) Create(ctx context.Context, request *comments.CreateRequest) (*comments.CreateResponse, error) {
	comment := model.Comment{}

	comment.ID = int(request.Comment.ID)
	comment.AuthorId = int(request.Comment.AuthorId)
	comment.BlogId = int(request.Comment.BlogId)
	comment.Text = request.Comment.Text
	comment.CreatedAt, _ = ProtoTimestampToTime(request.Comment.CreatedAt)
	comment.UpdatedAt, _ = ProtoTimestampToTime(request.Comment.UpdatedAt)

	handler.CommentService.Create(&comment)

	return &comments.CreateResponse{}, nil
}

func (handler *CommentHandler) Delete(ctx context.Context, request *comments.DeleteRequest) (*comments.DeleteResponse, error) {
	handler.CommentService.Delete(request.ID)
	return &comments.DeleteResponse{}, nil
}

func (handler *CommentHandler) Update(ctx context.Context, request *comments.UpdateRequest) (*comments.UpdateResponse, error) {
	comment := model.Comment{}

	comment.ID = int(request.Comment.ID)
	comment.AuthorId = int(request.Comment.AuthorId)
	comment.BlogId = int(request.Comment.BlogId)
	comment.Text = request.Comment.Text
	comment.CreatedAt, _ = ProtoTimestampToTime(request.Comment.CreatedAt)
	comment.UpdatedAt, _ = ProtoTimestampToTime(request.Comment.UpdatedAt)

	handler.CommentService.Update(&comment)

	return &comments.UpdateResponse{}, nil
}

func TimeToProtoTimestamp(t time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(t)
	return ts
}

func ProtoTimestampToTime(ts *timestamp.Timestamp) (time.Time, error) {
	return ptypes.Timestamp(ts)
}
