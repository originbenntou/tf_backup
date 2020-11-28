import { createMuiTheme } from '@material-ui/core/styles';
import red from '@material-ui/core/colors/red';

// Create a theme instance.
const theme = createMuiTheme({
  palette: {
    primary: {
      main: '#000000',
    },
    secondary: {
      main: '#FFF176',
    },
    error: {
      main: red.A400,
    },
    background: {
      default: '#E0E0E0',
    },
  },
});

export default theme;
