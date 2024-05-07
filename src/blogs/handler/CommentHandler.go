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

func (handler *CommentHandler) GetCommentById(ctx context.Context, request *comments.GetCommentByIdRequest) (*comments.GetCommentByIdResponse, error) {
	comment, _ := handler.CommentService.GetById(request.ID)

	commentResponse := comments.Comment{}
	commentResponse.ID = int32(comment.ID)
	commentResponse.AuthorId = int32(comment.AuthorId)
	commentResponse.BlogId = int32(comment.BlogId)
	commentResponse.Text = comment.Text
	commentResponse.CreatedAt = TimeToProtoTimestamp(comment.CreatedAt)
	commentResponse.UpdatedAt = TimeToProtoTimestamp(comment.UpdatedAt)

	ret := &comments.GetCommentByIdResponse{
		Comment: &commentResponse,
	}

	return ret, nil
}

func (handler *CommentHandler) GetCommentByBlogId(ctx context.Context, request *comments.GetCommentByBlogIdRequest) (*comments.GetCommentByBlogIdResponse, error) {
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

	ret := &comments.GetCommentByBlogIdResponse{
		Comments: commentResponse,
	}

	return ret, nil
}

func (handler *CommentHandler) GetAllComments(ctx context.Context, request *comments.GetAllCommentsRequest) (*comments.GetAllCommentsResponse, error) {
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

	ret := &comments.GetAllCommentsResponse{
		Comments: commentResponse,
	}

	return ret, nil
}

func (handler *CommentHandler) CreateComment(ctx context.Context, request *comments.CreateCommentRequest) (*comments.CreateCommentResponse, error) {
	comment := model.Comment{}

	comment.ID = int(request.Comment.ID)
	comment.AuthorId = int(request.Comment.AuthorId)
	comment.BlogId = int(request.Comment.BlogId)
	comment.Text = request.Comment.Text
	comment.CreatedAt, _ = ProtoTimestampToTime(request.Comment.CreatedAt)
	comment.UpdatedAt, _ = ProtoTimestampToTime(request.Comment.UpdatedAt)

	handler.CommentService.Create(&comment)

	return &comments.CreateCommentResponse{}, nil
}

func (handler *CommentHandler) DeleteComment(ctx context.Context, request *comments.DeleteCommentRequest) (*comments.DeleteCommentResponse, error) {
	handler.CommentService.Delete(request.ID)
	return &comments.DeleteCommentResponse{}, nil
}

func (handler *CommentHandler) UpdateComment(ctx context.Context, request *comments.UpdateCommentRequest) (*comments.UpdateCommentResponse, error) {
	comment := model.Comment{}

	comment.ID = int(request.Comment.ID)
	comment.AuthorId = int(request.Comment.AuthorId)
	comment.BlogId = int(request.Comment.BlogId)
	comment.Text = request.Comment.Text
	comment.CreatedAt, _ = ProtoTimestampToTime(request.Comment.CreatedAt)
	comment.UpdatedAt, _ = ProtoTimestampToTime(request.Comment.UpdatedAt)

	handler.CommentService.Update(&comment)

	return &comments.UpdateCommentResponse{}, nil
}

func TimeToProtoTimestamp(t time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(t)
	return ts
}

func ProtoTimestampToTime(ts *timestamp.Timestamp) (time.Time, error) {
	return ptypes.Timestamp(ts)
}
