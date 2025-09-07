package app

func NewUserServer() func() int {
	
	return func() int {
		return 1
	}
}

func run() error {
	// lis, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// srv := grpc.NewServer()

	// // userV1.RegisterUserServiceServer(srv, application.NewUserApplication())

	// err = srv.Serve(lis)
	// if err != nil {
	// 	return err
	// 	log.Fatal(err)
	// }
	return nil
}
