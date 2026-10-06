import React, { useEffect, useState } from 'react';

const themeKey = 'go-gpui:theme';

export default function ThemeToggle() {
  const [theme, setTheme] = useState(document.documentElement.dataset.theme || 'dark');
  useEffect(() => {
    document.documentElement.dataset.theme = theme;
    try { localStorage.setItem(themeKey, theme); } catch { /* Keep the theme for this visit. */ }
  }, [theme]);

  const next = theme === 'dark' ? 'light' : 'dark';
  return <button className="theme-toggle" type="button" onClick={() => setTheme(next)}
    aria-label={`Switch to ${next} theme`} title={`Switch to ${next} theme`}>
    <svg viewBox="0 0 24 24" width="18" height="18" fill="none" stroke="currentColor" strokeWidth="1.5" aria-hidden="true">
      {next === 'light' ? <><circle cx="12" cy="12" r="4" /><path d="M12 2v2m0 16v2M2 12h2m16 0h2M5 5l1.5 1.5m11 11L19 19M5 19l1.5-1.5m11-11L19 5" /></>
        : <path d="M20.5 14A9 9 0 0 1 10 3.5 9 9 0 1 0 20.5 14Z" />}
    </svg>
    <span>{next === 'light' ? 'Light' : 'Dark'}</span>
  </button>;
}
