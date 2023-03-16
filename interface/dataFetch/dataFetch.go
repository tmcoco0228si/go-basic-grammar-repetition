package main

import (
	"fmt"
	"math/rand"
	"time"
)

// DataStoreインターフェースの定義
type DataStore interface {
	FetchData() ([]string, error)
}

// Database構造体の定義
type Database struct{}

// Database構造体がDataStoreインターフェースを実装
func (db Database) FetchData() ([]string, error) {
	// データベースからデータを取得する実装
	data := []string{"Apple", "Banana", "Cherry"}
	return data, nil
}

// API構造体の定義
type API struct{}

// API構造体がDataStoreインターフェースを実装
func (api API) FetchData() ([]string, error) {
	// APIからデータを取得する実装
	data := []string{"Dog", "Elephant", "Fox"}
	return data, nil
}

func fetchDataFromStore(ds DataStore) ([]string, error) {
	data, err := ds.FetchData()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	db := Database{}
	api := API{}

	var ds DataStore

	// ランダムにデータストアを選択
	if rand.Intn(2) == 0 {
		ds = db
	} else {
		ds = api
	}

	data, err := fetchDataFromStore(ds)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Data:", data)
}
