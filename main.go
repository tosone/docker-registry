package main

import (
	"fmt"
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/legacy/tarball"
	"github.com/google/go-containerregistry/pkg/name"
)

func main() {
	var basic = &authn.Basic{
		Username: "tosone",
		Password: "testpassword",
	}
	crane.WithAuth(basic)
	fmt.Println(crane.Catalog("localhost:5000"))
	var img, err = crane.Pull("localhost:5000/release/alpine:3.12")
	fmt.Println("22", img, err, "123")
	var ref, _ = name.ParseReference("localhost:5000/release/alpine:3.12")
	fmt.Printf("%+v\n", ref)
	var file, _ = os.Create("alpine.tar")
	tarball.Write(ref, img, file)
	file.Close()
}
