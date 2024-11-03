import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import Stack from '@mui/material/Stack';
import HomeRoundedIcon from '@mui/icons-material/HomeRounded';
import AnalyticsRoundedIcon from '@mui/icons-material/AnalyticsRounded';
import PeopleRoundedIcon from '@mui/icons-material/PeopleRounded';
import InfoRoundedIcon from '@mui/icons-material/InfoRounded';
/* import AssignmentRoundedIcon from '@mui/icons-material/AssignmentRounded';
import SettingsRoundedIcon from '@mui/icons-material/SettingsRounded';
import HelpRoundedIcon from '@mui/icons-material/HelpRounded'; */

import { Link } from "react-router-dom";

const mainListItems = [
  { text: 'Home', to:'/', icon: <HomeRoundedIcon /> },
  { text: 'Analytics', to:'/simple', icon: <AnalyticsRoundedIcon /> },
  { text: 'Clients', to:'/details', icon: <PeopleRoundedIcon /> },
  /* { text: 'Tasks', icon: <AssignmentRoundedIcon /> }, */
];

const secondaryListItems = [
 /*  { text: 'Settings', icon: <SettingsRoundedIcon /> }, */
  { text: 'About', to:'/', icon: <InfoRoundedIcon /> },
  /* { text: 'Feedback', icon: <HelpRoundedIcon /> }, */
];

export default function MenuContent() {
  return (
    <Stack sx={{ flexGrow: 1, p: 1, justifyContent: 'space-between' }}>
      <List dense>
        {mainListItems.map((item, index) => (
          <ListItem key={index} disablePadding sx={{ display: 'block' }}>
            <ListItemButton selected={index === 0}>
              <ListItemIcon>{item.icon}</ListItemIcon>
              <ListItemText primary={<Link to={item.to}>{item.text}</Link>} />
            </ListItemButton>
          </ListItem>
        ))}
      </List>

      <List dense>
        {secondaryListItems.map((item, index) => (
          <ListItem key={index} disablePadding sx={{ display: 'block' }}>
            <ListItemButton>
              <ListItemIcon>{item.icon}</ListItemIcon>
              <ListItemText primary={<Link to={item.to}>{item.text}</Link>} />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
    </Stack>
  );
}
