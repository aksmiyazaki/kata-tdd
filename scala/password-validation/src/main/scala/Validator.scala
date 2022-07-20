class Validator {
  def validate(str: String): Boolean = {
    val lengthValidation = validateLength(str)
    val digitValidation = validateStringContainsNDigits(str)
    val capitalLetterValidation = validateStringContainsNCapitals(str)
    val specialDigitValidation = validateStringContainsNSpecialDigits(str)

    val errorMessage = concatenateValidations(lengthValidation, digitValidation, capitalLetterValidation, specialDigitValidation)

    if (errorMessage.length > 0) throw Validator.ValidationError(errorMessage)
    true
  }

  private def validateLength(validationTarget: String, minLength: Int = 8): Option[String] = {
    if (validationTarget.length < minLength) Some(Validator.VALIDATION_MESSAGE_LENGTH)
    else None
  }


  private def validateStringContainsNDigits(validationTarget: String, minDigits: Int = 2): Option[String] = {
    val digitNumber: Int = validationTarget.map(x => if (x.isDigit) 1 else 0).sum

    if (digitNumber < minDigits) Some(Validator.VALIDATION_MESSAGE_NUMBER_OF_DIGITS)
    else None
  }


  private def validateStringContainsNCapitals(validationTarget: String, minCapitals: Int = 1): Option[String] = {
    val capitalLetterNumber: Int = validationTarget.map(x => if (x.isUpper) 1 else 0).sum

    if (capitalLetterNumber < minCapitals) Some(Validator.VALIDATION_MESSAGE_NUMBER_OF_CAPITALS)
    else None
  }

  private def validateStringContainsNSpecialDigits(validationTarget: String, minSpecialDigits: Int = 1): Option[String] = {
    val specialDigitNumber: Int = validationTarget.map(x => if (x.isLetterOrDigit) 0 else 1).sum

    if (specialDigitNumber < minSpecialDigits) Some(Validator.VALIDATION_MESSAGE_NUMBER_OF_SPECIAL_DIGITS)
    else None
  }

  private def concatenateValidations(validations: Option[String]*): String = {
    validations.flatten.mkString("\n")
  }
}

object Validator {
  val VALIDATION_MESSAGE_LENGTH = "Password must be at least 8 characters"
  val VALIDATION_MESSAGE_NUMBER_OF_DIGITS = "The password must contain at least 2 numbers"
  val VALIDATION_MESSAGE_NUMBER_OF_CAPITALS = "The password must contain at least one capital letter"
  val VALIDATION_MESSAGE_NUMBER_OF_SPECIAL_DIGITS = "The password must contain at least one special character"

  case class ValidationError(private val message: String)
    extends Exception(message)

  def apply(): Validator = {
    new Validator()
  }
}
