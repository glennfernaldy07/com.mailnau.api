package service

import (
	"com.mailnau.api/common"
	comdb "com.mailnau.api/common/db"
	cerr "com.mailnau.api/common/errors"
	"com.mailnau.api/common/snap/snapauth"
	"com.mailnau.api/common/utils"
	"com.mailnau.api/config"
	_rmaDomain "com.mailnau.api/role-menu-action/domain"
	_roleDomain "com.mailnau.api/role/domain"
	"com.mailnau.api/user/domain"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type service struct {
	cfg       config.Config
	repo      domain.Repository
	roleSvc   _roleDomain.Service
	rmaSvc    _rmaDomain.Service
	cacheRepo domain.CacheRepository
	f         utils.LogFormatter
}

func NewService(cfg config.Config, repo domain.Repository, roleSvc _roleDomain.Service, rmaSvc _rmaDomain.Service, cacheRepo domain.CacheRepository) domain.Service {
	f := utils.NewLogFormatter("user.service")
	return &service{cfg: cfg, repo: repo, roleSvc: roleSvc, rmaSvc: rmaSvc, cacheRepo: cacheRepo, f: f}
}

func (s *service) LoginByEmail(ctx context.Context, req domain.LoginByEmail) (common.GeneralResponse, error) {
	//find email
	userModel, err := s.repo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		errMsg := fmt.Errorf("email not found, email=%s, error=%s", req.Email, err.Error())
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusUnauthorized, errMsg.Error(), errMsg)
	}
	// compare password
	if !utils.VerifyPassword(req.Password, userModel.Password) {
		errMsg := fmt.Errorf("invalid password")
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusUnauthorized, errMsg.Error(), errMsg)
	}

	//get role by user id
	role, err := s.roleSvc.GetRoleByID(ctx, userModel.RoleID)
	if err != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
	}

	//GET ROLE MENU ACTION
	roleMenuActions, err := s.rmaSvc.GetRoleMenuActionByRoleID(ctx, role.ID)
	if err != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
	}

	// generate token
	tokenExpTime := s.cfg.GetInt(config.TokenExpTime)
	token, errCreateToken := utils.CreateToken(userModel.ID.String(), int(tokenExpTime))
	if errCreateToken != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, errCreateToken.Error(), errCreateToken)
	}

	// store token
	dt := snapauth.AccessTokenResponse{
		AccessToken:    token,
		TokenType:      "bearer",
		ExpiresIn:      strconv.FormatInt(tokenExpTime, 10),
		AdditionalInfo: nil,
	}
	if err := s.cacheRepo.StoreAccessToken(ctx, userModel.ID.String(), dt); err != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
	}
	resp := domain.LoginDataResponse{
		Token:       token,
		UserID:      userModel.ID.String(),
		MenuActions: []_rmaDomain.RoleMenuAction{},
	}

	for _, v := range roleMenuActions {
		resp.MenuActions = append(resp.MenuActions, v)
	}
	// return token
	return common.GeneralResponse{Status: "200", Message: common.SuccessMessage, Data: resp}, nil
}

func (s *service) LoginByNIK(ctx context.Context, req domain.LoginByNIK) (common.GeneralResponse, error) {
	//find email
	userModel, err := s.repo.FindUserByNIK(ctx, req.NIK)
	if err != nil {
		errMsg := fmt.Errorf("nik not found, nik=%s, error=%s", req.NIK, err.Error())
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusUnauthorized, errMsg.Error(), errMsg)
	}
	// compare password
	if !utils.VerifyPassword(req.Password, userModel.Password) {
		errMsg := fmt.Errorf("invalid password")
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusUnauthorized, errMsg.Error(), errMsg)
	}

	//get role by user id
	role, err := s.roleSvc.GetRoleByID(ctx, userModel.RoleID)
	if err != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
	}

	//GET ROLE MENU ACTION
	roleMenuActions, err := s.rmaSvc.GetRoleMenuActionByRoleID(ctx, role.ID)
	if err != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
	}

	// generate token
	tokenExpTime := s.cfg.GetInt(config.TokenExpTime)
	token, errCreateToken := utils.CreateToken(userModel.ID.String(), int(tokenExpTime))
	if errCreateToken != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, errCreateToken.Error(), errCreateToken)
	}

	// store token
	dt := snapauth.AccessTokenResponse{
		AccessToken:    token,
		TokenType:      "bearer",
		ExpiresIn:      strconv.FormatInt(tokenExpTime, 10),
		AdditionalInfo: nil,
	}
	if err := s.cacheRepo.StoreAccessToken(ctx, userModel.ID.String(), dt); err != nil {
		return common.GeneralResponse{}, cerr.NewServiceErrorWrapper(http.StatusInternalServerError, err.Error(), err)
	}

	resp := domain.LoginDataResponse{
		Token:       token,
		UserID:      userModel.ID.String(),
		MenuActions: []_rmaDomain.RoleMenuAction{},
	}

	for _, v := range roleMenuActions {
		resp.MenuActions = append(resp.MenuActions, v)
	}
	// return token
	return common.GeneralResponse{Status: "200", Message: common.SuccessMessage, Data: resp}, nil
}

func (s *service) Register(ctx context.Context, req domain.RegisterRequest) (common.GeneralResponse, error) {
	if s.isValidEmail(ctx, req.Email) {
		errMsg := fmt.Errorf("email has been used, email=%s", req.Email)
		return common.GeneralResponse{}, errMsg
	}

	//CHECK ROLE
	role, err := s.roleSvc.GetRoleByID(ctx, req.RoleID)
	if err != nil {
		return common.GeneralResponse{}, err
	}
	if role.ID <= 0 {
		return common.GeneralResponse{}, errors.New("invalid role")
	}

	hashPass, err := utils.HashPassword(req.Password)
	if err != nil {
		errMsg := fmt.Errorf("err while hashing password, email=%s", req.Email)
		return common.GeneralResponse{}, errMsg
	}

	userModel := domain.User{
		Nik:      req.NIK,
		Email:    req.Email,
		Password: hashPass,
		RoleID:   role.ID,
		Status:   "Active",
		Base: comdb.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			CreatedBy: "SYSTEM", // TODO CHANGE TO USER WHO CREATE THE USER.
			UpdatedBy: "SYSTEM", // TODO CHANGE TO USER WHO CREATE THE USER.
		},
	}
	userModel, err = s.storeUser(ctx, userModel)
	if err != nil {
		errMsg := fmt.Errorf("error while store new user, err=%s", err.Error())
		return common.GeneralResponse{}, errMsg
	}

	resp := common.GeneralResponse{Status: "200", Message: common.SuccessMessage}
	return resp, nil
}

func (s *service) isValidEmail(ctx context.Context, email string) bool {
	var err error
	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return false
	}
	if user.ID.String() != "" {
		return false
	}
	return true
}

func (s *service) storeUser(ctx context.Context, model domain.User) (domain.User, error) {
	var userModel domain.User
	var err error
	if userModel, err = s.repo.StoreUser(ctx, model); err != nil {
		return domain.User{}, err
	}
	return userModel, nil
}

// DEPRECATED
func (s *service) storeUserRole(ctx context.Context, role domain.UserRole) error {
	if err := s.repo.StoreUserRole(ctx, role); err != nil {
		return err
	}
	return nil
}
