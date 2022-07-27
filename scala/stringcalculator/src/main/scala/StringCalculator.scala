import scala.annotation.tailrec

class StringCalculator() {
  def Add(str: String): Int = {
    if (str.isEmpty) 0
    else {
      val delimiters = getDelimiterFromInput(str)
      val strToProcess = getStringToBeProcessed(str)

      if (stringEndsWithDelimiter(delimiters, strToProcess))
        throw StringCalculator.DelimiterAtEndOfInput()
      else {
        val listOfDigits = getOnlyDigits(delimiters, strToProcess)

        processAddString(delimiters, strToProcess)
      }
    }
  }


  private def getDelimiterFromInput(str: String): List[String] = {
    if (isCustomDelimitedString(str)) {
      getSanitizedListOfDelimiters(str)
    } else {
      List(",", "\n")
    }
  }

  private def getSanitizedListOfDelimiters(str: String) = {
    List(str.substring(2, str.indexOf("\n")))
      .map(el => if (el.contains("|")) {
        el.replace("|", "\\|")
      } else el)
  }

  private def getStringToBeProcessed(str: String): String = {
    if (isCustomDelimitedString(str)) {
      str.substring(str.indexOf("\n") + 1)
    } else {
      str
    }
  }

  private def isCustomDelimitedString(str: String): Boolean = {
    str.startsWith("//")
  }

  private def stringEndsWithDelimiter(delimiters: List[String], strToProcess: String): Boolean = {
    delimiters contains strToProcess.substring(strToProcess.length - 1)
  }

  private def getDelimitersFromString(delimiters: List[String], strToProcess: String): List[Int] = {
    @tailrec
    def getDelimiters(strToProcess: String, curIdx: Int, curDelim:String, gotDelims: List[String]): List[String] = {
      val pass = curDelim
      if(!strToProcess.charAt(curIdx).isDigit) {
      }
    }
  }

  //  private def processAddString(delimiters: List[String], strToProcess: String): Int = {
  //    val listOfDigits: List[Int] = getDigitsFromString()
  //    val listOfDelimiters: List[String] = getListOf
  //  }
}

object StringCalculator {
  case class DelimiterAtEndOfInput() extends Exception

  def apply(): StringCalculator = {
    new StringCalculator()
  }
}