package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hossein-repo/myproj1/api/helper"
)

/*
type header struct {
	UserId  string
	Browser string
}
*/
/*
type personData struct {
	FirstName    string
	LastName     string

}
*/

type personData struct {
	FirstName    string `json:"first_name" binding:"required,alpha,min=4,max=10"`
	LastName     string `json:"last_name" binding:"required,alpha,min=6,max=20"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}


type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}


func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK,helper.GenerateBaseResponse("Working Test!",true,0))
}


//-----------------------------
/*
func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := personData{}
	c.ShouldBindJSON(&p)
	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder",
		"person": p,
	})
}

*/


func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := personData{}
	err := c.ShouldBindJSON(&p)


	if err != nil {
               c.AbortWithStatusJSON(http.StatusBadRequest,helper.GenerateBaseResponseWithValidationError(nil,false,-1,err))
		/*
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
          "validationErr": err.Error(),

		   })
		 
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder",
		"person": p,
	}) */


	return	
	}
	c.JSON(http.StatusOK,helper.GenerateBaseResponse(gin.H{
		"result": "BodyBinder",
		"person": p,
		}, true, 0))

}



/*

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type personData struct{
	FirstName: string 
	LastName : string
}

type TestHandler struct{
}

func NewTestHandler() *TestHandler{
	return &TestHandler{}
}

func(h *TestHandler) Test(c *gin.Context){
 c.JSON(http.StatusOK, gin.H{
"result":"Test",
 })
}

*/