
class ApiStore {
	status = $state('offline');
	message = $state('');

	setStatus(newStatus) {
		this.status = newStatus;
	}

	setMessage(newMessage) {
		this.message = newMessage;
	}
}

export const apiStore = new ApiStore();
