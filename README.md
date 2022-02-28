# How to use

```yaml
name: pr-tasklist-checker

on:
  pull_request:
    types:
      - opened
      - edited

jobs:
  tasklist-checker:
    runs-on: ubuntu-latest
    steps:
      - uses: sumally/github-check-tasklist-action

```
