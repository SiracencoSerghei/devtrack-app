export default {
	nav: { home: 'Home', orders: 'Logistics', about: 'About', health: 'Health' },
	home: {
		title: 'Welcome to DevTrack',
		welcome: 'Welcome',
		area_reserved: 'Restricted Area',
		guest_msg: 'To see registered users and access DevTrack features, you must have an account.',
		user_list: 'Here is the list of users registered in the system:',
		loading_or_empty: 'No users found or loading in progress...',
		active_roles: 'Your active roles in the system are:',
		select_department: 'Select the operational department you want to work on today.',
		access_denied_title: 'Access Denied',
		access_denied_msg:
			'You do not have the necessary permissions to access the Organization Panel.',
		back_to_dashboard: 'Back to Dashboard',
		active_roles_table: 'Active Roles'
	},
	server: {
		status: 'Server status:',
		message: 'Message:',
		checking: 'Checking...',
		ok: 'Online (OK)',
		offline: 'Server offline',
		error: 'Server error',
		msg_running: 'The server is running perfectly',
		offline_msg: 'Unable to connect to DevTrack backend'
	},
	errors: {
		no_token: 'Token not found. User not authenticated.',
		session_expired: 'Session expired. Please, log in again.',
		server_error: 'Server error',
		load_failed: 'Failed to load users:'
	},
	about: {
		title: 'About us',
		description:
			'DevTrack is a project monitoring application developed with Go and Svelte. It provides features to manage users, projects, and tasks, offering an intuitive interface to improve team productivity.',
		subtitle: 'This is the devtrack-app built with Go and Svelte!'
	},
	health: {
		title: 'System Status',
		status: 'Status:',
		error_detail: 'Error detail:',
		loading: 'Loading...',
		ok: 'Online (OK)',
		offline: 'Offline'
	},
	login_page: {
		title: 'Sign in to DevTrack',
		email_label: 'Email Address',
		password_label: 'Password',
		btn_submit: 'Sign In',
		no_account: "Don't have an account?",
		register_link: 'Register here',
		invalid_creds: 'Invalid credentials'
	},
	signup_page: {
		title: 'Create a new account',
		name_label: 'Full Name',
		name_placeholder: 'E.g. John Doe',
		email_label: 'Email Address',
		password_label: 'Password',
		btn_submit: 'Sign Up',
		has_account: 'Already have an account?',
		login_link: 'Login here',
		success_msg: 'Registration completed! Redirecting to login...',
		error_msg: 'Error during registration'
	},
	auth: {
		login: 'Login',
		signup: 'Sign Up',
		logout: 'Logout',
		welcome: 'Welcome',
		show_password: 'Show password',
		hide_password: 'Hide password'
	},
	dashboard: {
		system_infra: 'ERP Infrastructure',
		fleet_panel_title: '🏢 Organization Control Panel',
		driver_panel_title: '🚚 Driver Logistics Area',
		driver_panel_desc: 'Manage your trip information, documents, and cargo availability.',
		activation_required: '⚠️ Activation Required:',
		activation_desc:
			"Complete your logistics profile to be included in the company's active fleet.",
		license_label: 'License Number (Required)',
		phone_label: 'Company Phone',
		btn_activate: 'Activate Fleet Profile',
		license_num: 'License N°:',
		phone_num: 'Phone:',
		status_tip: 'Ready to receive orders and shipments from dispatchers.',
		operator: 'Operator',
		email_contact: 'Email Contact',
		loading_profile: 'Loading driver profile...',
		form_conn_error: 'Unable to connect to server',
		btn_guest_login: 'Access ERP',
		btn_guest_signup: 'Register new operator',
		driver_panel_desc_placeholder:
			'Access your trip data, logistics sheets, and manage your driving license.',
		fleet_panel_desc_placeholder:
			'View company staff, monitor active operators, and manage authorizations.'
	},
	orders: {
		title_page: '📦 Order Management',
		btn_open_search: '🔍 Search Order',
		title_create: '📦 Create New Order',
		title_list: '📋 Orders List',
		title_edit: '✏️ Edit Order',
		pickup_label: 'Pickup Address',
		pickup_placeholder: 'e.g., 123 Main St, New York',
		delivery_label: 'Delivery Address',
		delivery_placeholder: 'e.g., 456 Market St, Boston',
		btn_create: 'Create Order',
		btn_creating: 'Creating...',
		btn_edit: 'Edit',
		btn_save: 'Save Changes',
		btn_saving: 'Saving...',
		btn_cancel: 'Cancel',
		search_input_label: 'Order Number or Document Symbol',
		search_placeholder: 'e.g. ORD-1001 or code...',
		btn_search: 'Search',
		btn_searching: 'Searching...',
		not_found: 'Order not found',
		empty_list: 'No orders found. Create your first order above!',
		loading_orders: 'Loading orders...',
		found_result: 'Found Result:',
		action_col: 'Action',
		order_num_col: 'Order No.',
		status: 'Status',
		from: 'From',
		to: 'To',
		created_at: 'Created'
	}
};
