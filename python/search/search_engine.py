class SearchEngine:
    CITIES = ["Paris", "Budapest", "Skopje", "Rotterdam", "Valencia",
              "Vancouver", "Amsterdam", "Vienna", "Sydney", "New York City",
              "London", "Bangkok", "Hong Kong", "Dubai", "Rome", "Istanbul"]

    def __init__(self):
        self.__case_insensitive_data_base = {el.upper(): el for el in self.CITIES}

    def search(self, string_to_search):
        if self.__is_wildcard_search(string_to_search):
            return self.CITIES

        if self.__is_input_too_small(string_to_search):
            return None

        return self.case_insensitive_match_input_with_data_base(string_to_search)

    def case_insensitive_match_input_with_data_base(self, string_to_search):
        case_insensitive_search = string_to_search.upper()
        return [value
                for key, value
                in self.__case_insensitive_data_base.items() if case_insensitive_search in key]

    def __is_wildcard_search(self, input):
        return input == "*"

    def __is_input_too_small(self, input):
        return len(input) < 2
