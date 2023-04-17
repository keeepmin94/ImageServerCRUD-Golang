package api

import (
	"encoding/base64"
	"fmt"
	"homework/homework_2/internal/global"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestURI struct {
	ImageId string `uri:"id"`
}

type RequestFile struct {
	Image string `form:"image"` //물어보기 
	ImageNo string `json:"no"`  //물어보기
}

type Response struct {
	Res string `json:"res"`
}

func ReadImage(c *gin.Context) {
	reqURI := RequestURI{}
	if err := c.ShouldBindUri(&reqURI); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res:"file is not valid"})

		return
	}

	c.File("images/" + reqURI.ImageId)
}

func CreateImage(c *gin.Context) {
	req := RequestFile{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res:"file is not valid"})

		return
	}
	
	b64data := req.Image[strings.IndexByte(req.Image, ',')+1:] //binary데이터만 자르기 
	imageBytes, err := base64.StdEncoding.DecodeString(b64data) //디코딩 
	
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res:"imageBytes is failed"})

		return
	}
	
	err = ioutil.WriteFile("images/" + req.ImageNo, imageBytes, 0644) //.jpg //바이트 슬라이스 형태로 저장 (0644:파일 권한 모드)소유자가 읽기 및 쓰기 권한을 가지고, 그룹 및 모든 다른 사용자가 읽기 권한만 가지는 파일

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res:"writefile is failed"})

		return
	}

	global.MaxImage++;
	c.JSON(http.StatusOK, Response{Res: "success"})
	
}

func UpdateImage(c *gin.Context)  {

}

func DeleteImage(c *gin.Context)  {
	reqURI := RequestURI{}
	if err := c.ShouldBindUri(&reqURI); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res:"file is not valid"})

		return
	}

	if err := os.Remove("images/" +reqURI.ImageId); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res:"remove failed"})

		return
	}

	global.MaxImage--;
	c.JSON(http.StatusOK, Response{Res: "success"})
}