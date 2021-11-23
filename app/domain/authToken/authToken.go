package authToken

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type authToken struct {
	userID string
	token  string
}

const (
	// secret は openssl rand -base64 40 コマンドで作成した。TODO:環境変数化
	secret = "1ptTwOz/dcTcYVuzDbHS+S7nLCN9eO/LU+iSirrlTPyBkRl3lXbw0A=="

	// contextにセットするキー
	contextKey = "auth_token"

	// userIDKey はユーザーの ID を表す。
	userIDKey = "user_id"

	// iat と exp は登録済みクレーム名。それぞれの意味は https://tools.ietf.org/html/rfc7519#section-4.1 を参照。{
	iatKey = "iat"
	expKey = "exp"

	// lifetime は jwt の発行から失効までの期間を表す。
	lifetime = 30 * time.Minute
)

// userIDからjwtトークンを発行する
func New(userID string, now time.Time) (*authToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userIDKey: userID,
		iatKey:    now.Unix(),
		expKey:    now.Add(lifetime).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(secret))

	if err != nil {
		return nil, err
	}

	return &authToken{
		userID: userID,
		token:  tokenStr,
	}, nil
}

func SetToken(parents context.Context, t authToken) context.Context {
	return context.WithValue(parents, contextKey, &t)
}

func GetToken(ctx context.Context) (*authToken, error) {
	v := ctx.Value(contextKey)

	token, ok := v.(*authToken)
	if !ok {
		return nil, fmt.Errorf("token not found")
	}

	return token, nil
}

// トークンを検証しauthTokenを生成する
func GenerateAuthByToken(tokenString string) (authToken, error) {
	token, err := verifyToken(tokenString)
	if err != nil {
		return authToken{}, err
	}

	claims := token.Claims.(jwt.MapClaims)

	return authToken{
		userID: claims[userIDKey].(string),
		token:  token.Raw,
	}, nil

}

func (a authToken) GetUserID() string {
	return a.userID
}

func (a authToken) GetToken() string {
	return a.token
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	// jwtの検証（改ざん検知、期限切れ）
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
