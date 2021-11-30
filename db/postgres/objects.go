package postgres

import (
	"context"
	"log"

	"github.com/cecobask/redhat-coding-challenge/model"
	"github.com/jackc/pgx/v4"
)

// GetAllObjectsInBucket communicates with postgres to retrieve all objects in the specified bucket
func (db Database) GetAllObjectsInBucket(ctx context.Context, bucketName string) (*model.ObjectList, error) {
	log.Println("Database.GetAllObjectsInBucket() invoked")
	list := &model.ObjectList{
		Objects: make([]model.Object, 0),
	}
	rows, err := db.Conn.Query(ctx, getAllObjectsInBucket, bucketName)
	if err != nil {
		log.Println(err)
		return list, err
	}
	for rows.Next() {
		var object model.Object
		err := rows.Scan(&object.ID, &object.ObjectName, &object.ObjectExtension, &object.ObjectPath, &object.BucketName, &object.CreatedAt)
		if err != nil {
			log.Println(err)
			return list, err
		}
		list.Objects = append(list.Objects, object)
	}
	log.Println("Retrieved number of objects:", len(list.Objects))
	return list, nil
}

// GetObjectByBucketNameAndID communicates with postgres to retrieve an object based on bucket name and object id
func (db Database) GetObjectByBucketNameAndID(ctx context.Context, bucketName, objectID string) (*model.Object, error) {
	log.Println("Database.GetObjectByBucketNameAndID() invoked")
	object := &model.Object{}
	row := db.Conn.QueryRow(ctx, getObjectByBucketNameAndID, bucketName, objectID)
	err := row.Scan(&object.ID, &object.ObjectName, &object.ObjectExtension, &object.ObjectPath, &object.BucketName, &object.CreatedAt)
	if err == pgx.ErrNoRows {
		log.Println(ErrNoRows)
		return nil, ErrNoRows
	}
	return object, err
}

// CreateOrUpdateObject communicates with postgres to upsert an object in a specified bucket
func (db Database) CreateOrUpdateObject(ctx context.Context, object model.Object) (*string, error) {
	log.Println("Database.CreateOrUpdateObject() invoked")
	_, err := db.GetObjectByBucketNameAndID(ctx, object.BucketName, object.ID)
	if err != nil {
		_, err := db.Conn.Exec(ctx, createObject, object.ID, object.ObjectName, object.ObjectExtension, object.ObjectPath, object.BucketName)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		log.Println("Created object with ID", object.ID)
		return &object.ID, nil
	}
	_, err = db.Conn.Exec(ctx, updateObject, object.ID, object.ObjectName, object.ObjectExtension, object.ObjectPath, object.BucketName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Updated object with ID", object.ID)
	return &object.ID, nil
}

// DeleteObjectByBucketNameAndID communicates with postgres to delete an object based on bucket name and object id
func (db Database) DeleteObjectByBucketNameAndID(ctx context.Context, bucketName, objectID string) (*model.Object, error) {
	log.Println("Database.DeleteObjectByBucketNameAndID() invoked")
	object, err := db.GetObjectByBucketNameAndID(ctx, bucketName, objectID)
	if err == ErrNoRows {
		return nil, ErrNoRows
	}
	commandTag, err := db.Conn.Exec(ctx, deleteObjectByBucketNameAndID, bucketName, objectID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if commandTag.RowsAffected() == 0 {
		log.Println(ErrNoRows)
		return nil, ErrNoRows
	}
	log.Println("Deleted object with ID", objectID)
	return object, nil
}
