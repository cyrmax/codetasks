adjacent_digits: list[str] = [
    "08",  # 0
    "124",  # 1
    "2135",  # 2
    "326",  # 3
    "4157",  # 4
    "52468",  # 5
    "6539",  # 6
    "748",  # 7
    "85790",  # 8
    "986",  # 9
]


def get_pins(observed: str) -> list[str]:
    if len(observed) == 1:
        return adjacent_digits[int(observed)]
    results: list[str] = []
    first_digit = int(observed[0])
    for pincode in get_pins(observed[1:]):
        for digit in adjacent_digits[first_digit]:
            results.append(f"{digit}{pincode}")
    return results


if __name__ == "__main__":
    print("Enter pin and press enter or use ctrl+c to exit: ")
    while True:
        text = input("> ")
        if not text.isdigit():
            print("Error. Only digits are allowed")
            continue
        pins = get_pins(text)
        for pin in pins:
            print(pin)
