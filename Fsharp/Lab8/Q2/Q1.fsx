
#r "nuget: Microsoft.Data.Sqlite"

open Microsoft.Data.Sqlite

// Connect to SQLite database
let connection =
    new SqliteConnection("Data Source=Student.db")

connection.Open()

printfn "Database Connected Successfully!"

// Create Student table
let createTable = connection.CreateCommand()

createTable.CommandText <-
    "CREATE TABLE IF NOT EXISTS Student
    (Id INTEGER PRIMARY KEY, Name TEXT, Age INTEGER)"

createTable.ExecuteNonQuery() |> ignore

printfn "Student Table Created Successfully!"

// Remove old record before inserting
let clean = connection.CreateCommand()

clean.CommandText <-
    "DELETE FROM Student WHERE Id = 1"

clean.ExecuteNonQuery() |> ignore

// Function to display all students
let displayStudents() =
    let read = connection.CreateCommand()

    read.CommandText <- "SELECT * FROM Student"

    let reader = read.ExecuteReader()

    printfn "\nStudent Records:"

    while reader.Read() do
        printfn "ID: %d, Name: %s, Age: %d"
            (reader.GetInt32(0))
            (reader.GetString(1))
            (reader.GetInt32(2))

    reader.Close()

// CREATE operation
let insert = connection.CreateCommand()

insert.CommandText <-
    "INSERT INTO Student (Id, Name, Age)
     VALUES (1, 'Manju', 20)"

insert.ExecuteNonQuery() |> ignore

printfn "\nAfter CREATE:"
displayStudents()

// READ operation
printfn "\nAfter READ:"
displayStudents()

// UPDATE operation
let update = connection.CreateCommand()

update.CommandText <-
    "UPDATE Student SET Age = 21 WHERE Id = 1"

update.ExecuteNonQuery() |> ignore

printfn "\nAfter UPDATE:"
displayStudents()

// DELETE operation
let delete = connection.CreateCommand()

delete.CommandText <-
    "DELETE FROM Student WHERE Id = 1"

delete.ExecuteNonQuery() |> ignore

printfn "\nAfter DELETE:"
displayStudents()

// Close connection
connection.Close()

printfn "\nDatabase Connection Closed!"














