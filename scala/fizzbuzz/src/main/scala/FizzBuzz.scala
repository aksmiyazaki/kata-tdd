case class FizzBuzz() {
  def fizzbuzz(i: Int): String = {
    val fizzWhenMultipleOfThree = if (i % 3 == 0) "Fizz" else ""
    val buzzWhenMultipleOfFive = if (i % 5 == 0) "Buzz" else ""

    if ((fizzWhenMultipleOfThree + buzzWhenMultipleOfFive) != "") fizzWhenMultipleOfThree + buzzWhenMultipleOfFive
    else i.toString
  }
}
