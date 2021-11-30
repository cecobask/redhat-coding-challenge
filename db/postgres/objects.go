package postgres

import (
	"context"

	"github.com/cecobask/redhat-coding-challenge/model"
	"github.com/jackc/pgx/v4"
)

// GetAllObjectsInBucket communicates with postgres to retrieve all objects in the specified bucket
func (db Database) GetAllObjectsInBucket(ctx context.Context, bucketName string) (*model.ObjectList, error) {
	list := &model.ObjectList{
		Objects: make([]model.Object, 0),
	}
	rows, err := db.Conn.Query(ctx, getAllObjectsInBucket, bucketName)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var object model.Object
		err := rows.Scan(&object.ID, &object.ObjectName, &object.ObjectExtension, &object.ObjectPath, &object.BucketName, &object.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Objects = append(list.Objects, object)
	}
	return list, nil
}

// GetObjectByBucketNameAndID communicates with postgres to retrieve an object based on bucket name and object id
func (db Database) GetObjectByBucketNameAndID(ctx context.Context, bucketName, objectID string) (*model.Object, error) {
	object := &model.Object{}
	row := db.Conn.QueryRow(ctx, getObjectByBucketNameAndID, bucketName, objectID)
	err := row.Scan(&object.ID, &object.ObjectName, &object.ObjectExtension, &object.ObjectPath, &object.BucketName, &object.CreatedAt)
	switch err {
	case pgx.ErrNoRows:
		return object, ErrNoRows
	default:
		return object, err
	}
}
