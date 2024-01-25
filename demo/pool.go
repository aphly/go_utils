package main

import (
	"fmt"
	"github.com/aphly/go_utils"
	"io"
	"net/http"
)

func main() {
	p := go_utils.NewPool(3, 3)
	p.GoTask(func() {
		for i := 0; i < 10; i++ {
			task := go_utils.Task{
				Id: i,
				Job: func(t go_utils.Task) {
					fmt.Printf("Task %d \n", t.Id)
					fmt.Printf("Task %d is running\n", i)
					resp, err := http.Get("http://www.aphly.com")
					if err != nil {
						fmt.Println(err)
						return
					}
					defer resp.Body.Close()
					body, _ := io.ReadAll(resp.Body)
					fmt.Println(string(body))
				},
			}
			p.AddTask(task)
		}
	})
	p.Run()
}
