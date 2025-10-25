# Public Packages

This directory contains public Go packages that can be imported by external projects.

## Purpose

Code placed here is considered part of the public API and can be imported by other Go modules:

```go
import "github.com/example/project/pkg/packagename"
```

## Guidelines

- Only place stable, well-documented APIs here
- Maintain backward compatibility
- Keep internal implementation details in `/internal`
