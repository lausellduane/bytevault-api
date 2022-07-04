package main

type CodeFragment struct {
    ID              int64   `bson:"_id"`
    Title           string `bson:"title"`
    Description     string `bson:"description"`
    Notes           string   `bson:"notes"`
    Value           string  `bson:"value"`
    Tags            Tags   `bson:"tags"`
    Language        int64   `bson:"language"`
}
type CodeFragments []CodeFragment

type Tag struct {
    ID      int64   `bson:"id"`
    Label   string  `bson:"label"`
}
type Tags []Tag