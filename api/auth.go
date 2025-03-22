package api

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mcarre20/open_desk/util"
	"golang.org/x/crypto/bcrypt"
)

type contextKey int
const(
	contextUserId contextKey = iota
)
func (server *Server) LoginHandler(w http.ResponseWriter,r *http.Request){
	type UserCred struct{
		UserName string `json:"username"`
		Password string `jsong:"password"`
	}
	user := UserCred{}
	err := util.JsonDecode(r.Body,&user)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//get user from db by username
	dbUser, err := server.store.GetUserByUserName(r.Context(),user.UserName)
	if err != nil{
		msg:="error with username or password"
		util.RespondWithError(w,msg,http.StatusUnauthorized)
		return
	}

	//validate password
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.HashedPassword),[]byte(user.Password))
	if err != nil{
		msg := "error with username or password"
		util.RespondWithError(w,msg,http.StatusUnauthorized)
		return
	}

	//create jwt token
	token, err := createJWTToken(dbUser.ID,server.config.JWTSigningKey)
	if err != nil{
		msg := "server error"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}

	//todo: create refresh token

	//send jwt token
	err = util.RespondWithJson(w,http.StatusOK,struct{
		Token string `json:"jwt_token"`
	}{
		Token: token,
	})
	if err != nil{
		util.RespondWithError(w,"server error",http.StatusInternalServerError)
	}
}

func (server *Server) AuthMiddleWare(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		authHearder := r.Header.Get("Authorization")
		if authHearder == "" {
			msg := "authorization header cannot be blank"
			util.RespondWithError(w,msg,http.StatusUnauthorized)
			return
		}
		authHearderSplit := strings.Fields(authHearder)
		token := authHearderSplit[1]
		//validate token
		tokenClaims, err := validateJWToken(token,server.config.JWTSigningKey)
		if err != nil{
			msg := "invalide token"
			util.RespondWithError(w,msg,http.StatusUnauthorized)
			return
		}
		userId,err:= tokenClaims.GetSubject()
		if err != nil{
			msg := "error parsing token"
			util.RespondWithError(w,msg,http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(),contextUserId,userId)
		next.ServeHTTP(w,r.WithContext(ctx))
	})
}

//
func createJWTToken(userId uuid.UUID,signingKey string)(string, error){
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt: jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject: userId.String(),

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	
	signedToken, err := token.SignedString(signingKey)
	if err != nil{
		return "", errors.New("error signing token")
	}
	return signedToken, nil

}

//validat JWT token and return token's claim
func validateJWToken(s string,signingKey string) (jwt.Claims,error){
	token, err := jwt.Parse(s,func(token *jwt.Token)(interface{},error){
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil,errors.New("unexpected signing method")
		}
		return []byte(signingKey),nil;
	})
	if err != nil{
		return nil, err
	}
	claims := token.Claims
	tokenExp,err := claims.GetExpirationTime()
	if err != nil{
		return nil, err
	}
	if tokenExp.Before(time.Now()){
		return nil,errors.New("expired token")
	}
	return claims,nil
}