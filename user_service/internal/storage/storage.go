package storage

import (
	"database/sql"
	"log"
	"net"
	"os"
	"user_service/internal/infrastructura/redis"
	"user_service/internal/infrastructura/repository/postgres"
	"user_service/internal/service"
	userservice "user_service/service/user_service"
	"user_service/userproto"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func OpenSql(driverName, url string) (*sql.DB, error) {
	db, err := sql.Open(driverName, url)
	if err != nil {
		log.Println("failed to open database:", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println("Unable to connect to database:", err)
		return nil, err
	}

	return db, err
}

func Run() {
	db, err := OpenSql(os.Getenv("driver_name"), os.Getenv("postgres_url"))
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	defer db.Close()

	redisClient := redis.NewRedisClient("localhost:6379", "", 0)

	repo := postgres.NewUserPostgres(db)
	s := service.NewUserService(repo)
	user_handler := userservice.NewGrpcService(s, redisClient)
	server := grpc.NewServer()
	userproto.RegisterUserServiceServer(server, user_handler)

	lis, err := net.Listen("tcp", os.Getenv("server_url"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	log.Println("Server is listening on port ", os.Getenv("server_url"))
	if err = server.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}

}
