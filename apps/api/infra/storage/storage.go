package storage

// 定义 s3 文件存储的操作
// 未来 可拓展 其它支持s3存储协议的

type BucketObjectStorage interface {
	UpLoad() error
}
