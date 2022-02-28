# How to use

```yaml
name: pr-task-list-checker

on:
  pull_request:
    types:
      - opened
      - edited

jobs:
  e2e:
    runs-on: ubuntu-latest
    steps:
      - uses: sumally/github-check-tasklist-action

```
