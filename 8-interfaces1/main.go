package main

import (
	"fmt"
	"math"
)

// import (
// 	"encoding/json"
// 	"log/slog"
// 	"net/http"
// )

// type Account struct {
//   Username string
//   Email string
// }

// func handleCreateAccount(w http.ResponseWriter, r *http.Request) {
//   var account Account
//   if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
//     slog.Error("failed to decode the response body", "err", err)
//   }

//   if err := notifyAccountCreated(account); err != nil {
//     slog.Error("failed to notify account created", "err", err)
//     return
//   }

//   w.Header().Set("Content-Type", "application/json")
// }

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{4.5}
	r1 := rect{5, 6}
	shapes := []shape{&c1, &r1}

	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}
