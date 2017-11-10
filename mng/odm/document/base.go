package document

import "gopkg.in/mgo.v2/bson"

type (
	IIdentifier interface {
		GetId() bson.ObjectId
	}

	Base struct {
		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	}
)

func (b *Base) GetId() bson.ObjectId {
	return b.Id
}

func (b *Base) BeforeInsert() {
	b.initId()
}

func (b *Base) initId() {
	if b.Id == "" {
		b.Id = bson.NewObjectId()
	}
}
