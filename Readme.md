# Streakify

Streakify is a lightweight command-line tool written in Go that generates a GitHub-style contribution graph using your local Git repositories. It works independently of GitHub, allowing you to visualize your coding activity across personal projects, private repositories, GitLab, Bitbucket, and even repositories that have never been pushed online.

## Features

* Recursively scans directories for Git repositories
* Generates a GitHub-style contribution graph in the terminal
* Filters commits by author email
* Works entirely offline
* Ignores common dependency directories such as `node_modules` and `vendor`

## Installation

Clone the repository:

```bash
git clone https://github.com/DuttaNeel07/streakify.git
cd streakify
```

Build the project:

```bash
go build -o streakify
```

Or run it directly:

```bash
go run .
```

## Usage

### Add repositories

Scan a directory and register all Git repositories inside it.

```bash
streakify -add ~/Developer
```

You can run this command multiple times to add repositories from different locations.

### Generate your contribution graph

```bash
streakify -email you@example.com
```

Streakify will read the commit history from all registered repositories and count only the commits authored by the specified email.

## How it works

1. Scans directories recursively for Git repositories.
2. Stores the discovered repository paths locally.
3. Reads the commit history of each repository.
4. Filters commits using the provided author email.
5. Groups commits by day.
6. Renders a contribution graph directly in the terminal.

## Project Structure

```text
.
├── main.go      # CLI entry point
├── scan.go      # Repository discovery and storage
├── stats.go     # Commit processing and graph generation
├── go.mod
└── LICENSE
```

## Why Streakify?

GitHub's contribution graph only reflects activity that GitHub can associate with your account. Many developers spend time working on private repositories, local experiments, client projects, or repositories hosted on other platforms.

Streakify uses your local Git history as the source of truth, giving you a more complete picture of your coding activity regardless of where your repositories are hosted.

## Future Improvements

* Automatic detection of Git email
* Support for multiple author emails
* Configuration file support
* Interactive terminal interface
* Custom color themes
* Longest streak and contribution statistics
* Weekly and monthly summaries
* Export graphs as SVG or PNG

## Contributing

Contributions, bug reports, and feature suggestions are always welcome. If you'd like to improve Streakify, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
