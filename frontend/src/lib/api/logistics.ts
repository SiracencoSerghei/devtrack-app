import { client } from './client';
import { ENDPOINTS } from '../constants/api';
import type { Order } from '../types/api';

export const logisticsAPI = {
	createOrder: (data: any) => client.post<Order>(ENDPOINTS.LOGISTICS.ORDERS, data),
	getOrder: (id: string) => client.get<Order>(`${ENDPOINTS.LOGISTICS.ORDERS}?id=${id}`)
};
