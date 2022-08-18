from enum import Enum


class StringCalculatorState(Enum):
    EXPECTING_SEPARATOR = 1
    EXPECTING_DIGIT_OR_SEPARATOR = 2
    EXPECTING_DIGIT = 3


class StringCalculator:
    class MalformedString(Exception):
        pass

    def add(self, numbers):
        if self.__is_empty_string(numbers):
            return 0
        else:
            separators = [",", "\n"]
            contained_numbers = self.__segment_input(numbers, separators)
            return sum(contained_numbers)

    def __is_empty_string(self, numbers):
        return numbers == ""

    def __segment_input(self, numbers, separators):
        current_state = StringCalculatorState.EXPECTING_DIGIT
        current_digit_as_str = ""
        digits_as_list = []
        for element in numbers:
            if element.isdigit() and (StringCalculatorState.EXPECTING_DIGIT or StringCalculatorState.EXPECTING_DIGIT_OR_SEPARATOR):
                current_digit_as_str += element
                current_state = StringCalculatorState.EXPECTING_DIGIT_OR_SEPARATOR
            elif element in separators and StringCalculatorState.EXPECTING_DIGIT_OR_SEPARATOR:
                digits_as_list += int(current_digit_as_str)
                current_digit_as_str = ""
                current_state = current_state.EXPECTING_DIGIT
            else:
                raise StringCalculator.MalformedString

        return digits_as_list
