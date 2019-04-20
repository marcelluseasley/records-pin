# Pindrop Code Challenge

This is the API for the Pindrop code challenge. 

## Endpoints

### Record

* `/v1/record/create`

This endpoint accepts a JSON request with the following structure:

```json
{
    "phone_number":"444123568",
    "carrier":"T-Mobile",
    "score":94.5
}
```

The record is added to a Sqlite3 database, which is created if not already present. The database name is `pin_records.db`.

* `/v1/record/read/{phoneNumber}`
* `/v1/record/update/{phoneNumber}`
* `/v1/record/delete/{phoneNumber}`


### Records

* `/v1/records/listall`
* `/v1/records/deleteall`