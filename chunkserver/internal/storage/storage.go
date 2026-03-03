package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Storage struct{
	BasePath string
}

func New(basePath string) *Storage{
	_ = os.MkdirAll(basePath, os.ModePerm)
	return &Storage{BasePath: basePath}
}

func (s *Storage) chunkPath(id string) string{
	return filepath.Join(s.BasePath, "chunk_" + id)
}

func (s *Storage)WriteChunk(id string, r io.Reader) error{
	file, err := os.Create(s.chunkPath(id))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, r)
	return err
}

func (s *Storage) ReadChunk(id string) (*os.File, error){
	return os.Open(s.chunkPath(id))
}

func (s *Storage) DeleteChunk(id string) error{
	err:= os.Remove(s.chunkPath(id))
	if os.IsNotExist(err){
		return fmt.Errorf("chunk %s does not exist", id)
	}
	return err
}