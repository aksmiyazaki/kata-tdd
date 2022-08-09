class FizzBuzz:
    def fizzbuzz(self, number):
        res = ""
        if number % 3 == 0:
            res += "Fizz"
        if number % 5 == 0:
            res += "Buzz"

        if res == "":
            res = str(number)
        return res
