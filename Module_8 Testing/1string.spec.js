const { capitalize, reverseString } = require('./1string');

describe("String Utilities", () => {
    describe("capitalize", () => {
        it("should return an empty string for an empty input", () => {
            expect(capitalize("")).toBe("");
        });

        it("should capitalize a single-character word", () => {
            expect(capitalize("a")).toBe("A");
        });

        it("should capitalize a normal word", () => {
            expect(capitalize("hello")).toBe("Hello");
        });
    });

    describe("reverseString", () => {
        it("should return an empty string for an empty input", () => {
            expect(reverseString("")).toBe("");
        });

        it("should reverse a normal string", () => {
            expect(reverseString("hello")).toBe("olleh");
        });

        it("should handle palindromes correctly", () => {
            expect(reverseString("madam")).toBe("madam");
        });
    });
});
