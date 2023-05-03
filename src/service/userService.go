package service

import "context"

type TestAPIUsers struct{}

func (c *TestAPIUsers) Greetings(ctx context.Context) (string, error) {
	return "Hi arun.! Welcome to Paramount!!", nil
}
