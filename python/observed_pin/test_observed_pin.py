import unittest

import observed_pin


class TestObservedPin(unittest.TestCase):
    def test_get_pins_one_digit(self):
        # Testing all pins with one digit
        for i in range(10):
            self.assertEqual(
                observed_pin.get_pins(str(i)), observed_pin.adjacent_digits[i]
            )


if __name__ == "__main__":
    unittest.main()
