package mongo

import (
	"context"
	"fmt"
	"group-project/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//register//

func GetRegister(queryemail string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://IS766:66L8E4nVUyLBzWkn@is766cluster0.wvp1c.mongodb.net/IS766DB?retryWrites=true&w=majority"))
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("IS766FinalProject").Collection("users")
	cur, err := collection.Find(ctx, bson.D{{"email", queryemail}})
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer cur.Close(ctx)
	var registers []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			return []map[string]interface{}{}, err
		}
		// do something with result....
		register := make(map[string]interface{})
		for _, m := range result {
			register[m.Key] = m.Value
		}
		registers = append(registers, register)
	}
	if err := cur.Err(); err != nil {
		return []map[string]interface{}{}, err
	}
	return registers, nil
}

func CreateRegister(register model.CreateRegisterRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://IS766:66L8E4nVUyLBzWkn@is766cluster0.wvp1c.mongodb.net/IS766DB?retryWrites=true&w=majority"))
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("IS766FinalProject").Collection("users")
	registerBson, err := bson.Marshal(register)

	if err != nil {
		return err
	}
	_, err = collection.InsertOne(ctx, registerBson)

	if err != nil {
		return err
	}
	return nil
}

func UpdateRegister(updatedemail string, register model.CreateRegisterRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://IS766:66L8E4nVUyLBzWkn@is766cluster0.wvp1c.mongodb.net/IS766DB?retryWrites=true&w=majority"))
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("IS766FinalProject").Collection("users")

	_, err = collection.DeleteOne(ctx, bson.D{{"email", updatedemail}})
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(ctx, register)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRegister(deletedemail string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://IS766:66L8E4nVUyLBzWkn@is766cluster0.wvp1c.mongodb.net/IS766DB?retryWrites=true&w=majority"))
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("IS766FinalProject").Collection("users")

	_, err = collection.DeleteOne(ctx, bson.D{{"email", deletedemail}})
	if err != nil {
		return err
	}

	return nil
}

//search job//

func GetPosition(queryposition string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://IS766:66L8E4nVUyLBzWkn@is766cluster0.wvp1c.mongodb.net/IS766DB?retryWrites=true&w=majority"))
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("IS766FinalProject").Collection("jobs")
	cur, err := collection.Find(ctx, bson.D{{"position", primitive.Regex{Pattern: queryposition, Options: ""}}})
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer cur.Close(ctx)
	var jobs []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			return []map[string]interface{}{}, err
		}
		// do something with result....
		job := make(map[string]interface{})
		for _, m := range result {
			job[m.Key] = m.Value
		}
		jobs = append(jobs, job)
	}
	if err := cur.Err(); err != nil {
		return []map[string]interface{}{}, err
	}
	return jobs, nil
}

func GetCompany(querycompany string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://IS766:66L8E4nVUyLBzWkn@is766cluster0.wvp1c.mongodb.net/IS766DB?retryWrites=true&w=majority"))
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("IS766FinalProject").Collection("jobs")
	cur, err := collection.Find(ctx, bson.D{{"company", primitive.Regex{Pattern: querycompany, Options: ""}}})
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer cur.Close(ctx)
	var jobs []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			return []map[string]interface{}{}, err
		}
		// do something with result....
		job := make(map[string]interface{})
		for _, m := range result {
			job[m.Key] = m.Value
		}
		jobs = append(jobs, job)
	}
	if err := cur.Err(); err != nil {
		return []map[string]interface{}{}, err
	}
	return jobs, nil
}

func GetLocation(querylocation string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://IS766:66L8E4nVUyLBzWkn@is766cluster0.wvp1c.mongodb.net/IS766DB?retryWrites=true&w=majority"))
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("IS766FinalProject").Collection("jobs")
	cur, err := collection.Find(ctx, bson.D{{"location", primitive.Regex{Pattern: querylocation, Options: ""}}})
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer cur.Close(ctx)
	var jobs []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			return []map[string]interface{}{}, err
		}
		// do something with result....
		job := make(map[string]interface{})
		for _, m := range result {
			job[m.Key] = m.Value
		}
		jobs = append(jobs, job)
	}
	if err := cur.Err(); err != nil {
		return []map[string]interface{}{}, err
	}
	return jobs, nil
}

// upload file //

func InitiateMongoClient() *mongo.Client {
	var err error
	var client *mongo.Client
	uri := "mongodb+srv://IS766:66L8E4nVUyLBzWkn@is766cluster0.wvp1c.mongodb.net/IS766DB?retryWrites=true&w=majority"
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		fmt.Println(err.Error())
	}
	return client
}

//get file id//
func GetFileID(queryfilename string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://IS766:66L8E4nVUyLBzWkn@is766cluster0.wvp1c.mongodb.net/IS766DB?retryWrites=true&w=majority"))
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("myfiles").Collection("fs.files")
	cur, err := collection.Find(ctx, bson.D{{"filename", primitive.Regex{Pattern: queryfilename, Options: ""}}})
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer cur.Close(ctx)
	var fileids []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			return []map[string]interface{}{}, err
		}
		// do something with result....
		fileid := make(map[string]interface{})
		for _, m := range result {
			fileid[m.Key] = m.Value
		}
		fileids = append(fileids, fileid)
	}
	if err := cur.Err(); err != nil {
		return []map[string]interface{}{}, err
	}
	return fileids, nil
}

// update file id
func AddFileID(email string, fileid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://IS766:66L8E4nVUyLBzWkn@is766cluster0.wvp1c.mongodb.net/IS766DB?retryWrites=true&w=majority"))
	if err != nil {
		return err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("IS766FinalProject").Collection("users")
	_, err = collection.UpdateOne(
		ctx,
		bson.M{"email": email},
		bson.D{
			{"$set", bson.D{{"fileid", fileid}}},
		})

	if err != nil {
		return err
	}

	return nil
}

func GetFileFromFileId(fileId string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://IS766:66L8E4nVUyLBzWkn@is766cluster0.wvp1c.mongodb.net/IS766DB?retryWrites=true&w=majority"))
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("myfiles").Collection("fs.files")
	// cur, err := collection.Find(ctx, bson.D{{"_id", primitive.Regex{Pattern: fileId, Options: ""}}})
	objectId, _ := primitive.ObjectIDFromHex(fileId)
	cur, err := collection.Find(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer cur.Close(ctx)
	var fileids []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			return []map[string]interface{}{}, err
		}
		// do something with result....
		fileid := make(map[string]interface{})
		for _, m := range result {
			fileid[m.Key] = m.Value
		}
		fileids = append(fileids, fileid)
	}
	if err := cur.Err(); err != nil {
		return []map[string]interface{}{}, err
	}
	return fileids, nil

}
