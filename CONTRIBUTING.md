# Contributing Guide

[Instructions](https://contribute.cncf.io/maintainers/github/templates/required/contributing/#introduction)

* [New Contributor Guide](#contributing-guide)
  * [Ways to Contribute](#ways-to-contribute)
  * [Find an Issue](#find-an-issue)
  * [Ask for Help](#ask-for-help)
  * [Pull Request Lifecycle](#pull-request-lifecycle)
  * [Development Environment Setup](#development-environment-setup)
  * [Label Commits](#label-commits)
  * [Sign Your Commits](#sign-your-commits)
  * [Pull Request Checklist](#pull-request-checklist)

Welcome! We are glad that you want to contribute to our project! 💖

As you get started, you are in the best position to give us feedback on areas of
our project that we need help with including:

* Problems found during setting up a new developer environment
* Gaps in our Quickstart Guide or documentation
* Bugs in our automation scripts

If anything doesn't make sense, or doesn't work when you run it, please open a
bug report and let us know!

## Ways to Contribute

[Instructions](https://contribute.cncf.io/maintainers/github/templates/required/contributing/#ways-to-contribute)

We welcome many different types of contributions including:

* New features
* Builds, CI/CD
* Bug fixes
* Documentation
* Issue Triage
* Answering questions on Slack/Mailing List
* Web design
* Communications / Social Media / Blog Posts
* Release management

Not everything happens through a GitHub pull request. Please [contact us](mailto:casewylie@gmail.com) and let's discuss how we can work
together. 

## Find an Issue

[Instructions](https://contribute.cncf.io/maintainers/github/templates/required/contributing/#find-an-issue)

We have good first issues for new contributors and help wanted issues suitable
for any contributor. `good first issue` has extra information to
help you make your first contribution. `help wanted` are issues
suitable for someone who isn't a core maintainer and is good to move onto after
your first pull request.

Sometimes there won’t be any issues with these labels. That’s ok! There is
likely still something for you to work on like suggesting a feature request.

Once you see an issue that you'd like to work on, please post a comment saying
that you want to work on it. Something like "I want to work on this" is fine.

## Ask for Help

[Instructions](https://contribute.cncf.io/maintainers/github/templates/required/contributing/#ask-for-help)

The best way to reach us with a question when contributing is to ask on:

* The original github issue

## Pull Request Lifecycle

[Instructions](https://contribute.cncf.io/maintainers/github/templates/required/contributing/#pull-request-lifecycle)



## Development Environment Setup

[Instructions](https://contribute.cncf.io/maintainers/github/templates/required/contributing/#development-environment-setup)

- go 1.19

## Label Commits

Label commits as `BUGFIX|DOCS|FEATURE|TASK`

**_example_**:
```bash
git commit -s -am 'FEATURE: add customization for server port`
```

## Sign Your Commits

[Instructions](https://contribute.cncf.io/maintainers/github/templates/required/contributing/sign-your-commits)


### DCO
Licensing is important to open source projects. It provides some assurances that
the software will continue to be available based under the terms that the
author(s) desired. We require that contributors sign off on commits submitted to
our project's repositories. The [Developer Certificate of Origin
(DCO)](https://probot.github.io/apps/dco/) is a way to certify that you wrote and
have the right to contribute the code you are submitting to the project.

You sign-off by adding the following to your commit messages. Your sign-off must
match the git user and email associated with the commit.

    This is my commit message

    Signed-off-by: Your Name <your.name@example.com>

Git has a `-s` command line option to do this automatically:

    git commit -s -m 'This is my commit message'

If you forgot to do this and have not yet pushed your changes to the remote
repository, you can amend your commit with the sign-off by running 

    git commit --amend -s 


## Pull Request Checklist

When you submit your pull request, or you push new commits to it, our automated
systems will run some checks on your new code. We require that your pull request
passes these checks, but we also have more criteria than just that before we can
accept and merge it. We recommend that you check the following things locally
before you submit your code:

- commits are signed 
- all CI checks pass 
- the issue is referenced in the PR
