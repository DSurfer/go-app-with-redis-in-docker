package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/DSurfer/goredis/redisconn"
	"github.com/gin-gonic/gin"
)

func UsersHandlerGetKey(c *gin.Context) {
	red := redisconn.GetRedisConnection()
	key := c.Query("key")
	dat, err := red.Get(key).Result()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Не найдено"})
		return
	}
	c.IndentedJSON(http.StatusOK, dat)
}

func UsersHandlerSetKey(c *gin.Context) {
	redisdb := make(map[string]string)
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Некорректный запрос"})
		return
	}
	err = json.Unmarshal([]byte(body), &redisdb)

	if err != nil {
		return
	}

	red := redisconn.GetRedisConnection()
	_, err = red.Set(GetKeyFromMap(redisdb), redisdb[GetKeyFromMap(redisdb)], 0).Result()
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Error!"})
	}
	c.IndentedJSON(http.StatusOK, []byte(body))
}

func GetKeyFromMap(keys map[string]string) string {
	keysfrom := make([]string, 0, len(keys))
	for k := range keys {
		keysfrom = append(keysfrom, k)
	}
	return strings.Join(keysfrom, "")
}

func UsersHandlerDelKey(c *gin.Context) {
	redisdb := make(map[string]string)
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Некорректный запрос"})
		return
	}

	json.Unmarshal([]byte(body), &redisdb)

	red := redisconn.GetRedisConnection()
	_, err = red.Del(redisdb["key"]).Result()
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Error!"})
	}
}
