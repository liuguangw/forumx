package cache

import (
	"context"
	"fmt"
	"time"
)

//写入缓存的示例代码
func ExamplePutItem() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//key
	cacheKey := "hello"
	//value
	var cacheValue struct {
		A1 int
		B1 string
	}
	cacheValue.A1 = 666
	cacheValue.B1 = "hello"
	//设置缓存的失效时间
	expiredAt := time.Now().Add(5 * time.Minute)
	if err := PutItem(ctx, cacheKey, &cacheValue, expiredAt); err != nil {
		panic(err)
	}
}

//读取缓存的示例代码
func ExampleGetItem() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//key
	cacheKey := "hello"
	//声明一个匿名的结构
	var cachedItem struct {
		//字段名必须是ItemValue, 而且需要申明bson tag, 类型为所需读取的缓存数据类型指针
		ItemValue *struct {
			A1 int
			B1 string
		} `bson:"item_value"`
	}
	cacheExists, err := GetItem(ctx, cacheKey, &cachedItem)
	if err != nil {
		panic(err)
	}
	//缓存是否存在
	fmt.Println(cacheExists)
	//缓存的数据
	if cacheExists {
		fmt.Println(cachedItem.ItemValue)
	}
}

//删除缓存的示例代码
func ExampleDeleteItem() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//key
	cacheKey := "hello"
	if err := DeleteItem(ctx, cacheKey); err != nil {
		panic(err)
	}
}
