import * as React from 'react';
import * as ReactDOM from 'react-dom/client';
import { StyledEngineProvider } from '@mui/material/styles';
import './index.css'
import Dashboard from './Dashboard';
import { BrowserRouter as Router } from "react-router-dom";

ReactDOM.createRoot(document.querySelector("#root")!).render(
  <React.StrictMode>
    <StyledEngineProvider injectFirst>
      <Router>
        <Dashboard />
      </Router>
    </StyledEngineProvider>
  </React.StrictMode>
);