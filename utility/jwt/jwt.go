package jwt

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenExpireDuration = time.Hour * 24
	TokenFileDir        = "resource/token"
)

var (
	mySecret = []byte("star-project-secret-key-2026")
)

type UserClaims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type TokenManager struct {
	mu sync.RWMutex
}

var tokenManager = &TokenManager{}

func NewTokenManager() *TokenManager {
	return tokenManager
}

func (tm *TokenManager) GenerateToken(ctx context.Context, id uint, username, email string) (string, error) {
	claims := UserClaims{
		Id:       id,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "star-project",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySecret)
	if err != nil {
		g.Log().Errorf(ctx, "生成 Token 失败: %v", err)
		return "", err
	}

	return tokenString, nil
}

func (tm *TokenManager) ParseToken(ctx context.Context, tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})

	if err != nil {
		g.Log().Errorf(ctx, "解析 Token 失败: %v", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的 Token")
}

func (tm *TokenManager) SaveTokenToFile(ctx context.Context, username, token string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tokenDir := gfile.Join(gfile.Pwd(), TokenFileDir)
	if !gfile.Exists(tokenDir) {
		if mkdirErr := gfile.Mkdir(tokenDir); mkdirErr != nil {
			g.Log().Errorf(ctx, "创建 Token 目录失败: %v", mkdirErr)
			return mkdirErr
		}
	}

	tokenFile := gfile.Join(tokenDir, fmt.Sprintf("%s.json", username))

	tokenData := g.Map{
		"username":  username,
		"token":     token,
		"createdAt": time.Now().Format("2006-01-02 15:04:05"),
		"expireAt":  time.Now().Add(TokenExpireDuration).Format("2006-01-02 15:04:05"),
	}

	jsonContent, encodeErr := gjson.EncodeString(tokenData)
	if encodeErr != nil {
		g.Log().Errorf(ctx, "序列化 Token 数据失败: %v", encodeErr)
		return encodeErr
	}

	if putErr := gfile.PutContents(tokenFile, jsonContent); putErr != nil {
		g.Log().Errorf(ctx, "保存 Token 文件失败: %v", putErr)
		return putErr
	}

	g.Log().Infof(ctx, "Token 已保存到文件: %s", tokenFile)
	return nil
}

func (tm *TokenManager) GetTokenFromFile(ctx context.Context, username string) (string, error) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	tokenFile := gfile.Join(gfile.Pwd(), TokenFileDir, fmt.Sprintf("%s.json", username))

	if !gfile.Exists(tokenFile) {
		return "", errors.New("Token 文件不存在")
	}

	content := gfile.GetContents(tokenFile)
	if content == "" {
		//g.Log().Errorf(ctx, "读取 Token 文件失败: %v")
		return "", errors.New("读取 Token 文件失败")
	}

	data, decodeErr := gjson.DecodeToJson(content)
	if decodeErr != nil {
		g.Log().Errorf(ctx, "解析 Token 文件失败: %v", decodeErr)
		return "", decodeErr
	}

	token := data.Get("token").String()
	if token == "" {
		return "", errors.New("Token 数据无效")
	}

	return token, nil
}

func (tm *TokenManager) RemoveTokenFile(ctx context.Context, username string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tokenFile := gfile.Join(gfile.Pwd(), TokenFileDir, fmt.Sprintf("%s.json", username))

	if gfile.Exists(tokenFile) {
		if removeErr := gfile.Remove(tokenFile); removeErr != nil {
			g.Log().Errorf(ctx, "删除 Token 文件失败: %v", removeErr)
			return removeErr
		}
		g.Log().Infof(ctx, "Token 文件已删除: %s", tokenFile)
	}

	return nil
}
