package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type BackupInfoStore struct {
	backupInfoMap map[string]string
	mu            sync.RWMutex
	save          chan Record
}
type Record struct {
	Key, Info string
}

const saveQueueLength = 1000

func NewBackupInfoStore(filename string) *BackupInfoStore {
	store := &BackupInfoStore{backupInfoMap: make(map[string]string), save: make(chan Record, saveQueueLength),}
	err := store.load(filename)
	checkError("load file ", err)
	go store.saveLoop(filename)
	return store
}
func (s *BackupInfoStore) Set(key, info string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if value, present := s.backupInfoMap[key]; present {
		timeStr := strings.ReplaceAll(value, "-", "")
		newTimeStr := strings.ReplaceAll(info, "-", "")
		time, e := strconv.Atoi(timeStr)
		checkError("Atoi", e)
		newTime, e := strconv.Atoi(newTimeStr)
		checkError("Atoi", e)
		if time <= newTime {
			s.backupInfoMap[key] = info
			return true
		}
		return false
	} else {
		s.backupInfoMap[key] = info
		return true
	}
}
func (s *BackupInfoStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.backupInfoMap[key]
}
func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(fmt.Sprintf("%s Error ", msg), err)
	}
}
func (s *BackupInfoStore) load(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
	defer file.Close()
	checkError("Open File", err)
	d := json.NewDecoder(file)
	var r map[string]string
	if err = d.Decode(&r); err == nil {
		s.backupInfoMap = r
	}
	if err == io.EOF {
		return nil
	}
	return err
}
func (s *BackupInfoStore) Put(key, info string) bool {
	if set := s.Set(key, info); set {
		s.save <- Record{key, info}
		return true
	}
	return false
}

func (s *BackupInfoStore) saveLoop(filename string) {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0666)
	defer file.Close()
	checkError("Open File", err)
	encoder := json.NewEncoder(file)
	for {
		<-s.save
		file.Seek(0, 0)
		err := encoder.Encode(s.backupInfoMap)
		checkError("Saving to file", err)
	}
}

var store *BackupInfoStore

func main() {
	store = NewBackupInfoStore("./info.json")
	http.HandleFunc("/", putInfo)
	http.ListenAndServe(":8080", nil)
}
func putInfo(w http.ResponseWriter, r *http.Request) {
	record := Record{}
	bytes, e := ioutil.ReadAll(r.Body)
	checkError("read body", e)
	json.Unmarshal(bytes, &record)
	put := store.Put(record.Key, record.Info)
	w.Write([]byte(strconv.FormatBool(put)))

}
