package mongo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.org/anggi-susanto/mrt-go/config"
	"github.org/anggi-susanto/mrt-go/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WasteWaterRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewWasteWaterRepository(client *mongo.Client, config *config.MongoConfig) *WasteWaterRepository {
	collection := client.Database(config.Database).Collection(config.WasteWaterCollection)
	return &WasteWaterRepository{
		client:     client,
		collection: collection,
	}
}

func (r *WasteWaterRepository) Create(ctx context.Context, w *domain.WasteWaterData) error {
	_, err := r.collection.InsertOne(ctx, w)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r *WasteWaterRepository) GetAll(ctx context.Context, page, limit int) ([]domain.WasteWaterData, error) {
	skip := (page - 1) * limit
	filter := bson.D{}
	options := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))
	cursor, err := r.collection.Find(ctx, filter, options)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	var wastes []domain.WasteWaterData
	if err = cursor.All(ctx, &wastes); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return wastes, nil
}

func (r *WasteWaterRepository) GetByID(ctx context.Context, id string) (*domain.WasteWaterData, error) {
	var waste domain.WasteWaterData
	filter := bson.D{{Key: "_id", Value: id}}
	err := r.collection.FindOne(ctx, filter).Decode(&waste)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &waste, nil
}

func (r *WasteWaterRepository) Update(ctx context.Context, w *domain.WasteWaterData) error {
	filter := bson.D{{Key: "_id", Value: w.ID}}
	update := bson.D{{Key: "$set", Value: w}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r *WasteWaterRepository) Delete(ctx context.Context, id string) error {
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
