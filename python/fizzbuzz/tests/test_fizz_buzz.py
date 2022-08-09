import unittest

from fizz_buzz import FizzBuzz


class TestFizzBuzz(unittest.TestCase):
    processor = None

    def setUp(self):
        self.processor = FizzBuzz()

    def test_should_return_the_number_as_string_when_it_is_not_multiple_of_3_or_5(self):
        res = self.processor.fizzbuzz(1)

        self.assertEqual("1", res)

    def test_should_return_Fizz_when_it_is_a_multiple_of_3(self):
        res = self.processor.fizzbuzz(3)

        self.assertEqual("Fizz", res)

    def test_should_return_Buzz_when_it_is_a_multiple_of_5(self):
        res = self.processor.fizzbuzz(5)

        self.assertEqual("Buzz", res)

    def test_should_return_FizBuzz_when_it_is_a_multiple_of_3_and_5(self):
        res = self.processor.fizzbuzz(15)

        self.assertEqual("FizzBuzz", res)


if __name__ == '__main__':
    unittest.main()
