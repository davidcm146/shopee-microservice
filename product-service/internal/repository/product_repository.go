package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/davidcm146/shopee-microservice/product-service/internal/common"
	"github.com/davidcm146/shopee-microservice/product-service/internal/models"
	"github.com/davidcm146/shopee-microservice/product-service/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository interface {
	// Define product repository methods here
	FindAll(ctx context.Context) ([]*models.Product, error)
	FindByID(ctx context.Context, id string) (*models.Product, error)
	FindBySellerID(ctx context.Context, sellerID string) ([]*models.Product, error)
	Create(ctx context.Context, product *models.Product) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) (*models.Product, error)
	Delete(ctx context.Context, id string) error
	SoftDelete(ctx context.Context, id string) error
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(mongo *db.MongoConfig) ProductRepository {
	return &productRepository{
		collection: mongo.MongoDB.Collection("products"),
	}
}

func (r *productRepository) FindAll(ctx context.Context) ([]*models.Product, error) {
	filter := common.NotDeletedFilter()
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*models.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) FindByID(ctx context.Context, id string) (*models.Product, error) {
	// TODO
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var product models.Product
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)

	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindBySellerID(ctx context.Context, sellerID string) ([]*models.Product, error) {
	filter := bson.M{
		"sellerID":  sellerID,
		"isDeleted": false,
	}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*models.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	fmt.Println("Products by sellerID:", products)
	return products, nil
}

func (r *productRepository) Create(ctx context.Context, product *models.Product) (*models.Product, error) {
	// TODO
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	product.IsDeleted = false
	_, err := r.collection.InsertOne(ctx, product)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) Update(ctx context.Context, product *models.Product) (*models.Product, error) {
	objID, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"quantity":    product.Quantity,
			"category":    product.Category,
			"features":    product.Features,
			"updatedAt":   time.Now(),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated models.Product
	err = r.collection.FindOneAndUpdate(ctx, bson.M{"_id": objID, "isDeleted": false}, update, opts).Decode(&updated)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *productRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *productRepository) SoftDelete(ctx context.Context, id string) error {
	// TODO
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID, "isDeleted": false}
	update := bson.M{
		"$set": bson.M{
			"isDeleted": true,
			"updatedAt": time.Now(),
		},
	}
	_, err = r.collection.UpdateOne(ctx, filter, update)
	return err
}
