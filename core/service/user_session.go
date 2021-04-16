package service

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

//GetUserSessionByID 根据会话唯一标识获取会话信息
func GetUserSessionByID(sessionID string) (*models.UserSession, error) {
	itemInfo := new(models.UserSession)
	coll, err := db.Collection(userSessionCollectionName)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	filter := bson.M{
		"id": sessionID,
	}
	if err := coll.FindOne(ctx, filter).Decode(itemInfo); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return itemInfo, nil
}

//SaveUserSession 保存用户会话数据到数据库
func SaveUserSession(itemInfo *models.UserSession) error {
	itemInfo.UpdatedAt = time.Now()
	if itemInfo.ID == "" {
		return insertUserSession(itemInfo)
	}
	coll, err := db.Collection(userSessionCollectionName)
	if err != nil {
		return err
	}
	//update
	filter := bson.M{
		"id": itemInfo.ID,
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	if _, err := coll.UpdateOne(ctx, filter, bson.M{
		"$set": itemInfo,
	}); err != nil {
		return err
	}
	return nil
}

func insertUserSession(itemInfo *models.UserSession) error {
	//生成唯一且不存在于集合内的ID
	var sessionIDValid bool
	for !sessionIDValid {
		itemInfo.ID = generateSessionID()
		tmpSessionLog, err := GetUserSessionByID(itemInfo.ID)
		if err != nil {
			return err
		}
		sessionIDValid = tmpSessionLog == nil
	}
	itemInfo.CreatedAt = itemInfo.UpdatedAt
	//session生命周期
	sessionDuration := 15 * 24 * time.Hour
	itemInfo.ExpiredAt = time.Now().Add(sessionDuration)
	coll, err := db.Collection(userSessionCollectionName)
	if err != nil {
		return err
	}
	//insert
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	if _, err := coll.InsertOne(ctx, itemInfo); err != nil {
		return err
	}
	return nil
}

//generateSessionID 生成会话ID的算法
func generateSessionID() string {
	plainText := time.Now().Format(time.RFC3339Nano) + " / - / " + generateRandomString(30)
	return md5String(plainText)
}

//GetRequestSessionID 从客户端请求中,读取会话ID唯一标识符
func GetRequestSessionID(c *fiber.Ctx) string {
	authorizationValue := c.Get("Authorization", "")
	tokenPrefix := "Bearer"
	if strings.Index(authorizationValue, tokenPrefix) == 0 {
		return authorizationValue[len(tokenPrefix)+1:]
	}
	//从URL中读取
	return c.Query("sid", "")
}
