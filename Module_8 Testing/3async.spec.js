const { delayedGreeting } = require('./asyncUtils');

describe("Async Function: Delayed Greeting", () => {
    beforeEach(() => {
        jasmine.clock().install(); // Mock timers
    });

    afterEach(() => {
        jasmine.clock().uninstall();
    });

    it("should resolve with the correct greeting message", (done) => {
        delayedGreeting("Alice", 1000).then((message) => {
            expect(message).toBe("Hello, Alice!");
            done();
        });
        jasmine.clock().tick(1000); // Simulate timer advancement
    });

    it("should respect the delay", () => {
        const spy = jasmine.createSpy("greetingSpy");
        delayedGreeting("Alice", 1000).then(spy);
        expect(spy).not.toHaveBeenCalled();
        jasmine.clock().tick(1000);
        expect(spy).toHaveBeenCalled();
    });
});
