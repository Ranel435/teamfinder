export { matchers } from './matchers.js';

export const nodes = [
	() => import('./nodes/0'),
	() => import('./nodes/1'),
	() => import('./nodes/2'),
	() => import('./nodes/3'),
	() => import('./nodes/4'),
	() => import('./nodes/5'),
	() => import('./nodes/6'),
	() => import('./nodes/7'),
	() => import('./nodes/8'),
	() => import('./nodes/9'),
	() => import('./nodes/10'),
	() => import('./nodes/11'),
	() => import('./nodes/12')
];

export const server_loads = [];

export const dictionary = {
		"/(main)": [8,[3]],
		"/(forms)/auth": [5,[2]],
		"/(forms)/forget_password": [6,[2]],
		"/(main)/profile/account": [9,[3,4]],
		"/(main)/profile/forms": [10,[3,4]],
		"/(main)/profile/notifications": [11,[3,4]],
		"/(main)/profile/teams": [12,[3,4]],
		"/(forms)/reg": [7,[2]]
	};

export const hooks = {
	handleError: (({ error }) => { console.error(error) }),

	reroute: (() => {})
};

export { default as root } from '../root.svelte';