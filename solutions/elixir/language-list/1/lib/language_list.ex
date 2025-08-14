defmodule LanguageList do
  def new(), do: []

  def add(list, language), do: [language | list]

  def remove(list) do
    [_h | t] = list
    t
  end

  def first(list) do
    [h | _t] = list
    h
  end

  def count([]), do: 0
  def count([_ | t]), do: 1 + count(t)

  def functional_list?([]), do: false
  def functional_list?(["Elixir" | _]), do: true
  def functional_list?([_ | tail]), do: functional_list?(tail)
end
