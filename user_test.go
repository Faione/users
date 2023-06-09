package users

import (
	"encoding/json"
	"fmt"
	"os/user"
	"testing"
)

func TestListAll(t *testing.T) {
	users, err := ListAll(func(user *user.User) bool {
		return len(user.HomeDir) > 4 && user.HomeDir[:5] == "/home"
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(users) == 0 {
		t.Fail()
	}

	rlt, _ := json.MarshalIndent(users, "", "     ")
	fmt.Println(string(rlt))
}

func TestListLogged(t *testing.T) {
	loginUsers, err := ListLogged()

	if err != nil {
		t.Fatal(err)
	}

	rlt, _ := json.MarshalIndent(loginUsers, "", "     ")
	fmt.Println(string(rlt))
}

func ExampleListAll() {
	users, err := ListAll(func(user *user.User) bool {
		return len(user.HomeDir) > 4 && user.HomeDir[:5] == "/home"
	})
	if err != nil {
		return
	}

	if len(users) == 0 {
		return
	}

	rlt, _ := json.MarshalIndent(users, "", "     ")
	fmt.Println(string(rlt))
}

func ExampleListLogged() {
	loginUsers, err := ListLogged()

	if err != nil {
		return
	}

	rlt, _ := json.MarshalIndent(loginUsers, "", "     ")
	fmt.Println(string(rlt))
}
