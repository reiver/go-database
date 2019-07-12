/*
Package database_row represents a single row of a result from the database, by providing a type — ‘database_conclusion.Type’ — that can
be one of three kinds of values: ‘nothing’, ‘something’, ‘error’.


Example

Here is an example of how ‘database_row.Type’ might be used:

	var row database_row.Type
	
	// ...
	
	conclusion := row.Scan(&id, &givenName, &additionalName, &familyName)

	if err := conclusion.Return(); nil != err {
		return err
	}
*/
package database_row
