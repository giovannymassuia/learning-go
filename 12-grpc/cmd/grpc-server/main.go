package main

import (
	"12-grpc/internal/database"
	"12-grpc/internal/pb"
	"12-grpc/internal/service"
	"database/sql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	// grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	// enable reflection
	reflection.Register(grpcServer)

	// grpc server listen
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
