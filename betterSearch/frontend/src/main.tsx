import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App'; // Import your App component

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);

root.render(
  <React.StrictMode>
    <App /> {/* Render your App component */}
  </React.StrictMode>
);
