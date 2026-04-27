# omniagent-skills

[![Go Reference](https://pkg.go.dev/badge/github.com/plexusone/omniagent-skills.svg)](https://pkg.go.dev/github.com/plexusone/omniagent-skills)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Default skill pack for OmniAgent with 18 OpenClaw-compatible markdown skills.

## Installation

```bash
go get github.com/plexusone/omniagent-skills
```

## Usage

```go
import (
    "github.com/plexusone/omniagent/agent"
    skills "github.com/plexusone/omniagent-skills"
)

// Use all skills
agent, err := agent.New(config,
    agent.WithSkillPack(skills.Default().FS()),
)

// Or filter to specific skills
agent, err := agent.New(config,
    agent.WithSkillPack(skills.Default().FS()),
    agent.WithSkillIncludes("github", "weather", "tmux"),
)
```

## Included Skills

### Core Skills

| Skill | Description | Requirements |
|-------|-------------|--------------|
| github | GitHub CLI operations | `gh` |
| tmux | Terminal multiplexer control | `tmux` |
| weather | Weather information | `curl` |
| coding-agent | Sub-agent delegation | `claude`/`codex` |

### Utility Skills

| Skill | Description | Requirements |
|-------|-------------|--------------|
| notion | Notion API integration | `curl` |
| slack | Slack messaging | `curl` |
| trello | Trello boards | `curl` |
| gh-issues | GitHub issues | `gh` |
| summarize | URL/video summarization | `summarize` |
| openai-whisper-api | Speech-to-text | `curl` |
| xurl | URL utilities | `curl` |
| healthcheck | Service monitoring | `curl` |
| blogwatcher | RSS/blog monitoring | `curl` |
| gemini | Gemini API | `curl` |

### Meta Skills

| Skill | Description | Requirements |
|-------|-------------|--------------|
| session-logs | Session history | None |
| oracle | Structured queries | None |
| skill-creator | Meta-skill creation | None |
| goplaces | Location services | `curl` |

## Source

Skills are sourced from [OpenClaw](https://github.com/openclaw/openclaw) at commit `d4eb23652362a1b7d3fbcebd633a1c6f2a43c16f`.

## License

MIT License - see [LICENSE](LICENSE) for details.
