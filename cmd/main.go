package main

import "github.com/DevourTech/devourKV"

func main() {
	kvStore, err := devourKV.Init()
	if err != nil {
		panic(err)
	}

	err = kvStore.Run()
	if err != nil {
		panic(err)
	}
}
