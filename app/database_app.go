package app

import (
	"bytes"
	"encoding/json"
	"github.com/rahmaninsani/dns-api/helper"
	"github.com/rahmaninsani/dns-api/model/domain"
	"io"
	"log"
	"os"
	"path/filepath"
)

type DB struct {
	Path string
}

func NewDB(path string) *DB {
	db := &DB{
		Path: path,
	}

	db.createFileIfNotExist()

	return db
}

func (db *DB) createFolderIfNotExist() {
	dir := filepath.Dir(db.Path)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		log.Println("Directory not found, creating directory:", dir)

		err = os.MkdirAll(dir, 0755)
		helper.PanicIfError(err)

		log.Println("Directory created successfully:", dir)
	}
}

func (db *DB) createFileIfNotExist() {
	db.createFolderIfNotExist()

	_, err := os.Stat(db.Path)
	if os.IsNotExist(err) {
		log.Println("File not found, creating file:", db.Path)

		file, err := os.Create(db.Path)
		helper.PanicIfError(err)

		defer func(file *os.File) {
			err = file.Close()
			helper.PanicIfError(err)
		}(file)

		_, err = file.WriteString("[]")
		helper.PanicIfError(err)

		err = file.Sync()
		helper.PanicIfError(err)

		log.Println("File created successfully:", db.Path)
	}

	log.Println("Using file as database:", db.Path)
}

func (db *DB) readFile() *os.File {
	db.createFileIfNotExist()

	file, err := os.OpenFile(db.Path, os.O_RDWR, 0644)
	helper.PanicIfError(err)

	return file
}

func (db *DB) writeFile(securityCases []domain.SecurityCase) {
	db.TruncateDB()

	if len(securityCases) == 0 {
		securityCases = make([]domain.SecurityCase, 0)
	}

	securityCasesByte, err := json.MarshalIndent(securityCases, "", "    ")
	helper.PanicIfError(err)

	file := db.readFile()
	defer func(file *os.File) {
		err := file.Close()
		helper.PanicIfError(err)
	}(file)

	_, err = file.Write(securityCasesByte)
	helper.PanicIfError(err)

	err = file.Sync()
	helper.PanicIfError(err)

	log.Println("File written successfully")
}

func (db *DB) TruncateDB() {
	file := db.readFile()
	defer func(file *os.File) {
		err := file.Close()
		helper.PanicIfError(err)
	}(file)

	err := file.Truncate(0)
	helper.PanicIfError(err)
}

func (db *DB) getData() []domain.SecurityCase {
	file := db.readFile()
	defer func(file *os.File) {
		err := file.Close()
		helper.PanicIfError(err)
	}(file)

	securityCasesByte := make([]byte, 1024)
	for {
		n, err := file.Read(securityCasesByte)
		if err != io.EOF {
			helper.PanicIfError(err)
		}
		if n == 0 {
			break
		}
	}

	securityCasesByte = bytes.ReplaceAll(securityCasesByte, []byte("\x00"), []byte(""))
	securityCases := make([]domain.SecurityCase, 0)

	if len(securityCasesByte) == 0 {
		securityCasesByte = []byte("[]")
	}

	err := json.Unmarshal(securityCasesByte, &securityCases)
	helper.PanicIfError(err)

	return securityCases
}

func (db *DB) Create(securityCase domain.SecurityCase) error {
	securityCases := db.getData()

	securityCases = append(securityCases, securityCase)
	db.writeFile(securityCases)

	return nil
}

func (db *DB) Update(securityCase domain.SecurityCase) error {
	securityCases := db.getData()

	for i, sc := range securityCases {
		if sc.ID == securityCase.ID {
			securityCases[i] = securityCase
		}
	}

	db.writeFile(securityCases)

	return nil
}

func (db *DB) Delete(securityCaseID string) error {
	securityCases := db.getData()

	var newSecurityCases []domain.SecurityCase
	for _, sc := range securityCases {
		if sc.ID != securityCaseID {
			newSecurityCases = append(newSecurityCases, sc)
		}
	}

	db.writeFile(newSecurityCases)

	return nil
}

func (db *DB) Find() []domain.SecurityCase {
	return db.getData()
}

func (db *DB) FindOne(securityCaseId string) (domain.SecurityCase, error) {
	securityCases := db.getData()

	for _, sc := range securityCases {
		if sc.ID == securityCaseId {
			return sc, nil
		}
	}

	return domain.SecurityCase{}, nil
}
