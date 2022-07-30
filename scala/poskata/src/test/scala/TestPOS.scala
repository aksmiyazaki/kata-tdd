import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers

class TestPOS extends AnyFlatSpec with Matchers {
  val pos = KataPOS()

  it should "process single bar code" in {
    pos.scan("12345") shouldEqual 7.25
  }

  it should "return an error when barcode not found" in {
    try {
      pos.scan("99999")
    } catch {
      case e: KataPOS.BarCodeNotFound => e.getMessage shouldEqual "Error: barcode not found"
    }
  }

  it should "return an error when barcode is empty" in {
    try {
      pos.scan("")
    } catch {
      case e: KataPOS.EmptyBarCode => e.getMessage shouldEqual "Error: empty barcode"
    }
  }

  it should "return sum of products when total is called" in {
    pos.total("12345", "23456", "34567") shouldEqual 41.95
  }
}
