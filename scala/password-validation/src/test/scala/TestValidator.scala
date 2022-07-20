import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers

class TestValidator extends AnyFlatSpec with Matchers {
  val validator = Validator()

  def matchLengths(targetStringWithSeparator: String, listOfMessages:String*): Boolean = {
    val listLen: Int = listOfMessages.map(_.length).sum
    targetStringWithSeparator.replace("\n", "").length == listLen
  }

  it should "validate that a password has at least 8 chars" in {
    validator.validate("Ab|def78") shouldEqual true
  }

  it should "return an error if password doesnt have at least 8 chars" in {
    try {
      validator.validate("A|23456")
      fail
    }
    catch {
      case e: Validator.ValidationError => e.getMessage shouldEqual Validator.VALIDATION_MESSAGE_LENGTH
    }
  }

  it should "return an error if password doesnt have at least 2 numbers" in {
    try {
      validator.validate("A/cdefgh")
      fail
    }
    catch {
      case e: Validator.ValidationError => e.getMessage shouldEqual Validator.VALIDATION_MESSAGE_NUMBER_OF_DIGITS
    }
  }

  it should "return multiple errors in a single message" in {
    try {
      validator.validate("abcde")
      fail
    }
    catch {
      case e: Validator.ValidationError =>
        val msg = e.getMessage
        msg.contains(Validator.VALIDATION_MESSAGE_NUMBER_OF_DIGITS) shouldBe true
        msg.contains(Validator.VALIDATION_MESSAGE_LENGTH) shouldBe true
        msg.contains(Validator.VALIDATION_MESSAGE_NUMBER_OF_CAPITALS) shouldBe true
        msg.contains(Validator.VALIDATION_MESSAGE_NUMBER_OF_SPECIAL_DIGITS) shouldBe true
        matchLengths(msg,
          Validator.VALIDATION_MESSAGE_NUMBER_OF_DIGITS,
          Validator.VALIDATION_MESSAGE_LENGTH,
          Validator.VALIDATION_MESSAGE_NUMBER_OF_CAPITALS,
          Validator.VALIDATION_MESSAGE_NUMBER_OF_SPECIAL_DIGITS) shouldBe true
    }
  }

  it should "return an error when the password doesnt contain one capital letter" in {
    try {
      validator.validate("12|cdefghij")
      fail
    }
    catch {
      case e: Validator.ValidationError =>
        e.getMessage shouldEqual Validator.VALIDATION_MESSAGE_NUMBER_OF_CAPITALS
    }
  }

  it should "return an error when the password doesnt contain one special digit" in {
    try {
      validator.validate("12Bcdefghij")
      fail
    }
    catch {
      case e: Validator.ValidationError =>
        e.getMessage shouldEqual Validator.VALIDATION_MESSAGE_NUMBER_OF_SPECIAL_DIGITS
    }
  }
}
