package entity

import "gopkg.in/mgo.v2/bson"

// ID entity id
type ID bson.ObjectId

// String convert an ID in a string
func (i ID) String() string {
	return bson.ObjectId(i).Hex()
}

// StringToID convert a string to an ID
func StringToID(s string) ID {
	return ID(bson.ObjectIdHex(s))
}

// NewID create a new id
func NewID() ID {
	return StringToID(bson.NewObjectId().Hex())
}
