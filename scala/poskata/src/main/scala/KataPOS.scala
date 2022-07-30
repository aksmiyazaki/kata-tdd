final case class KataPOS() {
  val products: Map[String, Double] = Map("12345" -> 7.25, "23456" -> 12.50, "34567" -> 22.2)

  def scan(barCode: String): Double = {
    if (barCode.isEmpty) throw KataPOS.EmptyBarCode()

    if(products.contains(barCode)) products(barCode)
    else throw KataPOS.BarCodeNotFound()
  }

  def total(barCodes: String*): Double = {
    barCodes.map(x => this.scan(x)).sum
  }
}

object KataPOS {
  final case class BarCodeNotFound(private val msg: String ="Error: barcode not found") extends Exception(msg)
  final case class EmptyBarCode(private val msg: String ="Error: empty barcode") extends Exception(msg)
}