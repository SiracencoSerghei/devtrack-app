import { client } from './client';
import { ENDPOINTS } from '../constants/api';
import type { Order, CreateOrderRequest, UpdateOrderRequest } from '../types/api';

export const logisticsAPI = {

	createOrder: (data: CreateOrderRequest) => client.post<Order>(ENDPOINTS.LOGISTICS.ORDERS, data),

	getAllOrders: () => client.get<Order[]>(ENDPOINTS.LOGISTICS.ORDERS),

	searchOrders: (query: string) =>
		client.get<Order[]>(`${ENDPOINTS.LOGISTICS.ORDERS}?query=${encodeURIComponent(query)}`),

	updateOrder: (id: string, data: UpdateOrderRequest) =>
		client.patch<Order>(`${ENDPOINTS.LOGISTICS.ORDERS}/${id}`, data)
};
