package mongodb

import (
	"context"
	"errors"

	"github.com/briankliwon/microservices-product-catalog/product/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductModel struct {
	C *mongo.Collection
}

func (m *ProductModel) All() ([]models.Product, error) {
	ctx := context.TODO()
	mm := []models.Product{}

	ProductCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = ProductCursor.All(ctx, &mm)
	if err != nil {
		return nil, err
	}

	return mm, err
}

func (m *ProductModel) FindByID(id string) (*models.Product, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var Product = models.Product{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&Product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &Product, nil
}

func (m *ProductModel) Insert(Product models.Product) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), Product)
}

func (m *ProductModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
