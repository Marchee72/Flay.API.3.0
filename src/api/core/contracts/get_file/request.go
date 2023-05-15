package get_file

import "go.mongodb.org/mongo-driver/bson/primitive"

type Request struct {
	ContentRawId string `uri:"content_id" binding:"required"`
}

func (request *Request) ContentID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(request.ContentRawId)
	return id
}
