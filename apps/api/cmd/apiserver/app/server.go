package app

import (
	"github.com/gin-gonic/gin"
)

func NewAPIServer() func() int {
	// conf.InitConf(s)
	//logger.WithName("apiserver")

	// fmt.Printf("%v", s)

	//logger.SetVerbosityLevel(s.V)
	//conf.NewConf()
	//ctx := context.Background()
	//minioClient, err := minio.NewClient(ctx, "14.103.230.114:9000", "admin", "qwer1234")
	//if err != nil {
	//	fmt.Printf("error creating minio client: %v\n", err)
	//	return err
	//}
	//
	//err = minioClient.SetBucketPolicy(ctx, "nextadm")
	//fmt.Printf("set bucket policy: %v\n", err)

	//object, err := minioClient.GetObject(ctx, "nextadm", "dk21bsVJ.jpg")
	//if err != nil {
	//	fmt.Printf("error getting object: %v\n", err)
	//	return err
	//}
	//
	//defer object.Close()
	//ocalPath := filepath.Join(os.TempDir(), "nextadm.jpg")
	//fmt.Printf("writing %s\n", ocalPath)
	//localFile, err := os.Create(ocalPath)
	//if err != nil {
	//	fmt.Printf("error creating file: %v\n", err)
	//}
	//defer localFile.Close()
	//
	//if _, err := io.Copy(localFile, object); err != nil {
	//	fmt.Printf("error copying file: %v\n", err)
	//	return err
	//}

	//err = minioClient.CreateBucketIfNeed(ctx, "test-nextadmin", "cn-north-1")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	//config, err := NewConfig()
	//if err != nil {
	//	return err
	//}

	//server, err := createServerChai(config)
	//if err != nil {
	//	return err
	//}
	//
	//err = server.Run(":8080")
	//if err != nil {
	//	return err
	//}

	return func () int {
		 return 1
	}

}



func createServerChai(config *Config) (*gin.Engine, error) {
	r := gin.New()

	// 注册路由 --
	return r, nil
}
