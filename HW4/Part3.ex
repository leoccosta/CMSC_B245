# https://cs.brynmawr.edu/Courses/cs245/fall2022/HWS/hw4a.html 
# Write a program using function(s) in Enum to do the following. Given a list of strings, return only those strings that are
# 1. in odd numbered positions in the list; eg. [a,b,c,d,e] -> [b, d] as they are in positions 1 & 3 (start counting from 0)
# 2. are longer than 5 characters

defmodule Odd_Long do
    @moduledoc """
    This module uses functions in Enum to process a list of strings, returning only values at odd indices and that are longer than 5 characters.
    Instructions: https://cs.brynmawr.edu/Courses/cs245/fall2022/HWS/hw4a.html
    
    Created: Nov 2022 lcosta
    """

    @doc """
    Uses:
    • Enum.with_index() to add indices to the list
    • Enum.filter() to eliminate elements where the index is not odd and string length is not > 5
    • Enum.map() to remove the indices added by with_index().
    """
    def filter(lst) do
        lst
        # "Returns the enumerable with each element wrapped in a tuple alongside its index"
        |> Enum.with_index()
        # Returns only elements where the index is odd and string length is > 5:
        |> Enum.filter(fn x -> rem(elem(x, 1), 2) == 1 and String.length(elem(x, 0)) > 5 end) 
        # Returns only the values and discards index part of the tuple:
        |> Enum.map(fn x -> elem(x, 0) end)
    end
end

IO.inspect Odd_Long.filter(["maryy", "janee", "gracee", "lilyy", "alanaa", "quinnn"])