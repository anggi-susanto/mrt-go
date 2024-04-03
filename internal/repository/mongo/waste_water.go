package mongo

import (
	"context"

	"github.com/anggi-susanto/mrt-go/config"
	"github.com/anggi-susanto/mrt-go/domain"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// WasteWaterRepository is the implementation of the WasteWaterRepositoryInterface.
type WasteWaterRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// NewWasteWaterRepository creates a new WasteWaterRepository.
//
// The WasteWaterRepository is used to interact with the waste water collection in the database.
//
// Parameters:
// - client: a pointer to a mongo.Client.
// - config: a pointer to a config.MongoConfig.
// Returns a pointer to a WasteWaterRepository.
func NewWasteWaterRepository(client *mongo.Client, config *config.MongoConfig) *WasteWaterRepository {
	// Get the collection from the database
	collection := client.Database(config.Database).Collection(config.WasteWaterCollection)

	return &WasteWaterRepository{
		// The client used to interact with the database
		client: client,
		// The collection to interact with
		collection: collection,
	}
}

// Create adds a new waste water record to the database.
//
// ctx: the context in which the operation is performed.
// w: the waste water data request to be stored.
//
// Returns an error if the operation was not successful.
func (r *WasteWaterRepository) Create(ctx context.Context, w *domain.WastewaterDataRequest) error {
	// Insert the new waste water data into the database
	_, err := r.collection.InsertOne(ctx, w)
	if err != nil {
		// Log the error and return it
		logrus.Error(err)
		return err
	}
	return nil
}

// GetAll retrieves all waste water data with pagination from the WasteWaterRepository.
//
// ctx: the context for the operation.
// page: the page number for pagination.
// limit: the maximum number of items to return per page.
//
// Returns a list of waste water data and an error, if any.
func (r *WasteWaterRepository) GetAll(ctx context.Context, page, limit int) ([]domain.WasteWaterData, error) {
	// Calculate the skip value based on the page and limit values
	skip := (page - 1) * limit

	// Define the filter and options for the query
	filter := bson.D{} // empty filter to retrieve all documents
	options := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))

	// Execute the query and get a cursor
	cursor, err := r.collection.Find(ctx, filter, options)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// Decode all the documents in the cursor into a slice of WasteWaterData
	var wastes []domain.WasteWaterData
	if err = cursor.All(ctx, &wastes); err != nil {
		logrus.Error(err)
		return nil, err
	}

	// Return the slice of WasteWaterData and nil error
	return wastes, nil
}

// GetByID retrieves a WasteWaterData document by its ID.
//
// Parameters:
//
//	ctx - context.Context: the context for the operation
//	id - string: the ID of the document to retrieve
//
// Returns:
//
//	*domain.WasteWaterData - pointer to the retrieved WasteWaterData
//	error - nil if successful, error if not found or any other error occurs
func (r *WasteWaterRepository) GetByID(ctx context.Context, id string) (*domain.WasteWaterData, error) {
	// Define the filter for querying the document by its ID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	// Use the FindOne function to retrieve the document
	var waste domain.WasteWaterData
	if err := r.collection.FindOne(ctx, filter).Decode(&waste); err != nil {
		// Return nil and the error if the document was not found
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	// Return the pointer to the WasteWaterData and a nil error
	return &waste, nil
}

// Update updates a WasteWaterData in the WasteWaterRepository.
//
// ctx: the context for the operation.
// w: a pointer to the WasteWaterData to update.
//
// Returns an error if the operation was not successful.
func (r *WasteWaterRepository) Update(ctx context.Context, w *domain.WasteWaterData) error {
	// Define the filter for querying the document by its ID
	objectID, err := primitive.ObjectIDFromHex(w.ID.Hex())
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}

	// Define the update operation for the document
	update := bson.D{{Key: "$set", Value: w}}

	// Use the UpdateOne function to update the document
	_, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		// Log the error and return it
		logrus.Error(err)
		return err
	}

	// Return a nil error if the operation was successful
	return nil
}

// Delete removes a single document from the WasteWaterRepository collection using the provided context and ID.
// If the ID is not found, it returns a MongoDB exception with the NotFound error code.
// It returns an error if any other error occurs.
func (r *WasteWaterRepository) Delete(ctx context.Context, id string) error {
	// Define the filter for querying the document by its ID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}

	// Use the DeleteOne function to delete the document
	_, err = r.collection.DeleteOne(ctx, filter)

	// If an error occurs, log it and return it
	if err != nil {
		logrus.Error(err)
		return err
	}

	// Return a nil error if the operation was successful
	return nil
}
