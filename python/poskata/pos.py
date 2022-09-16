from product import Product


class POS():
    class ProductDoesNotExists(Exception):
        def __init__(self):
            super().__init__("barcode not found")

    class CannotScanEmptyBarCode(Exception):
        def __init__(self):
            super().__init__("empty bar code")

    def __init__(self):
        self.__products = None
        self.__build_pos_state()

    def __build_pos_state(self):
        self.__products = [
            Product("12345", 7.25),
            Product("23456", 12.50),
            Product("888", 0.25),
        ]

    def scan(self, bar_code):
        if len(bar_code) == 0:
            raise POS.CannotScanEmptyBarCode

        product = self.__find_product(bar_code)
        if product is not None:
            return self.__format_value(product.price)

    def __find_product(self, bar_code):
        for el in self.__products:
            if el.bar_code == bar_code:
                return el
        raise POS.ProductDoesNotExists

    def __format_value(self, price):
        return "${:.2f}".format(price)

    def total(self, *list_of_bar_codes):
        sum_of_prices = 0
        for bar_code in list_of_bar_codes:
            try:
                product = self.__find_product(bar_code)
                sum_of_prices += product.price
            except:
                pass
        return self.__format_value(sum_of_prices)