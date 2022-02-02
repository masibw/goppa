# goppa
Checks if there is any decrease in performance compared to the previous test. Currently, it simply notifies you if it takes more than 1.5 times longer than the previous time.

# Install

```bash
go install github.com/masibw/goppa@latest
```

# How to use
It is recommended to use a CI tool such as Github Actions rather than running it locally.

Basically, it can be executed as follows.
```Bash
go test ./... -v > prev.txt

// Suppose you make some changes.

go test ./... -v > current.txt
goppa --previous ./prev.txt --current ./current.txt
```

## Github Actions
Github Actions can be saved in a format such as `output-{commit-hash}` to compare with arbitrary test results.
For example, to compare the results of the current main branch with the results of a branch under development, you can create the following Github Actions.


I'll write later.