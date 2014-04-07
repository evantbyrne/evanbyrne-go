package service

import (
	"github.com/evantbyrne/evanbyrne-go/config"
	"github.com/evantbyrne/evanbyrne-go/model/dto"
	"github.com/evantbyrne/evanbyrne-go/util"
	"net/http"
	"net/url"
	"time"
)

func CreateUser(db *util.Database, email string, password string) {
	salt := util.Random(config.HashSaltLength)
	hashed := util.Hash(password, salt)
	user := dto.User{ Email: email, Password: hashed, Salt: salt }
	if err := db.Gorp().Insert(&user); err != nil {
		panic(err)
	}
}

func LoginUser(db *util.Database, response http.ResponseWriter, email string, password string) bool {
	var user dto.User;
	dbmap := db.Gorp()
	err := dbmap.SelectOne(&user, "select * from \"user\" where email = $1 limit 1", email)
	if err != nil || !util.ComparePassword(user.Password, password, user.Salt) {
		return false
	}

	if _, err := db.Connection.Exec("delete from user_session where user_email = $1", email); err != nil {
		panic(err)
	}

	secret := url.QueryEscape(util.Random(config.AuthSecretLength))
	session := dto.UserSession{ UserEmail: email, Secret: secret }
	if err := dbmap.Insert(&session); err != nil {
		panic(err)
	}

	expire := time.Now().AddDate(config.AuthCookieExpireYears, config.AuthCookieExpireMonths, config.AuthCookieExpireDays)
	cookie := http.Cookie{ Name: config.AuthCookieName, Value: session.Secret, Path: "/", Expires: expire }
	http.SetCookie(response, &cookie)

	return true
}

func LogoutUser(db *util.Database, request *http.Request, response http.ResponseWriter) {
	cookie, err := request.Cookie(config.AuthCookieName)
	if err == nil {
		_, err := db.Open().Exec("delete from user_session where secret = $1", cookie.Value)
		if err == nil {
			newCookie := http.Cookie{ Name: config.AuthCookieName, Value: "", Path: "/", Expires: time.Now() }
			http.SetCookie(response, &newCookie)
		}
	}
}

func ValidUserSession(db *util.Database, request *http.Request) bool {
	secret, err := request.Cookie(config.AuthCookieName)
	if err == nil {
		var res dto.UserSession;
		if err = db.Gorp().SelectOne(&res, "select * from user_session where secret = $1 limit 1", secret.Value); err == nil {
			return true
		}
	}
	
	return false
}