package models

import "context"

var Ctx = context.Background()

type Student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}
