package main

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	aa := ABC{"mingzi", "总监"}
	jj, _ := Marshal(aa)
	fmt.Println("jj = ", string(jj))

	dd := &ABC{}
	jsons := "{\"name\": \"你好\", \"title\": \"jingli\"}"
	err := Unmarshal([]byte(jsons), dd)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", dd)
	fmt.Println("==================")
	fmt.Printf("type of %T", aa)

}

type ABC struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

func Marshal(msg interface{}) ([]byte, error) {
	buf, err := json.Marshal(msg)
	return buf, errors.Wrapf(err, "%T Marshal Error", msg)
}

func Unmarshal(buffer []byte, msg interface{}) error {
	err := json.Unmarshal(buffer, msg)
	return errors.Wrapf(err, "%T Unmarshal Error: %s", msg, buffer)
}
