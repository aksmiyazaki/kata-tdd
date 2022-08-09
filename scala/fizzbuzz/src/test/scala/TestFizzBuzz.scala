import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers
import org.scalatest.BeforeAndAfter

class TestFizzBuzz extends AnyFlatSpec with Matchers with BeforeAndAfter {
  lazy val fizzBuzzHandler = FizzBuzz()


  "When the number is not a multiple of 3 or 5" should "return the number in string format" in {
    fizzBuzzHandler.fizzbuzz(1) shouldEqual "1"
  }

  "When the number is multiple of three" should "return Fizz" in {
    fizzBuzzHandler.fizzbuzz(3) shouldEqual "Fizz"
  }

  "When the number is a multiple of five" should "return Buzz" in {
    fizzBuzzHandler.fizzbuzz(5) shouldEqual "Buzz"
  }

  "When the number is a multiple of both three and Five" should "return FizzBuzz" in {
    fizzBuzzHandler.fizzbuzz(15) shouldEqual "FizzBuzz"
  }
}
