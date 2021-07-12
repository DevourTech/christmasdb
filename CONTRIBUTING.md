# Contributing to devourKV

Thank you so much for contributing to devourKV. We appreciate your time and help.
Here are some guidelines to help you get started.

## Filing a bug or feature

1. Before filing an issue, please check the existing issues to see if a
   similar one was already opened. If there is one already opened, feel free
   to comment on it.
2. If you believe you've found a bug, please provide detailed steps of
   reproduction, the version of devourKV and anything else you believe will be
   useful to help troubleshoot it (e.g. OS environment, environment variables,
   etc...). Also state the current behavior vs. the expected behavior.
3. If you'd like to see a feature or an enhancement please open an issue with
   a clear title and description of what the feature is and why it would be
   beneficial to the project and its users.

## General Rules

As much as possible, try to follow the existing format of markdown and code.

1. Don't forget to run `make test` and `make fmt` before submitting pull requests.
2. Make sure that 100% of your code is covered by tests.

## Conventions to be followed

1. Mention your **issue number** linked to the pull request while creating one :<br>
   `closes <Issue_No.>`
   
2. While creating a new branch, follow the naming conventions below :    
a. Adding a **new feature** :<br>
      `feature/kv- <Issue_No.>`

   b. Fixing a **bug** :<br>
      `bugfix/kv- <Issue_No.>`
   
3. While **committing the changes**, follow the commit message convention as below :<br>

   `#<Issue_No.> UPD : <commit message>`

## Quick steps to contribute

1. Fork the project.
2. Download your fork to your PC (`git clone https://github.com/your_username/devourKV && cd devourKV`)
3. Create your feature branch (`git checkout -b <branch name>`)
4. Make changes and run tests (`make test`)
5. Format the code (`make fmt`)   
6. Add them to staging (`git add .`)
7. Commit your changes (`git commit -m '#<Issue_No> UPD : <Commit Message>'`)
8. Push to the branch 
9. Create new pull request
   
