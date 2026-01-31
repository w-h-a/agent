package agent

import (
	"context"

	"github.com/w-h-a/agent/internal/service"
)

// TODO: support spaces
type Session struct {
	agent *service.Agent
	id    string
}

func (s *Session) ID() string {
	return s.id
}

func (s *Session) Ask(ctx context.Context, userInput string) (string, error) {
	return s.agent.Respond(ctx, "", s.id, userInput)
}

func (s *Session) Flush(ctx context.Context) error {
	return s.agent.Flush(ctx, s.id)
}
