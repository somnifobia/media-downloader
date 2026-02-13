package twitterdl

import (
    "fmt"
    "net/url"
    "path"
    "regexp"
    "strings"
)

var idRe = regexp.MustCompile(`(\d{5,})`)

func extractID(rawURL string) (string, error) {
    u, err := url.Parse(rawURL)
    if err != nil {
        return "", fmt.Errorf("invalid URL: %w", err)
    }

    // /user/status/1234567890123456789
    parts := strings.Split(strings.Trim(path.Clean(u.Path), "/"), "/")
    for _, p := range parts {
        if idRe.MatchString(p) {
            return idRe.FindString(p), nil
        }
    }

    return "", fmt.Errorf("did not find ID in %s", rawURL)
}
