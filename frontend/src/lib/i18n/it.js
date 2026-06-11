export default {
	nav: { home: 'Home', about: 'Chi Siamo', health: 'Stato' },
	home: {
		title: 'Benvenuto su DevTrack',
		welcome: 'Benvenuto',
		area_reserved: 'Area Riservata',
		guest_msg:
			'Per vedere gli utenti registrati ed accedere alle funzionalità di DevTrack, devi avere un account.',
		user_list: 'Ecco la lista degli utenti registrati nel sistema:',
		loading_or_empty: 'Nessun utente trovato o caricamento in corso...',
		active_roles: 'I tuoi ruoli attivi nel sistema sono:',
		select_department: 'Seleziona il dipartimento operativo su cui vuoi lavorare oggi.',
		access_denied_title: 'Accesso Negato',
		access_denied_msg:
			"Non disponi dei permessi necessari (Admin/Dispatcher) per accedere al Pannello dell'Organizzazione.",
		back_to_dashboard: 'Torna alla Dashboard',
		active_roles_table: 'Ruoli Attivi'
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
	auth: {
		login: 'Accedi',
		signup: 'Registrati',
		logout: 'Esci',
		welcome: 'Benvenuto',
		show_password: 'Mostra password',
		hide_password: 'Nascondi password'
	},
	dashboard: {
		system_infra: 'Infrastruttura ERP',
		fleet_panel_title: '🏢 Pannello di Controllo Organizzazione',
		driver_panel_title: '🚚 Area Logistica Conducente',
		driver_panel_desc:
			'Gestisci le tue informazioni di viaggio, documenti e reperibilità per i carichi.',
		activation_required: '⚠️ Attivazione Richiesta:',
		activation_desc:
			'Completa il tuo profilo logistico per essere inserito nella flotta attiva aziendale.',
		license_label: 'Numero Patente (Obbligatorio)',
		phone_label: 'Telefono Aziendale',
		btn_activate: 'Attiva Profilo Flotta',
		license_num: 'N° Patente:',
		phone_num: 'Telefono:',
		status_tip: 'Pronto per ricevere ordini e spedizioni dai dispatcher.',
		operator: 'Operatore',
		email_contact: 'Contatto Email',
		loading_profile: 'Caricamento profilo conducente...',
		form_conn_error: 'Impossibile connettersi al server',
		btn_guest_login: "Accedi all'ERP",
		btn_guest_signup: 'Registra nuovo operatore',
		driver_panel_desc_placeholder:
			'Accedi ai tuoi dati di viaggio, fogli di calcolo logistici e inserisci la tua patente.',
		fleet_panel_desc_placeholder:
			'Visualizza lo staff aziendale, monitora gli operatori attivi e gestisci le autorizzazioni.'
	},
};
