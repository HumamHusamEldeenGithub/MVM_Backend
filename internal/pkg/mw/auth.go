package mw

import (
	"context"
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	jwtToken = "jwt-token"
)

func TokenAuthorizer(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		userToken := md.Get(jwtToken)
		if len(userToken) > 0 {
			// TODO Get seceret from env var
			claims, err := jwt_manager.VerifyToken("secret", userToken[0])
			if err != nil {
				ctx = utils.UserIDToContext(ctx, "")
			} else {
				ctx = utils.UserIDToContext(ctx, claims.UserID)
			}
		}
	}

	return handler(ctx, req)
}
