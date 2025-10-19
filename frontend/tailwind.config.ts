import type { Config } from 'tailwindcss';

const config: Config = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    extend: {
      backgroundImage: {
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
        'gradient-conic':
          'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))',
      },
    },
  },
  plugins: [require('@tailwindcss/typography')],
  darkMode: 'class',
  safelist: [
    'z-[70]',
    'z-[71]',
    'z-[72]',
    'z-[73]',
    'z-[74]',
    'z-[75]',
    'z-[76]',
    'z-[77]',
    'z-[78]',
    'z-[79]',
    'z-[80]',
    'z-[81]',
    'z-[82]',
    'z-[83]',
    'z-[84]',
    'z-[85]',
    'z-[86]',
    'z-[87]',
    'z-[88]',
    'z-[89]',
    'z-[90]',
    'z-[91]',
    'z-[92]',
    'z-[93]',
    'z-[94]',
    'z-[95]',
    'z-[96]',
    'z-[97]',
    'z-[98]',
    'z-[99]',
    'z-[100]',
  ],
};
export default config;
