package agent

import (
	"context"
	"strings"

	"github.com/w-h-a/agent/generator"
	"github.com/w-h-a/agent/internal/service"
	"github.com/w-h-a/agent/retriever"
	toolprovider "github.com/w-h-a/agent/tool_provider"
)

type ADK struct {
	agent *service.Agent
}

// TODO: Space

func (a *ADK) NewSession(ctx context.Context, sessionId string) (*Session, error) {
	if len(strings.TrimSpace(sessionId)) == 0 {
		var err error
		sessionId, err = a.agent.CreateSession(ctx, "")
		if err != nil {
			return nil, err
		}
	}
	return &Session{
		agent: a.agent,
		id:    sessionId,
	}, nil
}

func (a *ADK) Close() error {
	// TODO: implement
	return nil
}

func New(
	ctx context.Context,
	retriever retriever.Retriever,
	generator generator.Generator,
	toolProviders map[string]toolprovider.ToolProvider,
	context int,
	systemPrompt string,
) *ADK {
	agent := service.NewAgent(
		retriever,
		generator,
		toolProviders,
		context,
		systemPrompt,
	)

	adk := &ADK{
		agent: agent,
	}

	return adk
}
