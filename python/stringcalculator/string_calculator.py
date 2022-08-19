from enum import Enum


class StringCalculatorState(Enum):
    EXPECTING_SEPARATOR = 1
    EXPECTING_DIGIT_OR_SEPARATOR = 2
    EXPECTING_DIGIT = 3


class StringCalculator:
    STANDARD_SEPARATORS = [',', "\n"]
    SANITIZED_SEPARATOR = '\0'

    class StringCalculatorErrors(Exception):
        def __init__(self, message=""):
            super().__init__(message)

    def __init__(self):
        self.__input = ""
        self.__standardized_input = ""
        self.__custom_separator = None
        self.__is_custom_separated = False
        self.__current_extraction_state = StringCalculatorState.EXPECTING_DIGIT
        self.__errors = []
        self.__extracted_digits = []

    def add(self, numbers):
        self.__set_state(numbers)

        if self.__is_empty_input():
            return 0
        else:
            self.__standardized_input = self.__standardize_input()
            self.__extract_contained_numbers()
            self.__validate_process()
            self.__check_occurred_errors()
            return sum(self.__extracted_digits)

    def __set_state(self, numbers):
        self.__input = numbers
        if self.__is_custom_separated_input():
            self.__custom_separator = self.__extract_custom_separator(numbers)
            self.__input = self.__extract_content_with_custom_separator(numbers)
            self.__is_custom_separated = True

    def __is_custom_separated_input(self):
        return self.__input.startswith("//")

    def __extract_custom_separator(self, numbers):
        separator_end = numbers.index("\n")
        return numbers[2:separator_end]

    def __extract_content_with_custom_separator(self, numbers):
        start_of_content = numbers.index("\n") + 1
        return numbers[start_of_content:]

    def __is_empty_input(self):
        return self.__input == ""

    def __standardize_input(self):
        if self.__is_custom_separated:
            return self.__input.replace(self.__custom_separator, StringCalculator.SANITIZED_SEPARATOR)
        else:
            return (self.__input
                    .replace(self.STANDARD_SEPARATORS[0], StringCalculator.SANITIZED_SEPARATOR)
                    .replace(self.STANDARD_SEPARATORS[1], StringCalculator.SANITIZED_SEPARATOR))

    def __extract_contained_numbers(self):
        current_digit_as_str = ""

        for idx, element in enumerate(self.__standardized_input):
            if self.__element_is_digit_when_it_is_expected(element):
                current_digit_as_str = self.__process_digit(current_digit_as_str, element)
            elif self.__element_is_separator_when_it_is_expected(element):
                current_digit_as_str = self.__process_separator(current_digit_as_str)
            else:
                current_digit_as_str = self.__treat_error_on_extraction(element, idx)

        if current_digit_as_str != "":
            self.__extracted_digits.append(int(current_digit_as_str))

    def __element_is_digit_when_it_is_expected(self, element):
        return ((element == "-" or element.isdigit()) and
                (self.__current_extraction_state in
                 [StringCalculatorState.EXPECTING_DIGIT, StringCalculatorState.EXPECTING_DIGIT_OR_SEPARATOR]))

    def __process_digit(self, current_digit_as_str, element):
        self.__current_extraction_state = StringCalculatorState.EXPECTING_DIGIT_OR_SEPARATOR
        return current_digit_as_str + element

    def __element_is_separator_when_it_is_expected(self, element):
        return (element == self.SANITIZED_SEPARATOR and
                self.__current_extraction_state == StringCalculatorState.EXPECTING_DIGIT_OR_SEPARATOR)

    def __process_separator(self, current_digit_as_str):
        self.__extracted_digits.append(int(current_digit_as_str))
        self.__current_extraction_state = StringCalculatorState.EXPECTING_DIGIT
        return ""

    def __treat_error_on_extraction(self, element, extract_idx):
        msg = ""
        if (self.__current_extraction_state == StringCalculatorState.EXPECTING_DIGIT_OR_SEPARATOR or
                self.__current_extraction_state == StringCalculatorState.EXPECTING_SEPARATOR):
            msg = f"'{self.__custom_separator}' expected but '{element}' found at position {extract_idx}"
        elif element == self.SANITIZED_SEPARATOR:
            msg = f"Cannot have two separators in sequence"
        self.__errors.append(msg)
        return ""

    def __validate_process(self):
        self.__filter_out_thousands()
        self.__check_trailing_separator()
        negatives = [str(x) for x in self.__extracted_digits if x < 0]
        if len(negatives) > 0:
            self.__errors.append(f"Negative number(s) not allowed: {','.join(negatives)}")

    def __filter_out_thousands(self):
        self.__extracted_digits = [x for x in self.__extracted_digits if x <= 1000]

    def __check_trailing_separator(self):
        if self.__standardized_input[-1] == self.SANITIZED_SEPARATOR:
            self.__errors.append("Input String terminated with Separator")

    def __check_occurred_errors(self):
        if len(self.__errors) > 0:
            raise self.StringCalculatorErrors("\n".join(self.__errors))
