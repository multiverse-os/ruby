package rbmarshal_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	rbmarshal "github.com/damonchen/rubymarshal"
	"testing"
)

// User user
type User struct {
	Name string `ruby:"name"`
	Age  int    `ruby:"age"`
}

// Profile profile
type Profile struct {
	User   User     `ruby:"user"`
	Job    string   `ruby:"job"`
	Time   int64    `ruby:"time"`
	Names  []string `ruby:"names"`
	Filter string   `ruby:"filter;key:string"`
}

func TestProfileMarshal(t *testing.T) {

	v := Profile{
		User: User{
			Name: "damon",
			Age:  18,
		},
		Job:  "programmer",
		Time: 1568104088,
	}
	buff := bytes.NewBufferString("")
	err := rbmarshal.NewEncoder(buff).Encode(&v)
	if err != nil {
		t.Error(err)
	}

	expected := "04087b0a3a09757365727b073a096e616d6549220a64616d6f6e063a0645543a0861676569173a086a6f6249220f70726f6772616d6d6572063b07543a0974696d656c2b07985e775d3a0a6e616d65735b0049220b66696c746572063b0754492200063b0754"
	s := hex.EncodeToString(buff.Bytes())
	if expected != s {
		fmt.Println(s)
		t.Error("expected not value")
	}
}

func TestProfileArray(t *testing.T) {
	type Ruby struct {
		profile []*Profile `ruby:"profile"`
	}

	ruby := Ruby{profile: []*Profile{
		&Profile{
			User: User{
				Name: "damon",
				Age:  18,
			},
			Job:  "programmer",
			Time: 1568104088,
		},
	}}
	buff := bytes.NewBufferString("")
	err := rbmarshal.NewEncoder(buff).Encode(&ruby)
	if err != nil {
		t.Error(err)
	}
	s := hex.EncodeToString(buff.Bytes())
	expected := "04087b063a0c70726f66696c655b067b083a09757365727b073a096e616d6549220a64616d6f6e063a0645543a0861676569173a086a6f6249220f70726f6772616d6d6572063b08543a0974696d656c2b07985e775d"
	if expected != s {
		t.Error("expected not value")
	}
}

func TestUnmarshal(t *testing.T) {
	value := "04087b083a09757365727b073a096e616d6549220a64616d6f6e063a0645543a0861676569173a086a6f6249220c70726f6772616d063b07543a0a6e616d65735b0749220631063b075449220632063b0754"
	bin, err := hex.DecodeString(value)
	if err != nil {
		t.Error(err)
	}

	v := Profile{}
	buff := bytes.NewBuffer(bin)
	err = rbmarshal.NewDecoder(buff).Decode(&v)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%#v", v)

}
