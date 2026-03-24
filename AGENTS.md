# Project Guidelines

This file defines workspace-wide instructions for coding agents (Copilot and other LLM assistants) working in this repository.

## Scope

- Applies to the entire repository.
- Keep changes focused and minimal.
- Preserve public APIs unless a breaking change is explicitly requested.

## Repository Map

This repository is a small helper library split into independent packages. Prefer extending existing packages instead of adding new top-level folders.

- `berr/` (`package berr`): Reusable error primitives for API/service layers.
	- Key files: `berr/api_error.go`, `berr/errorMap.go`, `berr/helpers.go`.
	- Main types: `BError`, `BErrorBase`, `ErrorMap`.
	- Use for: status-aware errors and field-level validation errors.
	- Do not: mix transport concerns (HTTP writer logic) into this package.

- `bsql/` (`package bsql`): JSON-friendly wrappers around `database/sql` nullable types.
	- Key file: `bsql/types.go`.
	- Main types: `NullString`, `NullInt16`, `NullInt32`, `NullInt64`, `NullFloat64`, `NullBool`, `NullByte`, `NullTime`.
	- Use for: API payloads where SQL nullables must encode to scalar-or-null JSON.
	- Do not: add query-building, DB connection, or ORM logic here.

- `utility/` (`package bns`): HTTP and request/response helpers used by handlers/controllers.
	- Key files: `utility/response.go`, `utility/request.go`, `utility/headers.go`, `utility/condition.go`.
	- Main functions/types: `Response`, `OkResponse`, `ErrResponse`, `HttpHandler`, `SafeDecodeBodyToJson`, `AppendHeader`, `CopyHeaderFrom`, `Ternary`.
	- Use for: response shaping, centralized HTTP error conversion, and header/request convenience utilities.
	- Note: folder name is `utility`, but import/package name is `bns`; preserve this convention.

- `tests/` (`package tests_test`): cross-package tests and behavior verification.
	- Current state: minimal coverage (`tests/response_test.go` is mostly a stub).
	- Prefer: table-driven tests validating public behavior over internals.

Dependency direction should stay simple:

- `berr` and `bsql` are base utility packages and should remain broadly reusable.
- `utility` may depend on `berr` for HTTP error handling.
- Avoid introducing reverse dependencies from `berr`/`bsql` back to `utility`.

## Build and Test

Use the standard Go toolchain commands from repository root:

- `go test ./...`
- `go build ./...`
- `go vet ./...`
- `go fmt ./...`

Run at least `go test ./...` after code changes.

## Architecture and Boundaries

- Keep `berr` focused on reusable error primitives and mapping validation-style errors.
- Keep `bsql` focused on JSON-friendly wrappers around `database/sql` nullable types.
- Keep HTTP/controller helpers in `utility` (`package bns`), including `Response` and `HttpHandler`.
- Avoid adding cross-package coupling unless required.

## Conventions

- Favor nil-safe behavior for utility helpers.
- Keep `MarshalJSON` wrappers in `bsql` returning `null` for nil/invalid values.
- For HTTP flows, prefer existing helpers in `utility/response.go` and `utility/request.go` over duplicating patterns.
- Keep error handling compatible with `berr.BError` and `berr.ErrorMap` handling in `HttpHandler`.
- Preserve package naming as-is, including the `utility/` folder using `package bns`.

## Testing Guidance

- Add or update tests for behavior changes.
- Prefer table-driven tests for helper utilities.
- Keep tests deterministic and avoid external network dependencies.

## Known Gaps and Pitfalls

- `utility/SearchParams.go` is currently a commented stub and not production-ready.
- Documentation is minimal; if API behavior changes, update `README.md` with concise usage notes.

## Change Checklist

Before finishing a task:

- Run `go test ./...`.
- Run `go vet ./...` when changing logic-heavy code.
- Ensure formatting is clean (`go fmt ./...`).
- Verify imports and package boundaries stay consistent.
