# Osmium - GitHub CLI extension

This is a GitHub CLI extension to integrate Osmium with GitHub CLI. Osmium is a prototype tool for exploring the dependencies between elements in a software system module by analyzing the evolution of the files that contain these elements.

## Overview

See [Osmium](https://github.com/zdrgeo/osmium).

## Usage

### Workflow

Osmium's workflow is organized into three stages: generate an analysis, generate at least one view of the analysis, explore the views.
The output of each stage is persisted and can be used mutiple times as input for the next stage. This enables iterative style of work - each stage can be repeated with the same input, but with different parameters to produce differenet outputs for the next stage. This also helps with time- and resource-intensive stages such as the analysis generation, where multiple API calls (which may also be subject of rate limits) are often required to obtain the necessary historical data from the source.

### Configuration

```
# config.env

BASEPATH=
SOURCE=github:pullrequest
```

### Commands to manipulate the DSM analyses

```
gh osmium analysis create
    --analysis-name -a
    --source -s {github:pullrequest, git:commit}
    [--source-option -o]
```
| Parameter | Default | Optional | Description |
|--|--|--|--|
| --analysis-name -a | | | Name of the analysis. |
| --source -s | github:pullrequest | | Source of the analysis - github:pullrequest or git:commit. |
| --source-option -o | | Yes | Options of the source. If the source is github:pullrequest - repository-owner and repository-name. If the source is git:commit - repository-url and repository-path. |

```
gh osmium analysis change
    --analysis-name -a
    --source -s {github:pullrequest, git:commit}
    [--source-option -o]
```
| Parameter | Default | Optional | Description |
|--|--|--|--|
| --analysis-name -a | | | Name of the analysis. |
| --source -s | github:pullrequest | | Source of the analysis - github:pullrequest or git:commit. |
| --source-option -o | | Yes | Options of the source. If the source is github:pullrequest - repository-owner and repository-name. If the source is git:commit - repository-url and repository-path. |

```
gh osmium analysis delete
    --analysis-name -a
```
| Parameter | Default | Optional | Description |
|--|--|--|--|
| --analysis-name -a | | | Name of the analysis. |

#### Examples

Generate an analysis model using the GitHub repository with owner 'scaleforce' and name 'tixets' as a source. Store the generated analysis model to the user home directory under the name "ticketing_tixets".
```
gh osmium analysis create \
    --analysis-name="ticketing_tixets" \
    --source="github:pullrequest" \
    --source-option="repository-owner=scaleforce,repository-name=tixets"
```

Alter the stored analysis model with name "ticketing_tixets". Use the same GitHub repository with owner "scaleforce" and name "tixets" as a source.
```
gh osmium analysis change \
    --analysis-name="ticketing_tixets" \
    --source="github:pullrequest" \
    --source-option="repository-owner=scaleforce,repository-name=tixets"
```

Remove the stored analysis model with name "ticketing_tixets" from the user home directory.
```
gh osmium analysis delete \
    --analysis-name="ticketing_tixets"
```

### Commands to manipulate the views

```
gh osmium view create
    --analysis-name -a
    --view-name -v
    [--node-name -n]
```
| Parameter | Default | Optional | Description |
|--|--|--|--|
| --analysis-name -a | | | Name of the analysis. |
| --view-name -v | | | Name of the view. |
| --node-name -n | | Yes | Names of the nodes. |

```
gh osmium view change
    --analysis-name -a
    --view-name -v
    [--node-name -n]
```
| Parameter | Default | Optional | Description |
|--|--|--|--|
| --analysis-name -a | | | Name of the analysis. |
| --view-name -v | | | Name of the view. |
| --node-name -n | | Yes | Names of the nodes. |

```
gh osmium view delete
    --analysis-name -a
    --view-name -v
```
| Parameter | Default | Optional | Description |
|--|--|--|--|
| --analysis-name -a | | | Name of the analysis. |
| --view-name -v | | | Name of the view. |

```
gh osmium view terminal render
    --analysis-name -a
    --view-name -v
    --span-name -s
    --x-node-start
    --y-node-start
    --node-count
```
| Parameter | Default | Optional | Description |
|--|--|--|--|
| --analysis-name -a | | | Name of the analysis. |
| --view-name -v | | | Name of the view. |
| --span-name -s | | | Name of the span. |
| --x-node-start | | | Start of the X nodes. |
| --y-node-start | | | Start of the Y nodes. |
| --node-count | | | Count of the nodes. |

```
gh osmium view web-browser render
    --analysis-name -a
    --view-name -v
    --span-name -s
```
| Parameter | Default | Optional | Description |
|--|--|--|--|
| --analysis-name -a | | | Name of the analysis. |
| --view-name -v | | | Name of the view. |
| --span-name -s | | | Name of the span. |

```
gh osmium view web-browser listen
    --analysis-name -a
    --view-name -v
```
| Parameter | Default | Optional | Description |
|--|--|--|--|
| --analysis-name -a | | | Name of the analysis. |
| --view-name -v | | | Name of the view. |

```
gh osmium view csv render
    --analysis-name -a
    --view-name -v
    --span-name -s
```
| Parameter | Default | Optional | Description |
|--|--|--|--|
| --analysis-name -a | | | Name of the analysis. |
| --view-name -v | | | Name of the view. |
| --span-name -s | | | Name of the span. |

#### Examples

Generate a view model based on the analysis with name "ticketing_tixets". Store the generated view model to the user home directory under the name "app". Scope the view model to nodes with names that match any of the the file paths "app/Controller/\*.php", "app/Service/\*.php" or "app/Repository/\*.php".
```
gh osmium view create \
    --analysis-name="ticketing_tixets" \
    --view-name="app" \
    --node-name="app/Controller/*.php" \
    --node-name="app/Service/*.php" \
    --node-name="app/Repository/*.php"
```

Alter the stored view model with name "app" based on the analysis with name "ticketing_tixets". Scope the view to the same nodes.
```
gh osmium view change \
    --analysis-name="ticketing_tixets" \
    --view-name="app" \
    --node-name="app/Controller/*.php" \
    --node-name="app/Service/*.php" \
    --node-name="app/Repository/*.php"
```

Remove the stored view model with name "app" based on the analysis with name "ticketing_tixets" from the user home directory.
```
gh osmium view delete \
    --analysis-name="ticketing_tixets" \
    --view-name="app"
```

Render the stored view model with name "app" based on the analysis with name "ticketing_tixets" to the terminal.
```
gh osmium view terminal render \
    --analysis-name="ticketing_tixets" \
    --view-name="app"
    --x-node-start=80
    --y-node-start=20
    --node-count=40
```

Render the stored view model with name "app" based on the analysis with name "ticketing_tixets" to HTML file.
```
gh osmium view web-browser render \
    --analysis-name="ticketing_tixets" \
    --view-name="app"
```

Start serving the rendered to HTML file view model with name "app" based on the analysis with name "ticketing_tixets".
```
gh osmium view web-browser listen \
    --analysis-name="ticketing_tixets" \
    --view-name="app"
```

Render the stored view model with name "app" based on the analysis with name "ticketing_tixets" to CSV file.
```
gh osmium view csv render \
    --analysis-name="ticketing_tixets" \
    --view-name="app"
```
