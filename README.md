# Orchestration Services
Backend services for handling VM sessions, environments' pipeline, and the API for them.

## Development
Architecture - https://www.figma.com/file/pr02o1okRpScOmNpAmgvCL/Architecture

### Subtrees
#### shared
Shared types are in a subtree made from https://github.com/devbookhq/shared repository.

The subtree commands you need for controling this repo are:
```bash
git subtree add --prefix shared https://github.com/devbookhq/shared.git master
```

```bash
git subtree pull --prefix shared https://github.com/devbookhq/shared.git master
```

```bash
git subtree push --prefix shared https://github.com/devbookhq/shared.git master
```
