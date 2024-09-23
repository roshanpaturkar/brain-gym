def roman_to_intefer(s: str) -> int:
    roman_map = {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000
    }
    result = 0
    prev = 0
    for i in range(len(s) - 1, -1, -1):
        num = roman_map[s[i]]
        if num < prev:
            result -= num
        else:
            result += num
        prev = num
    return result


print(roman_to_intefer('III'))  # 3
print(roman_to_intefer('IV'))  # 4
print(roman_to_intefer('IX'))  # 9
print(roman_to_intefer('LVIII'))  # 58
print(roman_to_intefer('MCMXCIV'))  # 1994
