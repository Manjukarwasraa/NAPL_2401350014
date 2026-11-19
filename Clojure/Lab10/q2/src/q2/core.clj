

;; Collection 1: List
(println "\n----- LIST OPERATIONS -----")

(def numbers '(1 2 3 4 5 6))

(println "Original List:" numbers)

;; 1. MAP - Square each element
(println "Squares:" (map #(* % %) numbers))

;; 2. FILTER - Select even numbers
(println "Even Numbers:" (filter even? numbers))

;; 3. REDUCE - Calculate sum
(println "Sum:" (reduce + numbers))

;; 4. TAKE - Get first 3 elements
(println "Take 3:" (take 3 numbers))

;; 5. DROP - Remove first 2 elements
(println "Drop 2:" (drop 2 numbers))

;; 6. THREADING MACRO ->>
;; Square numbers and then filter even squares
(println "Threading Result:"
         (->> numbers
              (map #(* % %))
              (filter even?)))


;; Collection 2: Vector
(println "\n----- VECTOR OPERATIONS -----")

(def marks [45 60 75 80 90])

(println "Original Vector:" marks)

;; 1. MAP - Add 5 to each mark
(println "Updated Marks:" (map #(+ % 5) marks))

;; 2. FILTER - Select marks greater than 60
(println "Marks Above 60:" (filter #(> % 60) marks))

;; 3. REDUCE - Find maximum mark
(println "Maximum Marks:" (reduce max marks))

;; 4. TAKE - Get first 3 marks
(println "Take 3:" (take 3 marks))

;; 5. DROP - Remove first 2 marks
(println "Drop 2:" (drop 2 marks))

;; 6. THREADING MACRO ->>
;; Add 5 to marks and then filter marks above 70
(println "Threading Result:"
         (->> marks
              (map #(+ % 5))
              (filter #(> % 70))))

(println "\nSequence transformations completed successfully!")
