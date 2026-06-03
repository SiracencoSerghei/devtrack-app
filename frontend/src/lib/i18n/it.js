export const it = {
	nav: { home: 'Home', about: 'Chi Siamo', health: 'Stato' },
	auth: { login: 'Accedi', signup: 'Registrati', logout: 'Esci', welcome: 'Ciao' },
	home: {
		title: 'Benvenuto su DevTrack',
		area_reserved: 'Area Riservata',
		guest_msg:
			'Per vedere gli utenti registrati ed accedere alle funzionalità di DevTrack, devi avere un account.',
		user_list: 'Ecco la lista degli utenti registrati nel sistema:',
		loading_or_empty: 'Nessun utente trovato o caricamento in corso...'
	},
	server: {
		status: 'Stato del server:',
		message: 'Messaggio:',
		checking: 'Verifica in corso...',
		ok: 'Funzionante (OK)',
		offline: 'Server offline',
		error: 'Errore del server',
		msg_running: 'Il server funziona perfettamente',
		offline_msg: 'Impossibile connettersi al backend di DevTrack'
	},
	errors: {
		no_token: 'Token non trovato. Utente non autenticato.',
		session_expired: 'Sessione scaduta. Per favore, effettua nuovamente il login.',
		server_error: 'Errore del server',
		load_failed: 'Impossibile caricare gli utenti:'
	},
	about: {
		title: 'Chi siamo',
		description:
			"DevTrack è un'applicazione di monitoraggio dei progetti sviluppata con Go e Svelte. Fornisce funzionalità per gestire utenti, progetti e attività, offrendo un'interfaccia intuitiva per migliorare la produttività del team.",
		subtitle: 'Questa è la devtrack-app costruita con Go e Svelte!'
	},
	health: {
		title: 'Stato del Sistema',
		status: 'Stato:',
		error_detail: 'Dettaglio errore:',
		loading: 'Caricamento...',
		ok: 'Funzionante (OK)',
		offline: 'Non in linea (offline)'
	},
	login_page: {
		title: 'Accedi a DevTrack',
		email_label: 'Indirizzo Email',
		password_label: 'Password',
		btn_submit: 'Accedi',
		no_account: 'Non hai un account?',
		register_link: 'Registrati qui',
		invalid_creds: 'Credenziali non valide'
	},
	signup_page: {
		title: 'Crea un nuovo account',
		name_label: 'Nome completo',
		name_placeholder: 'Es. Mario Rossi',
		email_label: 'Indirizzo Email',
		password_label: 'Password',
		btn_submit: 'Registrati',
		has_account: 'Hai già un account?',
		login_link: 'Accedi qui',
		success_msg: 'Registrazione completata! Verrai reindirizzato al login...',
		error_msg: 'Errore durante la registrazione'
	},
};
