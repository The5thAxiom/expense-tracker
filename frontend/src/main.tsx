import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import {Auth0Provider} from '@auth0/auth0-react';

const env = import.meta.env;

createRoot(document.getElementById('root')!).render(
  // <StrictMode>
    <Auth0Provider
      domain={env.VITE_AUTH0_DOMAIN}
      clientId={env.VITE_AUTH0_CLIENT_ID}
      authorizationParams={{
        redirect_uri: window.location.origin,
        audience: env.VITE_AUTH0_AUDIENCE,
        scope: 'crud:expenses'
      }}
    >
      <App />
    </Auth0Provider>
  //</StrictMode>
)
