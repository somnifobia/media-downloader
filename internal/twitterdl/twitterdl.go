package twitterdl

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
)

func Download(rawURL, outputDir string) error {
    // Ensure output directory exists
    if outputDir == "" {
        outputDir = "."
    }
    if err := os.MkdirAll(outputDir, 0o755); err != nil {
        return fmt.Errorf("failed to create output directory: %w", err)
    }

    // Build output template: twitter_<id>.%(ext)s inside outputDir
    tweetID, err := extractID(rawURL)
    if err != nil {
        return err
    }

    outTemplate := filepath.Join(outputDir, "twitter_"+tweetID+".%(ext)s")

    // yt-dlp command:
    // yt-dlp -o <template> <url>[web:103][web:105][web:108]
    cmd := exec.Command("yt-dlp", "-o", outTemplate, rawURL)

    // Attach stdio so user can see yt-dlp progress
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("yt-dlp failed: %w", err)
    }

    fmt.Printf("Twitter/X %s downloaded via yt-dlp to %s\n", tweetID, outputDir)
    return nil
}
