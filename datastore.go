// vessel-service/datastore.go
package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateClient 返回connect，因此调用该函数后记得用 defer client.Disconnect 断开连接
func CreateClient(ctx context.Context, uri string) (*mongo.Client, error) {
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return conn, nil
}
