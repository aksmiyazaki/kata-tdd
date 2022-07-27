case class SearchEngine() {
  val EXISTING_CITIES = Seq("Paris", "Budapest", "Skopje", "Rotterdam", "Valencia", "Vancouver", "Amsterdam", "Vienna",
    "Sydney", "New York City", "London", "Bangkok", "Hong Kong", "Dubai", "Rome", "Istanbul")

  def search(searchText: String): Seq[String] = {
    if (isAllCitiesQuery(searchText)) EXISTING_CITIES
    else if (tooFewCharactersToQuery(searchText)) Seq.empty
    else {
      EXISTING_CITIES.filter(s => s.toUpperCase.contains(searchText.toUpperCase))
    }
  }

  private def isAllCitiesQuery(text: String): Boolean = {
    text == "*"
  }

  private def tooFewCharactersToQuery(text: String): Boolean = {
    text.length < 2
  }
}
