defmodule Flattening do
    @moduledoc """
    This module allows for lists of any depth to be flattened. For example, [12, 6, [1, [2]],[3,4]] would become [12, 6, 1, 2, 3, 4].
    Instructions: https://cs.brynmawr.edu/Courses/cs245/fall2022/HWS/hw4a.html
    
    Created: Nov 2022 lcosta
    """
    
    @doc """
    A set of recursive functions that does the list flattening. If the list passed is empty, return the list. Otherwise, address the head and tail seperately and flatten them recursively.
    """
    def flatten([]) do
        [] # if the list is empty, return the list
    end
    
    def flatten([h|t]) do
        flattened = []

        # flatten the head of the list (if it is a list)
        # --> save it as a variable
        x = if is_list(h) do 
            flatten(h)
        else
            [h|flattened]
        end

        # flatten the tail of the list and append to flattened head
        flattened = x ++ flatten(t)

        # return the flattened head + the flattened tail
        flattened
    end
end

list = [12, 6, [1, [2]],[3,4]]
IO.inspect Flattening.flatten(list)