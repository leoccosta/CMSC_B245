defmodule Student do
    @moduledoc """
    This module contains a struct to represent a student (name and ID#) and has a function that takes in a list of names and returns a list of this type of struct.
    Instructions: https://cs.brynmawr.edu/Courses/cs245/fall2022/HWS/hw4a.html
    
    Created: Nov 2022 lcosta
    """
    
    defstruct id_number: nil, name: nil
    
    @doc """
    A set of recursive functions that goes through a list of students and generates Student structs.
    If a starting ID# is specified, the first student in that list is given that starting number while subsequent students are given the consecutive numbers. 
    If a starting ID# is specified, the first student is given an ID of 0 and subsequent students are given 1, 2, 3, etc. 
    """
    def register_students(lst) do
        # when unspecified, add 0 as the starting point
        register_students(lst, 0)
    end

    def register_students([a], id) do
        [%Student{id_number: id, name: a}]
    end
    def register_students([h|t], id) do
        [%Student{id_number: id, name: h}] ++ register_students(t, id + 1)
    end

    @doc """
    Tests the Student.register_students() function.
    """
    def main() do
        IO.inspect Student.register_students(["mary", "jane", "grace"])
        IO.inspect Student.register_students(["lily", "alana", "quinn"], 5501)
    end
end

defmodule Student_map do
    @moduledoc """
    This module is just like Student except instead of recursion, it uses the Enum.map() function to register students and adds the combine() function.
    """

    defstruct id_number: nil, name: nil

    @doc """
    Registers students from a list like in Student.register_students() except uses Enum.map() instead of recursion.
    """
    def register_students(lst) do
        lst = Enum.with_index(lst)
        Enum.map(lst, fn x -> %Student_map{id_number: elem(x, 1), name: elem(x, 0)} end)
    end

    def register_students(lst, id) do
        lst = Enum.with_index(lst, id)
        Enum.map(lst, fn x -> %Student_map{id_number: elem(x, 1), name: elem(x, 0)} end)
    end

    
    @doc """
    Combines 2 lists of registered students, eliminating duplicates (same name and ID#).
    """
    def combine(lst1, lst2) do
        lst1 ++ lst2
        |> Enum.sort()
        |> Enum.dedup()
    end

    @doc """
    Tests the Student_map.register_students() and Student_map.combine() functions.
    """
    def main() do
        IO.inspect Student_map.register_students(["mary", "jane", "grace"])
        IO.inspect Student_map.register_students(["lily", "alana", "quinn"], 5501)
        lst1 = Student_map.register_students(["lily", "alana", "quinn"], 5501)
        lst2 = Student_map.register_students(["lily", "lily", "alana", "quinn"], 5500)
        IO.inspect Student_map.combine(lst1, lst2)
    end
end

#Student.main()
Student_map.main()