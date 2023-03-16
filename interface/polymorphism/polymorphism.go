package main

import "fmt"

// Animalインターフェースの定義
type Animal interface {
    Speak() string
}

// Dog構造体の定義
type Dog struct{}

// Dog構造体がAnimalインターフェースを実装
func (d Dog) Speak() string {
    return "Woof!"
}

// Cat構造体の定義
type Cat struct{}

// Cat構造体がAnimalインターフェースを実装
func (c Cat) Speak() string {
    return "Meow!"
}

// Bird構造体の定義
type Bird struct{}

// Bird構造体がAnimalインターフェースを実装
func (b Bird) Speak() string {
    return "Tweet!"
}

func MakeAnimalsSpeak(animals []Animal) {
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}

func main() {
    dog := Dog{}
    cat := Cat{}
    bird := Bird{}

    // Animalインターフェースを実装した構造体のスライス
    animals := []Animal{dog, cat, bird}

    // 各動物のSpeak()メソッドを呼び出す
    MakeAnimalsSpeak(animals)
}
