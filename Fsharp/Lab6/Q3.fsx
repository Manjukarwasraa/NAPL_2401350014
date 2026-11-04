open System
open System.Text.RegularExpressions

let reverseString (str: string) =
    new string(Array.rev (str.ToCharArray()))

let isPalindrome (str: string) =
    str = reverseString str

let transformString (str: string) =
    if str |> Seq.forall Char.IsLetter then
        str.ToUpper()
    elif str |> Seq.forall Char.IsDigit then
        str + str
    else
        str

let isValidEmail (email: string) =
    Regex.IsMatch(email, @"^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$")

printfn " STRING PROCESSING PROGRAM "

printfn "\nTask 1: String Reversal and Palindrome Checking"

let strings = ["computer"; "Manju"; "level"]

for s in strings do
    printfn "String: %s" s
    printfn "Reversed String: %s" (reverseString s)
    printfn "Palindrome: %b" (isPalindrome s)

printfn "\nTask 2: String Transformation"

let inputs = ["python"; "5678"; "hello123"]

for s in inputs do
    printfn "Original String: %s" s
    printfn "Transformed String: %s" (transformString s)

printfn "\nTask 3: Email Validation"

let emails = [
    "manju@gmail.com"
    "student456@yahoo.com"
    "wrongemail"
]

for email in emails do
    if isValidEmail email then
        printfn "%s : Valid Email" email
    else
        printfn "%s : Invalid Email" email
