# Helm Charts
This repository holds all schub open-source helm charts. Each Chart should hold it's own decumentation inside it, this page contains contribution guidelines.

## Git Workflow
This repo uses [Trunk Based Development](https://trunkbaseddevelopment.com/#scaled-trunk-based-development). We will use the "scaled version" because we have a fully distrubuted and async workflow, so it can be complex to keep code pattern and style consistent without pairing sessions, in other words PR reviews are a must.

## Branching pattern
- `main`: The trunk branch.
- `feat/<description>`: Feature branches.
- `fix/<description>`: Fix branches.
- `refactor/<description>`: Refactor branches.
- `hotfix/<description>`: Hotfix branches.
- `chore/<description>`: Chore branches (for automations and optimizations).
- `docs/<description>`: Documentation branches.

## Commit guidelines

- Main should never receive direct commits.
- Commit messages should follow the pattern:
    - `feat: <summary>`
    - `breaking: <summary>`
    - `fix: <summary>`
    - `refactor: <summary>`
    - `chore: <summary>`
    - `docs: <summary>`

## Pull Request guidelines
- Pull requests should be reviewed by at least one Architect.
- Pull requests should be created from feature branch to main.
- After pipeline is passing and code is approved, it should be squashed and merged to main.
- For a release, a release tag should be created on the main branch.

## Running tests
You will need to install golang, check `test/go.sum` for the right minimum version.

Go into test folder and run `go test ./...` this should run all tests.
