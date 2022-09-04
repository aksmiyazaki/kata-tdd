class Validator:
    def __init__(self):
        self.__validation_errors = []

    def validate(self, password_to_validate):
        self.validate_password_length(password_to_validate)
        self.validate_password_number_of_digits(password_to_validate)
        self.validate_password_number_of_uppercase_letters(password_to_validate)

        if len(self.__validation_errors) > 0:
            return False, "\n".join(self.__validation_errors)
        return True, None

    def validate_password_number_of_uppercase_letters(self, password_to_validate):
        if sum(c.isupper() for c in password_to_validate) < 1:
            self.__validation_errors.append("Password must contain at least one capital letter")

    def validate_password_number_of_digits(self, password_to_validate):
        if sum(c.isdigit() for c in password_to_validate) < 2:
            self.__validation_errors.append("The password must contain at least 2 numbers")

    def validate_password_length(self, password_to_validate):
        if len(password_to_validate) < 8:
            self.__validation_errors.append("Password must be at least 8 characters")
