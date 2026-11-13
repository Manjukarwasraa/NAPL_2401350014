#r "nuget: MongoDB.Driver"

open MongoDB.Driver
open MongoDB.Bson

// MongoDB Atlas connection
let connectionString =
    "mongodb+srv://karwasramanju4_db_user:6s9x3jd1kyTd5Hx0@studentcluster.khnntwv.mongodb.net/"

let client = MongoClient(connectionString)

// Access database
let database = client.GetDatabase("StudentDB")

// Check database connection
let result =
    database.RunCommand<BsonDocument>(
        BsonDocument.Parse("{ ping: 1 }")
    )

printfn "MongoDB Atlas Connected Successfully!"

// Student class
type Student() =
    member val Id = 0 with get, set
    member val Name = "" with get, set
    member val Age = 0 with get, set

// Access Student collection
let collection =
    database.GetCollection<Student>("Students")

printfn "Student Collection Created/Accessed Successfully!"

// Function to display all students
let displayStudents() =
    let students =
        collection.Find(Builders<Student>.Filter.Empty).ToList()

    printfn "\nStudent Records:"

    for student in students do
        printfn "ID: %d, Name: %s, Age: %d"
            student.Id student.Name student.Age

// Remove previous record to avoid duplicate ID
let filter =
    Builders<Student>.Filter.Eq((fun (s: Student) -> s.Id), 1)

collection.DeleteOne(filter) |> ignore

// CREATE operation
let student = Student()
student.Id <- 1
student.Name <- "Manju"
student.Age <- 20

collection.InsertOne(student)

printfn "\nAfter CREATE:"
displayStudents()

// READ operation
printfn "\nAfter READ:"
displayStudents()

// UPDATE operation
let update =
    Builders<Student>.Update.Set((fun (s: Student) -> s.Age), 21)

collection.UpdateOne(filter, update) |> ignore

printfn "\nAfter UPDATE:"
displayStudents()

// DELETE operation
collection.DeleteOne(filter) |> ignore

printfn "\nAfter DELETE:"
displayStudents()

printfn "\nMongoDB CRUD Operations Completed Successfully!"
