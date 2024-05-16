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

func (handler *VoteHandler) GetVoteById(ctx context.Context, request *votes.GetVoteByIdRequest) (*votes.GetVoteByIdResponse, error) {
	vote, _ := handler.VoteService.GetById(request.ID)

	voteResponse := votes.Vote{}
	voteResponse.Id = int32(vote.ID)
	voteResponse.UserId = int32(vote.UserId)
	voteResponse.BlogId = int32(vote.BlogId)
	voteResponse.Type = votes.Vote_VoteType(vote.Type)

	ret := &votes.GetVoteByIdResponse{
		Vote: &voteResponse,
	}

	return ret, nil
}

func (handler *VoteHandler) GetAllVotes(ctx context.Context, request *votes.GetAllVotesRequest) (*votes.GetAllVotesResponse, error) {
	voteList, _ := handler.VoteService.GetAll()

	votesResponse := make([]*votes.Vote, len(*voteList))

	if voteList != nil && len(*voteList) > 0 {
		for i, vote := range *voteList {
			votesResponse[i] = &votes.Vote{
				Id:     int32(vote.ID),
				UserId: int32(vote.UserId),
				BlogId: int32(vote.BlogId),
				Type:   votes.Vote_VoteType(vote.Type),
			}
		}
	}

	ret := &votes.GetAllVotesResponse{
		Votes: votesResponse,
	}

	return ret, nil
}
