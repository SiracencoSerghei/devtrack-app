import { client } from './client';
import { ENDPOINTS } from '../constants/api';
import type { DriverProfile } from '../types/api';

// Строгий контракт для створення сутності водія флоту
export interface CreateDriverRequest {
	license_number: string;
	phone?: string;
}

export const fleetAPI = {
	createProfile: (data: CreateDriverRequest) =>
		client.post<DriverProfile>(ENDPOINTS.FLEET.DRIVERS, data),
	getProfile: () => client.get<DriverProfile>(ENDPOINTS.FLEET.DRIVERS)
};
