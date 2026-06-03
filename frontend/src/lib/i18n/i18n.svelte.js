import { it } from './it.js';
import { uk } from './uk.js';
import { en } from './en.js';

const translations = { it, uk, en };
let currentLang = $state(
	typeof localStorage !== 'undefined' ? localStorage.getItem('lang') || 'it' : 'it'
);

export const i18n = {
	get lang() {
		return currentLang;
	},

	set lang(newLang) {
		if (translations[newLang]) {
			currentLang = newLang;
			localStorage.setItem('lang', newLang);
		}
	},

	t(path) {
		const keys = path.split('.');
		let translation = translations[currentLang];

		for (const key of keys) {
			if (translation) {
				translation = translation[key];
			} else {
				return path;
			}
		}
		return translation || path;
	}
};
