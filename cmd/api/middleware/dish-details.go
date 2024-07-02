package middleware

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type BodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}



func DishDetails() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pathSlice := strings.Split(ctx.FullPath(), "/")
		if !(ctx.Query("details") == "true" && pathSlice[2] == "dishes") {
			return
		}
		
		bw := &BodyWriter{body: &bytes.Buffer{}, ResponseWriter: ctx.Writer}
		ctx.Writer = bw

		ctx.Next()
		data := bw.body.String()

		obj := make(map[string]interface{})

		json.Unmarshal([]byte(data), &obj)
		log.Println("--->", data)

		obj["name"] = "NEW_NAME"
		obj["work"] = "NEW_WORK"

		updatedBody, _ := json.Marshal(obj)

		bw.ResponseWriter.WriteString(string(updatedBody))
		bw.body.Reset()
	}
}
