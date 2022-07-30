import scala.annotation.tailrec

class StringCalculator() {
  def Add(str: String): Int = {
    if (str.isEmpty) 0
    else {
      val seqOfDelimiters = getDefinedDelimiters(str)

      val stringToProcess = getStringToProcess(str)
      val delimitersOnString = getDelimitersAtString(stringToProcess)

      checkIfStringHasDelimiterAtEnd(stringToProcess, delimitersOnString)
      checkIfStringHasNegativeNumbers(stringToProcess)

      if (delimitersOnString.forall(x => seqOfDelimiters.contains(x.delim))) {
        stringToProcess.split(buildSplitRegex(seqOfDelimiters)).map(_.toInt).filter(_ <= 1000).sum
      } else {
        throwFormattedException(seqOfDelimiters, delimitersOnString)
      }
    }
  }

  private def getDefinedDelimiters(str: String): Seq[String] = {
    if (str.startsWith("//")) {
      val d = str.substring(2, str.indexOf("\n"))
      Seq(d)
    } else Seq(",", "\n")
  }

  private def getStringToProcess(str: String): String = {
    if (str.startsWith("//")) str.substring(str.indexOf("\n") + 1)
    else str
  }

  private def getDelimitersAtString(str: String): Seq[Delimiter] = {
    @tailrec
    def findDelims(str: String, curSep: String, idx: Int, delims: Seq[Delimiter]): Seq[Delimiter] = {
      if (str.isEmpty) {
        val addedSep = checkCurrentSeparator(curSep, idx)
        concatenate(delims, addedSep)
      }
      else {
        val addedSep =
          if (str.charAt(0).isDigit) checkCurrentSeparator(curSep, idx)
          else None

        val nextSep =
          if (addedSep.isDefined || str.charAt(0).isDigit) ""
          else curSep + str.charAt(0)
        findDelims(str.substring(1), nextSep, idx + 1, concatenate(delims, addedSep))
      }
    }

    findDelims(str, "", 0, Seq.empty)
  }

  private def checkCurrentSeparator(curSep: String, idx: Int) = {
    if (curSep.nonEmpty)
      Some(Delimiter(curSep, idx - curSep.length))
    else None
  }

  private def concatenate(delims: Seq[Delimiter], addedSep: Option[Delimiter]) = {
    delims ++ addedSep
  }

  private def checkIfStringHasDelimiterAtEnd(stringToProcess: String, delimitersOnString: Seq[Delimiter]): Unit = {
    if (delimitersOnString.exists(_.position == stringToProcess.length - 1))
      throw StringCalculator.DelimiterAtEndOfInput()
  }

  private def checkIfStringHasNegativeNumbers(stringToProcess: String): Unit = {
    @tailrec
    def fetchNegatives(str: String, curNumber: String, negatives: Seq[String]): Seq[String] = {
      if (str.isEmpty) {
        if (curNumber.nonEmpty) negatives ++ Some(curNumber)
        else negatives
      }
      else {
        val newNumber =
          if (str.charAt(0) == '-' && curNumber.isEmpty) curNumber + "-"
          else if (str.charAt(0).isDigit && curNumber.nonEmpty) curNumber + str.charAt(0)
          else ""

        val addedNegative = if (newNumber.isEmpty && curNumber.nonEmpty) Some(curNumber) else None
        fetchNegatives(str.substring(1), newNumber, negatives ++ addedNegative)
      }
    }

    val negativeNumbers = fetchNegatives(stringToProcess, "", Seq.empty)
    if (negativeNumbers.nonEmpty)
      throw StringCalculator.NegativeNumbersNotAllowed(s"Negative number(s) " +
        s"not allowed: ${negativeNumbers.mkString(", ")}")
  }

  private def buildSplitRegex(seqOfDelimiters: Seq[String]): String = {
    seqOfDelimiters.map(x => if (x == "|") "\\|" else x).mkString("|")
  }

  private def throwFormattedException(seqOfDelimiters: Seq[String], delimitersOnString: Seq[Delimiter]): Nothing = {
    val firstUnexpectedDelim = delimitersOnString.filterNot(x => seqOfDelimiters.contains(x.delim)).head
    val expectedDelimiters = seqOfDelimiters.map(x => if (x == "\n") "\\n" else x).mkString(" or ")
    throw StringCalculator.UnexpectedDelimiter(
      s"$expectedDelimiters expected but ${firstUnexpectedDelim.delim}" +
        s" found at position ${firstUnexpectedDelim.position}"
    )
  }
}

object StringCalculator {
  case class DelimiterAtEndOfInput() extends Exception

  case class UnexpectedDelimiter(private val message: String) extends Exception(message)

  case class NegativeNumbersNotAllowed(private val message: String) extends Exception(message)

  def apply(): StringCalculator = {
    new StringCalculator()
  }
}