import React from 'react';
import { AppBar, Toolbar, Typography, CssBaseline, useScrollTrigger, Box, Container, Slide, TextField } from '@mui/material';

import logo from './logo.svg';
import './App.css';

const HideOnScroll = (props) => {
  const { children, window } = props;

  const trigger = useScrollTrigger({
    target: window ? window() : undefined,
  });

  return (
    <Slide appear={false} direction="down" in={!trigger}>
      {children}
    </Slide>
  );
};

const CustomNavBar = (props) => {
  return (
    <>
      <HideOnScroll {...props}>
        <AppBar>
          <Toolbar>
            <Typography variant="h6" component="div">
              Google Medium
            </Typography>
          </Toolbar>
        </AppBar>
      </HideOnScroll>
      <Toolbar />
    </>
  );
};

const SearchedItem = () => {
  return (
    <Box>
      <Typography>This is a searched item!</Typography>
    </Box>
  );
};

const SearchBar = () => {
  return (
    <Box
      component="form"
      noValidate
      autoComplete="off"
    >
      <TextField 
        id="outlined-basic" 
        label="Search" 
        variant="outlined" 
        sx={{ width: '100%' }} 
      />
    </Box>
  );
};

function App() {
  return (
    <>
      <CssBaseline />
      <CustomNavBar />
      <Container>
        <Box sx={{ my: 2 }}>
          <SearchBar />
        </Box>
      </Container>
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.js</code> and save to reload.
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>
      </div>
    </>
  );
}

export default App;
