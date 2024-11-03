import Box from '@mui/material/Box';
import Copyright from '../internals/components/Copyright';

import { Routes, Route, Outlet, Link } from "react-router-dom";

import Simple from "../page/Simple";
import Overview from "../page/Overview";
import Details from "../page/Details";

function NoMatch() {
  return (
    <div>
      <h2>Nothing to see here!</h2>
      <p>
        <Link to="/">Go to the home page</Link>
      </p>
    </div>
  );
}

export default function MainGrid() {
  return (
    <Box sx={{ width: '100%', maxWidth: { sm: '100%', md: '1700px' } }}>
      <Routes>
        <Route path="/" element={<Overview />} />
        <Route path="/simple" element={<Simple />} />
        <Route path="/details" element={<Details />} />
        <Route path="*" element={<NoMatch />} />
      </Routes>
      <Copyright sx={{ my: 4 }} />
    </Box>
  );
}
