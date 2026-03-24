# AGENTS

## Purpose

This repository is a personal archive of small Go snippets.
It is not a production codebase or a polished learning course.
Examples are meant to be quick memory aids that are easy to revisit later.

## Structure

- Organize content by topic first.
- Use short descriptive slugs for example folders.
- Keep each example self-contained in its own folder, usually with a single `main.go`.

## Snippet Style

- Prefer small, runnable examples over complete or reusable designs.
- Avoid over-engineering, shared helper packages, and unnecessary abstraction.
- Keep the happy path unless a tiny guard is needed to avoid confusion.
- Repetition is acceptable when it keeps examples easier to read in isolation.

## Comments

- Start each snippet with a short header comment explaining what it shows.
- Use inline comments only when they explain intent or Go-specific behavior.
- Do not comment obvious syntax or narrate every line.

## Validation

- The main standard is that code compiles and runs as expected.
- Formal tests are not required for this repo.
- Prefer `go build ./...` for a quick repo-wide check and `go run ./path/to/example` for spot checks.
