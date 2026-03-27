# Testing Ex01: Table-Driven Tests

Goal:

- practice tests for pure functions
- get comfortable with the table-driven style

Prompt:

Implement `NormalizeTags` and then write tests.

Expected behavior:

- surrounding spaces are trimmed
- everything is converted to lowercase
- empty strings are ignored
- the result must not contain duplicates

Minimum cases:

- empty input
- tags with uppercase letters and spaces
- duplicate tags

## Run tests

```bash
go test ./interview-prep/03-testing/ex01
```
