# KubeFS

This project is setup using GitOps pipelines that update source code when new tags are pushed. A description of pipeline events can be found below.

**Event Triggers**
- [On Pull Request](#on-pull-request)
- [On Tag Push](#on-tag-push)
- [On Promotion](#on-promotion)

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
