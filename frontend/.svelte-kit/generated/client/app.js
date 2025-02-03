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
	() => import('./nodes/12'),
	() => import('./nodes/13'),
	() => import('./nodes/14'),
	() => import('./nodes/15'),
	() => import('./nodes/16'),
	() => import('./nodes/17')
];

export const server_loads = [];

export const dictionary = {
		"/(main)": [8,[3]],
		"/(forms)/auth": [5,[2]],
		"/(main)/create_form": [9,[3]],
		"/(forms)/forget_password": [6,[2]],
		"/(main)/forms": [10,[3]],
		"/(main)/profile/account": [11,[3,4]],
		"/(main)/profile/achievements": [12,[3,4]],
		"/(main)/profile/forms": [13,[3,4]],
		"/(main)/profile/notifications": [14,[3,4]],
		"/(main)/profile/participate": [15,[3,4]],
		"/(main)/profile/teams": [16,[3,4]],
		"/(forms)/reg": [7,[2]],
		"/(main)/teams": [17,[3]]
	};

export const hooks = {
	handleError: (({ error }) => { console.error(error) }),
	
	reroute: (() => {}),
	transport: {}
};

export const decoders = Object.fromEntries(Object.entries(hooks.transport).map(([k, v]) => [k, v.decode]));

export const hash = false;

export const decode = (type, value) => decoders[type](value);

export { default as root } from '../root.svelte';