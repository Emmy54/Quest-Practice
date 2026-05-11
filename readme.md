# Quest-Practice

## Overview

`Quest-Practice` is a Go practice repository created to prepare for piscine checkpoints and strengthen algorithmic problem solving in Go. The repository contains a set of focused exercise files and working examples designed to improve familiarity with common coding questions and build confidence in Go syntax, control flow, and standard library usage.

## Repository Structure

- `main.go` - Central test runner for executing and validating exercise implementations.
- `go.mod` - Go module definition, including the dependency on `github.com/01-edu/z01`.
- Individual Go files (e.g. `countalpha.go`, `findprevprime.go`, `reverseStrCap/main.go`) - Each file contains an isolated exercise or challenge implementation.
- Subdirectories - Feature-specific Go programs with their own `main.go` entry points, such as `AddPrimeSum/`, `Fprime/`, `HiddenP/`, `Inter/`, `ReverseStrCap/`, `Union/`, and `WdMatch/`.

## How to Run

From the repository root, use the Go toolchain:

```bash
go run .
```

For subdirectory programs, navigate into the desired folder and run:

```bash
go run .
```

## Goals

- Practice Go fundamentals and idiomatic patterns
- Build familiarity with common coding problems and edge cases
- Develop reusable, testable solutions for algorithmic exercises

## Notes

- Each file is intended to represent a separate exercise or question.
- The main entry point can be used as a test harness for verifying implementations.
- This repository is focused on incremental skill-building and preparation for coding assessments.

## License

This repository is provided for personal study and practice.