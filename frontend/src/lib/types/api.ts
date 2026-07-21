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

// DTO для Auth
export interface SignUpRequest {
	name: string;
	email: string;
	password: string;
}

export interface LoginRequest {
	email: string;
	password: string;
}

// DTO для Fleet
export interface CreateDriverRequest {
	license_number: string;
	phone?: string;
}

export interface DriverProfile {
	id: string;
	user_id: string;
	license_number: string;
	phone: string;
	status: 'AVAILABLE' | 'IN_TRANSIT' | 'OFF_DUTY';
	updated_at: string;
}

// DTO для Core
export interface OnboardEmployeeRequest {
	first_name: string;
	last_name: string;
	phone: string;
	role_in_company: string;
	department_id?: string;
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

// DTO для Logistics
export interface CreateOrderRequest {
	pickup_address: string;
	delivery_address: string;
}

export interface Order {
	id: string;
	order_number?: string; // наприклад: ORD-1002
	customer_id: string;
	pickup_address: string;
	delivery_address: string;
	status: 'PENDING' | 'ASSIGNED' | 'IN_TRANSIT' | 'DELIVERED';
	created_at: string;
}

export interface CreateOrderRequest {
	pickup_address: string;
	delivery_address: string;
}

export interface UpdateOrderRequest {
	pickup_address?: string;
	delivery_address?: string;
	status?: 'PENDING' | 'ASSIGNED' | 'IN_TRANSIT' | 'DELIVERED';
}

export interface HealthResponse {
	status: 'ok' | 'offline';
}

export interface APIErrorResponse {
	error: string;
}
