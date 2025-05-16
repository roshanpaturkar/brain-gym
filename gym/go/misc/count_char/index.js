const countChar = (str) => {
    const charCount = {};
    for (let i = 0; i < str.length; i++) {
        const char = str[i];
        if (charCount[char]) {
        charCount[char]++;
        } else {
        charCount[char] = 1;
        }
    }
    return charCount;
}

const str = "hello world";
const result = countChar(str);
console.log(result);