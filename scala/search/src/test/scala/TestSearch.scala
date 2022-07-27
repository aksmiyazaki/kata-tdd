import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers

class TestSearch extends AnyFlatSpec with Matchers {
  val searchEngine = SearchEngine()

  it should "return empty when cant find a city city" in {
    searchEngine.search("Pary") shouldEqual Seq.empty
  }

  it should "return single element on seq when find a city" in {
    searchEngine.search("Paris") shouldEqual Seq("Paris")
  }

  it should "return empty when search text is less than 2 characters" in {
    searchEngine.search("P") shouldEqual Seq.empty
  }

  it should "return all occurrences of this text in the list of cities" in {
    searchEngine.search("Va") shouldEqual Seq("Valencia", "Vancouver")
  }

  it should "return all occurrences of this text in the list of cities in a case insensitive way" in {
    searchEngine.search("va") shouldEqual Seq("Valencia", "Vancouver")
  }

  it should "return all occurrences of cities with this text in any position" in {
    searchEngine.search("ape") shouldEqual Seq("Budapest")
  }

  it should "return all cities if the text is asterisk" in {
    searchEngine.search("*") shouldEqual Seq("Paris", "Budapest", "Skopje", "Rotterdam", "Valencia",
      "Vancouver", "Amsterdam", "Vienna", "Sydney", "New York City", "London", "Bangkok", "Hong Kong", "Dubai",
      "Rome", "Istanbul")
  }
}
