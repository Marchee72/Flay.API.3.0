package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FileRepository struct {
	Files *mongo.Collection
}

const CHUNK_SIZE = 4 * 1024 * 1024

func (repository *FileRepository) SaveFile(ctx context.Context, file []byte, contentId primitive.ObjectID) error {
	// Read the file contents into a byte slice
	chunks := repository.generateChunks(file)
	// Iterate over the file data and store each chunk
	for i, chunk := range chunks {
		chunkData := bson.M{
			"content_id": contentId, // Identifier for the file being stored
			"chunkIndex": i,         // Index of the current chunk
			"data":       chunk,     // Chunk data
		}

		_, err := repository.Files.InsertOne(ctx, chunkData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repository *FileRepository) GetFile(ctx context.Context, contentId primitive.ObjectID) ([]byte, error) {
	cursor, err := repository.Files.Find(context.Background(), bson.M{"content_id": contentId}, options.Find().SetSort(bson.D{{Key: "chunkIndex", Value: 1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reconstructedFile []byte
	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}

		chunk := result["data"].(primitive.Binary)
		reconstructedFile = append(reconstructedFile, chunk.Data...)
	}
	// The `reconstructedFile` now contains the complete file content
	return reconstructedFile, nil
}

func (repository *FileRepository) generateChunks(content []byte) [][]byte {
	var chunks [][]byte
	for i := 0; i < len(content); i += CHUNK_SIZE {
		end := i + CHUNK_SIZE
		if end > len(content) {
			end = len(content)
		}
		chunk := content[i:end]
		chunks = append(chunks, chunk)
	}
	return chunks
}
