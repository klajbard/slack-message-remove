# Slack message remover by query

## How it works
Sets up a connection using Slack OATH and User token then queries for specific inputs. If there is several results (at least 20) Slack automatically break it into pages. This app starts from the very last page and goes one by one. It stop when each messages from the query is deleted.

## Setup
Create `.ENV` file in project root folder
```
export SLACK_APP_TOKEN="xapp-1-S8TG645FS-1452313548648-z489trzaieuh89weiufh89sw4utsl89w4eurfholreu8s94euof89ywe8f9yoe32"
export SLACK_OAUTH_TOKEN="xoxp-234672973493-234672973493-234672973493-73z4ri48czje8rz89w4jur8owu4rc83k"
```

### Obtain tokens
1. Head to [https://api.slack.com/apps](https://api.slack.com/apps)
2. Choose or create your app for workspace
3. Chose **"Basic Information"** from the left side navigation
4. **\[SLACK_APP_TOKEN\]** Scroll down to **"App-Level Tokens"** and generate new token
3. Go to **"OAuth & Permissions"**
4. **\[SLACK_OAUTH_TOKEN\]** Copy the token from **"User OAuth Token"**

## How to use
Run the `main.go` or build an executable from it and pass a parameter as a query string

#### Deleting messages which contains the _"apple"_ word
```
go run main.go apple
```

#### Deleting messages which contains _"big red apple"_ words
```
go run main.go "big red apple"
```

#### Deleting messages which is sent by bot user (named: BOTUSER)
```
go run main.go "from:@BOTUSER"
```

#### Deleting messages which is sent to the channel "general"
```
go run main.go "in:general"
```

#### Deleting messages which was sent to the channel "general" by "BOTUSER" at a specific date which contains "big red apple" words
```
go run main.go "in:#general before:2021-04-26 after:2021-04-24 from:@BOTUSER big red apple"
```

## Known issues
#### Deleting message failed: slack rate limit exceeded, retry after 1s
> Sometimes Slack throws error message. In this case the solution is to create a bit higher delay between delete requests by assigning a higher value to `DELETE_DELAY` variable in `main.go`.
#### Not every message was deleted
> Receiving message which could be a result for already deleting query could interfer. Possible solution is to run it twice.
#### Terminal output shows empty pages which doesn't contain any message
> It could happen that Slack also count deleted messages in pagination so it shows more results than it should.
