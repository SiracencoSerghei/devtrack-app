import { client } from './client';
import { ENDPOINTS } from '../constants/api';
import type { AuthResponse, User } from '../types/api';

export const authAPI = {
	signup: (data: any) => client.post<User>(ENDPOINTS.AUTH.SIGNUP, data),
	login: (data: any) => client.post<AuthResponse>(ENDPOINTS.AUTH.LOGIN, data),
	getAll: () => client.get<User[]>(ENDPOINTS.AUTH.GET_ALL)
};
