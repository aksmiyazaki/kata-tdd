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

  it should "return the correct summation for a n-digit list" in {
    stringCalculator.Add("5,5 , 5, 5, 5 , 5") shouldEqual 30
  }

  it should "return the correct summation for a list delimited by new line" in {
    stringCalculator.Add("1,2\n3") shouldEqual 6
  }

  it should "raise an exception when there is a delimiter at the end of the list" in {
    assertThrows[StringCalculator.DelimiterAtEndOfInput] {
      stringCalculator.Add("1,2,")
    }
  }

  it should "process the first custom delimiter correctly" in {
    stringCalculator.Add("//;\n1;3") shouldEqual 4
  }

  it should "process the second custom delimiter correctly" in {
    stringCalculator.Add("//|\n1|2|3") shouldEqual 6
  }

  it should "process the third custom delimiter correctly" in {
    stringCalculator.Add("//sep\n2sep5") shouldEqual 7
  }

  it should "raise an exception with the position of the error character" in {
    assertThrows[StringCalculator.UnexpectedDelimiter] {
      stringCalculator.Add("//|\n1|2,3")
    }
  }
}
