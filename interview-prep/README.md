# Interview Prep

This directory contains an intensive preparation track for a Go technical interview with a backend/platform focus.

Goals:

- regain speed writing Go from memory
- practice HTTP, JSON, tests, caching, and concurrency
- rehearse platform-engineering interview discussions

Usage rules:

1. Move in order. Each block assumes the previous one.
2. `00-warmup/ex01-copy-by-hand` is the only fully solved exercise.
3. Everything after that uses starter code or small skeletons.
4. Try to solve each exercise on your own before bringing it back for review.
5. If you get stuck, read the exercise README before consulting external documentation.

Helpful commands:

```bash
go run ./interview-prep/00-warmup/ex01-copy-by-hand
go test ./...
```

Suggested path:

- Day 1: complete `00-warmup` + `01-go-core/ex01`
- Day 2: the rest of `01-go-core` + `02-http-json/ex01`
- Day 3: the rest of `02-http-json` + `03-testing/ex01`
- Day 4: the rest of `03-testing`
- Day 5: `04-cache-concurrency`
- Day 6: `05-mock-interview/mock01` + `06-platform-notes`
- Day 7: `05-mock-interview/mock02` + review weak points

Conventions:

- Write documentation and comments in English.
- Keep code identifiers in English.
- In backend exercises, aim for a consistent shape:
  - `handler`
  - `service`
  - `store` or `cache`
- Use the standard library by default.

Progress checklist:

- [ ] Warmup completed
- [ ] Go core completed
- [ ] First HTTP handler without checking docs
- [ ] First test with `httptest`
- [ ] In-memory cache with `sync.RWMutex`
- [ ] Exercise using `context.Context`
- [ ] 45-60 minute mock interview
- [ ] Review of platform/cloud notes
