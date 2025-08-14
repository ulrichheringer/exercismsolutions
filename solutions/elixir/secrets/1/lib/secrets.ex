defmodule Secrets do
  def secret_add(secret), do: &(&1 + secret)

  def secret_subtract(secret), do: &(&1 - secret)

  def secret_multiply(secret), do: &(&1 * secret)

  def secret_divide(secret), do: &(trunc(&1 / secret))

  def secret_and(secret), do: &(Bitwise.band(secret, &1))

  def secret_xor(secret), do: &(Bitwise.bxor(secret, &1))

  def secret_combine(secret_function1, secret_function2), do: &(secret_function2.(secret_function1.(&1)))
end
