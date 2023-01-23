# KubeFS

This project is setup using GitOps pipelines that update source code when new tags are pushed. A description of pipeline events can be found below.

**Event Triggers**
- [Pre-Push Hook](#pre-push-hook)
- [On Pull Request](#on-pull-request)
- [On Tag Push](#on-tag-push)
- [On Promotion](#on-promotion)

## Pre-Push Hook

Due to a limitation of GitHub Action, you must create the Server binary before pushing the code. `.git/hooks/pre-push`

```bash
#!/bin/sh

echo "Building server binary for arm64..."
cd server;
GOARCH=arm64 GOOS=linux go build -o kubefs ./cmd/kubefs;
mv kubefs build/kubefs;

git add build/kubefs;
git commit -s -m '[TASK] Commit new server binary for arm64';
```


## On Pull Request

- Check Commits
- Build and Test Code

## On Tag Push

- Compiles Server code into binary
- Tests Server code
- Unit Tests Frontend
- build a docker image and tags it with tag (Front and Server)
- Patches the image tags in the source code for staging folder ( Front and Server ) 
- Pushes back to git 
- Pushes Tag to Git

## On Promotion

- Patches the image tags in the source code for prod folder ( Front and Server )
