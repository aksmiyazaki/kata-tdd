import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers
import org.scalatest.BeforeAndAfter

class TestFizzBuzz extends AnyFlatSpec with Matchers with BeforeAndAfter {
  lazy val fizzBuzzHandler = FizzBuzz()


  "When the number is not a multiple of 3 nor 5" should "return the number in string format" in {
    fizzBuzzHandler.fizzbuzz(1) shouldEqual "1"
  }

  "When the number is multiple of 3" should "return Fizz" in {
    fizzBuzzHandler.fizzbuzz(3) shouldEqual "Fizz"
  }

  "When the number is a multiple of 5" should "return Buzz" in {
    fizzBuzzHandler.fizzbuzz(5) shouldEqual "Buzz"
  }

  "When the number is a multiple of both 3 and 5" should "return FizzBuzz" in {
    fizzBuzzHandler.fizzbuzz(15) shouldEqual "FizzBuzz"
  }
}
