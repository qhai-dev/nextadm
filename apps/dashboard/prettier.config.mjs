/** @type {import("prettier").Config}*/
const config = {
	proseWrap: 'never',
	printWidth: 120,
	tabWidth: 2,
	trailingComma: 'es5',
	bracketSpacing: true,
	singleQuote: true,
	jsxSingleQuote: true,
	useTabs: true,
	semi: true,
	// tailwindcss
	tailwindStylesheet: './styles/globals.css',
	tailwindFunctions: ['cva'],
	// sort imports
	importOrder: [
		'^(react/(.*)$)|^(react$)',
		'^(next/(.*)$)|^(next$)',
		'<THIRD_PARTY_MODULES>',
		'',
		'^types$',
		'',
		'^[./]',
	],
	importOrderSeparation: false,
	importOrderSortSpecifiers: true,
	importOrderBuiltinModulesToTop: true,
	importOrderParserPlugins: ['typescript', 'jsx', 'decorators-legacy'],
	importOrderMergeDuplicateImports: true,
	importOrderCombineTypeAndValueImports: true,
	plugins: ['@ianvs/prettier-plugin-sort-imports', 'prettier-plugin-tailwindcss'],
};

export default config;
