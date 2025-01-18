const { fetchAndDisplayUser } = require('./userUtils');

describe("Fetch and Display User", () => {
    it("should display the user name when fetch is successful", async () => {
        const mockService = {
            getUser: jasmine.createSpy("getUser").and.resolveTo({ name: "Alice" }),
        };
        const element = { textContent: "" };

        await fetchAndDisplayUser(mockService, 1, element);

        expect(element.textContent).toBe("Hello, Alice");
    });

    it("should display an error message for invalid user data", async () => {
        const mockService = {
            getUser: jasmine.createSpy("getUser").and.resolveTo({}),
        };
        const element = { textContent: "" };

        await fetchAndDisplayUser(mockService, 1, element);

        expect(element.textContent).toBe("Invalid user data");
    });

    it("should display an error message for fetch failure", async () => {
        const mockService = {
            getUser: jasmine.createSpy("getUser").and.rejectWith(new Error("Fetch error")),
        };
        const element = { textContent: "" };

        await fetchAndDisplayUser(mockService, 1, element);

        expect(element.textContent).toBe("Fetch error");
    });
});
