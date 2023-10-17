```
All commands for robinhood control

rhctl instruments -h
rhctl instruments list --limit 5
rhctl instruments list --limit 100 -o json
rhctl feed -h
rhctl feed list
rhctl feed list --output json
rhctl feed list popular_list --output json
rhctl orders list

Todo:

* Add account querying (easy)
* submit orders (easy)
* Recursive Query if next is not nil (to all queries)
* List/Get All instruments (does not require auth)
* Add tests for schema change, json locally
* Get the lists
* Get the most famous list
* Get by state/industry/or any other category like ipo


Improvements:
Fix output json, piping to jq breaks limit

```