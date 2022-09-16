import unittest

import pytest as pytest

from pos import POS


class MyTestCase(unittest.TestCase):
    pos = None

    def setUp(self) -> None:
        self.pos = POS()

    def test_should_return_value_when_product_exists(self):
        price_as_string = self.pos.scan("12345")
        self.assertEqual("$7.25", price_as_string)

    def test_should_return_another_value_when_product_exists(self):
        price_as_string = self.pos.scan("23456")
        self.assertEqual("$12.50", price_as_string)

    def test_should_fail_when_product_does_not_exists(self):
        with pytest.raises(POS.ProductDoesNotExists):
            _ = self.pos.scan("999")

    def test_should_fail_when_empty_bar_code(self):
        with pytest.raises(POS.CannotScanEmptyBarCode):
            _ = self.pos.scan("")

    def test_should_sum_all_products_when_exists(self):
        total_as_string = self.pos.total("12345", "23456", "888")
        self.assertEqual("$20.00", total_as_string)

    def test_should_ignore_products_when_they_doesnt_exists(self):
        total_as_string = self.pos.total("12345", "111", "888")
        self.assertEqual("$7.50", total_as_string)

    def test_should_ignore_all_products_when_they_doesnt_exists(self):
        total_as_string = self.pos.total("123", "111", "444")
        self.assertEqual("$0.00", total_as_string)

if __name__ == '__main__':
    unittest.main()
