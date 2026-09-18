# Required-env validator

```
envcheck.go
```
Check the test file next to the implementation to see how it actually works in practice.

Your app needs to crash immediately on boot if a required environment variable is missing. This utility handles that without pulling in extra modules.

It relies entirely on the Go standard library. No external services to run, no third-party dependencies to audit. Keeps the build fast and the dependency tree clean.