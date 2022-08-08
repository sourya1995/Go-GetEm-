package main

import (
	"sync"
	"encoding/gob"
	"io"
	"log"
	"os"
)

const saveQueueLength = 1000

type URLStore struct {
	urls map[string]string
	mu sync.RWMutex
	save chan record
}

type record struct {
	Key, URL string
}

func NewURLStore(filename string) *URLStore{
	s := &URLStore{urls: make(map[string]string),
				   save: make(chan record, saveQueueLength),}
	
	
	if err := s.load(filename); err != nil {
		log.Println("Error loading URLStore:", err)
	}
	go s.saveLoop(filename)
	return s
}

func (s *URLStore) Get(key string) string{
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *URLStore) Set(key, url string) string{
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.urls[key]; present{
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {  
	for {
		key := genKey(s.Count()) //generate short URL
		if ok := s.Set(key, url); ok{ //if we are able to set key, value for given URL, then return the key
			s.save <- record{key, url}
			return key
		}
	}

	panic("Shouldn't get here")
}

func(s *URLStore) load() error {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening URLStore:", err)
		return err
	}
	defer f.Close()
	
	d := json.NewDecoder(s.file)
	var err error
	for err == nil{
		var r record
		if err = d.Decoder(&r); err == nil {
			s.Set(r.Key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	log.Println("Error decoding URLStore:", err)
	return err

}

func (s *URLStore) saveLoop(filename string){
	f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening URLStore:", err)
	}
	defer f.Close()
	e := json.NewEncoder(f)
	for {
		r := <-s.save
		if err := e.Encode(r); err != nil {
			log.Println("Error saving to URLStore:", err)
		}
	}
} 