package service

import (
	"context"
	"github.com/LXJ0000/go-rpc/app/user/internal/domain"
	pb "github.com/LXJ0000/go-rpc/idl/pb/user"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer

	userRepo domain.UserRepository
}

func NewUserServiceServer(userRepo domain.UserRepository) *UserServiceServer {
	return &UserServiceServer{
		userRepo: userRepo,
	}
}

func (u *UserServiceServer) Login(ctx context.Context, req *pb.UserRequest) (*pb.UserDetailResponse, error) {
	resp, err := u.userRepo.GetByUserName(ctx, req.GetUserName())
	if err != nil {
		return &pb.UserDetailResponse{
			Code: http.StatusInternalServerError,
		}, err
	}
	return &pb.UserDetailResponse{
		User: &pb.UserResponse{
			UserId:   resp.UserID,
			NickName: resp.NickName,
			UserName: resp.UserName,
		},
		Code: http.StatusOK,
	}, nil
}

func (u *UserServiceServer) Register(ctx context.Context, req *pb.UserRequest) (*pb.UserCommonResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return &pb.UserCommonResponse{
			Code: http.StatusBadRequest,
			Msg:  "密码格式有误",
		}, err
	}
	user := &domain.User{
		UserName:       req.UserName,
		NickName:       req.NickName,
		UserID:         int64(1),     // TODO
		PasswordDigest: string(hash), // TODO
	}
	if err = u.userRepo.Create(ctx, user); err != nil {
		return &pb.UserCommonResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		}, err
	}
	return &pb.UserCommonResponse{
		Code: http.StatusOK,
		Msg:  "Success",
	}, nil
}
