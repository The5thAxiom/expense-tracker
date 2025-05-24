import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import {Auth0Provider} from '@auth0/auth0-react';


const domain = import.meta.env.VITE_DOMAIN
const clientId = import.meta.env.VITE_CLIENT_ID

createRoot(document.getElementById('root')!).render(
  // <StrictMode>
    <Auth0Provider
      domain={domain}
      clientId={clientId}
      authorizationParams={{
        redirect_uri: window.location.origin,
        audience: 'https://expenses.samridh.net/api',
        scope: 'crud:expenses'
      }}
    >
        <App />
    </Auth0Provider>
  //</StrictMode>
  
  
)
