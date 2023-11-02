package repo

import (
	"context"
	"errors"
	"fmt"
)

func CreateUser(ctx context.Context, id string) error {
	if id == "123" {
		return errors.New("ERROR")
	}
	fmt.Println("Nothing's gonna change my love for you")
	return nil
}

func Updateuser(ctx context.Context, id string) error {
	if id == "123" {
		return errors.New("ERROR")
	}
	fmt.Println("Nothing's gonna change my love for you")
	return nil
}
