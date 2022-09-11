package main

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type CodeFragment struct {
    ID              primitive.ObjectID `bson:"_id" json:"_id"`
    Title           string `bson:"title" json:"title"`
    Description     string `bson:"description" json:"description"`
    Notes           string `bson:"notes" json:"notes"`
    Value           string `bson:"value" json:"value"`
    Tags            Tags `bson:"tags" json:"tags"`
    Language        int64 `bson:"language" json:"language"`
}

type Tag struct {
    ID      int64   `bson:"id" json:"id"`
    Label   string  `bson:"label" json:"label"`
}
type Tags []Tag