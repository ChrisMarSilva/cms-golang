package user // repository_user

import (
	"context"
	"errors"

	entity "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/entities"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryMongo struct {
	ctx        context.Context
	db         *mongo.Database
	collection *mongo.Collection
}

// NewInstanceOfUserRepositoryMongo
func NewUserRepositoryMongo(ctx context.Context, db *mongo.Database) *UserRepositoryMongo {
	return &UserRepositoryMongo{
		ctx:        ctx,
		db:         db,
		collection: db.Collection("users"),
	}
}

func (repo UserRepositoryMongo) GetAll(ctx context.Context, users []*entity.User) (err error) {

	opts := options.Find()
	//opts.SetSort(bson.D{{"nome", -1}})
	opts.SetSort(bson.D{{Key: "nome", Value: -1}})
	//opts.SetSort(map[string]int{"nome": -1})
	//opts.SetSort(bson.M{"nome": 1})
	//opts.SetLimit(10)
	//options.SetLimit(int64(query.Limit))
	//opts.SetSkip(2)
	// options.SetSkip(int64((query.Page * query.Limit) - query.Limit))

	filter := bson.D{{}} // bson.D{} // bson.D{{"duration", bson.D{{"$gt", 24}}}}

	//cur, err := repo.collection.Find(repo.ctx, filter)
	cur, err := repo.collection.Find(repo.ctx, filter, opts)
	if err != nil {
		return err
	}
	defer cur.Close(repo.ctx)

	// NÃO INDICADO PARA MTOS REGISTROS POR CAUSA DA MEMORIA
	// if err = cur.All(repo.ctx, users); err != nil {
	// 	return err
	// }

	for cur.Next(repo.ctx) {
		var user entity.User
		err := cur.Decode(&user)
		if err != nil {
			return err
		}
		users = append(users, &user)
	}

	if err := cur.Err(); err != nil {
		return err
	}

	return nil
}

func (repo UserRepositoryMongo) Get(ctx context.Context, user *entity.User, id uuid.UUID) (err error) {

	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = repo.collection.FindOne(nil, filter).Decode(&user)
	if err != nil {
		return err
	}

	// paginacao
	// https://github.com/scalablescripts/go-search-mongodb/blob/main/main.go

	return nil
}

func (repo UserRepositoryMongo) Create(ctx context.Context, user *entity.User) (err error) {
	_, err = repo.collection.InsertOne(repo.ctx, user)
	if err != nil {
		return err
	}
	// result.InsertedID.(primitive.ObjectID).Hex(), nil
	return nil
}

func (repo UserRepositoryMongo) Update(ctx context.Context, user *entity.User) (err error) {

	// update := bson.M{}
	// if user.Nome != ""{
	// 	update["nome"] = user.Nome
	// }
	// if user.Status != ""{
	// 	update["status"] = user.Status
	// }
	// if len(update) == 0{
	// 	return nill
	// }
	// return bson.M{"$set": update}

	// filter := bson.D{primitive.E{Key: "_id", Value: user.ID}}
	// update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "nome", Value: user.Nome}}}}
	// t := &entity.User{}
	// //return repo.collection.UpdateByID(repo.ctx, user.ID, user)
	// return repo.collection.FindOneAndUpdate(repo.ctx, filter, update).Decode(t)

	id, _ := primitive.ObjectIDFromHex(user.ID.String())

	// res, err := repo.collection.ReplaceOne( repo.ctx, bson.M{"_id": id}, bson.M{"nome": user.Nome, "status": user.Status})
	// res, err := repo.collection.UpdateMany(repo.ctx, bson.M{"_id": id}, bson.D{{Key:"$set", Value: bson.D{{Key:"nome", Value: user.Nome}}}})
	res, err := repo.collection.UpdateOne(repo.ctx, bson.M{"_id": id}, bson.D{{Key: "$set", Value: bson.D{{Key: "nome", Value: user.Nome}}}})
	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return errors.New("No recodr were updated")
	}

	return nil

}

func (repo UserRepositoryMongo) Delete(ctx context.Context, id uuid.UUID) (err error) {

	// docID, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return err
	// }

	// filter := bson.M{"_id": docID}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	// res, err = repo.collection.DeleteMany(repo.ctx, bson.M{"duration": 25})
	// err = repo.collection.Drop(repo.ctx)

	res, err := repo.collection.DeleteOne(repo.ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("No recodr were deleted")
	}

	return nil
}

/*
type Pagination struct {
	Elements interface{} `json:"elements"`
	Metadata *Metadata   `json:"metadata"`
}

type Metadata struct {
	Total  int64 `json:"total"`
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Page   int   `json:"page"`
	Pages  int   `json:"pages"`
}




tx := r.db.Begin()
	err := tx.Debug().Model(&models.Product{}).Where("id = ?", product_id).Delete(&models.Product{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

func (r *productsRepositoryImpl) Count() (int64, error) {
	var c int64
	err := r.db.Debug().Model(&models.Product{}).Count(&c).Error
	return c, err
}

func (r *categoriesRepositoryImpl) Update(category *models.Category) error {
	tx := r.db.Begin()

	columns := map[string]interface{}{
		"description": category.Description,
		"updated_at":  time.Now(),
	}

	err := tx.Debug().Model(&models.Category{}).Where("id = ?", category.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}


*/
