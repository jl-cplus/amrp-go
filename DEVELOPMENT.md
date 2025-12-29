# Development Guidelines

> [!IMPORTANT]  
> All commits must include code formatting, and the code should be tested and built without errors before committing and pushing to the remote branch.

## Code Formatting

Run `go fmt` GO command-line tool to format the source code
```bash
$ go fmt ./...
```

## Commit Message

### Format
```
<type>: commit message
```

### Type
- **feat** - New feature.
- **fix** - Bug fix.
- **test** - Changes related to adding, modifying, or removing tests.
- **doc** - Changes related to adding, modifying, or removing documentation.
- **style** - Changes for formatting or indentation.
- **refactor** - Code changes excluding bug fixes and new features.
- **build** - Changes related to build tools and scripts.
- **task** - Routine tasks, such as updates dependency, setup project structure.
- **revert** - Revert previous commit.
- **merge** - For commits that are made as part of merging branches.