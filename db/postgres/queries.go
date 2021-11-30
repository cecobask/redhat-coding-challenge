package postgres

const (
	getAllObjectsInBucket         = `SELECT id, object_name, object_extension, object_path, bucket_name, created_at::text FROM objects WHERE bucket_name = $1;`
	getObjectByBucketNameAndID    = `SELECT id, object_name, object_extension, object_path, bucket_name, created_at::text FROM objects WHERE bucket_name = $1 AND id = $2;`
	createObject                  = `INSERT INTO objects (id, object_name, object_extension, object_path, bucket_name) VALUES ($1, $2, $3, $4, $5);`
	updateObject                  = `UPDATE objects SET object_name = $2, object_extension = $3, object_path = $4, bucket_name = $5, created_at = CURRENT_TIMESTAMP WHERE id = $1;`
	deleteObjectByBucketNameAndID = `DELETE FROM objects WHERE bucket_name = $1 AND id = $2;`
)
