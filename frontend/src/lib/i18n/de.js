export default {
	nav: {
		home: 'Startseite',
		orders: 'Logistik',
		about: 'Über uns',
		health: 'Status',
		orders: 'Logistik'
	},
	home: {
		title: 'Willkommen bei DevTrack',
		welcome: 'Willkommen',
		area_reserved: 'Eingeschränkter Zugriff',
		guest_msg:
			'Um registrierte Benutzer zu sehen und auf DevTrack-Funktionen zuzugreifen, müssen Sie ein Konto haben.',
		user_list: 'Hier ist die Liste der im System registrierten Benutzer:',
		loading_or_empty: 'Keine Benutzer gefunden oder Laden läuft...',
		active_roles: 'Ihre aktiven Rollen im System:',
		select_department: 'Wählen Sie die operative Abteilung, in der Sie heute arbeiten möchten.',
		access_denied_title: 'Zugriff verweigert',
		access_denied_msg:
			'Sie haben nicht die erforderlichen Rechte (Admin/Dispatcher), um auf das Organisations-Dashboard zuzugreifen.',
		back_to_dashboard: 'Zurück zum Dashboard',
		active_roles_table: 'Aktive Rollen'
	},
	server: {
		status: 'Serverstatus:',
		message: 'Nachricht:',
		checking: 'Prüfung...',
		ok: 'Funktioniert (OK)',
		offline: 'Server offline',
		error: 'Serverfehler',
		msg_running: 'Der Server läuft einwandfrei',
		offline_msg: 'Verbindung zum DevTrack-Backend fehlgeschlagen'
	},
	errors: {
		no_token: 'Kein Token gefunden. Benutzer ist nicht authentifiziert.',
		session_expired: 'Sitzung abgelaufen. Bitte melden Sie sich erneut an.',
		server_error: 'Serverfehler',
		load_failed: 'Fehler beim Laden der Benutzer:'
	},
	about: {
		title: 'Über uns',
		description:
			'DevTrack ist eine Projektmonitoring-Anwendung, die mit Go und Svelte entwickelt wurde. Sie bietet Funktionen zur Verwaltung von Benutzern, Projekten und Aufgaben mit einer intuitiven Benutzeroberfläche zur Steigerung der Teamproduktivität.',
		subtitle: 'Dies ist die devtrack-app, erstellt mit Go und Svelte!'
	},
	health: {
		title: 'Systemstatus',
		status: 'Status:',
		error_detail: 'Fehlerdetails:',
		loading: 'Laden...',
		ok: 'Funktioniert (OK)',
		offline: 'Offline'
	},
	login_page: {
		title: 'Anmelden bei DevTrack',
		email_label: 'E-Mail-Adresse',
		password_label: 'Passwort',
		btn_submit: 'Anmelden',
		no_account: 'Kein Konto?',
		register_link: 'Hier registrieren',
		invalid_creds: 'Ungültige Anmeldedaten'
	},
	signup_page: {
		title: 'Neues Konto erstellen',
		name_label: 'Vollständiger Name',
		name_placeholder: 'z. B. Max Mustermann',
		email_label: 'E-Mail-Adresse',
		password_label: 'Passwort',
		btn_submit: 'Registrieren',
		has_account: 'Bereits ein Konto?',
		login_link: 'Hier anmelden',
		success_msg: 'Registrierung erfolgreich! Sie werden zur Anmeldeseite weitergeleitet...',
		error_msg: 'Fehler bei der Registrierung'
	},
	auth: {
		login: 'Anmelden',
		signup: 'Registrieren',
		logout: 'Abmelden',
		welcome: 'Willkommen',
		show_password: 'Passwort anzeigen',
		hide_password: 'Passwort verbergen'
	},
	dashboard: {
		system_infra: 'ERP-Infrastruktur',
		fleet_panel_title: '🏢 Organisations-Dashboard',
		driver_panel_title: '🚚 Fahrer-Logistikbereich',
		driver_panel_desc:
			'Am Steuer? Verwalten Sie Fahrtdaten, Dokumente und Verfügbarkeit für Frachten.',
		activation_required: '⚠️ Aktivierung erforderlich:',
		activation_desc:
			'Vervollständigen Sie Ihr Logistikprofil, um in die aktive Flotte aufgenommen zu werden.',
		license_label: 'Führerscheinnummer (Erforderlich)',
		phone_label: 'Geschäftstelefon',
		btn_activate: 'Flottenprofil aktivieren',
		license_num: 'Führerschein-Nr.:',
		phone_num: 'Telefon:',
		status_tip: 'Bereit zum Empfang von Aufträgen und Lieferungen von Disponenten.',
		operator: 'Operator',
		email_contact: 'Kontakt-E-Mail',
		loading_profile: 'Fahrerprofil wird geladen...',
		form_conn_error: 'Verbindung zum Server nicht möglich',
		btn_guest_login: 'In ERP anmelden',
		btn_guest_signup: 'Neuen Operator registrieren',
		driver_panel_desc_placeholder:
			'Greifen Sie auf Fahrtdaten und Logistiktabellen zu und verwalten Sie Ihren Führerschein.',
		fleet_panel_desc_placeholder:
			'Unternehmenspersonal einsehen, aktive Operatoren überwachen und Zugriffsrechte verwalten.'
	},
	orders: {
		title_page: '📦 Auftragsverwaltung',
		btn_open_search: '🔍 Auftrag suchen',
		title_create: '📦 Neuen Auftrag erstellen',
		title_list: '📋 Auftragsliste',
		title_edit: '✏️ Auftrag bearbeiten',
		pickup_label: 'Abholadresse (Pickup)',
		pickup_placeholder: 'z.B. Hauptstraße 1, Berlin',
		delivery_label: 'Lieferadresse (Delivery)',
		delivery_placeholder: 'z.B. Goethestraße 10, München',
		btn_create: 'Erstellen',
		btn_creating: 'Erstellung...',
		btn_edit: 'Bearbeiten',
		btn_save: 'Änderungen speichern',
		btn_saving: 'Speichern...',
		btn_cancel: 'Abbrechen',
		search_input_label: 'Auftragsnummer oder Dokumentsymbol',
		search_placeholder: 'z.B. ORD-1001 oder Code...',
		btn_search: 'Suchen',
		btn_searching: 'Suche...',
		not_found: 'Auftrag nicht gefunden',
		empty_list: 'Keine Aufträge gefunden. Erstellen Sie oben Ihren ersten Auftrag!',
		loading_orders: 'Aufträge werden geladen...',
		found_result: 'Gefundenes Ergebnis:',
		action_col: 'Aktion',
		order_num_col: 'Auftrags-Nr.',
		status: 'Status',
		from: 'Von',
		to: 'Nach',
		created_at: 'Erstellt am'
	}
};
