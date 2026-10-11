# Moodle CLI

This is a tool for S.P.I.T students to use `moodle.spit.ac.in` in the terminal. Its really frustating to use that site and to miss out on a deadline because its lost in the 100s of other batch assignments. This tool gets you assignments relevant only to you.

## What it does

- Fetches your upcoming deadlines from Moodle (the same list as the dashboard Timeline)
- Hides assignments meant for other batches, so `E8-B2-T1` disappears if you're in A4
- Keeps whole-class assignments (ones with no batch tag, like `Assignment 8`)
- Shows each deadline with how much time is left, in a terminal UI

## Requirements

- Go 1.26+
- A S.P.I.T Moodle account

## Setup

### 1. Get your Moodle token

```
curl "https://moodle.spit.ac.in/login/token.php?username=YOUR_USERNAME&password=YOUR_PASSWORD&service=moodle_mobile_app"
```

You'll get back something like `{"token":"abc123...","privatetoken":"..."}`. Copy the `token` value.

> Your token works like your password. Don't share it or commit it. If it leaks, reset it in Moodle under **Preferences → Security keys**.

### 2. Create a `.env` file

```
MOODLE_TOKEN=your_token_here
MYBATCH=A4
```

`MYBATCH` is your batch (A1–D4).

### 3. Run

```
export $(cat .env)
go run .
```

The `export` only lasts for the current terminal window, so run it again in each new terminal.

## Controls

| Key | Action |
|---|---|
| `↑` / `k` | move up |
| `↓` / `j` | move down |
| `q` / `ctrl+c` | quit |

## How batch filtering works

Assignment names carry a batch tag in one of two places:

| Format | Example |
|---|---|
| between dashes | `E8-A4-T2 (FM)` |
| at the start | `A1-Application-...`, `A1W9T2` |

An assignment is hidden only if it has a batch tag that isn't yours. Anything without a tag is shown. If a professor uses a new naming format, the assignment will show up rather than get hidden, so you won't miss anything.

## Development

```
go vet ./...
go test ./...
```
