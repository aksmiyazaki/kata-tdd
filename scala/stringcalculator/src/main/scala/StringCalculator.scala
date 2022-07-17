case class StringCalculator() {
  def Add(str: String): Int = {
    if (str.isEmpty) 0
    else {
      val digits: Seq[Int] = str.split(",").map(_.trim.toInt)
      digits.sum
    }
  }
}
