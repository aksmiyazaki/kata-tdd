from enum import Enum


class StringCalculatorState(Enum):
    EXPECTING_SEPARATOR = 1
    EXPECTING_DIGIT_OR_SEPARATOR = 2
    EXPECTING_DIGIT = 3


class StringCalculator:
    STANDARD_SEPARATORS = [',', "\n"]
    SANITIZED_SEPARATOR = '\0'

    class MalformedString(Exception):
        def __init__(self, message=""):
            super().__init__(message)

    def __init__(self):
        self.__separator = StringCalculator.SANITIZED_SEPARATOR
        self.__input = ""
        self.__standardized_input = ""
        self.__custom_separator = None
        self.__is_custom_separated = False

    def add(self, numbers):
        self.__set_state(numbers)

        if self.__is_empty_input():
            return 0
        else:
            self.__standardized_input = self.__standardize_input()
            contained_numbers = self.__segment_input()
            return sum(contained_numbers)

    def __set_state(self, numbers):
        self.__input = numbers
        if self.__is_custom_separated_input():
            self.__custom_separator = self.__extract_custom_separator(numbers)
            self.__input = self.__extract_content_with_custom_separator(numbers)
            self.__is_custom_separated = True

    def __is_custom_separated_input(self):
        return self.__input.startswith("//")

    def __extract_custom_separators(self, numbers):
        separator_end = numbers.index("\n")
        return numbers[2:separator_end]

    def __extract_content_with_custom_separator(self, numbers):
        start_of_content = numbers.index("\n") + 1
        return numbers[start_of_content:]

    def __is_empty_input(self):
        return self.__input == ""

    def __segment_input(self):
        current_state = StringCalculatorState.EXPECTING_DIGIT
        current_digit_as_str = ""
        digits_as_list = []
        if self.__standardized_input[-1] == self.SANITIZED_SEPARATOR:
            raise StringCalculator.MalformedString

        for idx, element in enumerate(self.__standardized_input):
            if self.__element_is_digit_when_it_is_expected(current_state, element):
                current_digit_as_str += element
                current_state = StringCalculatorState.EXPECTING_DIGIT_OR_SEPARATOR
            elif self.__element_is_separator_when_it_is_expected(current_state, element):
                digits_as_list.append(int(current_digit_as_str))
                current_digit_as_str = ""
                current_state = current_state.EXPECTING_DIGIT
            else:
                msg = ""
                if (current_state == StringCalculatorState.EXPECTING_DIGIT_OR_SEPARATOR or
                    current_state == StringCalculatorState.EXPECTING_SEPARATOR):
                    msg = f"'{self.__custom_separator}' expected but '{element}' found at position {idx}"
                raise StringCalculator.MalformedString(msg)

        if current_digit_as_str != "":
            digits_as_list.append(int(current_digit_as_str))

        return digits_as_list

    def __element_is_digit_when_it_is_expected(self, current_state, element):
        return element.isdigit() and current_state in [StringCalculatorState.EXPECTING_DIGIT,
                                                       StringCalculatorState.EXPECTING_DIGIT_OR_SEPARATOR]

    def __element_is_separator_when_it_is_expected(self, current_state, element):
        return element == self.SANITIZED_SEPARATOR and current_state == StringCalculatorState.EXPECTING_DIGIT_OR_SEPARATOR

    def __standardize_input(self):
        if self.__is_custom_separated:
            return self.__input.replace(self.__custom_separator, StringCalculator.SANITIZED_SEPARATOR)
        else:
            return (self.__input
                    .replace(self.STANDARD_SEPARATORS[0], StringCalculator.SANITIZED_SEPARATOR)
                    .replace(self.STANDARD_SEPARATORS[1], StringCalculator.SANITIZED_SEPARATOR))

    def __extract_custom_separator(self, numbers):
        separator_end = numbers.index("\n")
        return numbers[2:separator_end]

    def __extract_content_with_custom_separator(self, numbers):
        start_of_content = numbers.index("\n") + 1
        return numbers[start_of_content:]
