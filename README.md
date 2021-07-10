# github_analytics

## Tasks:
- Top 10 active users sorted by amount of PRs created and commits pushed
- Top 10 repositories sorted by amount of commits pushed
- Top 10 repositories sorted by amount of watch events

## HOWTO:

Go to this repository and install package:

``` go install ```

Each task supported special command for that:

### Top 10 active users sorted by amount of PRs created and commits pushed

```github_analytics topActors```

### Top 10 repositories sorted by amount of commits pushed

```github_analytics topRepo -t commits```

### Top 10 repositories sorted by amount of watch events

```github_analytics topRepo -t watch```


Each command supported optional (positive) number of records. For example:

```github_analytics topRepo -t watch -n 3```

Result:

```
Top repositories
| ID: 230501783 | Repository Name: lihkg-backup/thread | Count: 331 |
| ID: 224857031 | Repository Name: otiny/up | Count: 222 |
| ID: 227725053 | Repository Name: ripamf2991/ntdtv | Count: 167 |

```

Learn more in help sections for each command.


Also you can use it without installation:

``` go run main.go topActors```

----
ETA by HR: 2-3 hours

Real spent time: 4 hours