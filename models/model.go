package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Admin struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Hint struct {
	ID      int    `json:"id" bson:"id"`
	Content string `json:"content" bson:"content"`
}

type Game struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Date    string             `json:"date" bson:"date"`
	Count   int                `json:"count" bson:"count"`
	Players []Player           `json:"players" bson:"players"`
	Admins  []Admin            `json:"admins" bson:"admins"`
	Hints   []Hint             `json:"hints" bson:"hints"`
	Score   []interface{}      `json:"score" bson:"score"`
	Courses []interface{}      `json:"courses" bson:"courses"`
}
