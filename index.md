# GitRelease
A tool to automate the process of generating release notes and fetching change logs of your projects in a git powered version control systems.

## Overview

In git projects generally while writing release notes for a new tag following steps are used:

* Title for the release is `Releasing <tag>`
* Description consists of a list of pull/merge requests merged between previous release and current tag.

The primary objective of this project is to automate the manual work required in above steps. Currently, this project supports following version control systems :
* [Github](https://github.com)
* [Gitlab](https://gitlab.com)  

## Setup Instructions

### Github Setup
Steps to generate personal access token on Github :

1. Sign in with your account credentials on [Github](https://github.com/login).
2. Visit [Personal Access Tokens](https://github.com/settings/tokens) and click on `Generate new token` button. A form will appear on your screen.
3. In note field type `Token for synrelease`.
4. Select `repo` scope for the token.
5. Click `Generate token` button at the bottom of form. 
6. Copy the token shown on your screen on form submission and use it while configuring this project. 

**NOTE : You can revoke the token anytime you want.**

### Gitlab Setup
Steps to generate personal access token on Gitlab :

1. Sign in with your account credentials on [Gitlab](https://gitlab.com/users/sign_in).
2. Visit [Personal Access Tokens](https://gitlab.com/profile/personal_access_tokens),
3. In `Name` field type `Token for synrelease`.
4. Set expiry date for the token in `Expires at` section. If you don't want the token to expire forever then leave that field as blank.
5. Select `api` and `read_api` as scope for the token.
6. Click `Create personal access token` button at the bottom of form.
7. Copy the token shown on your screen on form submission and use it while configuring this project. 

**NOTE : You can revoke the token anytime you want.**

## Support or Contact

Having trouble with the tool? Check out our [documentation](https://github.com/sychonet/gitrelease/blob/master/README.md) or raise an [issue](https://github.com/sychonet/gitrelease/issues) and we’ll help you sort it out. We will love to hear feedbacks about the codebase and about new features that can be added to it. For that as well please create an [issue](https://github.com/sychonet/gitrelease/issues). In case you want to contribute in this project then please fork this repository and open a pull request with your changes.
