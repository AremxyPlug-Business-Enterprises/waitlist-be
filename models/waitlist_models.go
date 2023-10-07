package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WaitlistEntry struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Email string             `bson:"email"`
}
