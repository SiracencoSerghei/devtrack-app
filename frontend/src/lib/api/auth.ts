import { client } from './client';
import { ENDPOINTS } from '../constants/api';
import type { AuthResponse, User, SignUpRequest, LoginRequest } from '../types/api';

export const authAPI = {
	signup: (data: SignUpRequest) => client.post<User>(ENDPOINTS.AUTH.SIGNUP, data),
	login: (data: LoginRequest) => client.post<AuthResponse>(ENDPOINTS.AUTH.LOGIN, data),
	getAll: () => client.get<User[]>(ENDPOINTS.AUTH.GET_ALL)
};
