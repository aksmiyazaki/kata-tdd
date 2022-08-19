import unittest

from string_calculator import StringCalculator


class MyTestCase(unittest.TestCase):
    calc = None

    def setUp(self):
        self.calc = StringCalculator()

    def test_should_return_0_when_empty_input(self):
        self.assertEqual(0, self.calc.add(""))

    def test_should_return_number_as_int_when_single_number_is_input(self):
        self.assertEqual(1, self.calc.add("1"))

    def test_should_return_sum_of_numbers_when_two_are_passed_separated_by_comma(self):
        self.assertEqual(11, self.calc.add("1,10"))

    def test_should_return_sum_of_numbers_when_multiple_are_passed_separated_by_comma(self):
        self.assertEqual(61, self.calc.add("1,10,20,30"))

    def test_should_return_sum_of_numbers_when_using_mixed_comma_and_newline_separators(self):
        self.assertEqual(31, self.calc.add("1,10\n20"))

    def test_should_fail_when_theres_consecutive_separators(self):
        try:
            self.calc.add("1,\n20")
            self.fail("Should've thrown an exception.")
        except StringCalculator.StringCalculatorErrors as err:
            self.assertEqual(str(err), "Cannot have two separators in sequence")

    def test_should_fail_when_theres_separator_in_the_end(self):
        try:
            self.calc.add("1,20,")
            self.fail("Should've thrown an exception.")
        except StringCalculator.StringCalculatorErrors as err:
            self.assertEqual(str(err), "Input String terminated with Separator")

    def test_should_add_numbers_when_using_custom_delimiters(self):
        self.assertEqual(4, self.calc.add("//sep\n1sep3"))

    def test_should_fail_when_invalid_custom_delimiter_is_used(self):
        try:
            self.calc.add("//|\n1|3,4")
            self.fail("Should've thrown an exception.")
        except StringCalculator.StringCalculatorErrors as err:
            self.assertEqual(str(err), "'|' expected but ',' found at position 3")

    def test_should_fail_when_negative_numbers_are_parametrized(self):
        try:
            self.calc.add("1,-4,-9")
            self.fail("Should've thrown an exception.")
        except StringCalculator.StringCalculatorErrors as err:
            self.assertEqual(str(err), "Negative number(s) not allowed: -4,-9")

    def test_should_return_multiple_error_messages_when_input_has_multiple_errors(self):
        try:
            self.calc.add("//|\n1|2,-3")
            self.fail("Should've thrown an exception.")
        except StringCalculator.StringCalculatorErrors as err:
            self.assertEqual(str(err),
                             "'|' expected but ',' found at position 3\nNegative number(s) not allowed: -3")

    def test_should_ignore_numbers_bigger_than_1000_when_input(self):
        self.assertEqual(self.calc.add("2,1001"), 2)


if __name__ == '__main__':
    unittest.main()
