# [Project Name]
[project description]

# Contributing
Here you will find some guidelines for contributing to this project

## Git Workflow
This repo uses [Trunk Based Development](https://trunkbaseddevelopment.com/#scaled-trunk-based-development). We will use the "scaled version" because we have a fully distrubuted and async workflow, so it can be complex to keep code pattern and style consistent without pairing sessions, in other words PR reviews are a must.

## Branching pattern
- `main`: The trunk branch.
- `feat/<jira-card-id>`: Feature branches.
- `fix/<jira-card-id>`: Fix branches.
- `refactor/<jira-card-id>`: Refactor branches.
- `hotfix/<jira-card-id>`: Hotfix branches.
- `chore/<jira-card-id>`: Chore branches (for automations and optimizations).
- `docs/<jira-card-id>`: Documentation branches.

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
