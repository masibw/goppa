# goppa
Checks if there is any decrease in performance compared to the previous test results. Currently, it simply notifies you if it takes more than 1.5 times longer than the previous time.

# Install

```bash
go install github.com/masibw/goppa@latest
```

# Quick Start
*It is recommended to use a CI tool such as GitHub Actions rather than running it locally.*

## Local
You will need to create a test result file for comparison.
```Bash
go test ./... -v > prev.txt

// Suppose you make some changes.

go test ./... -v > current.txt
```
Then, run goppa.
```Bash
goppa --previous ./prev.txt --current ./current.txt
```

## GitHub Actions
GitHub Actions can be saved in a format such as test-{commit-hash} to compare with arbitrary test results. For example, to compare the results of the current main branch with the results of a branch under development, you can create the following GitHub Actions.

In order to continuously compare the test results, we need to construct the following flow.

For the first run, we need to prepare the results of previous tests (e.g. main branch). So, add the following code and merge it into the main branch to create test results when pull request is merged.

```yaml
name: test
on:
  pull_request:
  push:
    branches:
      - main
jobs:
  test:
    name: Test local sources
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Sources
        uses: actions/checkout@v2
      - name: Setup Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.17.6
      - name: Restore cache
        uses: actions/cache@v2
        with:
          path: ~/go/pkg/mod
          key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
          restore-keys: |
            ${{ runner.os }}-go-
      - name: Get dependencies
        run: go mod download -x
      - name: Test
        run: go test -v ./...  > ./test-${{ github.sha }}.log
      - name: Save test results
        uses: actions/cache@v2
        with:
          key: test-${{ github.sha }}
          path: ./test-${{ github.sha }}.log
```

After the preparation is complete, you can add the action to run goppa when pull request is created.

```yaml
name: test
on:
  pull_request:
  push:
    branches:
      - main
jobs:
  test:
    name: Test local sources
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Sources
        uses: actions/checkout@v2
      - name: Setup Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.17.6
      - name: Restore cache
        uses: actions/cache@v2
        with:
          path: ~/go/pkg/mod
          key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
          restore-keys: |
            ${{ runner.os }}-go-
      - name: Get dependencies
        run: go mod download -x
      - name: Test
        run: go test -v ./...  > ./test-${{ github.sha }}.log
      - name: Save test results
        uses: actions/cache@v2
        with:
          key: test-${{ github.sha }}
          path: ./test-${{ github.sha }}.log

  goppa:
    if: ${{ github.event_name }} == "pull_request"
    needs: test
    runs-on: ubuntu-latest
    steps:
      - name: Setup Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.17.6
      - name: dump context
        run: echo '${{ toJSON(github) }}'
      - name: Install goppa
        run: go install github.com/masibw/goppa@latest
      - name: Load previous test result
        uses: actions/cache@v2
        with:
          key: test-${{ github.event.pull_request.base.sha }}
          path: ./test-${{ github.base.sha }}.log
      - name: Load current test result
        uses: actions/cache@v2
        with:
          key: test-${{ github.sha }}
          path: ./test-${{ github.sha }}.log
      - name: Run goppa
        run: goppa --previous ./test-${{ github.event.pull_request.base.sha }}.log --current ./test-${{ github.sha }}.log
```
## Flags
The following flags can be used.

| Flag     | Alias | Description                                               | Default |
|----------|-------|-----------------------------------------------------------|---------|
| previous | p     | previous test output file.(with go test -v option)        | None    |
| current  | c     | current test output file.(created with go test -v option) | None    |
| border   | b     | how many times slower than before to be detected.         | 1.5     | 
- When a pull request is created, a test is performed, and then the results are compared using goppa.
- When the pull request is merged, run the test again. (Because the hash of the latest commit will change)
