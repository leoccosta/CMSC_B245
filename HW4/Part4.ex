# https://cs.brynmawr.edu/Courses/cs245/fall2022/HWS/hw4a.html
# Almost totally missing from the Enum package are functions that are dependent on position in the list. So write two.

# map_index
# similar to Enum.map
# It takes two arguments: a list and a function.
# Unlike map, the input function should take two parameters: a value from the list and the index of the index of the item in the list.
# Here is an example of the usage of map_index

defmodule EnumIndex do
    @moduledoc """
    This module creates functions missing from the Enum package that are dependent on position in the list.
    Instructions: https://cs.brynmawr.edu/Courses/cs245/fall2022/HWS/hw4a.html
    
    Created: Nov 2022 lcosta
    """

    @doc """
    Based on Enum.map(), with the added functionality of processing a value in a list based on its placement in the list (i.e. index).
    """
    # do something with the value based on index
    def map_index(lst, f) do
        Enum.with_index(lst)
        |> Enum.map(fn x -> f.(elem(x, 0), elem(x, 1)) end)
        #alternatively: use t_map_tx_f(lst, f) https://cs.brynmawr.edu/Courses/cs245/fall2022/Topic06/higher_ex/higher.ex.txt 
    end

    # filter based on index
    @doc """
    Based on Enum.filter(), with the added functionality of filtering a list to only contain values based on its placement in the list (i.e. index) and/or other features.
    """
    def filter_index(lst, f) do
        Enum.with_index(lst)
        |> Enum.filter(fn x -> f.(elem(x, 0), elem(x, 1)) end)
        |> Enum.map(fn x -> elem(x, 0) end)
    end
end

IO.inspect EnumIndex.map_index(["a", "s", "d", "f", "g"], fn (val, idx) -> {val, idx} end)
# output: [{"a", 0}, {"s", 1}, {"d", 2}, {"f", 3}, {"g", 4}]

IO.inspect EnumIndex.filter_index(["a", "s", "d", "f", "g"], fn (val, idx) -> idx>1 and val>"e" end)
# output: ["f", "g"]