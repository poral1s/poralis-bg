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
	Number  int    `json:"number" bson:"number"`
	Content string `json:"content" bson:"content"`
}

type Course struct {
	ID         int         `json:"id" bson:"id"`
	Name       string      `json:"name" bson:"name"`
	Player     string      `json:"player" bson:"player"`
	StartPoint interface{} `json:"startpoint" bson:"startpoint"`
	GoalPoint  interface{} `json:"goalpoint" bson:"goalpoint"`
	Hints      []Hint      `json:"hints" bson:"hints"`
}

type Game struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Date    string             `json:"date" bson:"date"`
	Count   int                `json:"count" bson:"count"`
	Players []Player           `json:"players" bson:"players"`
	Admins  []Admin            `json:"admins" bson:"admins"`
	Courses []Course           `json:"courses" bson:"courses"`
}
