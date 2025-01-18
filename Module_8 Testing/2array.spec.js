const { getElement } = require('./arrayUtils');

describe("Array Index Error Handling", () => {
    it("should retrieve an element for a valid index", () => {
        const arr = [1, 2, 3];
        expect(getElement(arr, 1)).toBe(2);
    });

    it("should throw an error for a negative index", () => {
        const arr = [1, 2, 3];
        expect(() => getElement(arr, -1)).toThrowError("Index out of bounds");
    });

    it("should throw an error for an out-of-range index", () => {
        const arr = [1, 2, 3];
        expect(() => getElement(arr, 5)).toThrowError("Index out of bounds");
    });
});
