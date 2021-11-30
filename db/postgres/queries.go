package postgres

const (
	getAllObjectsInBucket      = "SELECT * FROM objects WHERE bucket_name = $1;"
	getObjectByBucketNameAndID = "SELECT * FROM objects WHERE bucket_name = $1 AND id = $2;"
)
