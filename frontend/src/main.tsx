import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import {Auth0Provider} from '@auth0/auth0-react';

console.log(window.location.origin)
createRoot(document.getElementById('root')!).render(
  // <StrictMode>
    <Auth0Provider
      domain='xxxx'
      clientId='xxxx'
      authorizationParams={{redirect_uri: window.location.origin}}
    >
        <App />
    </Auth0Provider>
  //</StrictMode>
  
  
)
