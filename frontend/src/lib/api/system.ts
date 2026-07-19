import { client } from './client';
import { ENDPOINTS } from '../constants/api';
import type { HealthResponse } from '../types/api';

export const systemAPI = {
	checkHealth: () => client.get<HealthResponse>(ENDPOINTS.SYSTEM.HEALTH)
};
