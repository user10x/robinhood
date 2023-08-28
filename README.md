```
All commands
go run cmd/rhctl/*.go instruments -h
go run cmd/rhctl/*.go instruments list --limit 5
go run cmd/rhctl/*.go instruments list --limit 100 -o json
go run cmd/rhctl/*.go feed -h
go run cmd/rhctl/*.go feed list
go run cmd/rhctl/*.go feed list --output json
go run cmd/rhctl/*.go feed list popular_list --output json
go run cmd/rhctl/*.go orders list

Todo:

* Add account querying (easy)
* submit orders (easy)
* Recursive Query if next is not nil (to all queries)
* List All instruments
* Add tests for schema change, json locally
* Get the lists
* Get the most famous list
* Get by state/industry/or any other category


Improvements:
Fix output json, piping to jq breaks limit

```