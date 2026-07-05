// Copyright 2025 John Wang. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package skills provides a skill pack with OpenClaw-compatible markdown skills.
//
// This pack bundles 18 markdown skills from OpenClaw for use with
// any omniskill-compatible agent. Skills are embedded at compile time
// and require no external dependencies at runtime.
//
// Usage:
//
//	import skills "github.com/plexusone/omniskill-pack"
//
//	agent, _ := agent.New(config,
//	    agent.WithSkillPack(skills.Default()),
//	)
package skills

import (
	"embed"
	_ "embed"

	"github.com/plexusone/omniskill/pack"
)

//go:embed skills/*
var skillsFS embed.FS

//go:embed VERSION
var version string

// Pack implements pack.SkillPack for the omniskill-pack skills.
type Pack struct{}

// Name returns the pack identifier.
func (Pack) Name() string { return "omniskill-pack" }

// Version returns the OpenClaw commit hash these skills were sourced from.
func (Pack) Version() string { return version }

// FS returns the embedded filesystem containing skills.
func (Pack) FS() embed.FS { return skillsFS }

// Default returns the default skill pack instance.
func Default() *Pack { return &Pack{} }

// Ensure Pack implements pack.SkillPack.
var _ pack.SkillPack = (*Pack)(nil)
