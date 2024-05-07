package handler

import (
	"blogs/proto/votes"
	"blogs/service"
	"context"
)

type VoteHandler struct {
	VoteService *service.VoteService
	votes.UnimplementedVotesServiceServer
}

func (handler *VoteHandler) GetById(ctx context.Context, request *votes.GetByIdRequest) (*votes.GetByIdResponse, error) {
	vote, _ := handler.VoteService.GetById(request.ID)

	voteResponse := votes.Vote{}
	voteResponse.ID = int32(vote.ID)
	voteResponse.UserId = int32(vote.UserId)
	voteResponse.BlogId = int32(vote.BlogId)
	voteResponse.Type = votes.Vote_VoteType(vote.Type)

	ret := &votes.GetByIdResponse{
		Vote: &voteResponse,
	}

	return ret, nil
}

func (handler *VoteHandler) GetAll(ctx context.Context, request *votes.GetAllRequest) (*votes.GetAllResponse, error) {
	voteList, _ := handler.VoteService.GetAll()

	votesResponse := make([]*votes.Vote, len(*voteList))

	if voteList != nil && len(*voteList) > 0 {
		for i, vote := range *voteList {
			votesResponse[i] = &votes.Vote{
				ID:     int32(vote.ID),
				UserId: int32(vote.UserId),
				BlogId: int32(vote.BlogId),
				Type:   votes.Vote_VoteType(vote.Type),
			}
		}
	}

	ret := &votes.GetAllResponse{
		Votes: votesResponse,
	}

	return ret, nil
}
