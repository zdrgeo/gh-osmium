# Osmium - GitHub CLI extension

This is Osmium - GitHub CLI extension. Osmium is a prototype tool for exploring the dependencies between elements in a software system module by analyzing the evolution of the files that contain these elements.

## Overview

See [Osmium](https://github.com/zdrgeo/osmium).

## Usage

Configuration

```
# config.env

SOURCE=github:pullrequest
GITHUB_TOKEN=
GITHUB_REPOSITORY_OWNER=
GITHUB_REPOSITORY_NAME=
```

Commands to [install the extension](https://docs.github.com/en/github-cli/github-cli/using-github-cli-extensions#installing-extensions)

```
gh extension install https://github.com/zdrgeo/gh-osmium

gh extension upgrade https://github.com/zdrgeo/gh-osmium
```

Commands to manipulate the DSM analyses

```
gh osmium analysis create --name="ticketing_tixets" --source="github:pullrequest"

gh osmium analysis change --name="ticketing_tixets" --source="github:pullrequest"

gh osmium analysis delete --name="ticketing_tixets"
```

Commands to manipulate the views

```
gh osmium view create --analysis-name="ticketing_tixets" --name="app" --node-name="app/Controller/*.php" --node-name="app/Service/*.php" --node-name="app/Repository/*.php"

gh osmium view change --analysis-name="ticketing_tixets" --name="app" --node-name="app/Controller/*.php" --node-name="app/Service/*.php" --node-name="app/Repository/*.php"

gh osmium view delete --analysis-name="ticketing_tixets" --name="app"

gh osmium view render --analysis-name="ticketing_tixets" --name="app"

gh osmium view listen --analysis-name="ticketing_tixets" --name="app"
```
