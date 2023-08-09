package repoimpl 

import (
	"context"
	models "go-module/model"
	repo "go-module/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go-module/config/mongodb"
	"go-module/libs/custom_type"
)

type UserRepoImpl struct {
	Db *mongo.Database
}

func NewUserRepo(db *mongo.Database) repo.UserRepo {
	return &UserRepoImpl {
		Db: db,
	}
}

func (mongo *UserRepoImpl) Count() (
	uint, error) {

	result, err := mongo.Db.Collection(
		mongodb.COLLECTION_USER).CountDocuments(
			context.Background(), bson.D{})
			
	res := uint(result)
	return res, err
}

func (mongo *UserRepoImpl) FindAll() (
	[]custom_type.ConciseData, error) {
	var users []custom_type.ConciseData
	cur, err := mongo.Db.Collection(
		mongodb.COLLECTION_USER).Find(
			context.Background(), bson.M{})
	
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		user := custom_type.ConciseData{}
		err = cur.Decode(&user)
		users = append(users, user)
	} 
		
	err = cur.Err()
	return users, err
}

func (mongo *UserRepoImpl) FindByField(field, value string) (
	models.User, error) {
	
	user := models.User{}
	result := mongo.Db.Collection(
		mongodb.COLLECTION_USER).FindOne(
			context.Background(), bson.M{field: value})
	// https://github.com/mongodb/mongo-go-driver/blob/master/bson/primitive/primitive.go

	err := result.Decode(&user)
	return user, err
}

func (mongo *UserRepoImpl) UpdateById(id string, 
	payload models.User) error {

	var user models.User
	filter := bson.M{"id": bson.M{"$eq": id}}
	update := bson.M{"$set": payload}
	result := mongo.Db.Collection(
		mongodb.COLLECTION_USER).FindOneAndUpdate(
			context.Background(), filter, update)

	err := result.Decode(&user)
	return err
}

func (mongo *UserRepoImpl) Insert(user models.User) (
	error) {
	bbytes, _ := bson.Marshal(user)

	_, err := mongo.Db.Collection(
		mongodb.COLLECTION_USER).InsertOne(
			context.Background(), bbytes)

	return err
}

func (mongo *UserRepoImpl) RemoveAll() error {
	_, err := mongo.Db.Collection(
		mongodb.COLLECTION_USER).DeleteMany(
			context.TODO(), bson.D{})

	return err
}

func (mongo *UserRepoImpl) CheckLoginInfo(email, password string) (
	models.User, error) {
	user := models.User{}

	result := mongo.Db.Collection(
		mongodb.COLLECTION_USER).FindOne(
			context.Background(), bson.M{"email": email, "password": password})
	
	err := result.Decode(&user)
	return user, err
}
