import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers
import org.scalatest.BeforeAndAfter

class TestFizzBuzz extends AnyFlatSpec with Matchers with BeforeAndAfter {
  lazy val fizzBuzzHandler = FizzBuzz()


  it should "return the number in string format" in {
    fizzBuzzHandler.fizzbuzz(1) shouldEqual "1"
  }

  it should "return Fizz if the number is a multiple of three" in {
    fizzBuzzHandler.fizzbuzz(3) shouldEqual "Fizz"
  }

  it should "return Buzz if the number is a multiple of Five" in {
    fizzBuzzHandler.fizzbuzz(5) shouldEqual "Buzz"
  }

  it should "return FizzBuzz if the number is a multiple of both Three and Five" in {
    fizzBuzzHandler.fizzbuzz(15) shouldEqual "FizzBuzz"
  }
}
