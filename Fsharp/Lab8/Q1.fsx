[<AbstractClass>]

type Shape() =
    abstract member Area : unit -> float

type Circle(radius: float) =
    inherit Shape()

    override this.Area() =
        3.14 * radius * radius

type Rectangle(length: float, width: float) =
    inherit Shape()

    override this.Area() =
        length * width


let circle = Circle(5.0)
let rectangle = Rectangle(4.0, 6.0)

printfn "Area of Circle = %f" (circle.Area())
printfn "Area of Rectangle = %f" (rectangle.Area())
