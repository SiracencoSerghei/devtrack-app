import it from './it.js';
import uk from './uk.js';
import en from './en.js';
import de from './de.js';

const dictionaries = { it, uk, en, de };

class I18nManager {
	
	lang = $state(typeof localStorage !== 'undefined' ? localStorage.getItem('lang') || 'it' : 'it');

	t(keyPath) {
		const keys = keyPath.split('.');
		let current = dictionaries[this.lang];

		for (const key of keys) {
			if (current && current[key] !== undefined) {
				current = current[key];
			} else {
				console.warn(`[i18n] Chiave mancante per la lingua "${this.lang}": ${keyPath}`);
				return keyPath;
			}
		}
		return current;
	}

	setLang(newLang) {
		if (dictionaries[newLang]) {
			this.lang = newLang;
			if (typeof localStorage !== 'undefined') {
				localStorage.setItem('lang', newLang);
			}
		}
	}
}

export const i18n = new I18nManager();
