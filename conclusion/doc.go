/*
Package database_conclusion provides a type — ‘database_conclusion.Type’ — that can be one of three kinds of values: ‘nothing’, ‘something’, ‘error’;
where ‘something’ is special (as compared to ‘option types’ in other programming languages) in that it doesn't value an inner value.


Nothing

The ‘nothing’ value for ‘database_conclusion.Type’ represents when nothing
has been loaded into ‘database_conclusion.Type’.

A variable of type ‘database_conclusion.Type’ that has not had anything
/assigned to it starts out with a value of ‘nothing’. I.e.,:

	var conclusion database_conclusion.Type

You can also assign a value of ‘nothing’ to a variable of ‘database_conclusion.Type’
with code similar to the following:

	conclusion := database_conclusion.Nothing()

This ‘database_conclusion.Nothing()’ can also be used for comparisons. Whether that be
in an if-statement:

	if database_conclusion.Nothing() == conclusion {
		// ...
	}

Or a switch-statement:

	switch conclusion {
	case database_conclusion.Nothing():
		// ...
	default:
		// ...
	}


Something

Type ‘database_conclusion.Type’ has a ‘something’ value that is a bit different than what you
might find with ‘option types’ in other programming languages.

With ‘database_conclusion.Type’ there isn't actually a value for ‘something’. It basically means
that everything was successful.

I.e., the variable was loaded, and there was no error.

You can also assign a value of ‘something’ to a variable of ‘database_conclusion.Type’
with code similar to the following:

	conclusion := database_conclusion.Something()

This ‘database_conclusion.Something()’ can also be used for comparisons. Whether that be
in an if-statement:

	if database_conclusion.Something() == conclusion {
		// ...
	}

 Or a switch-statement:

	switch conclusion {
	case database_conclusion.Something():
		// ...
	default:
		// ...
	}


Error

The ‘error’ value for ‘database_conclusion.Type’ represents when an error.

There was multiple ways of giving a ‘database_conclusion.Type’ a value of this kind:

	var err error
	
	// ...
	
	conclusion := database_conclusion.Err(err)

Or:

	conclusion := database_conclusion.Error("an error happened 😞")
Or:

	conclusion := database_conclusion.Errorf("an error happened: %s", msg)

Also, there is a special error that can be created for the when a variable having a
‘nothing’ value is an error (and thus is not loaded):

	conclusion := database_conclusion.ErrorNotLoaded("database_transaction: not loaded")

Why this special error is important is that it can be detected with a type switch. I.e.,:

	var conclusion database_conclusion.Type
	
	// ...
	
	if err := conclusion.Return(); nil != er {
		switch err.(type) {
		case database_conclusion.NotLoaded:
			//@TODO - not loaded
			return err
		default:
			//@TODO - a normal error
			return err
		}
	}

Example

Here is an example of how ‘database_conclusion.Type’ might be used:

	var conclusion database_conclusion.Type = database.Open(driverName, dataSourceName).QueryRow(query).Scan(&id)

	if err := conclusion.Return(); nil != err {
		return err
	}
*/
package database_conclusion
