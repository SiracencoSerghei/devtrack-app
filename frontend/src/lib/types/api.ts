export interface User {
	id: string;
	name: string;
	email: string;
	roles: string[];
	created_at: string;
}

export interface AuthResponse {
	token: string;
	user: User;
}

export interface DriverProfile {
	id: string;
	user_id: string;
	license_number: string;
	phone: string;
	status: 'AVAILABLE' | 'IN_TRANSIT' | 'OFF_DUTY';
	updated_at: string;
}

export interface EmployeeProfile {
	id: string;
	user_id: string;
	department_id?: string;
	first_name: string;
	last_name: string;
	phone: string;
	role_in_company: string;
	created_at: string;
}

export interface Order {
	id: string;
	customer_id: string;
	pickup_address: string;
	delivery_address: string;
	status: 'PENDING' | 'ASSIGNED' | 'IN_TRANSIT' | 'DELIVERED';
	created_at: string;
}

export interface HealthResponse {
	status: 'ok' | 'offline';
}

export interface APIErrorResponse {
	error: string;
}
