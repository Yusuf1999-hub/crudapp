package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "github.com/Yusuf-1999/crudapp/proto"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

var userConnection pb.UserProfilesClient

func CreateUser(c *gin.Context) {
	var newUser pb.UserProfile

	if err := c.ShouldBindJSON(&newUser); err != nil {
		return
	}

	req := &pb.CreateUserProfileRequest{
		UserProfile: &pb.UserProfile{
			Email:     newUser.Email,
			FirstName: newUser.FirstName,
			LastName:  newUser.LastName,
			Id:        newUser.Id,
		},
	}

	res, err := userConnection.CreateUserProfile(context.Background(), req)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, res)
}

func GetUser(c *gin.Context) {

	id := c.Param("userId")

	req := &pb.GetUserProfileRequest{
		Id: id,
	}

	res, err := userConnection.GetUserProfile(context.Background(), req)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, res)
}

func DeleteUser(c *gin.Context) {

	id := c.Param("userId")

	req := &pb.DeleteUserProfileRequest{Id: id}

	res, err := userConnection.DeleteUserProfile(context.Background(), req)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, res)
}

func UpdateUser(c *gin.Context) {

	var newUser pb.UserProfile

	if err := c.ShouldBindJSON(&newUser); err != nil {
		return
	}

	req := &pb.UpdateUserProfileRequest{
		UserProfile: &pb.UserProfile{
			Id:        newUser.Id,
			FirstName: newUser.FirstName,
			LastName:  newUser.LastName,
			Email:     newUser.Email,
		},
	}

	res, err := userConnection.UpdateUserProfile(context.Background(), req)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, res)
}

func ListUsers(c *gin.Context) {

	query := c.Param("query")

	req := &pb.ListUsersProfilesRequest{Query: query}

	res, err := userConnection.ListUsersProfiles(context.Background(), req)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, res)
}

func main() {

	fmt.Println("Helllo i am  a client")

	conn, err := grpc.Dial("localhost:9500", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer conn.Close()

	userConnection = pb.NewUserProfilesClient(conn)

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	router := gin.Default()

	router.POST("/users", CreateUser)
	router.GET("/users/query/:query", ListUsers)
	router.GET("/users/:userId", GetUser)
	router.PUT("/users/update", UpdateUser)
	router.DELETE("/users/:userId", DeleteUser)

	router.Run("localhost:5000")
}
