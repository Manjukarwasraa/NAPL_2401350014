
#r "nuget: MongoDB.Driver"

open MongoDB.Driver
open MongoDB.Bson
open MongoDB.Bson.Serialization.Attributes




type Student() =

    [<BsonId>]
    member val Id = "" with get, set

    member val Name = "" with get, set

    member val Marks = 0.0 with get, set

    member val Course = "" with get, set



let connectionString =
    "mongodb+srv://karwasramanju4_db_user:oXYcP79hzDJBrzoM@studentcluster.khnntwv.mongodb.net/"

let client = new MongoClient(connectionString)

let database = client.GetDatabase("StudentDB")

let collections =
    database.GetCollection<Student>("Students")



let result =
    database.RunCommand<BsonDocument>(
        BsonDocument.Parse("{ ping: 1 }")
    )

printfn "MongoDB Atlas Connected Successfully!"




let CreateStudent (student: Student) = async {

    do! collections.InsertOneAsync(student)
         |> Async.AwaitTask

    printfn "Student added successfully!"
}




let getStudents () = async {

    let! students =
        collections.Find(
            FilterDefinition<Student>.Empty
        ).ToListAsync()
        |> Async.AwaitTask

    return students |> Seq.toList
}




let updateStudent (id: string) (newMarks: float) = async {

    let filter =
        Builders<Student>.Filter.Eq(
            (fun s -> s.Id), id
        )

    let update =
        Builders<Student>.Update.Set(
            (fun s -> s.Marks), newMarks
        )

    let! result =
        collections.UpdateOneAsync(filter, update)
        |> Async.AwaitTask

    if result.MatchedCount > 0 then
        printfn "Student updated successfully!"
    else
        printfn "Student not found!"
}




let deleteStudent (id: string) = async {

    let filter =
        Builders<Student>.Filter.Eq(
            (fun s -> s.Id), id
        )

    let! result =
        collections.DeleteOneAsync(filter)
        |> Async.AwaitTask

    if result.DeletedCount > 0 then
        printfn "Student deleted successfully!"
    else
        printfn "Student not found!"
}




let filterStudent (students: Student list) =
    students
    |> List.filter (fun s -> s.Marks >= 60.0)




let getStudentName (students: Student list) =
    students
    |> List.map (fun s -> s.Name)




let displayStudents (students: Student list) =

    students
    |> List.iter (fun s ->
        printfn "ID: %s | Name: %s | Marks: %.1f | Course: %s"
            s.Id s.Name s.Marks s.Course
    )




let getGrade (marks: float) =

    match marks with

    | m when m >= 90.0 -> "A+"

    | m when m >= 75.0 -> "A"

    | m when m >= 60.0 -> "B"

    | _ -> "Fail"




let main () = async {

    

    let student1 =
        Student(
            Id = "101",
            Name = "Manju",
            Marks = 30.0,
            Course = "BTech"
        )

    let student2 =
        Student(
            Id = "102",
            Name = "Anjali",
            Marks = 70.0,
            Course = "BSC"
        )

    let student3 =
        Student(
            Id = "103",
            Name = "Anja",
            Marks = 60.0,
            Course = "BCom"
        )


 

    let studentList =
        [student1; student2; student3]


    // DELETE OLD RECORDS
    // Prevents duplicate IDs when running again

    do! deleteStudent "101"
    do! deleteStudent "102"
    do! deleteStudent "103"


    // CREATE
    for student in studentList do
        do! CreateStudent student


    // READ
    let! students = getStudents()

    displayStudents students


    // UPDATE

    do! updateStudent "101" 65.0

    let! updatedStudents = getStudents()

    displayStudents updatedStudents


    // FILTER
    let passedStudents =
        filterStudent updatedStudents

    displayStudents passedStudents


    // MAP


    let names =
        getStudentName updatedStudents

    names
    |> List.iter (fun name ->
        printfn "%s" name
    )


   
    printfn "\n--- STUDENT GRADES ---"

    updatedStudents
    |> List.iter (fun s ->
        printfn "%s : %s"
            s.Name (getGrade s.Marks)
    )


    // DELETE

   

    do! deleteStudent "102"


    // FINAL READ


    let! finalStudents = getStudents()

    displayStudents finalStudents


    printfn "\nProgram completed successfully!"
}



main () |> Async.RunSynchronously
