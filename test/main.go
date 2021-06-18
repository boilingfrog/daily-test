package main

import "golang.org/x/sync/errgroup"

type RecentlyHandleQuestionnaireBookOutput struct {
	RecommendBooks []string `json:"recommendBooks,omitempty"`
	UserTagDesc    string   `json:"userTagDesc,omitempty"`
	LikePercent    float64  `json:"likePercent,omitempty"` // 用户喜欢的百分比
	IsFirst        bool     `json:"isFirst"`
}

func main() {
	var output = &RecentlyHandleQuestionnaireBookOutput{}

	var eg errgroup.Group

	eg.Go(func() error {
		output.IsFirst = true
		return nil
	})

	eg.Go(func() error {
		output.LikePercent = 32
		return nil
	})

	if err := eg.Wait(); err != nil {
		return
	}

}
