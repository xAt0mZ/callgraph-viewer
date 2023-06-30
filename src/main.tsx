import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { App } from './App.tsx';

createRoot(document.getElementById('root') as HTMLElement).render(
  <StrictMode>
    <div style={{ width: '100vw', height: '100vh' }}>
      <App />
    </div>
  </StrictMode>
);
