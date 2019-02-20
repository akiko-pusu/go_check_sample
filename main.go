package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"./local"
)

var home   = os.Getenv("HOME")
var user   = os.Getenv("USER")
var gopath = os.Getenv("GOPATH")

func main() {
	log.Print(local.GlobalStr)
	// パッケージの非公開変数は呼び出せないので、ここをコメントアウトするとコンパイルできません
	// log.Print(local.localStr)

	if err := sample(); err != nil {
		log.Println("err", err)
		os.Exit(1)
	}

	// エラーの場合はここには到達しません
	foo()
	bar()
}


func doError(msg string) error {
	return errors.New(fmt.Sprintf("Errorを発生します: %s", msg))
}

func foo() {
	log.Print("Hello world from foo!")
}

func bar() {
	log.Print("Hello world from bar!")
}

func sample() error {
	if user != "takano" {
		// panic
		panic("$USERがtakanoではありません、panicで終了....。")

		//return doError("$USERがtakanoではありません、doErrror()でエラーを発生。")
	}
	log.Printf("Hello world from %s", user)
	return nil
}
