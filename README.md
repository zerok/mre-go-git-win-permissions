This repository contains (alongside this README) a single script file that has executable permissions.
It should serve as example for an issue in go-git when used on Windows where the permissions of that script are reset to 0644 (from 0755) and the file is therefore marked as modified when running `git status`.

## How to use

```
git clone https://github.com/zerok/mre-go-git-win-permissions.git
cd mre-go-git-win-permissions
go run .
```

