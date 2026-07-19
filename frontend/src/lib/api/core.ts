import { client } from './client';
import { ENDPOINTS } from '../constants/api';
import type { EmployeeProfile } from '../types/api';

export const coreAPI = {
	onboardEmployee: (data: any) => client.post<EmployeeProfile>(ENDPOINTS.CORE.EMPLOYEE, data),
	getEmployees: () => client.get<EmployeeProfile[]>(ENDPOINTS.CORE.EMPLOYEE)
};
