<script lang="ts">
	import { logisticsAPI } from '$lib/api/logistics';
	import { authStore } from '$lib/auth/auth.svelte';
	import { i18n } from '$lib/i18n/i18n.svelte.js';
	import { APIClientError } from '$lib/api/client';
	import type { Order } from '$lib/types/api';
	import Card from '$lib/components/ui/Card.svelte';
	import Badge from '$lib/components/ui/Badge.svelte';
	import Modal from '$lib/components/ui/Modal.svelte';

	// Стан списку всіх замовлень
	let orders = $state<Order[]>([]);
	let loadingOrders = $state(true);

	// Стан створення нового замовлення
	let pickupAddress = $state('');
	let deliveryAddress = $state('');
	let isCreating = $state(false);
	let createError = $state('');

	// Стан модального вікна ПОШУКУ
	let showSearchModal = $state(false);
	let searchQuery = $state('');
	let searchedOrders = $state<Order[]>([]); // Зберігає масив знайдених замовлень
	let searchError = $state('');
	let isSearching = $state(false);

	// Стан модального вікна РЕДАГУВАННЯ
	let showEditModal = $state(false);
	let editingOrder = $state<Order | null>(null);
	let editPickup = $state('');
	let editDelivery = $state('');
	let isUpdating = $state(false);
	let editError = $state('');

	// Завантаження всіх замовлень користувача
	async function loadOrders() {
		try {
			orders = await logisticsAPI.getAllOrders();
		} catch (err) {
			console.error('Failed to load orders', err);
		} finally {
			loadingOrders = false;
		}
	}

	// Обробка створення замовлення
	async function handleCreateOrder(e: SubmitEvent) {
		e.preventDefault();
		if (isCreating) return;

		createError = '';
		isCreating = true;

		try {
			await logisticsAPI.createOrder({
				pickup_address: pickupAddress,
				delivery_address: deliveryAddress
			});
			pickupAddress = '';
			deliveryAddress = '';
			await loadOrders(); // Оновлюємо таблицю
		} catch (err) {
			if (err instanceof APIClientError) {
				createError = err.message;
			} else {
				createError = i18n.t('server.offline_msg');
			}
		} finally {
			isCreating = false;
		}
	}

	// Універсальний пошук за будь-яким символом / текстом
	async function handleSearchOrder(e: SubmitEvent) {
		e.preventDefault();
		if (!searchQuery.trim() || isSearching) return;

		searchError = '';
		searchedOrders = [];
		isSearching = true;

		try {
			searchedOrders = await logisticsAPI.searchOrders(searchQuery.trim());
			if (searchedOrders.length === 0) {
				searchError = i18n.t('orders.not_found');
			}
		} catch (err) {
			if (err instanceof APIClientError) {
				searchError = err.message;
			} else {
				searchError = i18n.t('server.offline_msg');
			}
		} finally {
			isSearching = false;
		}
	}

	// Відкрити модальне вікно редагування
	function openEditModal(order: Order) {
		editingOrder = order;
		editPickup = order.pickup_address;
		editDelivery = order.delivery_address;
		editError = '';
		showEditModal = true;
	}

	// Зберегти відредаговані дані замовлення
	async function handleUpdateOrder(e: SubmitEvent) {
		e.preventDefault();
		if (!editingOrder || isUpdating) return;

		editError = '';
		isUpdating = true;

		try {
			await logisticsAPI.updateOrder(editingOrder.id, {
				pickup_address: editPickup,
				delivery_address: editDelivery
			});
			showEditModal = false;
			await loadOrders(); // Перезавантажуємо оновлену таблицю
		} catch (err) {
			if (err instanceof APIClientError) {
				editError = err.message;
			} else {
				editError = i18n.t('server.offline_msg');
			}
		} finally {
			isUpdating = false;
		}
	}

	function getBadgeVariant(status: Order['status']) {
		switch (status) {
			case 'DELIVERED': return 'success';
			case 'IN_TRANSIT': return 'warning';
			case 'ASSIGNED': return 'default';
			default: return 'danger';
		}
	}

	$effect(() => {
		if (authStore.isAuthenticated) {
			loadOrders();
		}
	});
</script>

{#if !authStore.isAuthenticated}
	<div class="card access-denied">
		<h2>🚫 {i18n.t('home.access_denied_title')}</h2>
		<p>{i18n.t('home.access_denied_msg')}</p>
		<a href="/login" class="btn btn-primary">{i18n.t('auth.login')}</a>
	</div>
{:else}
	<div class="orders-container">
		<!-- Верхня панель з кнопкою виклику пошуку -->
		<div class="top-bar">
			<h1>{i18n.t('orders.title_page')}</h1>
			
			<button class="btn btn-light search-trigger-btn" onclick={() => (showSearchModal = true)} type="button">
				{i18n.t('orders.btn_open_search')}
			</button>
		</div>

		<!-- Форма створення нового замовлення -->
		<div class="form-wrapper">
			<Card>
				<h3>{i18n.t('orders.title_create')}</h3>
				{#if createError}
					<div class="alert alert-danger">{createError}</div>
				{/if}

				<form onsubmit={handleCreateOrder} class="form-inline">
					<div class="form-group">
						<label for="pickup">{i18n.t('orders.pickup_label')}</label>
						<input type="text" id="pickup" bind:value={pickupAddress} required placeholder={i18n.t('orders.pickup_placeholder')} />
					</div>
					<div class="form-group">
						<label for="delivery">{i18n.t('orders.delivery_label')}</label>
						<input type="text" id="delivery" bind:value={deliveryAddress} required placeholder={i18n.t('orders.delivery_placeholder')} />
					</div>
					<button type="submit" class="btn btn-primary btn-align" disabled={isCreating}>
						{isCreating ? i18n.t('orders.btn_creating') : i18n.t('orders.btn_create')}
					</button>
				</form>
			</Card>
		</div>

		<!-- Таблиця списку замовлень -->
		<Card>
			<h3>{i18n.t('orders.title_list')}</h3>
			
			{#if loadingOrders}
				<p>{i18n.t('orders.loading_orders')}</p>
			{:else if orders.length === 0}
				<p class="empty-text">{i18n.t('orders.empty_list')}</p>
			{:else}
				<div class="table-responsive">
					<table class="erp-table">
						<thead>
							<tr>
								<th>{i18n.t('orders.order_num_col')}</th>
								<th>{i18n.t('orders.from')}</th>
								<th>{i18n.t('orders.to')}</th>
								<th>{i18n.t('orders.status')}</th>
								<th>{i18n.t('orders.created_at')}</th>
								<th>{i18n.t('orders.action_col')}</th>
							</tr>
						</thead>
						<tbody>
							{#each orders as order}
								<tr>
									<td>
										<strong>{order.order_number || `#${order.id.slice(0, 8)}`}</strong>
									</td>
									<td>{order.pickup_address}</td>
									<td>{order.delivery_address}</td>
									<td>
										<Badge type={getBadgeVariant(order.status)}>{order.status}</Badge>
									</td>
									<td>{new Date(order.created_at).toLocaleDateString()}</td>
									<td>
										<button class="btn-action-edit" onclick={() => openEditModal(order)} type="button">
											✏️ {i18n.t('orders.btn_edit')}
										</button>
									</td>
								</tr>
							{/each}
						</tbody>
					</table>
				</div>
			{/if}
		</Card>
	</div>
{/if}

<!-- МОДАЛЬНЕ ВІКНО ПОШУКУ (викликається за кнопкою "🔍 Пошук замовлення") -->
<Modal isOpen={showSearchModal} title={i18n.t('orders.btn_open_search')} onclose={() => (showSearchModal = false)}>
	<form onsubmit={handleSearchOrder} class="modal-form">
		<div class="form-group">
			<label for="modal-search">{i18n.t('orders.search_input_label')}</label>
			<input
				type="text"
				id="modal-search"
				bind:value={searchQuery}
				placeholder={i18n.t('orders.search_placeholder')}
				required
			/>
		</div>
		<button type="submit" class="btn btn-primary" disabled={isSearching}>
			{isSearching ? i18n.t('orders.btn_searching') : i18n.t('orders.btn_search')}
		</button>
	</form>

	{#if searchError}
		<div class="alert alert-danger mt-3">{searchError}</div>
	{/if}

	{#if searchedOrders.length > 0}
		<div class="searched-results">
			<h4>{i18n.t('orders.found_result')}</h4>
			{#each searchedOrders as order}
				<div class="searched-result-card">
					<p><strong>{i18n.t('orders.order_num_col')}:</strong> {order.order_number || order.id}</p>
					<p><strong>{i18n.t('orders.from')}:</strong> {order.pickup_address}</p>
					<p><strong>{i18n.t('orders.to')}:</strong> {order.delivery_address}</p>
					<p><strong>{i18n.t('orders.status')}:</strong> <Badge type={getBadgeVariant(order.status)}>{order.status}</Badge></p>
				</div>
			{/each}
		</div>
	{/if}
</Modal>

<!-- МОДАЛЬНЕ ВІКНО РЕДАГУВАННЯ -->
<Modal isOpen={showEditModal} title={i18n.t('orders.title_edit')} onclose={() => (showEditModal = false)}>
	{#if editError}
		<div class="alert alert-danger">{editError}</div>
	{/if}

	<form onsubmit={handleUpdateOrder} class="modal-form">
		<div class="form-group">
			<label for="edit-pickup">{i18n.t('orders.pickup_label')}</label>
			<input type="text" id="edit-pickup" bind:value={editPickup} required />
		</div>
		<div class="form-group">
			<label for="edit-delivery">{i18n.t('orders.delivery_label')}</label>
			<input type="text" id="edit-delivery" bind:value={editDelivery} required />
		</div>
		<div class="modal-actions">
			<button type="button" class="btn btn-light" onclick={() => (showEditModal = false)}>
				{i18n.t('orders.btn_cancel')}
			</button>
			<button type="submit" class="btn btn-primary" disabled={isUpdating}>
				{isUpdating ? i18n.t('orders.btn_saving') : i18n.t('orders.btn_save')}
			</button>
		</div>
	</form>
</Modal>

<style>
	.orders-container { max-width: 1100px; margin: 2rem auto; }
	.top-bar { display: flex; justify-content: space-between; align-items: center; margin-bottom: 1.5rem; }
	.top-bar h1 { margin: 0; font-size: 1.8rem; color: #0f172a; }
	.search-trigger-btn { font-size: 0.95rem; font-weight: 600; padding: 0.6rem 1.2rem; }
	
	.form-wrapper { margin-bottom: 1.5rem; }
	.form-inline { display: flex; gap: 1rem; flex-wrap: wrap; align-items: flex-end; }
	.form-inline .form-group { flex: 1; min-width: 240px; margin-bottom: 0; }
	.btn-align { height: 42px; }

	.table-responsive { overflow-x: auto; margin-top: 1rem; }
	.erp-table { width: 100%; border-collapse: collapse; text-align: left; font-size: 0.9rem; }
	.erp-table th { background: #f8fafc; padding: 0.75rem 1rem; color: #475569; border-bottom: 1px solid #e2e8f0; }
	.erp-table td { padding: 0.85rem 1rem; border-top: 1px solid #e2e8f0; color: #1e293b; }
	
	.btn-action-edit { background: #f1f5f9; border: 1px solid #cbd5e1; border-radius: 4px; padding: 0.3rem 0.6rem; font-size: 0.8rem; cursor: pointer; transition: background 0.2s; }
	.btn-action-edit:hover { background: #e2e8f0; }

	.modal-form { display: flex; flex-direction: column; gap: 1rem; margin-top: 0.5rem; }
	.modal-actions { display: flex; justify-content: flex-end; gap: 0.75rem; margin-top: 1rem; }
	.searched-results { margin-top: 1rem; display: flex; flex-direction: column; gap: 0.75rem; }
	.searched-result-card { background: #f8fafc; padding: 1rem; border-radius: 6px; border: 1px solid #e2e8f0; }
	.searched-result-card p { margin: 0.3rem 0; font-size: 0.875rem; }
	.empty-text { color: #64748b; padding: 1rem 0; }
	.mt-3 { margin-top: 1rem; }
</style>