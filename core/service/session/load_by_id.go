package session

import (
	"context"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//LoadByID 根据sessionID加载session会话
func LoadByID(ctx context.Context, sessionID string) (*models.UserSession, error) {
	itemInfo := new(models.UserSession)
	coll, err := db.Collection(common.UserSessionCollectionName)
	if err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}
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
