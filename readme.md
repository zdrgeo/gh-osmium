# Osmium - GitHub CLI extension

This is a GitHub CLI extension to integrate Osmium with GitHub CLI. Osmium is a prototype tool for exploring the dependencies between elements in a software system module by analyzing the evolution of the files that contain these elements.

## Overview

See [Osmium](https://github.com/zdrgeo/osmium).

## Usage

Configuration

```
# config.env

BASEPATH=
SOURCE=github:pullrequest
```

Commands to [install the extension](https://docs.github.com/en/github-cli/github-cli/using-github-cli-extensions#installing-extensions)

```
gh extension install https://github.com/zdrgeo/gh-osmium

gh extension upgrade https://github.com/zdrgeo/gh-osmium
```

Commands to manipulate the DSM analyses

```
gh osmium analysis create --analysis-name="ticketing_tixets" --source="github:pullrequest" --source-option="repository-owner=scaleforce,repository-name=tixets"

gh osmium analysis change --analysis-name="ticketing_tixets" --source="github:pullrequest" --source-option="repository-owner=scaleforce,repository-name=tixets"

gh osmium analysis delete --analysis-name="ticketing_tixets"
```

Commands to manipulate the views

```
gh osmium view create --analysis-name="ticketing_tixets" --view-name="app" --node-name="app/Controller/*.php" --node-name="app/Service/*.php" --node-name="app/Repository/*.php"

gh osmium view change --analysis-name="ticketing_tixets" --view-name="app" --node-name="app/Controller/*.php" --node-name="app/Service/*.php" --node-name="app/Repository/*.php"

gh osmium view delete --analysis-name="ticketing_tixets" --view-name="app"

gh osmium view terminal render --analysis-name="ticketing_tixets" --view-name="app"

gh osmium view web-browser render --analysis-name="ticketing_tixets" --view-name="app"

gh osmium view web-browser listen --analysis-name="ticketing_tixets" --view-name="app"

gh osmium view csv render --analysis-name="ticketing_tixets" --view-name="app"
```
