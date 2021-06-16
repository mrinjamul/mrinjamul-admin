package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/mrinjamul-admin/message"
)

// GetProjectsHandler returns all current todo items
func GetProjectsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "comming soon",
	})
}

// GetMessagesHandler returns all current todo items
func GetMessagesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, message.Get())
}

// DeleteMessageHandler will delete a specified message based on user http input
func DeleteMessageHandler(c *gin.Context) {
	todoID := c.Param("id")
	if err := message.Delete(todoID); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

// MarkAsReadHandler will complete a specified todo based on user http input
func MarkAsReadHandler(c *gin.Context) {
	todoItem, statusCode, err := convertHTTPBodyToMessage(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	if message.MarkAsRead(todoItem.ID) != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

func convertHTTPBodyToMessage(httpBody io.ReadCloser) (message.Message, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return message.Message{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToMessage(body)
}

func convertJSONBodyToMessage(jsonBody []byte) (message.Message, int, error) {
	var messageItem message.Message
	err := json.Unmarshal(jsonBody, &messageItem)
	if err != nil {
		return message.Message{}, http.StatusBadRequest, err
	}
	return messageItem, http.StatusOK, nil
}
