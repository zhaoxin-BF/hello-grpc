package service

import (
	"context"
	"fmt"
	"github.com/zhaoxin-BF/hello-grpc/golang/person/v1"
)

var PersonService = &personService{}

type personService struct {
	person.UnimplementedPersonServiceServer
}

func (p *personService) GetPerson(ctx context.Context, req *person.GetPersonRequest) (*person.GetPersonResponse, error) {

	fmt.Println(req.PersonName)

	var password = "111"

	user := person.Person{
		Name:     "boreas",
		Age:      18,
		Password: &password,
		Hobbies:  []string{"游泳", "健身", "其他"},
	}

	fmt.Print(user.GetName())
	return &person.GetPersonResponse{Person: &user, Message: "查询成功"}, nil
}

//func (p *personService) GetPerson(context context.Context, request *person.GetPersonRequest) (*person.GetPersonResponse, error) {
//
//	// 实现业务逻辑
//	fmt.Println(request)
//
//	user := person.Person{
//		Name:     "boreas",
//		Age:      18,
//		Password: nil,
//		Hobbies:  []string{"游泳", "健身", "其他"},
//	}
//
//	return &person.GetPersonResponse{Person: &user, Message: "查询成功"}, nil
//}
//
//func httpServer() error {
//	mux := runtime.NewServeMux()
//	server := &http.Server{
//		Addr:              "0.0.0.0:90",
//		Handler:           mux,
//		ReadHeaderTimeout: 1 * time.Second,
//		WriteTimeout:      90 * time.Second,
//		IdleTimeout:       120 * time.Second,
//	}
//	conn, err := grpc.DialContext(context.Background(), fmt.Sprintf("%s:%s", "127.0.0.1", "8002"))
//	if err != nil {
//		return err
//	}
//	err = networkdisk.RegisterControllerHandlerClient(context.Background(), mux, person.NewPersonServiceClient(conn))
//	if err != nil {
//		return err
//	}
//
//	if err = server.ListenAndServe(); err != nil {
//		return err
//	}
//	return nil
//}
