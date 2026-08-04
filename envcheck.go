// Package envcheck validates required environment variables, dependency-free.
package envcheck

func Missing(required []string, env map[string]string) []string {
	var out []string
	for _, k := range required {
		if env[k] == "" {
			out = append(out, k)
		}
	}
	return out
}
