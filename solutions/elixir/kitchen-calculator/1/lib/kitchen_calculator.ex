defmodule KitchenCalculator do
  def get_volume({_, numeric}), do: numeric

  def(to_milliliter(volume_pair)) do
    {m, q} = volume_pair

    cond do
      m == :cup ->
        {:milliliter, q * 240}

      m == :fluid_ounce ->
        {:milliliter, q * 30}

      m == :teaspoon ->
        {:milliliter, q * 5}

      m == :tablespoon ->
        {:milliliter, q * 15}

      m == :milliliter ->
        {
          :milliliter,
          q
        }

      true ->
        nil
    end
  end

  def from_milliliter(volume_pair, unit) do
    {_, milliliters} = volume_pair

    case unit do
      :cup -> {:cup, milliliters / 240}
      :fluid_ounce -> {:fluid_ounce, milliliters / 30}
      :teaspoon -> {:teaspoon, milliliters / 5}
      :tablespoon -> {:tablespoon, milliliters / 15}
      :milliliter -> {:milliliter, milliliters}
      _ -> nil
    end
  end

  def convert(volume_pair, unit) do
    volume_pair
    |> to_milliliter()
    |> from_milliliter(unit)
  end
end
