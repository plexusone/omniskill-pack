// Copyright 2025 John Wang. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package skills

import (
	"io/fs"
	"strings"
	"testing"

	osfs "github.com/grokify/oscompat/fs"
)

func TestPack(t *testing.T) {
	p := Default()

	if p.Name() != "omniagent-skills" {
		t.Errorf("Name() = %q, want %q", p.Name(), "omniagent-skills")
	}

	version := strings.TrimSpace(p.Version())
	if len(version) != 40 {
		t.Errorf("Version() = %q, want 40-char commit hash", version)
	}
}

func TestPackFS(t *testing.T) {
	p := Default()
	fsys := p.FS()

	// Expected skills (18 total)
	expectedSkills := []string{
		"github",
		"tmux",
		"weather",
		"coding-agent",
		"notion",
		"slack",
		"trello",
		"gh-issues",
		"summarize",
		"openai-whisper-api",
		"xurl",
		"healthcheck",
		"blogwatcher",
		"gemini",
		"session-logs",
		"oracle",
		"skill-creator",
		"goplaces",
	}

	for _, skill := range expectedSkills {
		path := "skills/" + skill + "/SKILL.md"
		_, err := fs.Stat(fsys, path)
		if err != nil {
			t.Errorf("skill %q not found: %v", skill, err)
		}
	}

	// Count total skills
	entries, err := fs.ReadDir(fsys, "skills")
	if err != nil {
		t.Fatalf("ReadDir(skills) failed: %v", err)
	}

	if len(entries) != len(expectedSkills) {
		t.Errorf("found %d skills, want %d", len(entries), len(expectedSkills))
	}
}

func TestSkillMDFormat(t *testing.T) {
	p := Default()
	fsys := p.FS()

	// Check that github skill has valid YAML frontmatter
	content, err := fs.ReadFile(fsys, "skills/github/SKILL.md")
	if err != nil {
		t.Fatalf("ReadFile(github/SKILL.md) failed: %v", err)
	}

	// Normalize line endings for cross-platform compatibility
	s := osfs.NormalizeLineEndings(string(content))
	if !strings.HasPrefix(s, "---\n") {
		t.Error("SKILL.md should start with YAML frontmatter (---)")
	}

	if !strings.Contains(s, "name: github") {
		t.Error("SKILL.md should contain 'name: github' in frontmatter")
	}
}
