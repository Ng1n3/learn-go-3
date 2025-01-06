package main

import (
	"fmt"

	// "github.com/gogo/protobuf/test/data"
	// "github.com/golang/protobuf/proto"
	 "google.golang.org/protobuf/proto"
	"github.com/Ng1n3/tooling/person"
)

// "fmt"
// "os"

// "k8s.io/apimachinery/pkg/runtime/serializer/yaml"
// yaml "sigs.k8s.io/yaml/goyaml.v3"

// var b = 20

// func main() {
// 	true := false
// 	a := 10
// 	b := 30
// 	if true {
// 		a := 20
// 		fmt.Println(a)
// 	}
// 	fmt.Println(a, b)

// }

//govulcheck
// func main() {
// 	info := Info{}
// 	err := yaml.Unmarshal([]byte(data), &info)
// 	if err != nil {
// 		fmt.Println("error: %v\n", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("%+v\n", info)
// }

//go generate
//go:generate protoc -I=. --go_out=. person.proto
func main() {
	p := &person.Person{
		Name: "bob Bobson",
		Id: 20,
		Email: "bob@bobson.com",
	}
	fmt.Println(p)
	protoBytes, _ := proto.Marshal(p)
	fmt.Println(protoBytes)
	var p2 person.Person
	proto.Unmarshal(protoBytes, &p2)
	fmt.Println(&p2)
}

//[lecture] -> goimportss can be a better formating tool instead of gofmt; goimports -l -w.

//[lecture] -> When you add linters to your build process, follow the old maxim “trust, but verify.” Since the kinds of issues that linters find are fuzzier, they sometimes have false positives and false negatives. This means you don’t have to make the changes that they suggest, but you should take the suggestions seriously. Go developers expect code to look a certain way and follow certain rules, and if your code does not, it sticks out. Use staticcheck. other linters are golangci-lint, revive

/**
revive
staticcheck
golandci-lint
govulncheck -> check vulnerabilites
*/

//[lecture] -> WHiile you could use go generate to run anything at all, it is most cmmonly used by developers to run tools that (unsuprisingly, given the name) generate source docde.
