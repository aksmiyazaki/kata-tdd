import unittest

from password_validation import Validator


class MyTestCase(unittest.TestCase):
    validator = None

    def setUp(self):
        self.validator = Validator()

    def test_should_validate_when_meets_all_constraints(self):
        self.__evaluate_validator_test_case("A,345678", True, None)

    def test_should_raise_error_when_password_is_less_than_8_chars(self):
        self.__evaluate_validator_test_case("A,34567", False, "Password must be at least 8 characters")

    def test_should_raise_error_when_theres_not_2_numbers(self):
        self.__evaluate_validator_test_case("A,bcdef7", False, "The password must contain at least 2 numbers")

    def test_should_raise_multiple_errors_when_fail_multiple_validations(self):
        self.__evaluate_validator_test_case("Abcdef",
                                            False,
                                            ("Password must be at least 8 characters\n"
                                             "The password must contain at least 2 numbers"))

    def test_should_raise_error_when_doesnt_contain_one_capital(self):
        self.__evaluate_validator_test_case("a,345678", False, "Password must contain at least one capital letter")

    def __evaluate_validator_test_case(self, password_input, expected_return, expected_error_message):
        res, constraints_not_met = self.validator.validate(password_input)
        self.assertEqual(res, expected_return)
        self.assertEqual(constraints_not_met, expected_error_message)


if __name__ == '__main__':
    unittest.main()
