import unittest

from search_engine import SearchEngine


class MyTestCase(unittest.TestCase):
    searcher = None
    ALL_CITIES = ["Paris", "Budapest", "Skopje", "Rotterdam", "Valencia",
                  "Vancouver", "Amsterdam", "Vienna", "Sydney", "New York City",
                  "London", "Bangkok", "Hong Kong", "Dubai", "Rome", "Istanbul"]

    def setUp(self) -> None:
        self.searcher = SearchEngine()

    def test_should_return_none_when_search_string_is_fewer_than_2_characters(self):
        self.searcher.search("s")

    def test_should_return_all_cities_starting_with_input_when_it_is_bigger_than_2_characters(self):
        self.assertEqual(["Valencia", "Vancouver"], self.searcher.search("Va"))

    def test_should_return_matching_cities_ignoring_case(self):
        self.assertEqual(["Valencia", "Vancouver"], self.searcher.search("VA"))

    def test_should_match_when_input_is_in_any_part_of_city_name(self):
        self.assertEqual(["Budapest"], self.searcher.search("ape"))

    def test_should_return_all_cities_when_asterisk_is_input(self):
        self.assertEqual(self.ALL_CITIES,
                         self.searcher.search("*"))


if __name__ == '__main__':
    unittest.main()
