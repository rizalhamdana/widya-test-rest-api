package main

import (
	"github.com/rizalhamdana/widya-test/config"
	"github.com/rizalhamdana/widya-test/middleware"
	"github.com/rizalhamdana/widya-test/modules/share"

	authHandler "github.com/rizalhamdana/widya-test/modules/api/auth/presenter"
	authRepository "github.com/rizalhamdana/widya-test/modules/api/auth/repo"
	authUsecase "github.com/rizalhamdana/widya-test/modules/api/auth/usecase"

	profileHandler "github.com/rizalhamdana/widya-test/modules/api/profile/presenter"
	profileRepository "github.com/rizalhamdana/widya-test/modules/api/profile/repo"
	profileUsecase "github.com/rizalhamdana/widya-test/modules/api/profile/usecase"
)

type Service struct {
	AuthHandler          *authHandler.HttpHandler
	ProfileHandler       *profileHandler.HttpHandler
	AuthMiddlewareCustom *middleware.AuthMiddlewareCustom
}

func MakeHandler() *Service {
	WritePostgres, _ := config.WritePostgres()
	ReadPostgres, _ := config.ReadPostgres()

	// define parent repository from shared
	repository := share.NewRepository(WritePostgres, ReadPostgres)

	authMiddlewareCustom := middleware.NewAuthMiddlewareCustom(repository)

	authRepo := authRepository.NewAuthRepositoryImpl(repository)
	authUsecase := authUsecase.NewAuthUsecaseImpl(authRepo)
	authHandler := authHandler.NewAuthHandler(authUsecase)

	profileRepo := profileRepository.NewProfileRepoImpl(repository)
	profileUsecase := profileUsecase.NewProfileUsecaseImpl(profileRepo)
	profileHandler := profileHandler.NewProfileHandler(profileUsecase)

	return &Service{
		AuthHandler:          authHandler,
		ProfileHandler:       profileHandler,
		AuthMiddlewareCustom: authMiddlewareCustom,
	}

}
