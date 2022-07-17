import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers

class TestStringCalculator extends AnyFlatSpec with Matchers {
  lazy val stringCalculator = StringCalculator()

  it should "return the correct summation for empty string" in {
    stringCalculator.Add("") shouldEqual 0
  }

  it should "return the correct summation for a single digit in the list" in {
    stringCalculator.Add("2") shouldEqual 2
  }

  it should "return the correct summation for a double digit in the list" in {
    stringCalculator.Add("2,3") shouldEqual 5
  }

  it should "return the correct summation for a double digit in the list with more than two digits" in {
    stringCalculator.Add("5,5 , 5, 5, 5 , 5") shouldEqual 30
  }
}
