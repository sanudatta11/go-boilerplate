package server

import "fmt"

func Init(env string) {
	r := NewRouter(env)
	err := r.Run(fmt.Sprintf(":%d", 8080))
	if err != nil {
		panic(err)
	}
}
