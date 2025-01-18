const { sendNotification } = require('./notificationUtils');

describe("Mocking: Notification Service", () => {
    it("should return success message when notification is sent successfully", () => {
        const mockService = {
            send: jasmine.createSpy("send").and.returnValue(true),
        };
        expect(sendNotification(mockService, "Test message")).toBe("Notification Sent");
        expect(mockService.send).toHaveBeenCalledWith("Test message");
    });

    it("should return failure message when notification fails", () => {
        const mockService = {
            send: jasmine.createSpy("send").and.returnValue(false),
        };
        expect(sendNotification(mockService, "Test message")).toBe("Failed to Send");
        expect(mockService.send).toHaveBeenCalledWith("Test message");
    });
});
