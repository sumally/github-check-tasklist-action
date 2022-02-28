# How to use

```yaml
name: pr-task-list-checker

on:
  pull_request:
    types:
      - opened
      - edited

jobs:
  pr-task-list-checker:
    runs-on: ubuntu-latest
    steps:
      - uses: sumally/github-check-tasklist-action@v1.0.0

```
