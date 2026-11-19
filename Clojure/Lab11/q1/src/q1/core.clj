(println "----- JAVA INTEROPERABILITY -----")

;; 1. Calling Java Static Methods

(println "\n1. Java Static Methods")

(println "Absolute Value:" (Math/abs -25))
(println "Square Root:" (Math/sqrt 64))
(println "Current Time:" (System/currentTimeMillis))


;; 2. Creating Java Object Instances

(println "\n2. Creating Java Objects")

(def my-list (new java.util.ArrayList))

(def my-map (new java.util.HashMap))

(println "ArrayList Created:" my-list)
(println "HashMap Created:" my-map)


;; 3. Invoking Instance Methods

(println "\n3. Instance Methods")

;; ArrayList methods
(.add my-list "Apple")
(.add my-list "Banana")
(.add my-list "Mango")

(println "ArrayList:" my-list)
(println "First Element:" (.get my-list 0))
(println "List Size:" (.size my-list))

;; HashMap methods
(.put my-map "Name" "Manju")
(.put my-map "Age" 20)

(println "HashMap:" my-map)
(println "Name:" (.get my-map "Name"))
(println "Map Size:" (.size my-map))


;; 4. Converting Clojure and Java Collections

(println "\n4. Collection Conversion")

;; Clojure Collection to Java Collection
(def clj-vector [10 20 30])

(def java-list
  (java.util.ArrayList. clj-vector))

(println "Clojure Vector:" clj-vector)
(println "Java ArrayList:" java-list)

;; Java Collection to Clojure Collection
(def new-clj-list
  (vec java-list))

(println "Converted Clojure Vector:" new-clj-list)

(println "\nJava Interoperability Demonstrated Successfully!")
