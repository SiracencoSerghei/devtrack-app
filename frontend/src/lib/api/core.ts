import { client } from './client';
import { ENDPOINTS } from '../constants/api';
import type { EmployeeProfile, OnboardEmployeeRequest } from '../types/api';

export const coreAPI = {
	onboardEmployee: (data: OnboardEmployeeRequest) =>
		client.post<EmployeeProfile>(ENDPOINTS.CORE.EMPLOYEE, data)
};
