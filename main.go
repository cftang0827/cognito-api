package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

// Env for some share resource
type Env struct {
	cognito *cognitoidentityprovider.CognitoIdentityProvider
}

func main() {
	engine := gin.Default()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	env := Env{cognito: cognitoidentityprovider.New(sess)}

	// Route path
	engine.POST("/set_password", env.setPassword)

	engine.Run()

}

func (e *Env) setPassword(c *gin.Context) {
	password := c.PostForm("password")
	userpoolid := c.PostForm("user_pool_id")
	username := c.PostForm("user_name")
	result, err := e.cognito.AdminSetUserPassword(&cognitoidentityprovider.AdminSetUserPasswordInput{Password: &password, UserPoolId: &userpoolid, Username: &username})

	if err != nil {
		mesg := err.Error()
		c.JSON(400, gin.H{
			"message": mesg,
		})

	} else {
		mesg := result.GoString()
		c.JSON(200, gin.H{
			"message": mesg,
		})
	}
}
