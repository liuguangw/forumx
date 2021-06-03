package tests

import (
	"context"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/app/service/cache"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//testCache 测试缓存功能
func testCache(t *testing.T) {
	itemKey := "hello"
	var itemValue struct {
		A1 int
		B1 string
	}
	//缓存不存在
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	cacheExists, err := cache.GetItem(ctx, itemKey, new(models.Cache))
	assert.NoError(t, err)
	assert.False(t, cacheExists)
	//写入缓存
	itemValue.A1 = 999
	itemValue.B1 = "world"
	err = cache.PutItem(ctx, itemKey, &itemValue, time.Now().Add(10*time.Second))
	assert.NoError(t, err)
	//读取缓存
	var cachedItem struct {
		ItemValue *struct {
			A1 int
			B1 string
		} `bson:"item_value"`
	}
	cacheExists, err = cache.GetItem(ctx, itemKey, &cachedItem)
	assert.NoError(t, err)
	assert.True(t, cacheExists)
	assert.Equal(t, itemValue.A1, cachedItem.ItemValue.A1)
	assert.Equal(t, itemValue.B1, cachedItem.ItemValue.B1)
	//删除不存在的key
	err = cache.DeleteItem(ctx, "foo")
	assert.NoError(t, err)
	//删除已存在的key
	err = cache.DeleteItem(ctx, itemKey)
	assert.NoError(t, err)
	//删除后,缓存不再存在
	cacheExists, err = cache.GetItem(ctx, itemKey, new(models.Cache))
	assert.NoError(t, err)
	assert.False(t, cacheExists)
}
