﻿package main

import (
	log "github.com/cihub/seelog"
	"encoding/json"
	"github.com/bottos-project/bottos/service/identity/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"github.com/asaskevich/govalidator"
	"strconv"
	"github.com/mojocn/base64Captcha"
	"os"
	"github.com/bottos-project/bottos/config"
	"regexp"
	//"github.com/gogo/protobuf/proto"
	//"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/proto"
	"github.com/bottos-project/crypto-go/crypto"
	errcode "github.com/bottos-project/bottos/error"
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	Client user.UserClient
}

func (u *User) GetVerify(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//var configCode = base64Captcha.ConfigCharacter{
	//	Height:             60,
	//	Width:              240,
	////	//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
	//	Mode:               base64Captcha.CaptchaModeNumber,
	//	ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
	//	ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
	//	IsShowHollowLine:   false,
	//	IsShowNoiseDot:     false,
	//	IsShowNoiseText:    false,
	//	IsShowSlimeLine:    false,
	//	IsShowSineLine:     false,
	//	CaptchaLen:         6,
	//}
	idKeyD, capD := base64Captcha.GenerateCaptcha("", base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	})
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": 1,
		"data": map[string]interface{}{
			"verify_key": idKeyD,
			"verify_data": base64stringD,
		},
		"msg": "OK",
	})

	rsp.Body = string(b)
	return nil
}


func (u *User) Register(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	var registerRequest user.RegisterRequest
	err := json.Unmarshal([]byte(req.Body), &registerRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	if config.Enable_verification {
		if !base64Captcha.VerifyCaptcha(registerRequest.VerifyCode, registerRequest.VerifyValue) {
			rsp.Body = errcode.ReturnError(1001)
			return nil
		}
	}

	match,err :=regexp.MatchString("^[1-5a-z.]{3,13}$",registerRequest.Account.Name)
	if err != nil {
		log.Error(err)
		return err
	}
	if !match {
		rsp.Body = errcode.ReturnError(1002)
		return nil
	}

	pubkey,err := hex.DecodeString(registerRequest.Account.Pubkey)
	if err != nil {
		log.Error(err)
		return err
	}

	signature,err := hex.DecodeString(registerRequest.User.Signatures)
	if err != nil {
		log.Error(err)
		return err
	}

	registerRequest.User.SetSignatures("")

	serializeData, err := proto.Marshal(registerRequest.User)
	if err != nil {
		log.Error(err)
		return err
	}
	h := sha256.New()
	h.Write([]byte(hex.EncodeToString(serializeData)))
	hash := h.Sum(nil)


	if !crypto.VerifySign(pubkey, hash, signature) {
		rsp.Body = errcode.ReturnError(1000)
		return nil
	}

	response, err := u.Client.Register(ctx, &registerRequest)
	if err != nil {
		return err
	}

	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"msg": response.Msg,
	})
	rsp.Body = string(b)
	return nil
}

func (s *User) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	header, _ := json.Marshal(req.Header)
	response, err := s.Client.Login(ctx, &user.LoginRequest{
		Body: req.Body,
		Header:string(header),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"token":response.Token,
	})
	rsp.Body = string(b)

	return nil
}

func (s *User) Logout(ctx context.Context, req *api.Request, rsp *api.Response) error {
	token := req.Header["Token"]

	if token == nil {
		rsp.StatusCode = 200
		b, _ := json.Marshal(map[string]interface{}{
			"code": "4001",
			"msg":"Token is nil",
		})
		rsp.Body = string(b)
		return nil
	}
	response, err := s.Client.Logout(ctx, &user.LogoutRequest{
		Token: token.Values[0],
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"msg":response.Msg,
	})
	rsp.Body = string(b)

	return nil
}

func (u *User) GetUserInfo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	//transfer to struct
	var queryRequest user.GetUserInfoRequest
	json.Unmarshal([]byte(body), &queryRequest)
	//Checkout data format
	ok, err := govalidator.ValidateStruct(queryRequest);
	if !ok {
		b, _ := json.Marshal(map[string]interface{}{
			"code": -7,
			"msg": err.Error(),
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}

	response, err := u.Client.GetUserInfo(ctx, &queryRequest)
	if err != nil {
		return err
	}

	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"msg": response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}

func (u *User) UpdateUserInfo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	//transfer to struct
	var updateUserInfoRequest user.UpdateUserInfoRequest
	json.Unmarshal([]byte(body), &updateUserInfoRequest)
	//Checkout data format
	ok, err := govalidator.ValidateStruct(updateUserInfoRequest);
	if !ok {
		b, _ := json.Marshal(map[string]interface{}{
			"code": -7,
			"msg": err.Error(),
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}

	response, err := u.Client.UpdateUserInfo(ctx, &updateUserInfoRequest)
	if err != nil {
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"msg": response.Msg,
	})

	rsp.Body = string(b)
	return nil
}

func (u *User) FavoriteMng(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//header, _ := json.Marshal(req.Header)
	response, err := u.Client.FavoriteMng(ctx, &user.FavoriteMngRequest{
		PostBody:   req.Body,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code":  response.Code,
		"msg": response.Msg,
		"data": response.Data,
	})
	rsp.Body = string(b)

	return nil
}
func (u *User) QueryFavorite(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	//transfer to struct
	var queryRequest user.QueryFavoriteRequest
	json.Unmarshal([]byte(body), &queryRequest)
	//Checkout data format

	ok, err := govalidator.ValidateStruct(queryRequest);
	if !ok {
		b, _ := json.Marshal(map[string]string{
			"code": "-7",
			"msg":  err.Error(),
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}

	response, err := u.Client.QueryFavorite(ctx, &queryRequest)
	if err != nil {
		return err
	}

	b, _ := json.Marshal(map[string]interface{}{
		"code": strconv.Itoa(int(response.Code)),
		"msg":  response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}
func (u *User) AddNotice(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := u.Client.AddNotice(ctx, &user.AddNoticeRequest{
		PostBody:   req.Body,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code":  response.Code,
		"msg": response.Msg,
		"data": response.Data,
	})
	rsp.Body = string(b)

	return nil
}
func (u *User) QueryNotice(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	//transfer to struct
	var queryRequest user.QueryNoticeRequest
	json.Unmarshal([]byte(body), &queryRequest)
	//Checkout data format
	ok, err := govalidator.ValidateStruct(queryRequest);
	if !ok {
		b, _ := json.Marshal(map[string]string{
			"code": "-7",
			"msg":  err.Error(),
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}

	response, err := u.Client.QueryNotice(ctx, &queryRequest)
	if err != nil {
		return err
	}

	b, _ := json.Marshal(map[string]interface{}{
		"code": strconv.Itoa(int(response.Code)),
		"msg":  response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}

func (u *User) GetAccount(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	//transfer to struct
	var queryRequest user.GetAccountRequest
	json.Unmarshal([]byte(body), &queryRequest)
	//Checkout data format
	ok, err := govalidator.ValidateStruct(queryRequest);
	if !ok {
		b, _ := json.Marshal(map[string]string{
			"code": "-7",
			"msg":  err.Error(),
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}

	response, err := u.Client.GetAccount(ctx, &queryRequest)
	if err != nil {
		return err
	}

	b, _ := json.Marshal(map[string]interface{}{
		"code": strconv.Itoa(int(response.Code)),
		"msg":  response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}
func (u *User) Transfer(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//header, _ := json.Marshal(req.Header)
	response, err := u.Client.Transfer(ctx, &user.TransferRequest{
		PostBody:   req.Body,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code":  response.Code,
		"msg": response.Msg,
		"data": response.Data,
	})
	rsp.Body = string(b)

	return nil
}
func (u *User) QueryTransfer(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	//transfer to struct
	var queryRequest user.QueryTransferRequest
	json.Unmarshal([]byte(body), &queryRequest)
	//Checkout data format
	ok, err := govalidator.ValidateStruct(queryRequest);
	if !ok {
		b, _ := json.Marshal(map[string]string{
			"code": "-7",
			"msg":  err.Error(),
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}

	response, err := u.Client.QueryTransfer(ctx, &queryRequest)
	if err != nil {
		return err
	}

	b, _ := json.Marshal(map[string]interface{}{
		"code": strconv.Itoa(int(response.Code)),
		"msg":  response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}

func (u *User) GetBlockInfo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := u.Client.GetBlockInfo(ctx, &user.GetBlockInfoRequest{})
	if err != nil {
		return err
	}
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"msg": response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}

func (u *User) GetDataBin(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body

	response, err := u.Client.GetDataBin(ctx, &user.GetDataBinRequest{
		Info:body,
	})
	if err != nil {
		return err
	}
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"msg": response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/api-user-log.xml")
	if err != nil{
		log.Error(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {
	service := micro.NewService(
		micro.Name("bottos.api.v3.user"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&User{Client: user.NewUserClient("bottos.srv.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		os.Exit(1)
	}
}