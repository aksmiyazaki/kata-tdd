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
        self.assertRaises(StringCalculator.MalformedString, self.calc.add("1,\n20"))

if __name__ == '__main__':
    unittest.main()
