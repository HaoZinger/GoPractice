package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

type record struct {
	Key  string          `json:"key"`
	Info MysqlBackupInfo `json:"info"`
}

type InfoStore struct {
	infos map[string]MysqlBackupInfo
	mu    sync.Mutex
	save  chan record
}

type MysqlBackupInfo struct {
	Name       string `json:"name" binding:"required"`
	Path       string `json:"path" binding:"required"`
	BackupDate string `json:"backupDate" binding:"required"`
	NFSIp      string `json:"nfsIp" binding:"required"`
}

var s *InfoStore

func main() {
	s = NewInfoStore("store.txt")
	router := setupRouter()
	router.Run(":8080")
}

func NewInfoStore(fileName string) *InfoStore {
	i := &InfoStore{
		infos: make(map[string]MysqlBackupInfo),
		save:  make(chan record, 20),
	}
	if e := i.load(fileName); e != nil {
		log.Println("Error loading InfoStore:", e)
	}
	go i.saveLoop(fileName)
	return i
}

func (s *InfoStore) load(fileName string) error {
	file, e := os.Open(fileName)
	if e != nil {
		return e
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var err error
	for err == nil {
		var r record
		if err = decoder.Decode(&r); err == nil {
			s.set(r.Key, r.Info)
		}
	}
	if err == io.EOF {
		return nil
	}
	log.Println("Error loading file:", err)
	return err
}

func (s *InfoStore) set(key string, info MysqlBackupInfo) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(strings.TrimSpace(info.BackupDate)) > 0 {
		s.infos[key] = info
	}
	return
}

func (s *InfoStore) put(key string, info MysqlBackupInfo) string {
	s.set(key, info)
	s.save <- record{key, info}
	return key
}

func (s *InfoStore) saveLoop(filename string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("openfile error:", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	for {
		<-s.save
		file.Seek(0, 0)
		for k, v := range s.infos {
			if err := encoder.Encode(record{k, v}); err != nil {
				log.Println("InfoStore:", err)
			}
		}
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
	router.POST("/:namespace", func(c *gin.Context) {
		namespace := c.Params.ByName("namespace")
		if len(namespace) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "namespace不能为空"})
			return
		}
		backupInfos := make([]MysqlBackupInfo, 5)
		if err := c.ShouldBindJSON(&backupInfos); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for _, v := range backupInfos {
			s.put(namespace+"_"+v.Name, v)
		}
		c.JSON(http.StatusCreated, s.infos)
		return
	})
	router.GET("/backupInfos", func(context *gin.Context) {

	})

	return router
}
