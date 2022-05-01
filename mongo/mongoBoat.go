package mongo

import (
	"context"
	"group-project/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//แสดง ตำแหน่งงานทั้งหมด ที่ user ทุกคน สมัครงาน หรือ ถูกใจไว้
func GetJobs() ([]map[string]interface{}, error) {
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

	collection := client.Database("IS766-Final-Project").Collection("jobbyuser")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer cur.Close(ctx)
	var Jobs []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			return []map[string]interface{}{}, err
		}
		Job := make(map[string]interface{})
		for _, m := range result {
			Job[m.Key] = m.Value
		}
		Jobs = append(Jobs, Job)
	}
	if err := cur.Err(); err != nil {
		return []map[string]interface{}{}, err
	}
	return Jobs, nil
}

//เพิ่ม job ที่เลือกแล้วลง DB 	//insert job ที่เลือกลง DB
func CreateJob(Job model.CreateJobRequest) error {
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
	collection := client.Database("IS766-Final-Project").Collection("jobbyuser")
	JobBson, err := bson.Marshal(Job)
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(ctx, JobBson)
	if err != nil {
		return err
	}
	return nil
}

//update status ของ Job โดยส่งใช้ Email link เป็นเงื่อนไข และเปลี่ยนเฉพาะ status ไป
//status = Register , Delete , Like
func UpdateJob(Status string, whereurllink string, whereemail string, Job model.CreateJobRequest) error {
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
	collection := client.Database("IS766-Final-Project").Collection("jobbyuser")
	//id, _ := primitive.ObjectIDFromHex("626405d650981a0d5cc6ad83")
	_, err = collection.UpdateOne(
		ctx,
		bson.M{"urllink": whereurllink, "email": whereemail},
		bson.D{
			{"$set", bson.D{{"status", Status}}},
		})

	if err != nil {
		return err
	}

	return nil
}

//ลบ Job ที่ ไม่ต้องการออกไปโดยใช้ ID
func DeleteJob(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}

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

	collection := client.Database("IS766-Final-Project").Collection("jobbyuser")

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

//แสดง ตำแหน่งงานทั้งหมดของแต่ละ User โดยหาตาม email
func GetUserJob(Email string) ([]map[string]interface{}, error) {
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

	collection := client.Database("IS766-Final-Project").Collection("jobbyuser")
	cur, err := collection.Find(ctx, bson.D{{"email", Email}})
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

//ค้าหาเฉพะบางตำแหน่งที่ Scraping มาเก็บไว้แล้ว ใน DB Table : Jobtest2

//ค้นหาตำแแหน่งานทั้งหมดที่ scrap ไว้ใน DB //
func GetPosition2() ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://is766:2hxK81IxuIiVbEL4@is766cluster0.7orlx.mongodb.net/is766db?retryWrites=true&w=majority"))
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("IS766-Final-Project").Collection("jobs")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer cur.Close(ctx)
	var Jobs []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			return []map[string]interface{}{}, err
		}
		Job := make(map[string]interface{})
		for _, m := range result {
			Job[m.Key] = m.Value
		}
		Jobs = append(Jobs, Job)
	}
	if err := cur.Err(); err != nil {
		return []map[string]interface{}{}, err
	}
	return Jobs, nil
}
