const { toggleVisibility } = require('./domUtils');

describe("Spying: DOM Manipulation", () => {
    it("should toggle visibility from none to block", () => {
        const element = { style: { display: "none" } };
        spyOnProperty(element.style, "display", "set").and.callThrough();

        toggleVisibility(element);

        expect(element.style.display).toBe("block");
    });

    it("should toggle visibility from block to none", () => {
        const element = { style: { display: "block" } };
        spyOnProperty(element.style, "display", "set").and.callThrough();

        toggleVisibility(element);

        expect(element.style.display).toBe("none");
    });
});
