export default {
	nav: { home: 'Home', about: 'About', health: 'Health' },
	home: {
		title: 'Welcome to DevTrack',
		area_reserved: 'Restricted Area',
		guest_msg: 'To see registered users and access DevTrack features, you must have an account.',
		user_list: 'Here is the list of users registered in the system:',
		loading_or_empty: 'No users found or loading in progress...'
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
		btn_guest_signup: 'Register new operator'
	}
};
