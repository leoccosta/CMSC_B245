INSTRUCTIONS: https://cs.brynmawr.edu/Courses/cs245/fall2022/HWS/hw2.html
=============

NAME:
=====
	Leo T. Costa

Programs Files:
===============
    Part 1:
        Part1.go (in directory Part1_go)
        not sure if go.mod (initialized using "go mod init aaa" in Terminal) is a program file

    Part 2:
        Part2.go (in directory Part2_go)
        again, not sure if go.mod (initialized using "go mod init aaa" in Terminal) is a program file
	
How to Compile:
===============
    Part 1:
       go run Part1.go
    Part 2:
       go run Part2.go
       
How to Run:
===========
    "go run [filename]" both compiles and runs
	
Reflection:
===========
    This was my first time sitting with Go for a longer period of time. I find some things, coming from Java, to be pretty annoying, such as the fact that the name goes before the type, but it is also really cool to learn about a language other than Java. The example code found on the course website was really helpful. Something that was tricky was understanding that func String(stru Launch) string {} doesn't work but func (stru Launch) String() string {} does and thus different ways in which functions interact with structures. It also initially broke my brain a little when I realized that changing a struct doesn’t change its copy but changing part of a slice changes the copy of the slice. I actually audibly gasped in surprise and my friend sitting next to me asked me if I was ok, which I think is very funny. Overall, this assignment felt like a good way to ease into the language and generally back into programming. 

Part 3 – Questions:
===========
From part 1: Explain your answer to "Are the original and the copy equal?" As a part of your explanation discuss how a copy of a struct differs from a similarly made copy of an object in Java?
    When the instance of Launch struct d (the copy) was first created, it was equivalent to the Launch struct named a (the original), which would be true whether d was a new instance or an alias. However, when I changed one field of d, that field in a was not changed. If d were an alias, "a" and "d" would just be two names for the same variable being stored in one place in memory. However, structs in Go by default create a new instance, meaning that it creates, in a new place in memory, a variable just like a, and assigns it to d. Thus, any changes happening to d happen only to d and not to a. Meanwhile, Java defaults to aliasing for objects and in the case of Java, changing d would change a. 

From part 2: Consider the original slice and the slice with the first 5 items. Are the items in position 1 of the two slices equal (the answer should be yes)? If you change the value of a field in the struct in position 1 of the original slice, is it still equal to the item in position 1 of the second slice? Explain why the answer to this question makes sense with respect to your answer to the first question.
    The first item in the original slice and in subSlice1 are indeed equal. However, in changing the value of a field in the struct in the first position of the original slice, the value of a field in the struct in the first position of the subslice also changes. While it is often appropriate to assume pass-by-value in Go, a slice in Go is actually an abstraction that uses arrays but makes them more versatile and useful. A slice references an underlying array, and thus when the copy is made, a copy of a reference is made, which ultimately means both names then still reference the same array. The blog post here (https://go.dev/blog/slices-intro) was helpful in clarifying this distinction for me.
