package main

import (
	chatgpt "GOproject/project1/a/gpt_token"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
)

func Cont(g *gin.Context) {
	Data := g.Request.Body
	read, _ := io.ReadAll(Data)
	postType := gjson.Get(string(read), "post_type").String()
	if postType == "message" {
		Message := gjson.Get(string(read), "message").String()
		if Message == "" {
			return
		} else {
			Chat := chatgpt.NewChat("填入token马上开冲")
			res, err := Chat.Send(Message)
			if err != nil {
				return
			}
			ms, err := json.Marshal(res.Message.Content.Parts)
			if err != nil {
				return
			}
			fmt.Println(string(ms))
			text := string(ms)
			text = strings.Replace(text, "[", "", -1)
			text = strings.Replace(text, "\"", "", -1)
			text = strings.Replace(text, "]", "", -1)
			text = strings.Replace(text, "\n", "", -1)
			text = strings.Replace(text, "\n\n", "", -1)

			g.JSON(http.StatusOK, gin.H{
				"reply": text,
			})
			//fmt.Println(res.Message.Content.Parts)
		}
	}
}
