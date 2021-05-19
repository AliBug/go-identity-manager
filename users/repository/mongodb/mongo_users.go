package mongodb

import (
	"context"
	"log"

	"github.com/alibug/go-identity-manager/domain"
	"github.com/alibug/go-identity-manager/users/repository/body"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoUsersRepository struct {
	userColl *mongo.Collection
}

// NewMongoUserRepository will create an object that represent the user.Repository interface
func NewMongoUserRepository(coll *mongo.Collection) domain.UsersRepository {
	return &mongoUsersRepository{coll}
}

func (m *mongoUsersRepository) ListUsers(ctx context.Context, limit int64, skip int64) ([]domain.User, error) {
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(skip)

	results := make([]domain.User, 0)
	cur, err := m.userColl.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem body.UserBody
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (m *mongoUsersRepository) DeleteUserByID(ctx context.Context, userID string) error {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = m.userColl.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	log.Printf("User with ID deleted: %s", userID)
	return nil
}
