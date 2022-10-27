package main

import "final_project/routers"

func main() {
	r := routers.StartApp()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
