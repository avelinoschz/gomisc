# gomisc

Personal repository of Go snippets.

The goal of this project is to keep a simple archive of small examples, quick experiments, and loose Go concepts that I want to preserve. It is not meant to be a library or a formal application. It is more like a place to store useful code for revisiting ideas, testing approaches, and documenting findings that I do not want to lose in Gists or scattered notes.

## What Is Here

- `algorithms/`: small algorithm-focused examples.
- `concurrency/`: small primitives and patterns with goroutines, channels, `select`, and cancellation signals.
- `data-structures/`: simple implementations of stack, queue, and linked list.
- `design-patterns/`: short examples of classic patterns applied in Go.
- `solid/`: small snippets centered around SOLID principles.

## Naming Convention

Folders use descriptive slugs so examples are easy to revisit later. The repo is organized by topic first, then by a short example name such as `channel-basic`, `done-channel`, or `fisher-yates`.

## How To Use It

Each folder contains self-contained examples, usually in a `main.go`, so they can be read quickly or run independently.

```bash
go run ./algorithms/fisher-yates
go run ./data-structures/stack
go run ./concurrency/primitives/channel-basic
go run ./solid/open-closed/polymorphism
```

## Note

These examples are meant to be lightweight references. Some prioritize clarity over robustness or edge-case validation, because their main purpose is to act as practical reminders of Go concepts.
