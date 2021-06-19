package main

import (
	"errors"
	"fmt"
)

type RecentlyHandleQuestionnaireBookOutput struct {
	RecommendBooks []string `json:"recommendBooks,omitempty"`
	UserTagDesc    string   `json:"userTagDesc,omitempty"`
	LikePercent    float64  `json:"likePercent,omitempty"` // 用户喜欢的百分比
	IsFirst        bool     `json:"isFirst"`
}

func main() {
	str, err := test1()
	defer errPrint(err)
	_ = str

	str2, err := test2()
	if err != nil {
		_ = err
	}
	_ = str2
}

func errPrint(err2 error) {
	fmt.Println(err2)
}

func test1() (string, error) {
	return "1", errors.New("err 1")
}

func test2() (string, error) {
	return "2", errors.New("err 2")
}
